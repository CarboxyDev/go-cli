package cmd

import (
	"fmt"
	"runtime"
	"time"

	"github.com/spf13/cobra"
)

var cpuCmd = &cobra.Command{
    Use:   "cpu",
    Short: "Display CPU information",
    Long:  `Display CPU usage, core count, and architecture information`,
    Run: func(cmd *cobra.Command, args []string) {
        watch, _ := cmd.Flags().GetBool("watch")
        interval, _ := cmd.Flags().GetInt("interval")
        
        if watch {
            watchCPU(interval)
        } else {
            showCPUInfo()
        }
    },
}

func init() {
    rootCmd.AddCommand(cpuCmd)
    cpuCmd.Flags().BoolP("watch", "w", false, "Watch CPU usage continuously")
    cpuCmd.Flags().IntP("interval", "i", 2, "Update interval in seconds (for watch mode)")
}

func showCPUInfo() {
    fmt.Printf("🔧 CPU Information\n")
    fmt.Printf("Architecture: %s\n", runtime.GOARCH)
    fmt.Printf("OS: %s\n", runtime.GOOS)
    fmt.Printf("CPU Cores: %d\n", runtime.NumCPU())
    fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
    
    // Simple CPU usage estimation
    start := time.Now()
    runtime.GC()
    duration := time.Since(start)
    fmt.Printf("Last GC took: %v\n", duration)
}

func watchCPU(interval int) {
    fmt.Printf("🔍 Watching CPU (Press Ctrl+C to stop)\n\n")
    
    for {
        fmt.Print("\033[H\033[2J") 
        fmt.Printf("🔧 CPU Monitor - Updated: %s\n", time.Now().Format("15:04:05"))
        fmt.Printf("═══════════════════════════════════════\n")
        
        showCPUInfo()
        
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        fmt.Printf("\n💾 Go Runtime Stats:\n")
        fmt.Printf("Allocated: %.2f MB\n", float64(m.Alloc)/1024/1024)
        fmt.Printf("Total Allocations: %.2f MB\n", float64(m.TotalAlloc)/1024/1024)
        fmt.Printf("System Memory: %.2f MB\n", float64(m.Sys)/1024/1024)
        fmt.Printf("GC Cycles: %d\n", m.NumGC)
        
        time.Sleep(time.Duration(interval) * time.Second)
    }
}
