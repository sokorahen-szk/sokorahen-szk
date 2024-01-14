package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sokorahen-szk/sokorahen-szk/cmd"
	"github.com/sokorahen-szk/sokorahen-szk/services"
)

type config struct {
	Resume          resume          `json:"resume"`
	PersonalHistory personalHistory `json:"personal_history"`
}

type resume struct {
	InputFile  string `json:"input_file"`
	OutputFile string `json:"output_file"`
}

type personalHistory struct {
	SpredSheetID string `json:"spredsheet_id"`
	InputFile    string `json:"input_file"`
	OutputFile   string `json:"output_file"`
}

func main() {
	cfg := readConfig()
	jsonCredentials := os.Getenv("GOOGLE_SECRET_CREDENTIALS")
	// 職務経歴書 -> ./outputs
	generateResume(cfg)
	//履歴書 -> ./outputs
	generatePersonalHistory(cfg, jsonCredentials)
}

// generateResume() 職務経歴書を生成する
func generateResume(cfg config) {
	fmt.Println("generate resume Start...")

	convertPDFService := services.NewConvertPDFService()
	err := convertPDFService.ConvertMarkdownToPDF(cfg.Resume.InputFile, cfg.Resume.OutputFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("generate resume Done!...")
}

// generatePersonalHistory() 履歴書を生成する
func generatePersonalHistory(cfg config, jsonCredentials string) {
	fmt.Println("generate personal history start...")

	spredsheetDownloadService := services.NewSpredSheetDownloadService(jsonCredentials)
	err := spredsheetDownloadService.Exec(
		context.Background(),
		cfg.PersonalHistory.SpredSheetID,
		cfg.PersonalHistory.InputFile,
	)
	if err != nil {
		log.Fatalln(err)
	}

	convertPDFService := services.NewConvertPDFService()
	err = convertPDFService.ConvertXLSXToPDF(cfg.PersonalHistory.InputFile, cfg.PersonalHistory.OutputFile)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("generate personal history Done!...")
}

// readConfig() コンフィグを読み込む
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
