package main

import (
	"flag"
	"fmt"
	"github.com/lenbo-ma/gokits/log"
	"os"
	"strings"
)

const (
	exec    = "ginpt"
	version = "1.0.0"
)

var (
	command  string
	rootPath string
)

var (
	appName   string
	appPort   int
	typesFile string
)

var commandMap map[string]*Command

type CommandFuncType func(name, desc string)
type Command struct {
	Name string
	Desc string
	Func CommandFuncType
}

func init() {
	// 初始化参数值
	flag.StringVar(&typesFile, "types", "types.yaml", "types config file")
	flag.StringVar(&appName, "name", "", "gin app name")
	flag.IntVar(&appPort, "port", 0, "gin server port")
}

func showHelp(name, desc string) {
	commands := make([]string, 0, len(commandMap))
	for _, v := range commandMap {
		commands = append(commands, fmt.Sprintf("%s\t%s", v.Name, v.Desc))
	}
	printHelp(fmt.Sprintf("Usage: %s <command> [options]", exec), commands,
		[]string{
			"-types\ttypes config file, default: types.yaml",
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
		"gen_gin_server": &Command{
			Name: "gen_gin_server",
			Desc: "generate gin server code from template",
			Func: genGinServer,
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
	flag.CommandLine.Parse(os.Args[2:])
	initCommands()
	parseRootDir()
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
