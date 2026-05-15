// Package notify 跨平台 OS 通知中心发送封装。
//
// 实现：macOS 走 osascript（未签名构建下来源会显示为 Script Editor），
// Windows 走 WinRT toast，Linux 走 notify-send（libnotify）。
// 通知失败不应阻塞主流程，调用方拿到 error 可以忽略。
package notify

import "github.com/gen2brain/beeep"

// Send 向系统通知中心发送一条通知。title 是粗体标题，body 是正文。
func Send(title, body string) error {
	return beeep.Notify(title, body, "")
}
