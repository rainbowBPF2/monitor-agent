package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func main() {

	//executeRemoteFile()

	//ParseAppConf()

}
func executeRemoteFile() {
	var url = "http://10.45.11.194:45520/test.sh"
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err.Error())
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	//os
	//cmd
	//exec
	cmd := exec.Command("sh", "-c", string(data))
	out, err := cmd.Output()
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Print(string(out))
}
