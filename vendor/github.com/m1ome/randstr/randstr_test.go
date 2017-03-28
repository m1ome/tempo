package randstr

import "testing"

func TestGetByte(t *testing.T) {
	a := GetByte(100)
	if len(a) != 100 {
		t.Fatal("Error length missmatch")
	}
}

func TestGetString(t *testing.T) {
	a := GetString(100)
	if len(a) != 100 {
		t.Fatal("Error length missmatch")
	}
}

func BenchmarkGetByte(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetByte(32)
	}
}

func BenchmarkGetString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetString(32)
	}
}
