package core

import (
	"os"
	"testing"
)

var inf Info = Make_info(
	"https://saimei.ftp.acc.umu.se/debian-cd/current/amd64/iso-cd/debian-11.0.0-amd64-netinst.iso",
	os.Getenv("HOME")+"/Downloads",
	"debian.iso", 4)

func BenchmarkDownload_Seq(b *testing.B) {
	Download_Seq(inf, true)
}

func BenchmarkDownload(b *testing.B) {
	Download(inf, true)
}
