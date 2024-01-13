package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	bf "github.com/russross/blackfriday/v2"
	"github.com/sokorahen-szk/sokorahen-szk/cmd"
)

type config struct {
	Resume resume `json:"resume"`
}

type resume struct {
	InputFile        string `json:"input_file"`
	OutputFile       string `json:"output_file"`
	HTMLTemplateFile string `json:"html_template_file"`
}

const tmpHTMLFileName = "tmp.html"

func main() {
	generateResume(readConfig())
}

func readConfig() config {
	cmdArg := cmd.NewCmdArg()
	if !cmdArg.Validate([]string{"config"}) {
		log.Fatal("invalid arguments error")
	}

	configFilePath, err := cmdArg.Get("config")
	if err != nil {
		log.Fatal(err)
	}

	var config config
	cmd.ReadConfig(*configFilePath, &config)

	return config
}

func generateResume(config config) error {
	mdData, err := ioutil.ReadFile(config.Resume.InputFile)
	if err != nil {
		return err
	}

	markdown_html := bf.Run(mdData)
	htmlTemplate, _ := ioutil.ReadFile(config.Resume.HTMLTemplateFile)
	html := strings.Replace(string(htmlTemplate), "#body#", string(markdown_html), 1)

	err = ioutil.WriteFile(tmpHTMLFileName, []byte(html), 0644)
	if err != nil {
		return err
	}

	cmd := exec.Command("wkhtmltopdf", tmpHTMLFileName, config.Resume.OutputFile)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("stdoutStderr: %+v, failed to generate pdf err: %+v", stdoutStderr, err)
	}

	return nil
}
