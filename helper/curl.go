package helper

import (
	"io"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

//strings.NewReader("name=cjb")
func HttpPost(url string, data io.Reader) (body []byte, err error) {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		data)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
