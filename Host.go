package main

import (
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

var titleRe = regexp.MustCompile(`(?i)<title>(.*?)</title>`)

var tr = http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}

var httpClient = http.Client{
	Timeout:   time.Duration(5 * time.Second),
	Transport: &tr,
}

type Host struct {
	Url string
}

func getTitle(data []byte) string {
	m := titleRe.FindSubmatch(data)
	if m == nil {
		return ""
	}
	if len(m) != 2 {
		return ""
	}
	return string(m[1][:])
}

func getHash(data []byte) string {
	hasher := sha256.New()
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}

func (h Host) Get() (*Footprint, error) {
	req, err := http.NewRequest("GET", h.Url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "slowpoke")
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	hash := getHash(body)
	title := getTitle(body)
	return &Footprint{Url: h.Url, Hash: hash, Title: title}, nil
}
