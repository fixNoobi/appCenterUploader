package internal

import (
	"bytes"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func (a *AppStruct) CreateApp() {

	log.Info("Start Create app.")
	req, err := http.NewRequest("POST", a.ApiUri+"/release_uploads", nil)
	if err != nil {
		log.Errorf("Error create app. Error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Api-Token", a.ApiToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("Error Create app. Error: %v", err)
	}
	if err := json.NewDecoder(resp.Body).Decode(&a.ReleaseInit); err != nil {
		log.Errorf("Error: %v", err)
	}
	log.Info("End Create app.")
	log.Debugf("ID %v", a.ReleaseInit.UploadID)
	log.Debugf("Upload Url: %v", a.ReleaseInit.UploadURL)

	defer resp.Body.Close()
}

func (a *AppStruct) UploadAppToAppCenter() {
	log.Info("Start upload app.")
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "multipart/form-data").
		EnableTrace().
		SetFile("ipa", a.AppFile).
		Post(a.ReleaseInit.UploadURL)
	if err != nil {
		log.Errorf("Error upload file to AppCenter. Error: %v", err)
	}
	log.Info("End upload app.")
	result := string([]byte(resp.Body()))
	log.Debugf("Body: %v", result)
}

func (a *AppStruct) PublishAppFromAppCenter() {

	log.Info("Start public app.")
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "multipart/form-data").
		SetHeader("X-Api-Token", a.ApiToken).
		SetBody(`{"status": "committed"}`).
		EnableTrace().
		Patch(a.ApiUri + "/release_uploads/" + a.ReleaseInit.UploadID)

	if err = json.NewDecoder(bytes.NewReader(resp.Body())).Decode(&a.StatusRelease); err != nil {
		log.Errorf("Error: %v", err)
	}
	log.Info("End public app.")
	log.Debugf("ID %v", a.StatusRelease.ReleaseID)
	log.Debugf("Url: %v", a.StatusRelease.ReleaseURL)
}
