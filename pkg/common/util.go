package common

import (
	"fmt"
	"strings"
)

func GetConnectionString(connInfo ConnectionInfo) string {
	return fmt.Sprintf("%s://%s:%d", connInfo.Protocol, connInfo.Host, connInfo.Port)
}

func GeneralizeGpuName(name string) string {
	// replace RTX to none
	name = strings.Replace(name, "RTX", "", -1)
	// replace Radeon to none
	name = strings.Replace(name, "Radeon", "", -1)
	// replace RX to none
	name = strings.Replace(name, "RX", "", -1)

	return strings.Trim(name, " ")
}

func GeneralizePoolAddress(source string) string {
	strs := strings.Split(source, "://")
	if len(strs) < 2 {
		return source
	}
	return strs[1]
}

func GeneralizeAlgorithm(source string) string {
	return strings.ToLower(source)
}

func ExtractAddress(source string) string {
	return strings.Split(source, ".")[0]
}

func ExtractWorkerNameFromAddress(source string) string {
	strs := strings.Split(source, ".")
	if len(strs) < 2 {
		return source
	}
	return strings.ToUpper(strings.Join(strs[1:], "."))
}
