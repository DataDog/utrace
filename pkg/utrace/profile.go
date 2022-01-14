package utrace

import (
	"io"
	"time"

	"github.com/google/pprof/profile"
)

// Inspired from https://github.com/felixge/pprofutils/blob/main/internal/legacy/text.go#L34
// Thanks @felixge

func (u *UTrace) DumpProfile(w io.Writer) error {
	var stackTraces []*StackTrace
	for _, funcStackTraces := range u.matchingFuncStackTraces {
		for _, stackTrace := range funcStackTraces {
			stackTraces = append(stackTraces, stackTrace)
		}
	}
	profile, err := CreateAllocationPProf(stackTraces)
	if err != nil {
		return err
	}

	profile.Write(w)

	return nil
}

// assumption is that the map contains all the call stacks
func CreateAllocationPProf(stackTraces []*StackTrace) (*profile.Profile, error) {
	// TODO : fill this as file names
	//		Name:     "heap",
	//   	Filename: "heap.pprof",

	var (
		functionID = uint64(1)
		// locationID = uint64(1)
		p = &profile.Profile{
			TimeNanos: time.Now().UnixNano(),
			SampleType: []*profile.ValueType{{
				Type: "alloc_objects",
				Unit: "count",
			},
			{
				Type: "alloc_space",
				Unit: "bytes",
			}},
			// Without his, Delta.Convert() fails in profile.Merge(). Perhaps an
			// issue that's worth reporting upstream.
			PeriodType: &profile.ValueType{},
		}
	)

	pprofFunctions := make(map[string]uint64)

	m := &profile.Mapping{ID: 1, HasFunctions: true}
	p.Mapping = []*profile.Mapping{m}

	for _, stackTrace := range stackTraces {
		sample := &profile.Sample{}
		// sample values should match what was initialized as sample types up top
		sample.Value = append(sample.Value, int64(stackTrace.Count))
		sample.Value = append(sample.Value, int64(stackTrace.Value))

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

		// append this sample
		p.Sample = append(p.Sample, sample)
	}
	return p, nil
}
