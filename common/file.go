package common

import "os"

func IsDirExists(path string) bool {
	var _, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
