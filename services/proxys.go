package services

import (
	"fmt"
	"github.com/mDNSService/onvif"
	"github.com/mDNSService/onvif-viewer/config"

	"github.com/urfave/cli/v2"
	"log"
	"time"
)

func Run(c *cli.Context) error {
	fmt.Println("config:", config.ConfigModel)
	for _, deviceConf := range config.ConfigModel.OnvifDevices {
		dev, err := onvif.NewDevice(deviceConf.XAddr)
		if err != nil {
			log.Println(err)
			continue
		}
		dev.Authenticate(deviceConf.UserName, deviceConf.Password)
		go RegRtspProxy(dev, deviceConf.Name, deviceConf.XAddr)
	}
	RegOnvifCameraManager()
	for {
		time.Sleep(time.Hour)
	}
	return nil
}
