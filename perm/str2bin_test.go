package perm

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkStr2Bin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Str2Bin("101111")
	}
}

func BenchmarkStrconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := strconv.ParseInt("101111", 2, 64)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
