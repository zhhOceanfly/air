package main

import (
	"os/exec"
	"strings"
)

var airVersion string
var goVersion string

func init() {
	airVersion = "v1.27.4"
	goVersion,_ = GetGoVersion()
}
func GetGoVersion() (string, error) {
	//获取go版本号
	goVersions, err := exec.Command("bash", "-c", "go version").CombinedOutput()
	goVersionArr := strings.Split(string(goVersions), " ")
	if err != nil || len(goVersionArr) < 3 || len(goVersionArr[2]) < 3 {
		return "", err
	}
	return goVersionArr[2][2:], nil
}