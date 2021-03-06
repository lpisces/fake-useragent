package downloader

import (
	"io"
	"net"
	"net/http"
	"time"
)

type Download struct {
	Delay   time.Duration
	Timeout time.Duration
}

func (d *Download) Get(url string) (io.ReadCloser, error) {
	time.Sleep(d.Delay)
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: d.Timeout,
			}).DialContext,
		},
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
