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

func NewAddDownloadWindow(
	app fyne.App,
	dm *downloader.DownloadsManager,
) *AddDownloadWindow {
	window := app.NewWindow("Add download")
	window.SetContent(container.NewVBox(
		widget.NewLabel("This is a child window!"),
		widget.NewButton("Add", func() {
		}),
	))

	window.Resize(fyne.NewSize(300, 200))
	window.Show()
	return &AddDownloadWindow{
		Window:           window,
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
	mainWindow := app.NewWindow("YouGopher")
	addDownloadButton := widget.NewButton("Add download", func() {
		addDownloadWindow := NewAddDownloadWindow(app, dm)
		addDownloadWindow.Show()
	})
	downloadsList := widget.NewList(
		func() int {
			return len(dm.Downloads)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(i int, o fyne.CanvasObject) {
			download := dm.Downloads[i]
			o.(*widget.Label).SetText(download.ID)
		},
	)
	mainContainer := container.NewVBox(
		layout.NewSpacer(), // Отступ между списком и кнопкой
		addDownloadButton,
	)
	mainWindow.SetContent(mainContainer)
	mainWindow.Resize(fyne.NewSize(400, 300))
	return &MainWindow{
		Window:            mainWindow,
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
