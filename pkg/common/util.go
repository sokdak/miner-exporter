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
