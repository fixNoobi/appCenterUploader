package internal

type StructReleaseInit struct {
	UploadID  string `json:"upload_id"`
	UploadURL string `json:"upload_url"`
}

type StructRelease struct {
	ReleaseID  string `json:"release_id"`
	ReleaseURL string `json:"release_url"`
}

type SymbolUploads struct {
	Build          string `json:"build"`
	ClientCallback string `json:"client_callback"`
	FileName       string `json:"file_name"`
	SymbolType     string `json:"symbol_type"`
	Version        string `json:"version"`
	Status         string `json:"status"`
}

type StructSymbolUploadStatus struct {
	ExpirationDate string `json:"expiration_date"`
	SymbolUploadID string `json:"symbol_upload_id"`
	UploadURL      string `json:"upload_url"`
}

type AppStruct struct {
	ApiUri             string
	ApiOwner           string
	ApiNameApp         string
	ApiToken           string
	AppFile            string
	AppSumbolFile      string
	AppBuildNumber     string
	AppVersionNumber   string
	ReleaseInit        StructReleaseInit
	StatusRelease      StructRelease
	SymbolUploadStatus StructSymbolUploadStatus
	LogLevel           string
}
