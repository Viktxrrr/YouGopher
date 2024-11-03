package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Viktxrrr/YouGopher/internal/settings"
	"github.com/Viktxrrr/YouGopher/internal/utils"
	"github.com/kkdai/youtube/v2"
)

type DownloadsManager struct {
	HTTPClient *http.Client
	YTClient   *youtube.Client
	Settings   *settings.Settings
	Downloads  []Download
}

func NewDownloadsManager(
	httpClient *http.Client,
	settings *settings.Settings,
) *DownloadsManager {
	return &DownloadsManager{
		HTTPClient: httpClient,
		Settings:   settings,
		Downloads:  []Download{},
	}
}

func (dm *DownloadsManager) AddDownload(d Download) {
	dm.Downloads = append(dm.Downloads, d)
}

func (dm *DownloadsManager) FindDownloadById(id string) (d Download, err error) {
	for _, d := range dm.Downloads {
		if d.ID == id {
			return d, err
		}
	}
	err = fmt.Errorf("Download with id %s not found", id)
	return d, err
}

func (dm *DownloadsManager) StartDownload(
	id string,
) (err error) {
	d, err := dm.FindDownloadById(id)
	if err != nil {
		return
	}

	if d.SelectedFormat == nil {
		return fmt.Errorf("format not specified for download with id %s", id)
	}

	downloadURL, err := dm.HTTPClient.Get(d.SelectedFormat.URL)
	if err != nil {
		return fmt.Errorf("can't get download url for video: %v", err)
	}
	defer downloadURL.Body.Close()

	video, err := d.GetVideo()
	if err != nil {
		return
	}

	outputDir := dm.Settings.Download.DefaultDownloadsDirPath
	if d.DestinationDir != "" {
		outputDir = d.DestinationDir
	}

	destinationFile, err := utils.GetOutputFile(
		d.SelectedFormat,
		outputDir,
		video.Title,
	) // TODO: Добавить проверку существования файла
	outFile, err := os.Create(destinationFile)
	if err != nil {
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, downloadURL.Body)
	if err != nil {
		return
	}

	fmt.Printf("Downloaded video %s to %s", d.URL, destinationFile)
	return
}
