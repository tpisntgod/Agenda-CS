package cookie

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var writeFilePath = "src/github.com/bilibiliChangKai/Agenda-CS/cli/network/cookie/cookie.json"
var mycookie http.Cookie

func init() {
	writeFilePath = filepath.Join(*GetGOPATH(), writeFilePath)
}

func ReadCookie() {
	data, err := ioutil.ReadFile(writeFilePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &mycookie)
	if err != nil {
		//panic(err)
	}
}

func WriteCookie(cookie *http.Cookie) {
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
	os.Remove(writeFilePath)
}
func GetCookie() *http.Cookie {
	ReadCookie()
	return &mycookie
}
func ExistCookie() bool {
	_, err := os.Stat(writeFilePath)
	if err == nil {
		return true
	}
	return false
}

//GetGOPATH 获得用户环境的gopath
func GetGOPATH() *string {
	var sp string
	if runtime.GOOS == "windows" {
		sp = ";"
	} else {
		sp = ":"
	}
	goPath := strings.Split(os.Getenv("GOPATH"), sp)
	for _, v := range goPath {
		if _, err := os.Stat(filepath.Join(v, "/src/github.com/bilibiliChangKai/Agenda-CS/cli/network/cookie/cookie.go")); !os.IsNotExist(err) {
			return &v
		}
	}
	return nil
}
