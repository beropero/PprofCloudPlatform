package client

import (
	"bytes"
	"fmt"

	"github.com/google/pprof/profile"
)

// PrintProfileText 解析并打印CPU Profile的文本信息
func PrintProfileText(profileData []byte) error {
	// 使用 github.com/google/pprof 包解析 Profile 数据
	prof, err := profile.Parse(bytes.NewReader(profileData))
	if err != nil {
		return fmt.Errorf("failed to parse profile: %v", err)
	}
	fmt.Printf("prof: %v\n", prof)
	// // 打印 Profile 的基本信息
	// fmt.Println("=== CPU Profile Information ===")
	// fmt.Printf("Sample Count: %d\n", len(prof.Sample))
	// fmt.Printf("Duration: %v\n", prof.DurationNanos)
	// fmt.Printf("Period: %v\n", prof.Period)
	// fmt.Println()

	// // 打印每个样本的调用栈
	// fmt.Println("=== Call Stacks ===")
	// for i, sample := range prof.Sample {
	// 	fmt.Printf("Sample %d:\n", i+1)
	// 	fmt.Printf("  CPU Time: %v\n", sample.Value) // 通常是 CPU 时间
	// 	fmt.Println("  Call Stack:")
	// 	for _, loc := range sample.Location {
	// 		// 打印每个调用栈位置的信息
	// 		for _, line := range loc.Line {
	// 			fmt.Printf("    %s (File: %s, Line: %d)\n", line.Function.Name, line.Function.Filename, line.Line)
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	return nil
}
