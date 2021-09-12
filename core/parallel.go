package core

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

type Info struct {
	url          string
	NAME         string
	Location     string
	Length       int64
	Pieces       int
	Piece_length int64
}

type piece struct {
	index int
	start int64
	end   int64
	done  bool
}

func Make_info(url string, name string, location string, pieces int) Info {
	len := get_length(url)

	return Info{url, name, location, len, pieces, len / int64(pieces)}

}

func get_length(url string) int64 {

	res, err := http.Head(url)
	if err != nil {
		panic(err)
	}
	contentlength := res.ContentLength
	// fmt.Printf("ContentLength:%v\n", contentlength)
	return contentlength

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func download_Piece(wg *sync.WaitGroup, p *piece, inf Info, verbose bool) {
	defer wg.Done()
	os.MkdirAll(inf.Location+"/tmp", 0777)
	addr := inf.Location + "/tmp/dat" + fmt.Sprint(p.index)
	f, err := os.Create(addr)
	check(err)
	defer f.Close()

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 20
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
	if verbose {
		fmt.Printf("Downloaded %s size %d\n", addr, size)
	}
}

func make_pieces(inf Info) []piece {
	pieces := make([]piece, inf.Pieces)
	for i := 0; i < inf.Pieces; i++ {
		if i == 0 {
			pieces[i].start = 0
		} else {
			pieces[i].start = pieces[i-1].end + 1
		}

		pieces[i].index = i
		pieces[i].done = false
		if i == inf.Pieces-1 {
			pieces[i].end = inf.Length - 1
		} else {
			pieces[i].end = pieces[i].start + inf.Piece_length - 1
		}

	}
	return pieces
}

func merge(inf Info, pieces []piece) error {
	f, err := os.OpenFile(inf.Location+"/"+inf.NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	for i := 0; i < inf.Pieces; i++ {
		if !pieces[i].done {
			log.Fatal("Piece not downloaded")
			return fmt.Errorf("Piece %d not downloaded", i)
		}
		addr := inf.Location + "/tmp/dat" + fmt.Sprint(i)
		data, err := ioutil.ReadFile(addr)
		check(err)
		if _, err := f.Write(data); err != nil {
			log.Fatal(err)
			return err
		}
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func Download(inf Info, verbose bool) error {
	//make pieces
	pieces := make_pieces(inf)
	//download pieces
	var wg sync.WaitGroup
	for i := 0; i < inf.Pieces; i++ {
		wg.Add(1)
		go download_Piece(&wg, &pieces[i], inf, verbose)
	}
	wg.Wait()

	//Merge Pieces
	if verbose {
		fmt.Printf("Merging Pieces\n")
	}
	err := merge(inf, pieces)
	//cleanup
	os.RemoveAll(inf.Location + "/tmp")
	if err != nil {
		return err
	}

	return nil

}
