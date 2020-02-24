package crypto

import (
	// "encoding/hex"
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {

	_b, _ := GenerateHash("ministor")

	_is := CompareHash(_b, "ministor")

	if !_is {
		t.Errorf("CompareMd5 error")
	}

	// fmt.Println(hex.EncodeToString(_b))
	fmt.Println(string(_b))
}

func BenchmarkHash(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GenerateHash("13810167616")
	}

}
