//go:build darwin

package notify

/*
#cgo CFLAGS: -x objective-c -fmodules -Wno-deprecated-declarations
#cgo LDFLAGS: -framework Foundation
#import <Foundation/Foundation.h>

// 用 NSUserNotification（macOS 11 起 deprecated，但截至 macOS 26 仍可用）。
// 相比 UNUserNotificationCenter 不需要代码签名 + entitlement，对未签名构建友好。
// 来源名（"Gridea Pro"）和图标自动从 Bundle 的 Info.plist 取，无需手动指定。
static void sendMacOSNotification(const char* title, const char* body) {
    @autoreleasepool {
        NSUserNotification *n = [[NSUserNotification alloc] init];
        n.title = [NSString stringWithUTF8String:title];
        n.informativeText = [NSString stringWithUTF8String:body];
        [[NSUserNotificationCenter defaultUserNotificationCenter] deliverNotification:n];
    }
}
*/
import "C"

import "unsafe"

func sendPlatform(title, body string) error {
	cTitle := C.CString(title)
	cBody := C.CString(body)
	defer C.free(unsafe.Pointer(cTitle))
	defer C.free(unsafe.Pointer(cBody))
	C.sendMacOSNotification(cTitle, cBody)
	return nil
}
