//go:build windows

package notify

import (
	toast "git.sr.ht/~jackmordaunt/go-toast/v2"
)

// sendPlatform 走 WinRT toast，AUMID 决定角标和来源归属。
// 完整的 "Gridea Pro" 图标 + 名字显示要求 Start Menu 里有一个 .lnk 快捷方式
// 设置了同样的 AUMID 属性（NSIS 安装脚本应处理）。未注册时 toast 仍能弹出，
// 标题里写了 AppID 字符串 + 通用图标。
func sendPlatform(title, body string) error {
	n := toast.Notification{
		AppID:               appDisplayName,
		Title:               title,
		Body:                body,
		Icon:                appIconPath(),
	}
	// AUMID 在 toast XML 的 launch 属性 / activationType 之外，由 AppID 字段表示。
	// go-toast v2 直接用 AppID 调 RegisterActivator + ToastNotifier，
	// 没注册过的 AUMID 也能弹（Win10/11 都行），只是图标显示降级。
	return n.Push()
}
