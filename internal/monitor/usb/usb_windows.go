package usb

import (
	"fmt"
	"time"
	"encoding/json"
	"os/exec"
	"examguard/internal/event"
	"examguard/internal/logging"
)

type USBDevice struct {
	Name     string
	DeviceID string
}


func GetUSBDevices() ([]USBDevice, error) {

	psCommand := `
	$devices = @(
		Get-CimInstance Win32_PnPEntity |
		Where-Object { $_.PNPDeviceID -like 'USB\*' } |
		Select-Object @{
			Name = 'Name'
			Expression = { $_.Name }
		}, @{
			Name = 'DeviceID'
			Expression = { $_.PNPDeviceID }
		}
	)
	ConvertTo-Json -InputObject $devices -Compress
	`
	cmd := exec.Command("powershell.exe", "-NoProfile", "-NonInteractive", "-Command", psCommand)
	output, err := cmd.Output()
	if err != nil { return nil, err }

	var devices []USBDevice
	err = json.Unmarshal(output, &devices)
	if err != nil { return nil, err}

	return devices, nil
}

func MonitorUSB() {

	devices, err := GetUSBDevices()
	if err != nil {
		fmt.Println("USB scan failed:", err)
		return
	}

	knownDevices := make(map[string]USBDevice)
	for _, device := range devices {
		knownDevices[device.DeviceID] = device
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		devices, err := GetUSBDevices()
		if err != nil {
			fmt.Println("USB scan failed:", err)
			continue
		}

		currentDevices := make(map[string]USBDevice)

		for _, device := range devices {
			currentDevices[device.DeviceID] = device
		}

		for id, device := range currentDevices {
			if _, exists := knownDevices[id]; !exists {
				evt := event.Event{
					Time:    time.Now(),
					Source:  "USB",
					Type:    "USB_DEVICE_ADDED",
					Details: fmt.Sprintf("%s (ID: %s)", device.Name, device.DeviceID),
				}

				if err := logging.LogEvent(evt); err != nil {
					fmt.Println("Logging failed:", err)
					continue
				}

				fmt.Println("USB connected:", device.Name)
			}
		}

		for id, device := range knownDevices {
			if _, exists := currentDevices[id]; !exists {
				evt := event.Event{
					Time:    time.Now(),
					Source:  "USB",
					Type:    "USB_DEVICE_REMOVED",
					Details: fmt.Sprintf("%s (ID: %s)", device.Name, device.DeviceID),
				}

				if err := logging.LogEvent(evt); err != nil {
					fmt.Println("Logging failed:", err)
					continue
				}

				fmt.Println("USB disconnected:", device.Name)
			}
		}

		knownDevices = currentDevices
	}
}