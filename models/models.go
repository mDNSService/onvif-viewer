package models

type ConfigModel struct {
	//this DVR name
	Name         string
	OnvifDevices []OnvifDeviceConfig
}

type OnvifDeviceConfig struct {
	//this Camera name
	Name     string
	XAddr    string
	UserName string
	Password string
}

type OperationResp struct {
	Code int
	Msg  string
}

type OnvifDeviceConfigListResp struct {
	Code         int
	Msg          string
	OnvifDevices []OnvifDeviceConfig
}

type ServiceInfo struct {
	Instance string   `json:"instance"`
	Service  string   `json:"service"`
	Domain   string   `json:"domain"`
	Port     int      `json:"port"`
	HostName string   `json:"host_name"`
	Ip       string   `json:"ip"`
	Text     []string `json:"text"`
}
