package settings

type DownloadSettings struct {
	DefaultDownloadsDirPath string
}

type Settings struct {
	Download *DownloadSettings
}
