package create_file

import (
	"fmt"
	log "github.com/lwydyby/logrus"
	"os"
	"path"
	"strings"
)

/**
 * Created by Chris on 2021/9/20.
 */

/**
循环创建文件夹功能
*/

const logDri = "./log/"

func InitCreateDirectory() {
	if err := CreateDirectoryLoop(logDri); err != nil {
		log.Error(err)
	}
}

func PathFileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDirectory(path string) error {
	exist, err := PathFileExist(path)
	if err != nil {
		log.Errorf("Get directory error: %v", err)
		return err
	}
	if exist {
		log.Infof("The directory already exists : %v", path)
	} else {
		//log.Infof("The directory does not exist, creating it: %v", path)
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			log.Errorf("Failed to create the directory, because %v", err)
		} else {
			log.Infof("The directory was created successfully: %v", path)
		}
	}
	return nil
}

func CreateDirectoryLoop(pathName string) error {
	var isLinux bool
	var paths []string
	if strings.Contains(pathName, "/") {
		paths = strings.Split(pathName, "/")
		isLinux = true
	} else {
		paths = strings.Split(pathName, "\\")
		isLinux = false
	}
	var tempPath string
	if isLinux && paths[0] != "." {
		tempPath = path.Join(tempPath, "/")
	}
	for index := 0; index < len(paths); index++ {
		if len(paths[index]) == 0 {
			continue
		}
		tempPath = path.Join(tempPath, paths[index])
		if err := CreateDirectory(tempPath); err != nil {
			return err
		}
	}
	return nil
}

func Test() {
	//测试Windows系统
	//_dir := "D:\\Code\\go_echo_logrus\\gzFiles2\\xx"
	//_dir := "D:\\Code1\\gzFiles23\\xx"
	_dir := ".\\Code1\\gzFiles23\\xx"

	//测试Linux系统
	//_dir := "./Code2/gzFiles23/xx"
	//_dir := "/Code2/gzFiles23/xx"

	//开始创建文件夹
	if err := CreateDirectoryLoop(_dir); err != nil {
		fmt.Println(err)
	}
}
