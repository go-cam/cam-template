package main

import (
	"fmt"
	"github.com/go-cam/cam/core/camConstants"
	"github.com/go-cam/cam/core/camModels"
	"github.com/go-cam/cam/core/camUtils"
	"os"
	"strings"
)

func main() {
	cam := newCam()
	cam.run()
}

// the struct of camCommand command
type camCommand struct {
}

func newCam() *camCommand {
	c := new(camCommand)
	return c
}

// run command
func (c *camCommand) run() {
	argLen := len(os.Args)
	if argLen < 2 {
		c.help()
		return
	}

	action := os.Args[1]
	switch action {
	case "init":
		c.init()
	case "help":
		fallthrough
	default:
		c.help()
		break
	}
}

// init .env file
func (c *camCommand) init() {
	c.initCamJsonFile()
	c.initEnvFile()
}

// init camCommand.json file
func (c *camCommand) initCamJsonFile() {
	cam := new(camModels.Cam)
	cam.Modules = []*camModels.CamModule{}

	// common
	commonModule := new(camModels.CamModule)
	commonModule.Name = "common"
	commonModule.Type = camConstants.CamModuleTypeLib
	cam.Modules = append(cam.Modules, commonModule)

	// server
	serverModule := new(camModels.CamModule)
	serverModule.Name = "server"
	serverModule.Type = camConstants.CamModuleTypeApp
	cam.Modules = append(cam.Modules, serverModule)

	jsonBytes := camUtils.Json.Encode(cam)
	err := camUtils.File.WriteFile("cam.json", jsonBytes)
	camUtils.Error.Panic(err)
}

// init .env file
func (c *camCommand) initEnvFile() {
	rootPath := camUtils.File.GetRunPath()
	envFileContent := c.getEnvFileContent()
	envFileContent = strings.Replace(envFileContent, "[ROOT_PATH]", rootPath, -1)
	bytes := []byte(envFileContent)

	cam := c.getCam()
	for _, module := range cam.Modules {
		if module.Type != camConstants.CamModuleTypeApp {
			continue
		}

		filename := rootPath + "/" + module.Name + "/.env"
		err := camUtils.File.WriteFile(filename, bytes)
		camUtils.Error.Panic(err)
	}
}

// .env file content
func (c *camCommand) getEnvFileContent() string {
	return `
# project config
ROOT_PATH = [ROOT_PATH]

# database config
DB_DRIVER_NAME = mysql
DB_HOST = localhost
DB_PORT = 3306
DB_NAME = camCommand
DB_USERNAME = root
DB_PASSWORD = 123456
`
}

// get cam struct form file, cam.json
func (c *camCommand) getCam() *camModels.Cam {
	bytes, err := camUtils.File.ReadFile("cam.json")
	camUtils.Error.Panic(err)

	cam := new(camModels.Cam)
	camUtils.Json.DecodeToObj(bytes, cam)

	return cam
}

// print help info
func (c *camCommand) help() {
	fmt.Println("Cam help:")
	fmt.Println("    help        print help info")
	fmt.Println("    init        init project")
}
