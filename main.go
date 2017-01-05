package main

import (
	
	"os/exec"
	"flag"
	"fmt"
	"io/ioutil"
	"io"
	"bytes"
	"strings"
	"time"
	"github.com/jlaffaye/ftp"
	"encoding/json"
)

func uploadByFTP(host string, login string, pass string, dir string, filename string, data io.Reader) {
	conn, err := ftp.Connect(host)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Quit()
	if err = conn.Login(login, pass); err != nil {
		fmt.Println(err.Error())
		return
	}
	if err = conn.ChangeDir(dir); err != nil {
		fmt.Println(err.Error())
		return
	}
	if err = conn.Stor(filename, data); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func dump(config Config) {
	cmd := exec.Command("mysqldump", "-u", config.Mysql_user_name, "-p"+config.Mysql_password, "--databases", strings.Join(config.Mysql_databases, " "))
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	filename := time.Now().Format("2006-01-02_15:04:05") + ".sql"
	err = ioutil.WriteFile(config.Path + filename, out.Bytes(), 0777)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bytesReader := bytes.NewReader(out.Bytes())
	for _, conf := range config.Ftp_configs {
		uploadByFTP(conf.Host, conf.Login, conf.Pass, conf.Path, filename, bytesReader)
	}
}

func main() {
	configFile := flag.String("config", "config.json", "--config=/path/config.json")
	flag.Parse()
	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	config := Config{};
	if err = json.Unmarshal(data, &config); err != nil {
		fmt.Println(err.Error())
		return
	}
	dump(config)
}