package utrace

import (
	"debug/elf"
	"fmt"
	"strings"
	manager "github.com/DataDog/ebpf-manager"
	"github.com/DataDog/gopsutil/process"
)

type SymbolInfo struct {
	Symbol elf.Symbol
	File   *elf.File
	Path   string
	Offset SymbolAddr
}

func ListProcMaps(pid int) (*[]process.MemoryMapsStat, error) {
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

func ListSymbolsFromProcMaps(procMaps *[]process.MemoryMapsStat, symbols *[]SymbolInfo) error {
	for _, m := range *procMaps {
		f, syms, err := manager.OpenAndListSymbols(m.Path)
		if err != nil {
			fmt.Println("[error] ListSymbols: Skipping", m.Path)
			continue
		}
		fmt.Printf("Processing %50s, Start: %x, End: 0x%x, Offset: 0x%x\n", m.Path, m.StartAddr, m.EndAddr, m.Offset)

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
			*symbols = append(*symbols, SymbolInfo{sym, f, m.Path, addr})
		}
	}

	return nil
}

func ListSymbolsFromPID(pid int) (*[]SymbolInfo, error) {
	procMaps,err := ListProcMaps(pid)
	if err != nil {
		return nil, err
	}
	var symbols []SymbolInfo
	err = ListSymbolsFromProcMaps(procMaps, &symbols)
	if err != nil {
		return nil, err
	}
	return &symbols, nil
}