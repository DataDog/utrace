package utrace

import (
	"io"
	"time"

	"github.com/google/pprof/profile"
)

// Inspired from https://github.com/felixge/pprofutils/blob/main/internal/legacy/text.go#L34
// Thanks @felixge

func (u *UTrace) DumpProfile(w io.Writer) error {
	return nil
}

// assumption is that the map contains all the call stacks
func CreateAllocationPProf(allocationMap map[CombinedID]*StackTrace) (*profile.Profile, error) {
	// TODO : fill this as file names
	//		Name:     "heap",
	//   	Filename: "heap.pprof",

	var (
		functionID = uint64(1)
		// locationID = uint64(1)
		p = &profile.Profile{
			TimeNanos: time.Now().UnixNano(),
			SampleType: []*profile.ValueType{{
				Type: "samples",
				Unit: "count",
			}},
			// Without his, Delta.Convert() fails in profile.Merge(). Perhaps an
			// issue that's worth reporting upstream.
			PeriodType: &profile.ValueType{},
		}
	)

	pprofFunctions := make(map[string]uint64)

	m := &profile.Mapping{ID: 1, HasFunctions: true}
	p.Mapping = []*profile.Mapping{m}

	// If we want to add a new sample type
	// add the sample types
	// p.SampleType = append(p.SampleType, &profile.ValueType{Type: "alloc_space", Unit: "bytes"})

	for _, stackTrace := range allocationMap {
		sample := &profile.Sample{}
		// sample values should match what was initialized as sample types up top
		sample.Value = append(sample.Value, int64(stackTrace.Count))

		for _, stackTraceNode := range stackTrace.UserStacktrace {
			pprofFuncId, ok := pprofFunctions[stackTraceNode.Symbol.Name]
			if !ok {
				pprofFuncId = functionID
				functionID++
				pprofFunctions[stackTraceNode.Symbol.Name] = pprofFuncId
			}
			function := &profile.Function{
				ID:   pprofFuncId,
				Name: stackTraceNode.Symbol.Name,
			}
			p.Function = append(p.Function, function)

			location := &profile.Location{
				ID:      pprofFuncId,
				Mapping: m,
				Line:    []profile.Line{{Function: function}},
			}

			p.Location = append(p.Location, location)

			sample.Location = append(sample.Location, location)
		}
		for _, stackTraceNode := range stackTrace.KernelStackTrace {
			pprofFuncId, ok := pprofFunctions[stackTraceNode.Symbol.Name]
			if !ok {
				pprofFuncId = functionID
				functionID++
				pprofFunctions[stackTraceNode.Symbol.Name] = pprofFuncId
			}

			function := &profile.Function{
				ID:   pprofFuncId,
				Name: stackTraceNode.Symbol.Name,
			}
			p.Function = append(p.Function, function)
			location := &profile.Location{
				ID:      pprofFuncId,
				Mapping: m,
				Line:    []profile.Line{{Function: function}},
			}

			p.Location = append(p.Location, location)

			sample.Location = append(sample.Location, location)
		}

		// append this sample
		p.Sample = append(p.Sample, sample)
	}
	return p, nil
}
