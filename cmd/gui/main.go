package main

import (
	"net/http"

	"github.com/Viktxrrr/YouGopher/internal/downloader"
	"github.com/Viktxrrr/YouGopher/internal/settings"
	"github.com/Viktxrrr/YouGopher/internal/uiapp"
	"github.com/kkdai/youtube/v2"
)

func main() {
	settings := &settings.Settings{
		Download: &settings.DownloadSettings{
			DefaultDownloadsDirPath: ".",
		},
	}
	ytClient := &youtube.Client{}
	httpClient := &http.Client{}
	dm := downloader.NewDownloadsManager(ytClient, httpClient, settings)
	app := uiapp.NewAppUI(dm, settings)
	app.Run()
}
