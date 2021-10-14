package ph_go

import "os"

func GetEnv(key, defVal string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		val = defVal
	}
	return val
}
