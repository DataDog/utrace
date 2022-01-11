package utrace

import (
	"debug/elf"
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"

	manager "github.com/DataDog/ebpf-manager"
	"github.com/DataDog/gopsutil/process"
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

func ListSymbols(procMaps *[]process.MemoryMapsStat, symbolsCache map[SymbolAddr]elf.Symbol) error {
	for _, m := range *procMaps {
		_, syms, err := manager.OpenAndListSymbols(m.Path)
		if err != nil {
			fmt.Println("[error] ListSymbols: Skipping", m.Path)
			continue
		}
		fmt.Println("Start: ", m.StartAddr, ", end:", m.EndAddr, ", offset:", m.Offset)

		// from the entire list of symbols, only keep the functions that match the provided pattern
		for _, sym := range syms {
			if len(sym.Version) > 0 {
				// skip dynamic symbols
				continue
			}

			if sym.Value < m.Offset {
				// errors we still need to understand
				fmt.Println("[error] ListSymbols:  Skipping symbol", sym)
				continue
			}
			addr := SymbolAddr(sym.Value) - SymbolAddr(m.Offset) + SymbolAddr(m.StartAddr)
			symbolsCache[addr] = sym
		}
	}

	return nil
}

func TestListSymbols(t *testing.T) {
	pid := os.Getpid()
	procMaps, err := ListProcmaps(pid)
	require.NoError(t, err)
	symbolsCache := make(map[SymbolAddr]elf.Symbol)
	ListSymbols(procMaps, symbolsCache)
	fmt.Println("map:", symbolsCache)
}
