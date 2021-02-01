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
	tokenEls := document.Root().FindElements("//Envelope/Body/GetProfilesResponse/Profiles")
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
	tokenEl := document.Root().FindElement("//Envelope/Body/GetStreamUriResponse/MediaUri/Uri")
	uri = tokenEl.Text()
	return
}
