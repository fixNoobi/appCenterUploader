package internal

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

const API_URL = "https://api.appcenter.ms/v0.1/apps/"

func Run() {
	kingpin.Parse()
	LogInit()
	newConfig().run()
}

func newConfig() *AppStruct {

	a := AppStruct{
		ApiToken:         *API_TOKEN_FLAG,
		ApiOwner:         *API_OWNER_FLAG,
		ApiNameApp:       *API_NAMEAPP_FLAG,
		AppFile:          *APP_FILE_FLAG,
		AppVersionNumber: *APP_VERSION_NUMBER_FLAG,
		AppBuildNumber:   *APP_BUILD_NUMBER_FLAG,
		AppSumbolFile:    *APP_SUMBOL_FILE_FLAG,
		ApiUri:           API_URL + *API_OWNER_FLAG + "/" + *API_NAMEAPP_FLAG,
		LogLevel:         *LOG_LEVEL,
	}
	return &a
}

func (a *AppStruct) run() {
	a.CreateApp()
	a.UploadAppToAppCenter()
	a.PublishAppFromAppCenter()
	a.PrepearSymbolUpload()
	a.UploadSymbols()
	a.CommittedSymbols()
}
