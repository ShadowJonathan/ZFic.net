package ZFic

var webroot string
var fileroot string
var Serv *ServerConfig

func Load() (bool, string) {
	webroot = "../ZFic/WebPages"
	fileroot = "..ZFic/Files"
	Serv = &ServerConfig{
		Archiveworker: false,
		Stories:       len(GetArchive()),
		AWQ:           make(chan ARequest, 100),
	}
	return true, ""
}
