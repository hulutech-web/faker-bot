package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

type MyLogger struct {
}

func (m *MyLogger) Print(message string) {

}
func (m *MyLogger) Trace(message string) {

}
func (m *MyLogger) Debug(message string) {

}
func (m *MyLogger) Info(message string) {

}
func (m *MyLogger) Warning(message string) {

}
func (m *MyLogger) Error(message string) {

}
func (m *MyLogger) Fatal(message string) {

}

var myLogger = &MyLogger{}

func main() {
	// Create an instance of the app structure
	app := NewApp()
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "评论区座谈会",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Logger:                           myLogger,
		LogLevel:                         logger.DEBUG,
		LogLevelProduction:               logger.ERROR,
		BackgroundColour:                 &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:                        app.startup,
		EnableFraudulentWebsiteDetection: false,
		EnableDefaultContextMenu:         false,
		ErrorFormatter:                   func(err error) any { return err.Error() },
		Bind: []interface{}{
			app,
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: false,
		},
		// Mac平台特定选项
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 true,
				HideToolbarSeparator:       false,
			},
			Appearance:           mac.NSAppearanceNameAccessibilityHighContrastVibrantLight,
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			About: &mac.AboutInfo{
				Title:   "评论区座谈会",
				Message: "抖音评论区座谈会",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
