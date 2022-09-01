//go:build !linux && !freebsd && !netbsd && !openbsd && !windows && !darwin && !js
// +build !linux,!freebsd,!netbsd,!openbsd,!windows,!darwin,!js

package reminder

import "errors"

func Start() error {
	return errors.New("reminder is not supported on this OS yet")
}

func Stop() error {
	return errors.New("reminder is not supported on this OS yet")
}
