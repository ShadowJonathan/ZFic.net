package ZFic

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Archive(w http.ResponseWriter, r *http.Request) {

}

func GetArchive() ([]*Story, error) {
	var stories []*Story
	Stories, err := ioutil.ReadFile(fileroot + "/Stories.Ar")
	if err != nil {
		// stuff here
	}
	err = json.Unmarshal(Stories, stories)
	if err == nil {
		return stories, err
	}
	log.Fatal(err)
	return stories, err
}
