package core

import (
	"io"
	"net/http"
	"os"
)

func Download_Seq(inf Info, verbose bool) error {
	resp, err := http.Get(inf.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(inf.Location + "/" + inf.NAME)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}
