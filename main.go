package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/domego/gokits/log"
	"github.com/domego/gopt/gens/common"
)

const (
	exec    = "gopt"
	version = "1.0.0"
)

var (
	command  string
	rootPath string
)

var (
	appName    string
	appPort    int
	typesFile  string
	errorsFile string
	ormFile    string
	apiFile    string
	template   string
)

var commandMap map[string]*Command

type CommandFuncType func(name, desc string)
type Command struct {
	Name string
	Desc string
	Func CommandFuncType
}

func init() {
	genutils.Asset = Asset
	// 初始化参数值
	flag.StringVar(&errorsFile, "errors", "errors.yaml", "errors config file")
	flag.StringVar(&typesFile, "types", "types.yaml", "types config file")
	flag.StringVar(&ormFile, "orm", "db.yaml", "database config file")
	flag.StringVar(&apiFile, "api", "api.yaml", "gin controller config file")
	flag.StringVar(&appName, "name", "", "gin app name")
	flag.StringVar(&template, "template", "vue", "api.js template")
	flag.IntVar(&appPort, "port", 0, "gin server port")
}

func showHelp(name, desc string) {
	commands := make([]string, 0, len(commandMap))
	for _, v := range commandMap {
		commands = append(commands, fmt.Sprintf("%s\t%s", v.Name, v.Desc))
	}
	printHelp(fmt.Sprintf("Usage: %s <command> [options]", exec), commands,
		[]string{
			"-errors\terrors config file, default: errors.yaml",
			"-types\ttypes config file, default: types.yaml",
			"-orm\tdatabase config file, default: db.yaml",
			"-api\tgin controller config file, default: api.yaml",
			"-template\tapi.js template, default: vue",
		}, nil)
}

func initCommands() {
	for i, v := range os.Args {
		switch i {
		case 1:
			command = v
		}
	}

	// 初始化命令列表
	commandMap = map[string]*Command{
		"version": &Command{
			Name: "version",
			Desc: "show version",
			Func: showVersion,
		},
		"help": &Command{
			Name: "help",
			Desc: "show help",
			Func: showHelp,
		},
		"gen_types": &Command{
			Name: "gen_types",
			Desc: "generate structs code from yaml config file",
			Func: genTypes,
		},
		"gen_errors": &Command{
			Name: "gen_errors",
			Desc: "generate structs code from yaml config file",
			Func: genErrors,
		},
		"gen_gin_server": &Command{
			Name: "gen_gin_server",
			Desc: "generate gin server code from template",
			Func: genGinServer,
		},
		"gen_orm": &Command{
			Name: "gen_orm",
			Desc: "generate database orm code from template",
			Func: genORM,
		},
		"gen_gin_api": &Command{
			Name: "gen_gin_api",
			Desc: "generate gin controller code from template",
			Func: genGinController,
		},
		"gen_js_api": &Command{
			Name: "gen_js_api",
			Desc: "generate api.js code from template",
			Func: genJavascriptAPI,
		},
		"gen_api_doc": &Command{
			Name: "gen_api_doc",
			Desc: "generate api_doc document from template",
			Func: genAPIDoc,
		},
	}
}

func checkArgs() bool {
	if command == "" {
		showHelp("help", commandMap["help"].Desc)
		return false
	} else if command == "gen_gin_server" {
		if appName == "" {
			log.Fatalf("name can not be empty")
		}
		if appPort <= 0 {
			log.Fatalf("port can not be empty")
		}
	}
	return true
}

func printHelp(usage string, commands, options, examples []string) {
	fmt.Println()
	fmt.Println(usage)
	if len(commands) > 0 {
		fmt.Println("\nCommands:\n")
		for _, s := range commands {
			fmt.Println(fmt.Sprintf("\t%s", s))
		}
	}

	if len(options) > 0 {
		fmt.Println("\nOptions:\n")
		for _, s := range options {
			fmt.Println(fmt.Sprintf("\t%s", s))
		}
	}

	if len(examples) > 0 {
		fmt.Println("\nExamples:\n")
		for _, s := range examples {
			fmt.Println(fmt.Sprintf("\t%s", s))
		}
	}

	fmt.Println()
}

func showVersion(name, desc string) {
	fmt.Println(version)
}

func parseRootDir() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("%s", err)
	}
	pathList := strings.Split(os.Getenv("GOPATH"), ":")
	for _, path := range pathList {
		if strings.HasPrefix(currentDir, path) {
			var prefix string
			if strings.HasSuffix(path, "/") {
				prefix = path + "src/"
			} else {
				prefix = path + "/src/"
			}
			currentDir = currentDir[len(prefix):]
		}

		rootPath = currentDir
	}
}

func main() {
	parseRootDir()
	initCommands()
	if len(os.Args) < 2 {
		showHelp("help", commandMap["help"].Desc)
		return
	}
	flag.CommandLine.Parse(os.Args[2:])
	if !checkArgs() {
		return
	}
	c := commandMap[command]
	if c == nil {
		showHelp("help", commandMap["help"].Desc)
		return
	} else {
		c.Func(c.Name, c.Desc)
	}
}
