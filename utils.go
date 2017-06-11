package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/lenbo-ma/gokits/log"
)

func mkDirIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

func templateExecute(bs []byte, targetPath string, model interface{}) {
	var err error
	var f *os.File
	t := template.New("gpt")
	if t, err = t.Funcs(funcMap).Parse(string(bs)); err != nil {
		log.Fatalf("%s", err)
	}
	if f, err = os.Create(targetPath); err != nil {
		log.Fatalf("%s", err)
	}
	if err = t.Execute(f, model); err != nil {
		log.Fatalf("%s", err)
	}
	if f != nil {
		f.Close()
	}
}

func updateFileName(n string) string {
	if n == "bin/god" {
		return fmt.Sprintf("bin/%s-god", appName)
	}

	if n == "bin/god-linux" {
		return fmt.Sprintf("bin/%s-god-linux", appName)
	}

	if strings.HasSuffix(n, ".tmpl") {
		return strings.Replace(n, ".tmpl", "", 1)
	}
	return n
}
