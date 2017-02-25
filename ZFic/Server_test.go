package ZFic

import (
	"testing"
	"fmt"
)

var pagetoload = "../Zfic/Files/AoZ/" + "Mainpage.icfg"

func BenchmarkLoadPage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HandlePage(pagetoload)
	}
}
