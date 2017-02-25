package ZFic

import (
	"fmt"
)

//noinspection GoUnusedGlobalVariable
var Serv *ServerConfig
var ZFServ *HttpServer
var DB = &[]*User{}
var Sesses *map[string]*Session

//noinspection GoUnusedExportedFunction
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
	ZFServ = &HttpServer{
		port:    "9000",
		address: "",
	}
	return ZFServ, err
}
