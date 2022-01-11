package utrace

import (
	"fmt"
	"os"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"

	"github.com/DataDog/gopsutil/process"
	//manager "github.com/DataDog/ebpf-manager"
)

func ListProcmaps(pid int) (*[]process.MemoryMapsStat, error) {
	p, err := process.NewProcess(int32(pid))
	if err != nil {
		return nil, err
	}
	procMaps, err := p.MemoryMaps(true)
	if err != nil {
		return nil, err
	}
	filteredProcMaps := []process.MemoryMapsStat{}
	for _, m := range *procMaps {
		if m.Permission.Execute && !strings.HasPrefix(m.Path, "[") {
			fmt.Print(m, "\n")
			filteredProcMaps = append(filteredProcMaps, m)
		}		
	}

	return &filteredProcMaps, nil
}

func ListSymbols(procMaps *[]process.MemoryMapsStat) (error) {
	return nil
}

func TestListSymbols(t *testing.T) {
	pid := os.Getpid()
	_, err := ListProcmaps(pid)
	require.NoError(t, err)
}
