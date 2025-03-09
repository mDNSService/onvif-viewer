package services

import (
	"encoding/json"
	"fmt"
	"github.com/OpenIoTHub/utils/net"
	"github.com/mDNSService/onvif"
	"github.com/mDNSService/onvif-viewer/config"
	"github.com/mDNSService/onvif-viewer/models"
	"log"
	"net/http"
)

func RegOnvifCameraManager() (err error) {
	http.HandleFunc("/", Index)
	http.HandleFunc("/list", GetOnvifCameraList)
	http.HandleFunc("/add", AddOnvifCamera)
	http.HandleFunc("/delete", DeleteOnvifCamera)
	//port, err := nettool.GetOneFreeTcpPort()
	port := 34324
	fmt.Println("Http ListenAndServe on:", port)

	var txts = map[string]string{
		"name":                 "OnvifCameraManager",
		"model":                "com.iotserv.services.OnvifCameraManager",
		"author":               "Farry",
		"email":                "newfarry@126.com",
		"home-page":            "https://github.com/OpenIoTHub",
		"firmware-respository": "https://github.com/mDNSService",
		"firmware-version":     "1.0",
	}
	nettool.RegistermDNSService(txts, port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	return
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OnvifCameraManager"))
}

func GetOnvifCameraList(w http.ResponseWriter, r *http.Request) {
	for _, dev := range config.ConfigModel.OnvifDevices {
		log.Println(dev.Name)
	}
	var onvifDeviceConfigListResp = models.OnvifDeviceConfigListResp{
		Code:         0,
		Msg:          "",
		OnvifDevices: config.ConfigModel.OnvifDevices,
	}
	data, err := json.Marshal(onvifDeviceConfigListResp)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	_, err = w.Write(data)
	if err != nil {
		log.Println(err)
	}
}

func AddOnvifCamera(w http.ResponseWriter, r *http.Request) {
	var operationResp = new(models.OperationResp)
	query := r.URL.Query()
	//Name     string
	Name := query["Name"][0]
	//XAddr    string
	XAddr := query["XAddr"][0]
	//UserName string
	UserName := query["UserName"][0]
	//Password string
	Password := query["Password"][0]
	//add device
	dev, err := onvif.NewDevice(XAddr)
	if err != nil {
		log.Println(err)
		operationResp.Code = 1
		operationResp.Msg = err.Error()
		out, _ := json.Marshal(operationResp)
		w.Write(out)
		return
	}
	dev.Authenticate(UserName, Password)
	go RegRtspProxy(dev, Name, XAddr)
	//write config file
	config.ConfigModel.OnvifDevices = append(config.ConfigModel.OnvifDevices, models.OnvifDeviceConfig{
		Name:     Name,
		XAddr:    XAddr,
		UserName: UserName,
		Password: Password,
	})
	err = config.WriteConfigFile(config.ConfigModel, "")
	if err != nil {
		log.Println(err)
		operationResp.Code = 1
		operationResp.Msg = err.Error()
		out, _ := json.Marshal(operationResp)
		w.Write(out)
		return
	}
	//
	out, _ := json.Marshal(operationResp)
	w.Write(out)
}

func DeleteOnvifCamera(w http.ResponseWriter, r *http.Request) {
	var operationResp = new(models.OperationResp)
	query := r.URL.Query()
	XAddr := query["XAddr"][0]
	//delete device
	UnRegRtspProxy(XAddr)
	//write config file
	for k, v := range config.ConfigModel.OnvifDevices {
		if v.XAddr == XAddr {
			config.ConfigModel.OnvifDevices = append(config.ConfigModel.OnvifDevices[:k], config.ConfigModel.OnvifDevices[k+1:]...)
		}
	}
	err := config.WriteConfigFile(config.ConfigModel, "")
	if err != nil {
		log.Println(err)
		operationResp.Code = 1
		operationResp.Msg = err.Error()
		out, _ := json.Marshal(operationResp)
		w.Write(out)
		return
	}
	//
	out, _ := json.Marshal(operationResp)
	w.Write(out)
}
