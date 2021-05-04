package main

func main() {
	info := make_info(
		"https://nl34.seedr.cc/ff_get/972039482/Rick.and.Morty.S05E07.720p.WEBRip.x264-BAE.mkv?st=pwXieCnQodFquEjH5HFWsg&e=1628065458",
		"yts.jpg", 20)
	Download(info)
	Download_Seq(info)

}
