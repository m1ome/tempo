package tempo

import (
	"os"
	"strings"
	"testing"
)

func TestDir(t *testing.T) {
	tmp := Dir()

	fi, err := os.Stat(tmp)
	if err != nil {
		t.Fatal(err)
	}

	if !fi.IsDir() {
		t.Fatal("Not directory")
	}

	if fi.Mode()&os.ModePerm == 0 {
		t.Fatal("Wrong mode")
	}
}

func TestFile(t *testing.T) {
	tmp := File(".png")

	if !strings.HasSuffix(tmp, ".png") {
		t.Fatal("Not contains extension")
	}
}
