package tempo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/m1ome/randstr"
)

func dir() string {
	tmp := os.TempDir()
	tmp = filepath.Join(tmp, randstr.GetString(32))

	return tmp
}

// File return file name that lays in /tmp OS directory
//
// In e.g.:
// f := File(".png")
//
// Will return:
// /private/var/folders/3x/jf5977zase79drglr7rk0tq4d00000gn/T/4f502b9eds5ba0e89451617bf9f971bb.png
func File(extension string) string {
	return fmt.Sprintf("%s.%s", dir(), extension)
}

// Dir return & create random tmp directory (chmod 0777)
//
// In e.g.:
// f := Dir()
//
// Will return:
// /private/var/folders/3x/jf5977zase79drglr7rk0tq4d00000gn/T/4f502b9eds5ba0e89451617bf9f971bb
func Dir() string {
	dir := dir()
	os.MkdirAll(dir, os.ModePerm)

	return dir
}
