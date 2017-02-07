package ZFic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Archive(w http.ResponseWriter, r *http.Request) {

}

func GetArchive() ([]*Story, error) {
	var stories []*Story
	var Emerr error
	Stories, err := ioutil.ReadFile(fileroot + "/Stories.Ar")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(Stories, stories)
	if err == nil {
		return stories, err
	} else {
		fmt.Println(err)
		return stories, Emerr
	}
	return stories, err
}
