package downloader

import (
	"fmt"
	"strings"

	"github.com/Viktxrrr/YouGopher/internal/utils"
	"github.com/kkdai/youtube/v2"
)

type VideoData struct {
	URL             string
	Video           *youtube.Video
	SelectedQuality string
	SelectedCodec   string
}

func NewVideoData(url string) *VideoData {
	return &VideoData{
		URL: url,
	}
}

func (v *VideoData) Initialize(client *youtube.Client, doneChan chan bool, errChan chan error) {
	go v.fetchVideo(client, doneChan, errChan)
}

func (v *VideoData) fetchVideo(client *youtube.Client, done chan bool, errChan chan error) {
	video, err := client.GetVideo(v.URL)
	if err != nil {
		errChan <- err
		return
	}
	v.Video = video
	done <- true
}

func (v *VideoData) GetSelectedFormat() (f youtube.Format, err error) {
	if v.SelectedQuality == "" || v.SelectedCodec == "" {
		return youtube.Format{}, fmt.Errorf("quality or codec is empty")
	}
	for _, f = range v.Video.Formats {
		if f.QualityLabel == v.SelectedQuality && strings.Contains(f.MimeType, v.SelectedCodec) {
			return f, nil
		}
	}
	return youtube.Format{}, nil
}

func (v *VideoData) GetQualities() (qualities []string) {
	qualities = []string{}
	for _, f := range v.Video.Formats {
		if f.QualityLabel == "" || utils.Contains(qualities, f.QualityLabel) {
			continue
		}
		qualities = append(qualities, f.QualityLabel)
	}
	return
}

func (v *VideoData) GetVideoCodecsForQuality(quality string) (codecs []string) {
	codecs = []string{}
	for _, f := range v.Video.Formats {
		parts := strings.SplitN(f.MimeType, ";", 2) // for example video/mp4; codecs="avc..."
		codecsPart := strings.TrimSpace(parts[1])
		codec := strings.Trim(strings.SplitN(codecsPart, "=", 2)[1], "\"")
		if utils.Contains(codecs, codec) || f.QualityLabel != quality {
			continue
		}
		codecs = append(codecs, codec)
	}
	return
}
