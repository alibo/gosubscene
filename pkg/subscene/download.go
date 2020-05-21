package subscene

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func Download(token string, f func(contentDisposition, contentType string, size int, reader io.Reader)) error {
	client := http.Client{
		Timeout: time.Second * 10,
	}

	url := "https://subscene.com/subtitles/farsi_persian-text/" + token

	res, err := client.Get(url)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()

		if res.StatusCode == http.StatusNotFound {
			return fmt.Errorf(http.StatusText(http.StatusNotFound))
		}
		return fmt.Errorf("http status is %d", res.Status)
	}

	contentDisposition := res.Header.Get("Content-Disposition")
	contentType := res.Header.Get("Content-Type")
	size, _ := strconv.Atoi(res.Header.Get("Content-Length"))
	f(contentDisposition, contentType, size, res.Body)

	return nil
}
