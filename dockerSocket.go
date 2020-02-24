package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tv42/httpunix"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var channel chan string

func checkSock(path string) (*http.Response, error) {

	logrus.Debug("[-] Checking Sock for HTTP: " + path)
	u := &httpunix.Transport{
		DialTimeout:           100 * time.Millisecond,
		RequestTimeout:        1 * time.Second,
		ResponseHeaderTimeout: 1 * time.Second,
	}
	u.RegisterLocation("dockerd", path)
	var client = http.Client{
		Transport: u,
	}
	resp, err := client.Get("http+unix://dockerd/info")

	if resp == nil {
		return nil, err
	}
	return resp, nil
}

func walkpath(path string, info os.FileInfo, err error) error {
	if err != nil {
		logrus.Debug(err)
	} else {
		switch mode := info.Mode(); {
		case mode&os.ModeSocket != 0:
			logrus.Debug("Valid Socket: " + path)
			resp, err := checkSock(path)
			if err == nil {
				if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
					channel <- path
				} else {
					logrus.Debug("Invalid Docker Socket: " + path)
				}
				defer resp.Body.Close()
			} else {
				logrus.Debug("Invalid Docker Socket: " + path)
			}
		default:
			if false {
				logrus.Debug("No Socket: " + path)
			}

		}
	}
	return nil
}

// GetValidSockets searches for valid docker-sockets.
// Search starts from startPath
// Values are returned through socktesChannel
func GetValidSockets(startPath string, socketsChannel chan string) {
	channel = socketsChannel
	err := filepath.Walk(startPath, walkpath)
	logrus.Debug("Filewalk complete")
	if err != nil {
		logrus.Debug(err)
	}
	close(socketsChannel)
}
