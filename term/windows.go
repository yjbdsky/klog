//go:build windows
// +build windows

package term

import (
	"fmt"
	"syscall"
)

var (
	kernel32    *syscall.LazyDLL  = syscall.NewLazyDLL(`kernel32.dll`)
	proc        *syscall.LazyProc = kernel32.NewProc(`SetConsoleTextAttribute`)
	CloseHandle *syscall.LazyProc = kernel32.NewProc(`CloseHandle`)

	// 给字体颜色对象赋值
	FontColor Color = Color{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
)

type Color struct {
	black        int // 黑色
	blue         int // 蓝色
	green        int // 绿色
	cyan         int // 青色
	red          int // 红色
	purple       int // 紫色
	yellow       int // 黄色
	light_gray   int // 淡灰色（系统默认值）
	gray         int // 灰色
	light_blue   int // 亮蓝色
	light_green  int // 亮绿色
	light_cyan   int // 亮青色
	light_red    int // 亮红色
	light_purple int // 亮紫色
	light_yellow int // 亮黄色
	white        int // 白色
}

// 输出有颜色的字体
func ColorPrint(s string, i int) string {
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
	defer CloseHandle.Call(handle)
	return s
}

func Whitef(format string, a ...interface{}) string {
	return ColorPrint(fmt.Sprintf(format, a), FontColor.white)
}

func Magentaf(format string, a ...interface{}) string {
	return ColorPrint(fmt.Sprintf(format, a), FontColor.purple)
}

func Redf(format string, a ...interface{}) string {
	return ColorPrint(fmt.Sprintf(format, a), FontColor.red)
}

func Yellowf(format string, a ...interface{}) string {
	return ColorPrint(fmt.Sprintf(format, a), FontColor.yellow)
}

func Bluef(format string, a ...interface{}) string {
	return ColorPrint(fmt.Sprintf(format, a), FontColor.blue)
}
