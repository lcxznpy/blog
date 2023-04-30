package main

import (
	"encoding/json"
	"fmt"
	//"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
	"os"
)

const file = "/go_blog/blog_server/models/res/err_code.json"

type ErrMap map[string]string

func main() {
	byteData, err := os.ReadFile(file)
	if err != nil {
		logrus.Error(err)
		return
	}
	var errMap = ErrMap{}
	err = json.Unmarshal(byteData, &errMap)
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(errMap)
}
