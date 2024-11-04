package downloader

import (
	"github.com/google/uuid"
	"github.com/kkdai/youtube/v2"
)

type Download struct {
	Client           *youtube.Client
	ID               string
	Title            string
	FormatToDownload youtube.Format
	DestinationDir   string
}

func NewDownload(
	client *youtube.Client,
	url string,
	title string,
	format youtube.Format,
	destinationDir string,
) *Download {
	return &Download{
		Client:           client,
		ID:               uuid.New().String(),
		FormatToDownload: format,
		DestinationDir:   destinationDir,
		Title:            title,
	}
}
