//go:build linux || windows

package notify

import (
	_ "embed"
	"os"
	"path/filepath"
	"sync"
)

// appicon.png 通过 //go:embed 嵌入二进制，首次调用时解出到用户缓存目录。
// Windows / Linux 的通知 API 都需要一个文件系统路径才能把图标挂上去；
// macOS 直接走 Bundle Info.plist 不需要这步。
//
//go:embed appicon.png
var embeddedIconPNG []byte

var (
	iconPathOnce sync.Once
	cachedIcon   string
)

// appIconPath 返回写入到用户缓存目录的图标文件路径，失败返回空串
// （此时通知仍能发，只是没自定义图标）。
func appIconPath() string {
	iconPathOnce.Do(func() {
		cacheDir, err := os.UserCacheDir()
		if err != nil {
			return
		}
		dir := filepath.Join(cacheDir, "gridea-pro")
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return
		}
		p := filepath.Join(dir, "appicon.png")
		if _, statErr := os.Stat(p); os.IsNotExist(statErr) {
			if err := os.WriteFile(p, embeddedIconPNG, 0o644); err != nil {
				return
			}
		}
		cachedIcon = p
	})
	return cachedIcon
}
