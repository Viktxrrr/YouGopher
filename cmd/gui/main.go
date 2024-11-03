package main

import (
	"net/http"

	"github.com/Viktxrrr/YouGopher/internal/downloader"
	"github.com/Viktxrrr/YouGopher/internal/settings"
	"github.com/Viktxrrr/YouGopher/internal/uiapp"
)

func main() {
	settings := &settings.Settings{
		Download: &settings.DownloadSettings{
			DefaultDownloadsDirPath: ".",
		},
	}
	httpClient := &http.Client{}
	dm := downloader.NewDownloadsManager(httpClient, settings)
	app := uiapp.NewAppUI(dm, settings)
	app.Run()
}
