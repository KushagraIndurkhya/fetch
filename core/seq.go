package core

import (
	"io"
	"net/http"
	"os"
)

func Download_Seq(inf Info) error {
	resp, err := http.Get(inf.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(inf.location)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}
