package service

import (
	"jnpdf/schema"

	"io/ioutil"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

var systemURL = beego.AppConfig.String("systemURL")

// GetPDFCertificate is get back the remote data
func GetPDFCertificate(id string) (certificate schema.Information, err error) {
	request := httplib.Get(systemURL + "/certificate/GetPDFCertificate/" + id)
	request.Header("Content-Type", "application/json")
	request.Header("Accept-Encoding", "gzip, deflate")

	response, err := request.DoRequest()
	if err != nil {
		beego.Error(err)
		return certificate, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		beego.Error(err)
		return certificate, err
	}
	var replyInformation schema.ReplyInformation
	json.Unmarshal(contents, &replyInformation)

	contents, err = json.Marshal(&replyInformation.Result[0])
	if err != nil {
		beego.Error(err)
		return certificate, err
	}
	json.Unmarshal(contents, &certificate)

	return certificate, nil
}

// GetImage is get back the remote data
func GetImage(id string) (contents []byte, err error) {
	request := httplib.Get(systemURL + "/image/GetImage/" + id)
	//request.Header("Content-Type", "application/json")
	//request.Header("Accept-Encoding", "gzip, deflate")

	response, err := request.DoRequest()
	if err != nil {
		beego.Error(err)
		return contents, err
	}

	contents, err = ioutil.ReadAll(response.Body)
	if err != nil {
		beego.Error(err)
		return contents, err
	}

	return contents, nil
}
