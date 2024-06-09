package env

import (
	"os"
	"strconv"
)

func LookupStringEnv(envVar string, defaultValue string) string {
	envVarValue, ok := os.LookupEnv(envVar)
	if !ok {
		return defaultValue
	}
	return envVarValue
}

func LookupIntEnv(envVar string, defaultValue int) int {
	envVarValue, ok := os.LookupEnv(envVar)
	if !ok {
		return defaultValue
	}
	intVal, _ := strconv.Atoi(envVarValue)
	return intVal
}

func LookupInt64Env(envVar string, defaultValue int64) int64 {
	envVarValue, ok := os.LookupEnv(envVar)
	if !ok {
		return defaultValue
	}
	int64Val, _ := strconv.ParseInt(envVarValue, 10, 64)
	return int64Val
}

func LookupInt16Env(envVar string, defaultValue int16) int16 {
	envVarValue, ok := os.LookupEnv(envVar)
	if !ok {
		return defaultValue
	}
	int16Val, _ := strconv.ParseInt(envVarValue, 10, 16)
	return int16(int16Val)
}
