package ZFic

import (
	"fmt"
)

var Serv *ServerConfig
var ZFIC *HttpServer
var DB = &[]*User{}
var Sesses *map[string]*Session

func Load() (*HttpServer, error) {
	Archive, err := GetArchive()
	LoadDataBase()
	LoadSessions()
	StartReaper()
	if err != nil {
		fmt.Println(err.Error())
	}
	Serv = &ServerConfig{
		Archiveworker: false,
		Stories:       int64(len(Archive)),
		AWQ:           make(chan ARequest, 100),
	}
	ZFIC = &HttpServer{
		port:    "80",
		address: "",
	}
	return ZFIC, err
}
