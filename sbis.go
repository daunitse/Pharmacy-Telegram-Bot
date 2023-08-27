package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

var coneectIDSbis = "id"
var passKeySbis = "passkey"
var secretKey = "secretkey"
var token = "token"
var sid = "sid"
var pointIdlEN = "pointidlen"
var pricelistid = "pricelistid"
var now = time.Now()
var formTime = now.Format("2006-01-02")
var drugs string
var answToClient = ""

func sbis() {

	client := &http.Client{}
	pointid := "pointId=" + pointIdlEN + "&"
	priceListId := "priceListId=" + pricelistid + "&"
	searchString := "searchString=" + url.QueryEscape(drugs)
	withBalance := "withBalance=true&"
	noStopList := "noStopList=true&"

	url := "https://api.sbis.ru/retail/nomenclature/list?" + pointid + priceListId + noStopList + withBalance + searchString

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("creating new request: %s", err)
		return
	}
	req.Header.Add("X-SBISAccessToken", token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("sending request: %s", err)
		return
	}

	defer resp.Body.Close()
	resp_body, _ := io.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	drugsQTLine := gjson.Get(string(resp_body), "nomenclatures.#")
	if drugsQTLine.String() != "" {
		wrightDrugsQT, err := strconv.Atoi(drugsQTLine.String())
		if err != nil {
			log.Fatal(err)
		}
		drugsQT := wrightDrugsQT - 1
		drugsCost := gjson.Get(string(resp_body), "nomenclatures."+strconv.Itoa(drugsQT)+".cost")
		drugsBalance := gjson.Get(string(resp_body), "nomenclatures."+strconv.Itoa(drugsQT)+".balance")
		drugsName := gjson.Get(string(resp_body), "nomenclatures."+strconv.Itoa(drugsQT)+".name")

		for i := 0; i != drugsQT; {
			if drugsBalance.String() != "" {
				answToClient += drugsName.String() + " стоит " + drugsCost.String() + " Рублей. Осталось " + drugsBalance.String() + " шт\n\n"
			}
			drugsQT = drugsQT - 1
			drugsCost = gjson.Get(string(resp_body), "nomenclatures."+strconv.Itoa(drugsQT)+".cost")
			drugsBalance = gjson.Get(string(resp_body), "nomenclatures."+strconv.Itoa(drugsQT)+".balance")
			drugsName = gjson.Get(string(resp_body), "nomenclatures."+strconv.Itoa(drugsQT)+".name")
		}

	}

}
