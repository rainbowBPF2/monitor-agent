package utill

import (
	"agent/conf"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var cMap = make(map[string]string)
var mutex = sync.Mutex{}

func GetRemoteFileBody(filePath string) string {

	log.Print("begin download file:", filePath)
	resp, err := http.Get(filePath)

	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	fileBody := string(data)

	conMap := ParseAppConf()
	if conMap[conf.Encrypted] == "true" {

		fileBody = strings.Replace(fileBody, "\n", "", -1)
		fileBody = strings.Replace(fileBody, "\r", "", -1)
		fileBody, _ = DecryptAesData(fileBody)
	}

	return fileBody
}

func ParseAppConf() map[string]string {

	var configPath string

	mutex.Lock()
	defer mutex.Unlock()

	if len(cMap) != 0 {
		return cMap
	}

	home, _ := os.Getwd()
	configPath = filepath.Join(home, "src/agent/conf", "app.conf")

	if !FileExists(configPath) {
		configPath = filepath.Join(home, "conf", "app.conf")

		if !FileExists(configPath) {
			panic("app.conf not found!" + configPath)
		}
	}

	cFile, _ := os.Open(configPath)
	fileBytes, _ := ioutil.ReadAll(cFile)
	fileBody := string(fileBytes)

	fileBody = strings.Replace(fileBody, "\r", "", -1)
	bodyArr := strings.Split(fileBody, "\n")

	for _, body := range bodyArr {
		if body != "" && !strings.HasPrefix(body, "#") {
			arrV := strings.Split(body, "=")
			if len(arrV) == 2 {
				cMap[arrV[0]] = arrV[1]
			}
		}
	}

	return cMap
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
