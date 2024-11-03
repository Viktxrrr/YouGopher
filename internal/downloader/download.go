package downloader

import (
	"fmt"

	"github.com/kkdai/youtube/v2"
)

type Download struct {
	Client         *youtube.Client
	ID             string
	URL            string
	DestinationDir string
	SelectedFormat *youtube.Format
}

func (d *Download) FindFormat(
	mimeType string,
	quality string,
	video bool,
	audio bool,
) (*youtube.Format, error) {
	formats, err := d.FetchFormats()
	if err != nil {
		return nil, err
	}
	filteredFormats := formats.Select(func(f youtube.Format) bool {
		return f.MimeType == mimeType && f.Quality == quality // TODO: Подумать над проверкой
	})

	if video {
		filteredFormats = formats.Select(func(f youtube.Format) bool {
			return f.Quality != ""
		})
	}

	if audio {
		filteredFormats = formats.Select(func(f youtube.Format) bool {
			return f.AudioQuality != ""
		})
	}

	if len(filteredFormats) == 0 {
		return nil, fmt.Errorf("format with specified parameters %s not found", mimeType)
	}

	return &filteredFormats[0], nil
}

func (d *Download) SetFormat(
	format *youtube.Format,
) {
	d.SelectedFormat = format
}

func (d *Download) GetVideo() (video *youtube.Video, err error) {
	video, err = d.Client.GetVideo(d.URL)
	if err != nil {
		return
	}
	return
}

func (d *Download) FetchFormats() (formats *youtube.FormatList, err error) {
	video, err := d.GetVideo()
	if err != nil {
		return
	}
	formats = &video.Formats
	return
}
