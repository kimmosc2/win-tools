package function

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"os/exec"
)

// ChangePassword 利用cmd,net user修改密码
// main里直接copy过来了,今天太晚了,有空再改
func ChangePassword(username, password string) error {
	command := exec.Command("net", "user","BuTn","123456")
	stdoutPipe, err := command.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderrPipe, err := command.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = command.Start()
	if err != nil {
		log.Fatal(err)
	}


	reader := transform.NewReader(stdoutPipe, simplifiedchinese.GBK.NewDecoder())
	all, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(all))


	errReader := transform.NewReader(stderrPipe, simplifiedchinese.GBK.NewDecoder())
	errMsg, err := ioutil.ReadAll(errReader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(errMsg))
}
