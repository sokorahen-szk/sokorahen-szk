package services

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	bf "github.com/russross/blackfriday/v2"
)

type convertPDFService struct {
}

const (
	tmpHTMLFileName                        string      = "tmp.html"
	convertPDFServiceDefaultFilePermission fs.FileMode = 0644
)

func NewConvertPDFService() convertPDFService {
	return convertPDFService{}
}

func (s convertPDFService) ConvertMarkdownToPDF(inputFileName string, outputFileName string) error {
	if err := s.convertMarkdownToHTML(inputFileName); err != nil {
		return err
	}

	cmd := exec.Command("wkhtmltopdf", "--enable-local-file-access", "--print-media-type", tmpHTMLFileName, outputFileName)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("stdoutStderr: %+v, failed to generate pdf err: %+v", stdoutStderr, err)
	}

	return nil
}

func (s convertPDFService) ConvertXLSXToPDF(inputFileName string, outputFileName string) error {
	cmd := exec.Command("soffice", "--headless", "--language=ja", "--convert-to", "pdf", "--outdir", filepath.Dir(outputFileName), inputFileName)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("stdoutStderr: %+v, failed to generate pdf err: %+v", stdoutStderr, err)
	}

	return nil
}

func (s convertPDFService) convertMarkdownToHTML(inputFileName string) error {
	mdData, err := os.ReadFile(inputFileName)
	if err != nil {
		return err
	}

	markdownHTML := bf.Run(mdData)

	html := fmt.Sprintf(`
        <html>
            <head>
                <meta charset="UTF-8" />
				<link href="./src/files/style.css" rel="stylesheet" />
            </head>
            <body>%s</body>
        </html>`, string(markdownHTML))
	err = os.WriteFile(tmpHTMLFileName, []byte(html), convertPDFServiceDefaultFilePermission)
	if err != nil {
		return err
	}

	return nil
}
