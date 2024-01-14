package services

import (
	"context"
	"errors"
	"io/fs"
	"io/ioutil"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
)

type spredSheetDownloadService struct {
	credentials []byte
}

const spredSheetDownloadServiceDefaultFilePermission fs.FileMode = 0644

func NewSpredSheetDownloadService(jsonCredentials string) spredSheetDownloadService {
	return spredSheetDownloadService{
		credentials: []byte(jsonCredentials),
	}
}

func (s spredSheetDownloadService) Exec(ctx context.Context, spredsheetID string, outputFileName string) error {
	config, err := google.JWTConfigFromJSON(s.credentials, drive.DriveScope)
	if err != nil {
		return errors.New("Unable to parse client secret file to config")
	}

	client := config.Client(ctx)
	srv, err := drive.New(client)
	if err != nil {
		return errors.New("Unable to retrieve Drive client")
	}

	file, err := srv.Files.Get(spredsheetID).Do()
	if err != nil {
		return errors.New("Unable to retrieve file metadata")
	}

	url := file.ExportLinks["application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"]
	resp, err := client.Get(url)
	if err != nil {
		return errors.New("Unable to retrieve file")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Unable to read file")
	}

	err = ioutil.WriteFile(outputFileName, data, spredSheetDownloadServiceDefaultFilePermission)
	if err != nil {
		return errors.New("failed write file")
	}

	return nil
}
