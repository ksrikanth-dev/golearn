package main

import (
	"fmt"
	"strings"
)

// -------- Helper Functions --------

// Validate IP address (simple check for format, not full validation)
func validateIP(ip string) bool {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}
	return true
}

// Count active devices
func countActive(status map[string]bool) int {
	count := 0
	for _, v := range status {
		if v {
			count++
		}
	}
	return count
}

// Find device with longest name
func longestName(devices []string) string {
	longest := devices[0]
	for _, d := range devices {
		if len(d) > len(longest) {
			longest = d
		}
	}
	return longest
}

// Safe lookup for device status
func getStatus(name string, status map[string]bool) (bool, error) {
	val, ok := status[name]
	if !ok {
		return false, fmt.Errorf("device %s not found", name)
	}
	return val, nil
}

// Generate dummy device configuration
func generateConfig(name, ip string, active bool) []string {
	config := []string{
		fmt.Sprintf("hostname %s", name),
		fmt.Sprintf("interface eth0"),
		fmt.Sprintf("  ip address %s/24", ip),
		"  no shutdown",
		"enable secret admin123",
		"line vty 0 4",
		"  login local",
		"  transport input ssh",
	}
	if !active {
		// mark device as shutdown in dummy config
		config = append(config, "shutdown")
	}
	return config
}

// -------- Main Program --------

func main() {
	var n int
	fmt.Print("Enter number of devices: ")
	fmt.Scanln(&n)

	devices := make([]string, 0, n)
	ips := make(map[string]string)
	status := make(map[string]bool)
	configs := make(map[string][]string)

	for i := 0; i < n; i++ {
		var name, ip string
		var active string

		fmt.Printf("\nEnter name of device %d: ", i+1)
		fmt.Scanln(&name)

		fmt.Printf("Enter IP address of %s: ", name)
		fmt.Scanln(&ip)

		if !validateIP(ip) {
			fmt.Println("Invalid IP format, try again!")
			i-- // redo this iteration
			continue
		}

		fmt.Printf("Is %s active? (yes/no): ", name)
		fmt.Scanln(&active)

		devices = append(devices, name)
		ips[name] = ip
		status[name] = strings.ToLower(active) == "yes"

		// generate dummy config immediately
		configs[name] = generateConfig(name, ip, status[name])
	}

	// Show results
	fmt.Println("\n--- Device Configuration Summary ---")
	for _, d := range devices {
		activeStatus := "Inactive"
		if status[d] {
			activeStatus = "Active"
		}
		fmt.Printf("Device: %s | IP: %s | Status: %s\n", d, ips[d], activeStatus)
	}

	// Some analysis
	fmt.Println("\n--- Analysis ---")
	fmt.Printf("Total devices: %d | Active devices: %d\n", len(devices), countActive(status))
	fmt.Printf("Device with longest name: %s\n", longestName(devices))

	// Safe status lookup
	fmt.Println("\n--- Safe Lookup Demo ---")
	checkDevice := devices[0] // first device
	val, err := getStatus(checkDevice, status)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Status of %s: %v\n", checkDevice, val)
	}

	// Print dummy configurations
	fmt.Println("\n--- Generated Device Configurations ---")
	for _, d := range devices {
		fmt.Printf("\nConfiguration for %s:\n", d)
		for _, line := range configs[d] {
			fmt.Println(line)
		}
	}
}

/*
Sample Output:

Enter number of devices: 2

Enter name of device 1: Router1
Enter IP address of Router1: 192.168.1.1
Is Router1 active? (yes/no): yes

Enter name of device 2: Switch01
Enter IP address of Switch01: 192.168.1.10
Is Switch01 active? (yes/no): no

--- Device Configuration Summary ---
Device: Router1 | IP: 192.168.1.1 | Status: Active
Device: Switch01 | IP: 192.168.1.10 | Status: Inactive

--- Analysis ---
Total devices: 2 | Active devices: 1
Device with longest name: Switch01

--- Safe Lookup Demo ---
Status of Router1: true

--- Generated Device Configurations ---

Configuration for Router1:
hostname Router1
interface eth0
  ip address 192.168.1.1/24
  no shutdown
enable secret admin123
line vty 0 4
  login local
  transport input ssh

Configuration for Switch01:
hostname Switch01
interface eth0
  ip address 192.168.1.10/24
  no shutdown
enable secret admin123
line vty 0 4
  login local
  transport input ssh
shutdown
*/
