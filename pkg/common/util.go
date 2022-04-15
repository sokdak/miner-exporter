package common

import (
	"fmt"
	"strings"
)

func GetConnectionString(connInfo ConnectionInfo) string {
	return fmt.Sprintf("%s://%s:%d", connInfo.Protocol, connInfo.Host, connInfo.Port)
}

func GeneralizeGpuName(name string) string {
	return strings.Trim(strings.Replace(name, "RTX", "", -1), " ")
}
