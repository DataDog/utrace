package utrace

import (
	"debug/elf"
	"fmt"
	"os"
	"testing"
	// "github.com/google/pprof/profile"
	"github.com/stretchr/testify/require"
)

func MakeSimpleStack() map[CombinedID]*StackTrace {
	stackTraces := make(map[CombinedID]*StackTrace)
	combinedID := CombinedID(1)<<32 + CombinedID(2)
	stackTrace := NewStackTrace(1)
	stackTraces[combinedID] = stackTrace

	symbolFoo := elf.Symbol{Name: "Foo"}
	symbolBar := elf.Symbol{Name: "Bar"}
	symbolMain := elf.Symbol{Name: "Main"}
	symbolKMalloc := elf.Symbol{Name: "kmalloc"}

	stackTrace.UserStacktrace = append(stackTrace.UserStacktrace, StackTraceNode{
		Type:   User,
		Symbol: symbolFoo,
		FuncID: 0,
		Offset: 0,
	})
	stackTrace.UserStacktrace = append(stackTrace.UserStacktrace, StackTraceNode{
		Type:   User,
		Symbol: symbolBar,
		FuncID: 1,
		Offset: 0,
	})
	stackTrace.UserStacktrace = append(stackTrace.UserStacktrace, StackTraceNode{
		Type:   User,
		Symbol: symbolMain,
		FuncID: 2,
		Offset: 0,
	})
	stackTrace.KernelStackTrace = append(stackTrace.KernelStackTrace, StackTraceNode{
		Type:   User,
		Symbol: symbolKMalloc,
		FuncID: 3,
		Offset: 0,
	})

	return stackTraces
}

func MakeRepeatedStack() map[CombinedID]*StackTrace {
	stackTraces := make(map[CombinedID]*StackTrace)
	combinedID := CombinedID(1)<<32 + CombinedID(2)
	stackTrace := NewStackTrace(1)
	stackTraces[combinedID] = stackTrace

	symbolFoo := elf.Symbol{Name: "Foo"}
	symbolBar := elf.Symbol{Name: "Bar"}
	symbolKMalloc := elf.Symbol{Name: "kmalloc"}

	stackTrace.UserStacktrace = append(stackTrace.UserStacktrace, StackTraceNode{
		Type:   User,
		Symbol: symbolFoo,
		FuncID: 0,
		Offset: 0,
	})
	stackTrace.UserStacktrace = append(stackTrace.UserStacktrace, StackTraceNode{
		Type:   User,
		Symbol: symbolBar,
		FuncID: 1,
		Offset: 0,
	})
	stackTrace.UserStacktrace = append(stackTrace.UserStacktrace, StackTraceNode{
		Type:   User,
		Symbol: symbolFoo,
		FuncID: 0,
		Offset: 0,
	})
	stackTrace.KernelStackTrace = append(stackTrace.KernelStackTrace, StackTraceNode{
		Type:   User,
		Symbol: symbolKMalloc,
		FuncID: 3,
		Offset: 0,
	})

	return stackTraces
}
func TestConvertFakeStacks(t *testing.T) {
	simpleStackTraces := MakeSimpleStack()
	var stackTraces []*StackTrace
	for _, stackTrace := range simpleStackTraces {
		stackTraces = append(stackTraces, stackTrace)
	}
	repeatedStackTraces := MakeRepeatedStack()
	for _, stackTrace := range repeatedStackTraces {
		stackTraces = append(stackTraces, stackTrace)
	}

	p, err := CreateAllocationPProf(stackTraces)
	require.NoError(t, err)

	file, err := os.Create("test.proto")
	if err != nil {
		fmt.Println("couldn't create ")
	} else {
		err := p.Write(file)
		require.NoError(t, err)
	}
}
