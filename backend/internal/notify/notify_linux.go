//go:build linux

package notify

import (
	"github.com/esiqveland/notify"
	"github.com/godbus/dbus/v5"
)

// sendPlatform 走 D-Bus org.freedesktop.Notifications，显式传 app_name + app_icon。
// 覆盖所有遵循 freedesktop.org 规范的桌面环境（GNOME / KDE / XFCE / Cinnamon /
// MATE / LXQt / dunst / mako / pantheon 等）。无通知守护进程的环境（纯 CLI /
// 服务器）会拿到 D-Bus 错误。
func sendPlatform(title, body string) error {
	conn, err := dbus.SessionBus()
	if err != nil {
		return err
	}
	notifier, err := notify.New(conn)
	if err != nil {
		return err
	}
	defer notifier.Close()
	_, err = notifier.SendNotification(notify.Notification{
		AppName:       appDisplayName,
		AppIcon:       appIconPath(),
		Summary:       title,
		Body:          body,
		ExpireTimeout: notify.ExpireTimeoutSetByNotificationServer,
	})
	return err
}
