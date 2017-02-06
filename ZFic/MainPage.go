package ZFic

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	MPD, err := ioutil.ReadFile(MP)
	if err != nil {
		panic(err)
	}
	var icfg = &TopLayer{}
	err = json.Unmarshal(MPD, icfg)
	if err != nil {
		log.Fatal(err)
	}
	TL, err := ioutil.ReadFile(tp + icfg.TP)
	if err != nil {
		log.Fatal(err)
	}
	var total []byte
	if icfg.Inception {
		total = HandleLayer(icfg.I, TL)
	} else {
		total = TL
	}
	err = ioutil.WriteFile(Mainpagehtml, total, 9001)
	if err != nil {
		log.Fatal(err)
	}
	http.ServeFile(w, r, Mainpagehtml)
}

func HandleLayer(L *sublayer, UL []byte) []byte {
	LAYER, err := ioutil.ReadFile(ZFroot + L.TP)
	if err != nil {
		log.Fatal(err)
	}
	var total []byte
	if L.IsMD {
		total = bytes.Replace(UL, []byte("TPDATA"), GetMD(LAYER), 1)
	} else {
		total = bytes.Replace(UL, []byte("TPDATA"), LAYER, 1)
	}
	if L.Inception {
		return HandleLayer(L.I, total)
	}
	return total
}
