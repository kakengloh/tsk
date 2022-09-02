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

	// Check if scheduled task exists
	err = exec.Command("schtasks", "/query", "/tn", "TskReminder").Run()
	// Delete scheduled task if exists
	if err == nil {
		err = exec.Command("schtasks", "/delete", "/tn", "TskReminder", "/f").Run()
		if err != nil {
			return fmt.Errorf("failed to delete scheduled task: %w", err)
		}
	}

	// Create scheduled task
	err = exec.Command("schtasks", "/create", "/sc", "MINUTE", "/mo", "1", "/tn", "TskReminder", "/tr", binPath+" notify").Run()
	if err != nil {
		return fmt.Errorf("failed to create scheduled task: %w", err)
	}

	// Run scheduled task
	err = exec.Command("schtasks", "/run", "/tn", "TskReminder").Run()
	if err != nil {
		return fmt.Errorf("failed to run scheduled task: %w", err)
	}

	return nil
}

func Stop() error {
	// End scheduled task
	err := exec.Command("schtasks", "/end", "/tn", "TskReminder").Run()
	if err != nil {
		return fmt.Errorf("failed to end scheduled task: %w", err)
	}

	return nil
}
