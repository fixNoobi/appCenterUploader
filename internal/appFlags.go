package internal

import "gopkg.in/alecthomas/kingpin.v2"

var (
	API_OWNER_FLAG   = kingpin.Flag("api_owner", "Verbose mode.").Short('o').Required().Envar("APPCENTER_API_OWNER").String()
	API_NAMEAPP_FLAG = kingpin.Flag("api_nameapp", "Verbose mode.").Short('a').Required().Envar("APPCENTER_API_NAMEAPP").String()
	API_TOKEN_FLAG   = kingpin.Flag("api_token", "Verbose mode.").Short('t').Required().Envar("APPCENTER_API_TOKEN").String()

	APP_FILE_FLAG           = kingpin.Flag("app_file", "Verbose mode.").Short('f').Required().Envar("APPCENTER_APP_FILE").String()
	APP_SUMBOL_FILE_FLAG    = kingpin.Flag("app_sumbol_file", "Verbose mode.").Short('s').Required().Envar("APPCENTER_APP_SUMBOL_FILE").String()
	APP_BUILD_NUMBER_FLAG   = kingpin.Flag("app_build", "Verbose mode.").Short('b').Required().Envar("APPCENTER_APP_BUILD_NUMBER").String()
	APP_VERSION_NUMBER_FLAG = kingpin.Flag("app_version", "Verbose mode.").Short('n').Required().Envar("APPCENTER_APP_VERSION_NUMBER").String()

	LOG_LEVEL = kingpin.Flag("log-level", "Verbose mode.").Short('l').Default("INFO").String()
)
