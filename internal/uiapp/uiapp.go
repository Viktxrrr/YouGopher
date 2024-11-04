package uiapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Viktxrrr/YouGopher/internal/downloader"
	"github.com/Viktxrrr/YouGopher/internal/settings"
)

type AddDownloadWindow struct {
	Window         fyne.Window
	URLEntry       *widget.Entry
	FormatDropdown *widget.Select
	VoiceCheckbox  *widget.Check
	VideoCheckbox  *widget.Check
	ConfirmButton  *widget.Button

	DownloadsManager *downloader.DownloadsManager
}

type MainWindow struct {
	Window            fyne.Window
	DownloadsList     *widget.List
	AddDownloadButton *widget.Button

	DownloadsManager *downloader.DownloadsManager
}

type SettingsWindow struct {
	Settings *settings.Settings
}

type AppUI struct {
	App               fyne.App
	MainWindow        *MainWindow
	AddDownloadWindow *AddDownloadWindow
	SettingsWindow    *SettingsWindow

	DownloadsManager *downloader.DownloadsManager
	Settings         *settings.Settings
}

func CreateURLEntry() (urlEntry *widget.Entry) {
	urlEntry = widget.NewEntry()
	urlEntry.SetPlaceHolder("Video URL")
	return
}

func NewAddDownloadWindow(app fyne.App, dm *downloader.DownloadsManager) *AddDownloadWindow {
	w := app.NewWindow("Add download")
	urlEntry := CreateURLEntry()
	qualitySelect := widget.NewSelect(nil, nil)
	codecSelect := widget.NewSelect(nil, nil)

	var vd *downloader.VideoData

	urlEntry.OnChanged = func(url string) {
		go HandleURLChange(url, dm, &vd, qualitySelect)
	}

	qualitySelect.OnChanged = func(quality string) {
		if vd == nil {
			return
		}
		go HandleQualityChange(vd, qualitySelect, codecSelect)
	}

	codecSelect.OnChanged = func(codec string) {
		if vd == nil {
			return
		}
		go HandleCodecChange(vd, codecSelect)
	}

	addButton := widget.NewButton("Add", func() {
		if vd == nil {
			return
		}
		go StartDownloadOnButtonClick(vd, dm)
		w.Close()
	})

	container := container.NewVBox(
		urlEntry,
		qualitySelect,
		codecSelect,
		addButton,
	)
	w.SetContent(container)
	w.Resize(fyne.NewSize(600, 300))
	return &AddDownloadWindow{
		Window:           w,
		DownloadsManager: dm,
	}
}

func (w *AddDownloadWindow) Show() {
	w.Window.Show()
}

func NewMainWindow(
	app fyne.App,
	dm *downloader.DownloadsManager,
) *MainWindow {
	w := app.NewWindow("YouGopher")
	addDownloadButton := widget.NewButton("Add download", func() {
		addDownloadWindow := NewAddDownloadWindow(app, dm)
		addDownloadWindow.Show()
	})
	downloadsList := widget.NewList(
		func() int {
			return len(dm.Downloads)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i int, o fyne.CanvasObject) {
			download := dm.Downloads[i]
			o.(*widget.Label).SetText(download.ID)
		},
	)
	mainContainer := container.NewVBox(
		downloadsList,
		layout.NewSpacer(),
		addDownloadButton,
	)
	w.SetContent(mainContainer)
	w.Resize(fyne.NewSize(800, 600))
	return &MainWindow{
		Window:            w,
		DownloadsList:     downloadsList,
		AddDownloadButton: addDownloadButton,
		DownloadsManager:  dm,
	}
}

func NewAppUI(
	dm *downloader.DownloadsManager,
	settings *settings.Settings,
) *AppUI {
	app := app.New()
	mainWindow := NewMainWindow(app, dm)
	return &AppUI{
		App:               app,
		MainWindow:        mainWindow,
		DownloadsManager:  dm,
		Settings:          settings,
		SettingsWindow:    nil,
		AddDownloadWindow: nil,
	}
}

func (a *AppUI) Run() {
	a.MainWindow.Window.ShowAndRun()
}
