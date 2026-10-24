package main

import (
	"fmt"
	"os"
	"regexp"
)

func getOs() string {
	var osVerPattern = regexp.MustCompile(`(?m)^NAME="([^"]+)"$`)

	osVer, err := os.ReadFile("/etc/os-release")
	if err != nil {
		fmt.Println("Failed to read os Version!")
	}

	matches := osVerPattern.FindStringSubmatch(string(osVer))
	if len(matches) > 1 {
		return matches[1]
	} else {
		return ""
	}
}

func getOsLogo(OperatingSystem string) string {
	OsImageFiles := map[string]string{
		"elementary OS": "./art/ElementaryOS.txt",
		"Ubuntu":        "./art/Ubuntu.txt",
	}

	if logoPath, ok := OsImageFiles[OperatingSystem]; ok {
		if asciiArt, err := os.ReadFile(logoPath); err != nil {
			fmt.Println("Failed to get Ascii Art, the art file is probably missing!")
			return ""
		} else {
			return string(asciiArt)
		}

	} else {
		return ""
	}
}

func getHardwareInfos(OperatingSystem string) map[string]string {
	var hardwareInfos = make(map[string]string)

	if OperatingSystem == "elementary OS" {

		//CPU Part
		var cpuInfoPattern = regexp.MustCompile(`(?m)^model name    :"?([^"\r\n]+)"?$`)
		cpuInfo, CpuErr := os.ReadFile("/proc/cpuinfo")
		if CpuErr != nil {
			fmt.Println("Failed to read CPU Infos!")
		}

		matchesCpu := cpuInfoPattern.FindStringSubmatch(string(cpuInfo))
		if len(matchesCpu) > 0 {
			hardwareInfos["cpu"] = matchesCpu[1]
		} else {
			hardwareInfos["cpu"] = "Not Found"
		}

		// Memory Part
		var memTotalInfoPattern = regexp.MustCompile(`(?m)^MemTotal:"?([^"\r\n]+)"?$`)
		var memAvailableInfoPattern = regexp.MustCompile(`(?m)^MemFree:"?([^"\r\n]+)"?$`)

		memInfo, memErr := os.ReadFile("/proc/meminfo")
		if memErr != nil {
			fmt.Println("Failed to read memory info!")
		}

		matchesMemTotal := memTotalInfoPattern.FindStringSubmatch(string(memInfo))
		if len(matchesMemTotal) > 0 {
			hardwareInfos["memTotal"] = matchesMemTotal[1]
		} else {
			hardwareInfos["memTotal"] = "Not Found"
		}

		matchesMemAvailable := memAvailableInfoPattern.FindStringSubmatch(string(memInfo))
		if len(matchesMemAvailable) > 0 {
			hardwareInfos["memAvailable"] = matchesMemAvailable[1]
		} else {
			hardwareInfos["memAvailable"] = "Not Found"
		}
	}

	return hardwareInfos
}

func main() {
	var osVersion = getOs()

	if osVersion != "" {
		fmt.Println(getOsLogo(osVersion) + "\n\n")
		fmt.Println("OS: ", osVersion)
	} else {
		fmt.Println("OS: Not found")
		os.Exit(1)
	}

	var hardwareInfos = getHardwareInfos(osVersion)
	hardwareInfoString := fmt.Sprintf("CPU : %s\nTotal Memory : %s\nFree Memory : %s", hardwareInfos["cpu"], hardwareInfos["memTotal"], hardwareInfos["memAvailable"])

	fmt.Println(hardwareInfoString)
}
