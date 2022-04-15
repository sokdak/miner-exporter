package common

import (
	"fmt"
)

func GetConnectionString(connInfo ConnectionInfo) string {
	return fmt.Sprintf("%s://%s:%d", connInfo.Protocol, connInfo.Host, connInfo.Port)
}
