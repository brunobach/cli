package create

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/brunobach/cli/internal/config"
	"github.com/brunobach/cli/internal/pkg/helper"
	"github.com/brunobach/cli/internal/tpl"
)

type Create struct {
	ProjectName                 string
	CreateType                  string
	FilePath                    string
	FileName                    string
	FileNameTitleLower          string
	FileNameFirstChar           string
	FileNameCamelCase           string
	FileNameCamelCaseTitleLower string
	IsFull                      bool
	TplPath                     string
}

func NewCreate() *Create {
	return &Create{}
}

func Run(cfg *config.Cfg) {
	c := NewCreate()
	c.ProjectName = helper.GetProjectName(".")
	c.CreateType = cfg.CreateType
	c.TplPath = cfg.TplPath
	c.FilePath, c.FileName = filepath.Split(cfg.Args[0])
	c.FileName = strings.ReplaceAll(strings.ToUpper(string(c.FileName[0]))+c.FileName[1:], ".go", "")
	c.FileNameTitleLower = strings.ToLower(string(c.FileName[0])) + c.FileName[1:]
	c.FileNameFirstChar = string(c.FileNameTitleLower[0])
	c.FileNameCamelCase = helper.CamelCase(c.FileName)
	c.FileNameCamelCaseTitleLower = strings.ToLower(string(c.FileNameCamelCase[0])) + c.FileNameCamelCase[1:]

	switch c.CreateType {
	case "controller":
		c.genFile()
	default:
		log.Fatalf("Invalid handler type: %s", c.CreateType)
	}

}
func (c *Create) genFile() {
	filePath := c.FilePath

	switch c.CreateType {
	case "controller":
		filePath = "src/application/controllers/"
	}

	f := createFile(filePath, strings.ToLower(c.FileName+"_"+c.CreateType)+".go")
	if f == nil {
		log.Printf("warn: file %s%s %s", filePath, strings.ToLower(c.FileName)+".go", "already exists.")
		return
	}
	defer f.Close()
	var t *template.Template
	var err error
	if c.TplPath == "" {
		t, err = template.ParseFS(tpl.CreateTemplateFS, fmt.Sprintf("create/%s.tpl", c.CreateType))
	} else {
		t, err = template.ParseFiles(path.Join(c.TplPath, fmt.Sprintf("%s.tpl", c.CreateType)))
	}
	if err != nil {
		log.Fatalf("create %s error: %s", c.CreateType, err.Error())
	}
	err = t.Execute(f, c)
	if err != nil {
		log.Fatalf("create %s error: %s", c.CreateType, err.Error())
	}
	log.Printf("Created new %s: %s", c.CreateType, filePath+strings.ToLower(c.FileName)+".go")

}

func createFile(dirPath, filename string) *os.File {
	filePath := filepath.Join(dirPath, filename)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create dir %s: %v", dirPath, err)
	}
	stat, _ := os.Stat(filePath)
	if stat != nil {
		return nil
	}
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", filePath, err)
	}

	return file
}
