package common

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"
)

func GetConnectionString(connInfo ConnectionInfo) string {
	return fmt.Sprintf("%s://%s:%d", connInfo.Protocol, connInfo.Host, connInfo.Port)
}

func GeneralizeGpuName(name string) string {
	// replace GeForce to none
	name = strings.Replace(name, "GeForce", "", -1)
	// replace RTX to none
	name = strings.Replace(name, "RTX", "", -1)
	// replace Radeon to none
	name = strings.Replace(name, "Radeon", "", -1)
	// replace RX to none
	name = strings.Replace(name, "RX", "", -1)
	// replace whitespace to none
	name = strings.Replace(name, " ", "", -1)

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

func GetNonValueInsteadIfNotPresent(value int) int {
	if value == -1 || value == 0 {
		return ValueNotSet
	}
	return value
}

func DeepCopy(src, dist interface{}) (err error) {
	buf := bytes.Buffer{}
	if err = gob.NewEncoder(&buf).Encode(src); err != nil {
		return
	}
	return gob.NewDecoder(&buf).Decode(dist)
}
