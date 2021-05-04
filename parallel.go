package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/hashicorp/go-retryablehttp"
)

type info struct {
	url          string
	location     string
	length       int64
	pieces       int
	piece_length int64
}

type piece struct {
	index int
	start int64
	end   int64
	done  bool
}

func make_info(url string, location string, pieces int) info {
	len := get_length(url)

	return info{url, location, len, pieces, len / int64(pieces)}

}

func get_length(url string) int64 {

	res, err := http.Head(url)
	if err != nil {
		panic(err)
	}
	contentlength := res.ContentLength
	fmt.Printf("ContentLength:%v\n", contentlength)
	return contentlength

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func download_Piece(wg *sync.WaitGroup, p *piece, inf info) {
	defer wg.Done()
	addr := "./tmp/dat" + fmt.Sprint(p.index)
	f, err := os.Create(addr)
	check(err)
	defer f.Close()

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10
	retryClient.Logger = nil

	client := retryClient.StandardClient()

	req, _ := http.NewRequest("GET", inf.url, nil)
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", p.start, p.end))
	res, _ := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	size, err := io.Copy(f, res.Body)
	defer f.Close()
	p.done = true
	fmt.Printf("Downloaded a file %s with size %d\n", addr, size)
}

func make_pieces(inf info) []piece {
	pieces := make([]piece, inf.pieces)
	for i := 0; i < inf.pieces; i++ {
		if i == 0 {
			pieces[i].start = 0
		} else {
			pieces[i].start = pieces[i-1].end + 1
		}

		pieces[i].index = i
		pieces[i].done = false
		if i == inf.pieces-1 {
			pieces[i].end = inf.length - 1
		} else {
			pieces[i].end = pieces[i].start + inf.piece_length - 1
		}

	}
	return pieces
}

func merge(inf info, pieces []piece) {
	f, err := os.OpenFile("try.mkv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < inf.pieces; i++ {
		if !pieces[i].done {
			log.Fatal("Piece not downloaded")
		}
		addr := "./tmp/dat" + fmt.Sprint(i)
		data, err := ioutil.ReadFile(addr)
		check(err)
		if _, err := f.Write(data); err != nil {
			log.Fatal(err)
		}
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}
func Download(inf info) {

	//make pieces
	pieces := make_pieces(inf)

	//download pieces
	var wg sync.WaitGroup
	for i := 0; i < inf.pieces; i++ {
		wg.Add(1)
		go download_Piece(&wg, &pieces[i], inf)
	}
	wg.Wait()

	//Merge Pieces
	merge(inf, pieces)

	fmt.Println("Download complete")
}
