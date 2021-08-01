/*
Copyright © 2021 GUILLAUME FOURNIER

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utrace

import (
	"debug/elf"
	"regexp"
	"sort"
	"time"

	"github.com/pkg/errors"
)

const (
	// MaxUserSymbolsCount is the max number of symbols probed at the same time
	MaxUserSymbolsCount = uint32(100)
	// MaxKernelSymbolsCount is the max number of kernel symbols probed at the same time
	MaxKernelSymbolsCount = uint32(500)
)

var (
	// NotEnoughDataErr indicates that the data retrieved from the perf map is not long enough
	NotEnoughDataErr = errors.New("not enough data")
	// EmptyPatternsErr indicates that both the kernel function and userspace symbol patterns are empty
	EmptyPatternsErr = errors.New("empty patterns")
	// EmptyBinaryPathErr indicates that a userspace symbol pattern was provided but without a binary path
	EmptyBinaryPathErr = errors.New("empty binary path")
	// NoPatternProvidedErr indicates that no function pattern was provided
	NoPatternProvidedErr = errors.New("no function pattern was provided")
)

// StackID is a unique identifier used to select a stack trace
type StackID uint32

// CombinedID is a unique identifier used to select a combined stack trace (user and kernel)
type CombinedID uint64

// FuncID is the id of a function traced in kernel space
type FuncID int32

// SymbolAddr is the address of a symbol
type SymbolAddr uint64

// PathMax - Maximum path length of the binary path handled by utrace
const PathMax = 350

// Options contains the parameters of UTrace
type Options struct {
	Binary            string
	FuncPattern       *regexp.Regexp
	KernelFuncPattern *regexp.Regexp
	Latency           bool
	StackTraces       bool
}

func (o Options) check() error {
	if o.FuncPattern == nil && o.KernelFuncPattern == nil {
		return EmptyPatternsErr
	}
	if o.FuncPattern != nil && len(o.Binary) == 0 {
		return EmptyBinaryPathErr
	}
	return nil
}

var (
	// UserSymbolNotFound is used to notify that a userspace symbol could not be resolved
	UserSymbolNotFound = elf.Symbol{Name: "[user_symbol_not_found]"}
	// KernelSymbolNotFound is used to notify that a kernel space symbol could not be resolved
	KernelSymbolNotFound = elf.Symbol{Name: "[kernel_symbol_not_found]"}
)

func symbolNotFound(funcType FuncType) elf.Symbol {
	if funcType == Kernel {
		return KernelSymbolNotFound
	}
	return UserSymbolNotFound
}

// Report contains the statistics generated by UTRace
type Report struct {
	stackTracerCount StackTraceCount

	orderedByLatency []Func
	orderedByHits    []Func
	functions        map[FuncID]Func
	duration         time.Duration
}

// NewReport instanciates a new Report
func NewReport(duration time.Duration) Report {
	return Report{
		functions: make(map[FuncID]Func),
		duration:  duration,
	}
}

// GetStackTraceCount returns the total number of stack traces collected from the kernel
func (r *Report) GetStackTraceCount() StackTraceCount {
	return r.stackTracerCount
}

// GetDuration returns the duration of the trace
func (r *Report) GetDuration() time.Duration {
	return r.duration
}

// GetFunc returns a Func by its FuncID
func (r *Report) GetFunc(id FuncID, funcType FuncType) Func {
	if id == -1 {
		return NewFunc(symbolNotFound(funcType))
	}
	ret, ok := r.functions[id]
	if !ok {
		return NewFunc(symbolNotFound(funcType))
	}
	return ret
}

// GetFunctionsByHits returns the list of traced functions ordered by their hits count
func (r *Report) GetFunctionsByHits() []Func {
	if len(r.orderedByHits) == 0 {
		for _, f := range r.functions {
			r.orderedByHits = append(r.orderedByHits, f)
		}
		sort.Sort(orderByHits(r.orderedByHits))
	}
	return r.orderedByHits
}

// GetFunctionsByLatency returns the list of traced functions ordered by their latency
func (r *Report) GetFunctionsByLatency() []Func {
	if len(r.orderedByLatency) == 0 {
		for _, f := range r.functions {
			r.orderedByLatency = append(r.orderedByLatency, f)
		}
		sort.Sort(orderByLatency(r.orderedByLatency))
	}
	return r.orderedByLatency
}

// StackTraceCount holds the amount of traces that were collected or lost
type StackTraceCount struct {
	Kernel     uint64
	LostKernel uint64
	User       uint64
	LostUser   uint64
}

// FuncType is the type of a traced function
type FuncType string

const (
	// Kernel is used for kernel symbols
	Kernel FuncType = "kernel"
	// User is used for user space symbols
	User FuncType = "user"
)

// Func contains the data collected by utrace for a function
type Func struct {
	Type           FuncType
	Symbol         elf.Symbol
	Count          uint64
	AverageLatency time.Duration

	stackTraces   map[CombinedID]*StackTrace
	orderedByHits []*StackTrace
}

// NewFunc instanciates a new Func
func NewFunc(symbol elf.Symbol) Func {
	return Func{
		Symbol:      symbol,
		stackTraces: make(map[CombinedID]*StackTrace),
	}
}

// GetStackTrace returns a stack trace from its StackID
func (f *Func) GetStackTrace(stackID CombinedID) *StackTrace {
	return f.stackTraces[stackID]
}

// GetStackTracesByHits returns the list of stack traces by hits count
func (f *Func) GetStackTracesByHits() []*StackTrace {
	if len(f.orderedByHits) == 0 {
		for _, trace := range f.stackTraces {
			f.orderedByHits = append(f.orderedByHits, trace)
		}
		sort.Sort(orderTraceByHits(f.orderedByHits))
	}
	return f.orderedByHits
}

type orderByHits []Func

func (obh orderByHits) Len() int           { return len(obh) }
func (obh orderByHits) Swap(i, j int)      { obh[i], obh[j] = obh[j], obh[i] }
func (obh orderByHits) Less(i, j int) bool { return obh[i].Count > obh[j].Count }

type orderByLatency []Func

func (obl orderByLatency) Len() int           { return len(obl) }
func (obl orderByLatency) Swap(i, j int)      { obl[i], obl[j] = obl[j], obl[i] }
func (obl orderByLatency) Less(i, j int) bool { return obl[i].AverageLatency > obl[j].AverageLatency }

// StackTrace contains the ordered list of symbols of a stack trace
type StackTrace struct {
	Count            int
	UserStacktrace   []StackTraceNode
	KernelStackTrace []StackTraceNode
}

type orderTraceByHits []*StackTrace

func (otbh orderTraceByHits) Len() int           { return len(otbh) }
func (otbh orderTraceByHits) Swap(i, j int)      { otbh[i], otbh[j] = otbh[j], otbh[i] }
func (otbh orderTraceByHits) Less(i, j int) bool { return otbh[i].Count > otbh[j].Count }

// NewStackTrace returns a new StackTrace initialized with the provided count
func NewStackTrace(count int) *StackTrace {
	return &StackTrace{
		Count: count,
	}
}

// StackTraceNode represents a node of a stack trace
type StackTraceNode struct {
	Type   FuncType
	FuncID FuncID
	Offset SymbolAddr
	Symbol elf.Symbol
}

type kernelCounter struct {
	CumulatedTime uint64
	Count         uint64
}

// TraceEvent is a kernel trace event retrieved from a perf map
type TraceEvent struct {
	Pid           uint32
	Tid           uint32
	UserStackID   StackID
	KernelStackID StackID
	FuncID        FuncID
}

// UnmarshalBinary unmarshals the binary representation of a trace event
func (te *TraceEvent) UnmarshalBinary(data []byte) (int, error) {
	if len(data) < 24 {
		return 0, NotEnoughDataErr
	}

	te.Pid = ByteOrder.Uint32(data[0:4])
	te.Tid = ByteOrder.Uint32(data[4:8])
	te.UserStackID = StackID(ByteOrder.Uint32(data[8:12]))
	te.KernelStackID = StackID(ByteOrder.Uint32(data[12:16]))
	te.FuncID = FuncID(ByteOrder.Uint32(data[16:20]))
	return 24, nil
}
