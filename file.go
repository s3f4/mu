package mu

import "os"

// DirExists returns boolean value
// that shows whether there is a directory with given string
func DirExists(dir string) bool {
	stat, err := os.Stat(dir)
	return !os.IsNotExist(err) && stat.IsDir()
}
