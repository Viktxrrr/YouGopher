package uiapp

import (
	"fyne.io/fyne/v2/widget"
	"github.com/Viktxrrr/YouGopher/internal/downloader"
	"github.com/Viktxrrr/YouGopher/internal/utils"
)

func StartDownloadOnButtonClick(vd *downloader.VideoData, dm *downloader.DownloadsManager) {
	selectedFormat, err := vd.GetSelectedFormat()
	if err != nil {
		panic(err)
	}
	d := downloader.NewDownload(
		dm.YTClient,
		vd.URL,
		vd.Video.Title,
		selectedFormat,
		".",
	)
	dm.AddDownload(*d)
	dm.StartDownload(d.ID)
}

func HandleQualityChange(
	vd *downloader.VideoData,
	qualitySelect *widget.Select,
	codecSelect *widget.Select,
) {
	vd.SelectedQuality = qualitySelect.Selected
	codecs := vd.GetVideoCodecsForQuality(qualitySelect.Selected)
	UpdateCodecSelectOnQualityChange(codecSelect, codecs)
}

func UpdateCodecSelectOnQualityChange(
	codecSelect *widget.Select,
	codecs []string,
) {
	codecSelect.SetOptions(codecs)
	if len(codecs) > 0 {
		codecSelect.SetSelected(codecs[0])
	}
}

func UpdateQualitySelect(qualitySelect *widget.Select, qualities []string) {
	qualitySelect.SetOptions(qualities)
	if len(qualities) > 0 {
		qualitySelect.SetSelected(qualities[0])
	}
}

func HandleURLChange(
	url string,
	dm *downloader.DownloadsManager,
	vd **downloader.VideoData,
	qualitySelect *widget.Select,
) {
	if !utils.IsValidYouTubeURL(url) {
		return
	}
	videoData := downloader.NewVideoData(url)
	errChan := make(chan error)
	doneChan := make(chan bool)
	go videoData.Initialize(dm.YTClient, doneChan, errChan)
	select {
	case err := <-errChan:
		panic(err)
	case <-doneChan:
		*vd = videoData
		UpdateQualitySelect(qualitySelect, videoData.GetQualities())
	}
}

func HandleCodecChange(vd *downloader.VideoData, codecSelect *widget.Select) {
	vd.SelectedCodec = codecSelect.Selected
}
