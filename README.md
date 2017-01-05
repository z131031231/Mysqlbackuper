# Mysqlbackuper
Fast export of Mysql DB to your FTP server

### Configuration file
```json 
{
	"mysql_user_name": "LOGIN",
	"mysql_password": "PASS",
	"mysql_databases": ["DB1", "DB2"],
	"path": "/path/to/your/local/backup/folder/",
	"ftp_configs": [
		{
			"host": "HOST",
			"login": "LOGIN",
			"pass": "PASS",
			"path": "/path/to/backup/folder/on/your/server/"
		}
	]
}
```

### How to use
1. Edit configuration file
2. Build the tool
2. ./mysqlbackuper --config=/path/to/your/config/file
