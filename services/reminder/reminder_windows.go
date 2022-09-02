//go:build windows && !darwin && !linux && !freebsd && !netbsd && !openbsd && !js
// +build windows,!darwin,!linux,!freebsd,!netbsd,!openbsd,!js

package reminder

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/kakengloh/tsk/driver"
)

func createVbsPayload(binPath string) string {
	return fmt.Sprintf(`
		Set WshShell = CreateObject("WScript.Shell")
		WshShell.Run "%s notify", 0
		Set WshShell = Nothing
	`, binPath)
}

func Start() error {
	// Find tsk path
	binPath, err := os.Executable()
	if err != nil {
		return err
	}

	// Find tsk data dir
	dataPath, err := driver.GetDataDir()
	if err != nil {
		return err
	}

	// Find wscript path
	wscriptPath, err := exec.LookPath("wscript.exe")
	if err != nil {
		return fmt.Errorf("unable to find wscript: %w", err)
	}

	// Build VBScript path
	vbScriptPath := path.Join(dataPath, "TskReminder.vbs")

	// Check if scheduled task exists
	err = exec.Command("schtasks", "/query", "/tn", "TskReminder").Run()
	// Delete scheduled task if exists
	if err == nil {
		err = Stop()
		if err != nil {
			return err
		}
	}

	// Create VBScript file
	f, err := os.OpenFile(vbScriptPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open VBScript file: %w", err)
	}
	defer f.Close()
	_, err = f.Write([]byte(createVbsPayload(binPath)))
	if err != nil {
		return fmt.Errorf("failed to create VBScript: %w", err)
	}

	// Create scheduled task
	err = exec.Command("schtasks", "/create", "/sc", "MINUTE", "/mo", "1", "/tn", "TskReminder", "/tr", wscriptPath+" "+vbScriptPath).Run()
	if err != nil {
		return fmt.Errorf("failed to create scheduled task: %w", err)
	}

	return nil
}

func Stop() error {
	// Delete scheduled task
	err := exec.Command("schtasks", "/delete", "/tn", "TskReminder", "/f").Run()
	if err != nil {
		return fmt.Errorf("failed to delete scheduled task: %w", err)
	}

	return nil
}
