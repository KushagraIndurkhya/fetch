package core

import (
	"testing"
)

var inf Info = Make_info(
	"https://saimei.ftp.acc.umu.se/debian-cd/current/amd64/iso-cd/debian-11.0.0-amd64-netinst.iso",
	"debian.iso", 4)

func BenchmarkDownload_Seq(b *testing.B) {
	Download_Seq(inf)
}

func BenchmarkDownload(b *testing.B) {
	Download(inf)
}
