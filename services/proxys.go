package services

import (
	"fmt"
	"github.com/iotdevice/zeroconf"
	"github.com/mDNSService/onvif"
	"github.com/mDNSService/onvif-viewer/config"
	"github.com/mDNSService/onvif-viewer/models"
	"github.com/mDNSService/onvif-viewer/utils"
	//device "github.com/mDNSService/onvif/device"
	"github.com/mDNSService/onvif/media"
	onvif2 "github.com/mDNSService/onvif/xsd/onvif"
	"github.com/satori/go.uuid"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var Servers = make(map[string]*zeroconf.Server)

func Run(c *cli.Context) error {
	fmt.Println("config:", config.ConfigModel)
	for _, deviceConf := range config.ConfigModel.OnvifDevices {
		dev, err := onvif.NewDevice(deviceConf.XAddr)
		if err != nil {
			log.Println(err)
			continue
		}
		dev.Authenticate(deviceConf.UserName, deviceConf.Password)
		go ProxyAndRegRtsp(dev)
	}
	for {
		time.Sleep(time.Hour)
	}
	return nil
}

func ProxyAndRegRtsp(dev *onvif.Device) {
	//测试
	//log.Println("dev.GetServices():")
	//log.Println(dev.GetServices())
	//getCapabilities := device.GetCapabilities{Category: "All"}
	//resp, err := dev.CallMethod(getCapabilities)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//bytes, _ := ioutil.ReadAll(resp.Body)
	//log.Println("GetCapabilities:")
	//log.Println(string(bytes))
	//第一步：获取GetProfiles，从中获取token
	resp, err := dev.CallMethod(media.GetProfiles{})
	if err != nil {
		log.Println(err)
		return
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	//log.Println(string(bytes))
	tokens, err := utils.GetTokenFromGetProfiles(string(bytes))
	if err != nil {
		log.Println(err)
		return
	}
	if len(tokens) < 1 {
		log.Println("没有找到token:len(tokens)<1")
		return
	}
	log.Println(tokens[len(tokens)-1])
	//第二步：使用上一步的token获取GetStreamUri，获取视频流的地址
	resp, err = dev.CallMethod(media.GetStreamUri{
		StreamSetup: onvif2.StreamSetup{
			Stream:    "RTP-Unicast",
			Transport: onvif2.Transport{Protocol: "HTTP", Tunnel: nil},
		},
		ProfileToken: onvif2.ReferenceToken(tokens[len(tokens)-1]),
	})
	if err != nil {
		log.Println(err)
		return
	}
	bytes, _ = ioutil.ReadAll(resp.Body)
	//log.Println("GetStreamUri:", string(bytes))
	uri, err := utils.GetUriFromGetMediaUri(string(bytes))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(uri)
	URL, err := url.Parse(uri)
	if err != nil {
		log.Println(err)
		return
	}
	//TODO 第三部解析流视频的地址，并将端口使用mdns注册
	uuidStr := uuid.NewV4().String()
	host := strings.Split(URL.Host, ":")[0]
	port, err := strconv.Atoi(URL.Port())
	if err != nil {
		log.Println(err)
		return
	}
	username := URL.User.Username()
	password, _ := URL.User.Password()
	//log.Println(URL.Path)
	//log.Println(URL.Scheme)
	log.Println(host, port, username, password)
	newEntry := &models.ServiceInfo{
		Instance: uuidStr,
		Service:  "_iotdevice._tcp",
		Domain:   "local",
		Port:     port,
		HostName: uuidStr,
		Ip:       host,
		Text: []string{
			"name=Onvif Camera",
			"model=com.iotserv.services.vlc.player",
			"mac=unknown",
			fmt.Sprintf("id=%s", uuidStr),
			"author=Farry",
			"email=newfarry@126.com",
			"home-page=https://github.com/mDNSService/onvif-viewer",
			"firmware-respository=https://github.com/mDNSService/onvif-viewer",
			fmt.Sprintf("firmware-version=%s", "1.0"),
			//addition
			fmt.Sprintf("scheme=%s", URL.Scheme),
			fmt.Sprintf("path=%s", URL.Path),
			fmt.Sprintf("username=%s", username),
			fmt.Sprintf("password=%s", password),
		},
	}

	server, err := zeroconf.RegisterProxy(newEntry.Instance, newEntry.Service, newEntry.Domain,
		newEntry.Port, newEntry.HostName, []string{newEntry.Ip}, newEntry.Text, nil)
	if err != nil {
		log.Println(err)
		return
	}
	Servers[uuidStr] = server
}
