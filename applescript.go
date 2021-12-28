package menuet

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa

#import <Cocoa/Cocoa.h>

#ifndef __APPLESCRIPT_H_H__
#import "applescript.h"
#endif

bool execute(const char* s);

*/
import "C"
import (
	"strings"
	"unsafe"
)

func (a *Application) ExecuteAppleScript(script string) bool {
	if strings.TrimSpace(script) == "" {
		return false
	}

	cstr := C.CString(script)
	result := C.execute(cstr)
	C.free(unsafe.Pointer(cstr))
	return bool(result)
}
