package utrace

import (
	"fmt"
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
			SampleType: []*profile.ValueType{
				{
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

	pprofLocations := make(map[FuncID]*profile.Location)

	m := &profile.Mapping{ID: 1, HasFunctions: true}
	p.Mapping = []*profile.Mapping{m}

	for _, stackTrace := range stackTraces {
		sample := &profile.Sample{}
		// sample values should match what was initialized as sample types up top
		sample.Value = append(sample.Value, int64(stackTrace.Count))
		sample.Value = append(sample.Value, int64(stackTrace.Value))

		for _, stackTraceNode := range stackTrace.UserStacktrace {
			currentLocation, ok := pprofLocations[stackTraceNode.FuncID]
			if !ok {
				fmt.Println("Creating loc for ", stackTraceNode.Symbol.Name)

				currentFunc := &profile.Function{
					ID:   functionID,
					Name: stackTraceNode.Symbol.Name,
				}
				p.Function = append(p.Function, currentFunc)

				currentLocation = &profile.Location{
					ID:      functionID,
					Mapping: m,
					Address: stackTraceNode.Symbol.Value,
					Line:    []profile.Line{{Function: p.Function[len(p.Function)-1]}},
				}
				functionID++

				p.Location = append(p.Location, currentLocation)
				// cache this location
				pprofLocations[stackTraceNode.FuncID] = currentLocation
			}
			fmt.Println("Adding ", currentLocation)

			sample.Location = append(sample.Location, currentLocation)
		}

		// append this sample
		p.Sample = append(p.Sample, sample)
		p.CheckValid()
	}

	return p, nil
}
