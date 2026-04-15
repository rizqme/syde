//go:build unix

package daemon

import "syscall"

// detachedSysProc returns the platform-specific SysProcAttr that makes
// the spawned syded survive the parent CLI exiting. On Unix we want a
// new session so it isn't killed with the parent terminal.
func detachedSysProc() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{Setsid: true}
}
