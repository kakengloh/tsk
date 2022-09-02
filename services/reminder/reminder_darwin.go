//go:build darwin && !windows && !linux && !freebsd && !netbsd && !openbsd
// +build darwin,!windows,!linux,!freebsd,!netbsd,!openbsd

package reminder

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/kakengloh/tsk/driver"
)

const launchAgentFileName = "com.tsk.tskd.plist"

func createLaunchAgentPayload(binPath, dataPath string) string {
	return fmt.Sprintf(`
		<?xml version="1.0" encoding="UTF-8"?>
		<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
		<plist version="1.0">
			<dict>
				<key>Label</key>
				<string>com.tsk.tskd</string>
				<key>ProcessType</key>
				<string>Interactive</string>
				<key>Disabled</key>
				<false />
				<key>RunAtLoad</key>
				<true />
				<key>KeepAlive</key>
				<dict>
					<key>SuccessfulExit</key>
					<false/>
				</dict>
				<key>LaunchOnlyOnce</key>
				<false />
				<key>StartCalendarInterval</key>
				<dict>
					<key>Second</key>
					<integer>0</integer>
				</dict>
				<key>Program</key>
				<string>%s</string>
				<key>ProgramArguments</key>
				<array>
					<string>%s</string>
					<string>notify</string>
				</array>
			</dict>
		</plist>
	`, binPath, binPath)
}

func getLaunchAgentsDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("$HOME directory not found: %w", err)
	}

	dir := path.Join(home, "Library", "LaunchAgents")

	_, err = os.Stat(dir)
	if err != nil {
		return "", fmt.Errorf("`LaunchAgents` directory not found: %w", err)
	}

	return dir, nil
}

func Start() error {
	// Find tsk executable path
	binPath, err := os.Executable()
	if err != nil {
		return err
	}

	// Find tsk data dir
	dataPath, err := driver.GetDataDir()
	if err != nil {
		return err
	}

	// Build launch agent XML payload
	payload := createLaunchAgentPayload(binPath, dataPath)

	// Find LaunchAgents dir
	launchAgentsDir, err := getLaunchAgentsDir()
	if err != nil {
		return err
	}

	// Create launch agent file
	f, err := os.OpenFile(path.Join(launchAgentsDir, launchAgentFileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to create launch agent: %w", err)
	}
	defer f.Close()
	_, err = f.Write([]byte(payload))
	if err != nil {
		return fmt.Errorf("failed to create launch agent: %w", err)
	}

	// Load launch agent
	err = exec.Command("launchctl", "load", "-w", path.Join(launchAgentsDir, launchAgentFileName)).Run()
	if err != nil {
		return fmt.Errorf("failed to start launch agent: %w", err)
	}

	return nil
}

func Stop() error {
	// Find LaunchAgents dir
	launchAgentsDir, err := getLaunchAgentsDir()
	if err != nil {
		return fmt.Errorf("launch agent has not created: %w", err)
	}

	// Unload launch agent
	err = exec.Command("launchctl", "unload", "-w", path.Join(launchAgentsDir, launchAgentFileName)).Run()
	if err != nil {
		return fmt.Errorf("failed to stop launch agent: %w", err)
	}

	return nil
}
