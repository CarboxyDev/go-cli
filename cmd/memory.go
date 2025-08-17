package cmd

import (
	"fmt"
	"runtime"
	"time"

	"github.com/spf13/cobra"
)

var memoryCmd = &cobra.Command{
    Use:   "memory",
    Short: "Display memory usage information",
    Long:  `Display memory usage, allocation statistics, and garbage collection info`,
    Run: func(cmd *cobra.Command, args []string) {
        watch, _ := cmd.Flags().GetBool("watch")
        interval, _ := cmd.Flags().GetInt("interval")
        detailed, _ := cmd.Flags().GetBool("detailed")
        
        if watch {
            watchMemory(interval, detailed)
        } else {
            showMemoryInfo(detailed)
        }
    },
}

func init() {
    rootCmd.AddCommand(memoryCmd)
    memoryCmd.Flags().BoolP("watch", "w", false, "Watch memory usage continuously")
    memoryCmd.Flags().IntP("interval", "i", 2, "Update interval in seconds")
    memoryCmd.Flags().BoolP("detailed", "d", false, "Show detailed memory statistics")
}

func showMemoryInfo(detailed bool) {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    fmt.Printf("💾 Memory Information\n")
    fmt.Printf("═══════════════════════\n")
    fmt.Printf("Allocated Memory: %.2f MB\n", bToMB(m.Alloc))
    fmt.Printf("Total Allocations: %.2f MB\n", bToMB(m.TotalAlloc))
    fmt.Printf("System Memory: %.2f MB\n", bToMB(m.Sys))
    fmt.Printf("Heap Memory: %.2f MB\n", bToMB(m.HeapAlloc))
    fmt.Printf("Stack Memory: %.2f MB\n", bToMB(m.StackSys))
    
    if detailed {
        fmt.Printf("\n🔍 Detailed Statistics:\n")
        fmt.Printf("Heap Objects: %d\n", m.HeapObjects)
        fmt.Printf("GC Cycles: %d\n", m.NumGC)
        fmt.Printf("Last GC: %v ago\n", time.Since(time.Unix(0, int64(m.LastGC))))
        fmt.Printf("Next GC Target: %.2f MB\n", bToMB(m.NextGC))
        fmt.Printf("Pause Time: %v\n", time.Duration(m.PauseTotalNs))
    }
}

func watchMemory(interval int, detailed bool) {
    fmt.Printf("🔍 Watching Memory (Press Ctrl+C to stop)\n\n")
    
    for {
        fmt.Print("\033[H\033[2J") 
        fmt.Printf("💾 Memory Monitor - Updated: %s\n", time.Now().Format("15:04:05"))
        fmt.Printf("═══════════════════════════════════════\n")
        
        showMemoryInfo(detailed)
        
        time.Sleep(time.Duration(interval) * time.Second)
    }
}

func bToMB(b uint64) float64 {
    return float64(b) / 1024 / 1024
}
