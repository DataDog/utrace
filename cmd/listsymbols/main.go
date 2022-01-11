package main

import (
	"debug/elf"
	"flag"
	"fmt"
	"os"

	"github.com/Gui774ume/utrace/pkg/utrace"
)

func main() {
	pidPtr := flag.Int("pid", -1, "pid of process")
	flag.Parse()
	
	pid := *pidPtr
	if pid == -1 {
		pid = os.Getpid()
	}
	symbols, err := utrace.ListSymbolsFromPID(pid)
	if err != nil {
		fmt.Errorf("Error: ", err)
		os.Exit(1)
	}

	for _, symInfo := range *symbols {
		if elf.ST_TYPE(symInfo.Symbol.Info) == elf.STT_FUNC {
			fmt.Printf("addr: %x, file: %s sym:", symInfo.Offset, symInfo.Path)
			fmt.Println(symInfo.Symbol, int(symInfo.Symbol.Section))
		}
	}
}
