package utils

import (
	"github.com/beevik/etree"
	"log"
)

func GetTokenFromGetProfiles(doc string) (tokens []string, err error) {
	document := etree.NewDocument()
	err = document.ReadFromString(doc)
	if err != nil {
		return
	}
	tokenEls := document.Root().FindElements("//SOAP-ENV:Envelope/SOAP-ENV:Body/trt:GetProfilesResponse/trt:Profiles")
	//log.Println(len(tokenEls))
	for _, el := range tokenEls {
		//log.Println(el.SelectAttr("token").Value)
		tokens = append(tokens, el.SelectAttr("token").Value)
	}
	log.Println(tokens)
	return
}
func GetUriFromGetMediaUri(doc string) (uri string, err error) {
	document := etree.NewDocument()
	err = document.ReadFromString(doc)
	if err != nil {
		return
	}
	tokenEl := document.Root().FindElement("//SOAP-ENV:Envelope/SOAP-ENV:Body/trt:GetStreamUriResponse/trt:MediaUri/tt:Uri")
	uri = tokenEl.Text()
	return
}
