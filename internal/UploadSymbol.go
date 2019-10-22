package internal

import (
	"bytes"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"path"
)

func (a *AppStruct) PrepearSymbolUpload() {
	log.Info("Start prepear symbol upload.")
	client := resty.New()

	_, fileName := path.Split(a.AppSumbolFile)
	data2, err := json.Marshal(SymbolUploads{
		Build:          a.AppBuildNumber,
		SymbolType:     "AndroidProguard",
		Version:        a.AppVersionNumber,
		ClientCallback: "string",
		FileName:       fileName,
	})

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Api-Token", a.ApiToken).
		EnableTrace().
		SetBody(data2).
		Post(a.ApiUri + "/symbol_uploads")
	if err != nil {
		log.Errorf("Error ipload: %v", err)
	}

	if err = json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&a.SymbolUploadStatus); err != nil {
		log.Errorf("Error: %v", err)
	}
	log.Info("End prepear symbol.")
	log.Debugf("ID %v", a.SymbolUploadStatus.SymbolUploadID)
	log.Debugf("Url: %v", a.SymbolUploadStatus.UploadURL)
}

func (a *AppStruct) UploadSymbols() {

	log.Info("Start upload symbol.")
	client := resty.New()

	resp, err := client.R().
		SetHeader("x-ms-blob-type", "BlockBlob").
		SetHeader("Context-Type", "text/plain; charset=UTF-8").
		SetHeader("X-Api-Token", a.ApiToken).
		EnableTrace().
		//SetBody(data2).
		SetFile("file_name", a.AppSumbolFile).
		Put(a.SymbolUploadStatus.UploadURL)
	if err != nil {
		log.Errorf("Error ipload: %v", err)
	}
	log.Info("End upload app.")
	log.Debugf("Body       :\n", resp)
}

func (a *AppStruct) CommittedSymbols() {

	log.Info("Start CommittedSymbols app.")
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "multipart/form-data").
		SetHeader("X-Api-Token", a.ApiToken).
		SetBody(`{"status": "committed"}`).
		EnableTrace().
		Patch(a.ApiUri + "/symbol_uploads/" + a.SymbolUploadStatus.SymbolUploadID)

	if err = json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&a.StatusRelease); err != nil {
		log.Errorf("Error: %v", err)
	}
	log.Info("End public app.")
	log.Debugf("ID %v", a.StatusRelease.ReleaseID)
	log.Debugf("Url: %v", a.StatusRelease.ReleaseURL)

}
