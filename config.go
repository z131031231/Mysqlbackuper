package main

type Config struct {
	Mysql_user_name string
	Mysql_password string
	Mysql_databases []string
	Path string
	Ftp_configs []Ftp_config
}

type Ftp_config struct {
	Host string
	Login string
	Pass string
	Path string
}