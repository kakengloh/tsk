//go:build windows && !darwin && !linux && !freebsd && !netbsd && !openbsd && !js
// +build windows,!darwin,!linux,!freebsd,!netbsd,!openbsd,!js

package reminder

import (
	"fmt"
	"os"
	"os/exec"
)

func Start() error {
	binPath, err := os.Executable()
	if err != nil {
		return err
	}

	err = exec.Command("schtasks", "/create", "/sc", "MINUTE", "/mo", "1", "/ts", "tsk reminder", "/tr", binPath, "notify").Run()
	if err != nil {
		return fmt.Errorf("failed to create scheduled task: %w", err)
	}

	err = exec.Command("schtasks", "/run", "/tn", "tsk reminder").Run()
	if err != nil {
		return fmt.Errorf("failed to run scheduled task: %w", err)
	}

	return nil
}

func Stop() error {
	err := exec.Command("schtasks", "/end", "/tn", "tsk reminder").Run()
	if err != nil {
		return fmt.Errorf("failed to end scheduled task: %w", err)
	}

	return nil
}
