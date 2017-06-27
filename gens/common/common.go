package genutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/domego/gokits/log"
	"github.com/lenbo-ma/ginpt/gens/funcs"
)

var (
	// Asset asset 用于读取二进制文件内容
	Asset func(name string) ([]byte, error)
	// Values 变量集合
	Values map[string]interface{}
)

func init() {
	Values = make(map[string]interface{})
}

// SetValues 设置变量
func SetValues(values map[string]interface{}) {
	for k, v := range values {
		Values[k] = v
	}
}

// InitDirs 初始化指定文件夹
func InitDirs(dirs []string) {
	// mkdirs
	for _, dir := range dirs {
		if err := MkDirIfNotExists(dir); err != nil {
			log.Fatalf("%s", err)
		}
	}
}

// CopyFiles copy file
func CopyFiles(files []string) {
	// copy files
	for _, n := range files {
		bs, err := Asset(n)
		if err != nil {
			log.Fatalf("%s", err)
		}
		n = UpdateFileName(n)
		var mode os.FileMode = 0644
		if strings.Contains(n, "bin/") || strings.HasSuffix(n, ".sh") {
			mode = 0755
		}
		if err = ioutil.WriteFile(n, bs, mode); err != nil {
			log.Fatalf("%s", err)
		}
	}
}

// GenFiles 根据模板生成内容
func GenFiles(files []string, model map[string]interface{}) {
	for _, n := range files {
		GenFile(n, model)
	}
}

// Genfile 根据模板生成文件(模板和输出路径相同)
func GenFile(filePath string, model interface{}) {
	GenFileWithTargetPath(filePath, filePath, model)
}

// GenFile  根据模板生成文件(指定输出路径)
func GenFileWithTargetPath(templatePath, targetPath string, model interface{}) {
	log.Infof("start gen file %s", targetPath)
	var (
		bs  []byte
		err error
	)
	if bs, err = Asset(templatePath); err != nil {
		log.Fatalf("%s", err)
	}
	targetPath = UpdateFileName(targetPath)
	TemplateExecute(bs, targetPath, model)
}

// MkDirIfNotExists 如果指定文件夹不存在则创建
func MkDirIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

// TemplateExecute 根据模板生成内容并写入文件
func TemplateExecute(bs []byte, targetPath string, model interface{}) {
	var err error
	var f *os.File
	t := template.New("gpt")
	if t, err = t.Funcs(funcs.FuncMap).Parse(string(bs)); err != nil {
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

// UpdateFileName 对部分文件名做特殊处理
func UpdateFileName(n string) string {
	if n == "bin/god" {
		if Values["AppName"] == "" {
			log.Fatalf("genutils.Values[\"AppName\"] can not be empty")
		}
		return fmt.Sprintf("bin/%s-god", Values["AppName"])
	}

	if n == "bin/god-linux" {
		if Values["AppName"] == "" {
			log.Fatalf("genutils.Values[\"AppName\"] can not be empty")
		}
		return fmt.Sprintf("bin/%s-god-linux", Values["AppName"])
	}

	if strings.HasSuffix(n, ".tmpl") {
		return strings.Replace(n, ".tmpl", "", 1)
	}
	return n
}
