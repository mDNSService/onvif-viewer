package main

import (
	"fmt"
	"github.com/mDNSService/onvif-viewer/services"

	"github.com/mDNSService/onvif"
	"github.com/mDNSService/onvif-viewer/config"
	"github.com/urfave/cli/v2"
	"log"
	"net"
	"os"
)

var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

func main() {
	myApp := cli.NewApp()
	myApp.Name = "onvif-viewer"
	myApp.Usage = "-c [config file path]"
	myApp.Version = buildVersion(version, commit, date, builtBy)
	myApp.Commands = []*cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "init config file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "config",
					Aliases:     []string{"c"},
					Value:       config.ConfigFilePath,
					Usage:       "config file path",
					EnvVars:     []string{"ConfigFilePath"},
					Destination: &config.ConfigFilePath,
				},
			},
			Action: func(c *cli.Context) error {
				config.LoadSnapcraftConfigPath()
				config.InitConfigFile()
				return nil
			},
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "run without config file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "xaddr",
					Aliases:     []string{"x"},
					Value:       config.ConfigModel.OnvifDevices[0].XAddr,
					Usage:       fmt.Sprintf("onvif xaddr eg: %s", config.ConfigModel.OnvifDevices[0].XAddr),
					EnvVars:     []string{"xaddr"},
					Destination: &config.ConfigModel.OnvifDevices[0].XAddr,
				},
				&cli.StringFlag{
					Name:        "username",
					Aliases:     []string{"u"},
					Value:       config.ConfigModel.OnvifDevices[0].UserName,
					Usage:       fmt.Sprintf("onvif username eg: %s", config.ConfigModel.OnvifDevices[0].UserName),
					EnvVars:     []string{"username"},
					Destination: &config.ConfigModel.OnvifDevices[0].UserName,
				},
				&cli.StringFlag{
					Name:        "password",
					Aliases:     []string{"p"},
					Value:       config.ConfigModel.OnvifDevices[0].Password,
					Usage:       fmt.Sprintf("onvif password eg: %s", config.ConfigModel.OnvifDevices[0].Password),
					EnvVars:     []string{"password"},
					Destination: &config.ConfigModel.OnvifDevices[0].Password,
				},
			},
			Action: func(c *cli.Context) error {
				return services.Run(c)
			},
		},
		{
			Name:    "find",
			Aliases: []string{"f"},
			Usage:   "find onvif divice",
			Action: func(c *cli.Context) error {
				itfs, err := net.Interfaces()
				if err != nil {
					return err
				}
				for _, itf := range itfs {
					devices := onvif.GetAvailableDevicesAtSpecificEthernetInterface(itf.Name)
					for _, dev := range devices {
						fmt.Println("===new device:===")
						fmt.Println("GetServices", dev.GetServices())
					}
				}
				return nil
			},
		},
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "test this command",
			Action: func(c *cli.Context) error {
				fmt.Println("ok")
				return nil
			},
		},
	}
	myApp.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Value:       config.ConfigFilePath,
			Usage:       "config file path",
			EnvVars:     []string{"ConfigFilePath"},
			Destination: &config.ConfigFilePath,
		},
	}
	myApp.Action = func(c *cli.Context) error {
		config.LoadSnapcraftConfigPath()
		_, err := os.Stat(config.ConfigFilePath)
		if err != nil {
			config.InitConfigFile()
		}
		config.UseConfigFile()
		return services.Run(c)
	}
	err := myApp.Run(os.Args)
	if err != nil {
		log.Println(err.Error())
	}
}

func buildVersion(version, commit, date, builtBy string) string {
	var result = version
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}
	return result
}
