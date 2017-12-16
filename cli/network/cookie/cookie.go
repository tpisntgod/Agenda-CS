package cookie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var writeFilePath = "src/github.com/bilibiliChangKai/Agenda-CS/network/cookie/cookie.json"
var mycookie http.Cookie

func ReadCookie() {
	writeFilePath = filepath.Join(os.Getenv("GOPATH"), writeFilePath)
	data, err := ioutil.ReadFile(writeFilePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &mycookie)
	if err != nil {
		//panic(err)
	}
	fmt.Println(mycookie.Name)
	result2print, _ := json.MarshalIndent(mycookie, "", "    ")
	fmt.Print(string(result2print) + "\n")
}

func WriteCookie(cookie *http.Cookie) {
	writeFilePath := filepath.Join(os.Getenv("GOPATH"), writeFilePath)
	b, err := json.Marshal(cookie)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(writeFilePath, b, 0644)
	if err != nil {
		panic(err)
	}
}

func DeleteCookie() {
	writeFilePath = filepath.Join(os.Getenv("GOPATH"), writeFilePath)
	os.Remove(writeFilePath)
}
func GetCookie() *http.Cookie {
	ReadCookie()
	return &mycookie
}
func ExistCookie() bool {
	writeFilePath = filepath.Join(os.Getenv("GOPATH"), writeFilePath)
	_, err := os.Stat(writeFilePath)
	if err == nil {
		return true
	}
	return false
}
