package main

import (
	"encoding/json"
	"fmt"

	"net/http"

	util "github.com/krisztiansala/golang-helloworld/internal"
	log "github.com/sirupsen/logrus"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, helloMessage(""))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, helloMessage(name))
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonResp := versionResponse(GitHash, GitProject)
	w.Write(jsonResp)
}

func helloMessage(name string) string {
	name = util.ParseName(name)
	if name == "" {
		return "Hello Stranger"
	}
	return fmt.Sprintf("Hello %s", name)
}

// Write test for this, specify the json expected in the test and convert it into a byte array when comparing
func versionResponse(gitHash string, project string) []byte {
	resp := make(map[string]string)
	resp["hash"] = gitHash
	resp["project"] = project
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Errorf("Error happened in JSON marshal. Err: %s", err)
		return nil
	}
	fmt.Println(string(jsonResp))
	return jsonResp
}
