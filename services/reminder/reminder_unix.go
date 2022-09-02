//go:build linux || freebsd || netbsd || openbsd
// +build linux freebsd netbsd openbsd

package reminder

import (
	"fmt"
	"os/exec"
)

func Start() error {
	// Find tsk executable
	binPath, err := exec.LookPath("tsk")
	if err != nil {
		return err
	}

	// Add crontab entry
	line := fmt.Sprintf(`(crontab -l; echo "* * * * * %s notify") | sort | uniq | crontab -`, binPath)
	err = exec.Command("bash", "-c", line).Run()
	if err != nil {
		return fmt.Errorf("failed to add crontab entry: %w", err)
	}

	return nil
}

func Stop() error {
	// Find tsk executable
	binPath, err := exec.LookPath("tsk")
	if err != nil {
		return err
	}

	// Remove crontab entry
	line := fmt.Sprintf(`(crontab -l; echo "* * * * * %s notify") | grep -v %s |  sort | uniq | crontab -`, binPath, binPath)
	err = exec.Command("bash", "-c", line).Run()
	if err != nil {
		return fmt.Errorf("failed to remove crontab entry: %w", err)
	}

	return nil
}
