package models

import (
	"dreamEbagPapers/helper"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	curlIdClient  *http.Client
	curlIdClient2 *http.Client
)

type CurlReseponId struct {
	F_id string `json:"F_id"`
}

type CurlReseponIntId struct {
	F_id int `json:"F_id"`
}

func abc() {
	curlIdClient = &http.Client{
		Transport: &http.Transport{
			//			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression:    true,
			ResponseHeaderTimeout: time.Second * 0,
		},
	}
	curlIdClient2 = &http.Client{
		Transport: &http.Transport{
			//			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression:    true,
			ResponseHeaderTimeout: time.Second * 0,
		},
	}
}

type MSnowflakCurl struct {
}

//获取发号器发出的ID(string类型,20位)
func (u *MSnowflakCurl) GetId() (id string) {
	id = ""

	uri := Config.SnowFlakDomain + "/v1/snowflak/id"
	method := "GET"
	req, _ := http.NewRequest(method, uri, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authentication", Config.SnowFlakAuthUser+":"+helper.Md5(Config.SnowFlakAuthUserSecurity))

	client := curlIdClient

	//log request
	// var logObj *MLog
	// logObj.LogSnowflakCurlRequest(uri, method, map[string]string{})

	resp, err := client.Do(req)
	idObj := CurlReseponId{}
	if err == nil {
		defer resp.Body.Close()
		bodyByte, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			dec := json.NewDecoder(strings.NewReader(string(bodyByte)))
			dec.UseNumber()
			dec.Decode(&idObj)
		}
		if resp.Status == "200 OK" {
			id = idObj.F_id
		}
		//log response
		// logObj.LogSnowflakCurlResponse(idObj, resp.Header, resp.Status)
	} else {
		//log err
		// logObj.LogErr2("snowflak module", err, "")
	}

	return
}

//获取发号器发出的ID(int类型,16位)
func (u *MSnowflakCurl) GetIntId() (id int) {
	id = 0

	uri := Config.SnowFlakDomain + "/v1/snowflak/intId"
	method := "GET"
	req, _ := http.NewRequest(method, uri, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authentication", Config.SnowFlakAuthUser+":"+helper.Md5(Config.SnowFlakAuthUserSecurity))

	client := curlIdClient2

	// //log request
	// var logObj *MLog
	// logObj.LogSnowflakCurlRequest(uri, method, map[string]string{})

	resp, err := client.Do(req)
	idObj := CurlReseponIntId{}
	if err == nil {
		defer resp.Body.Close()
		bodyByte, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			dec := json.NewDecoder(strings.NewReader(string(bodyByte)))
			dec.UseNumber()
			dec.Decode(&idObj)
		}
		if resp.Status == "200 OK" {
			id = idObj.F_id
		}
		//log response
		// logObj.LogSnowflakCurlResponse(idObj, resp.Header, resp.Status)
	} else {
		//log err
		// logObj.LogErr2("snowflak module", err, "")
	}

	return
}
