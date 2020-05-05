package ui

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type configitem struct {
	localitems   map[string]string
	sftpitems    map[string]string
	Encryptitems map[string]string
}

type itemstruct struct {
	Localbackup_Compress    string
	Localbackup_storagepath string
	Localbackup_includes    []string
	Localbackup_excludes    []string

	Sftpbackup_includes       []string
	Sftpbackup_excludes       []string
	Sftpbackup_Compress       string
	Sftpbackup_Connecttimeout string
	Sftpbackup_Username       string
	Sftpbackup_password       string
	Sftpbackup_hostip         string
	Sftpbackup_storagepath    string

	Encryptbackup_Compress    string
	Encryptbackup_Encrypt     string
	Encryptbackup_password    string
	Encryptbackup_salt        string
	Encryptbackup_storagepath string
	Encrptbackup_includes     []string
	Encryptbackup_excludes    []string
}

func SliceFolder(folder string) []string {
	folders := strings.Split(folder, "\n")
	return folders
}

func ParseText(item configitem) {
	yaml := `models:
    local_backup:
      compress_with:
          type: {{.Localbackup_Compress}}
      store_with:
          type: local
          keep: 15
          path: {{.Localbackup_storagepath}}
      archive:
          includes:{{range  $k, $v := .Localbackup_includes}}
            - {{$v}}{{end}}    
          excludes:{{range  $k, $v := .Localbackup_excludes}}
            - {{$v}}{{end}}
        
    sftp_backup:
      compress_with:
          type: {{.Sftpbackup_Compress}}
      store_with:
          type: ftp
          keep: 15
          path: {{.Sftpbackup_storagepath}}
          host: {{.Sftpbackup_hostip}}
          port: 21
          timeout: {{.Sftpbackup_Connecttimeout}}
          username: {{.Sftpbackup_Username}}
          password: {{.Sftpbackup_password}}
      archive:
          includes:{{range  $k, $v := .Sftpbackup_includes}}
            - {{$v}}{{end}}      
          excludes:{{range  $k, $v := .Sftpbackup_excludes}}
            - {{$v}}{{end}}
      
    Encrypt_backup:
      compress_with:
          type: {{.Encryptbackup_Compress}}
      encrypt_with:
          type: {{.Encryptbackup_Encrypt}}
          password: {{.Encryptbackup_password}}
          salt: {{.Encryptbackup_salt}}
          openssl: true
      store_with:
          type: local
          keep: 15
          path: {{.Encryptbackup_storagepath}}
      archive:
          includes:{{range  $k, $v := .Encrptbackup_includes}}
            - {{$v}}{{end}}
          excludes:{{range  $k, $v := .Encryptbackup_excludes}}
            - {{$v}} {{end}}
    `

	Localbackup_includes := SliceFolder(item.localitems["backup_include"])
	Localbackup_excludes := SliceFolder(item.localitems["backup_excludes"])
	Sftpbackup_includes := SliceFolder(item.sftpitems["backup_include"])
	Sftpbackup_excludes := SliceFolder(item.sftpitems["backup_excludes"])
	Encrptbackup_includes := SliceFolder(item.Encryptitems["backup_include"])
	Encryptbackup_excludes := SliceFolder(item.Encryptitems["backup_excludes"])

	var itemstruct itemstruct
	itemstruct.Localbackup_includes = Localbackup_includes
	itemstruct.Localbackup_excludes = Localbackup_excludes
	itemstruct.Sftpbackup_includes = Sftpbackup_includes
	itemstruct.Sftpbackup_excludes = Sftpbackup_excludes
	itemstruct.Encrptbackup_includes = Encrptbackup_includes
	itemstruct.Encryptbackup_excludes = Encryptbackup_excludes

	/* map[Compress:tgz backup_excludes:  backup_include:  storge_path:/tmp/gobackup]
	   map[Compress:tgz Connect_timeout: Username: backup_excludes:  backup_include:  host_ip: password: storge_path:/tmp/gobackup]
	   map[Compress:tgz Encrypt:Openssl Encrypt_password:  Encrypt_salt:true backup_excludes:  backup_include:  storge_path:/tmp/gobackup] */

	itemstruct.Localbackup_Compress = item.localitems["Compress"]
	itemstruct.Localbackup_storagepath = item.localitems["storge_path"]

	itemstruct.Sftpbackup_hostip = item.sftpitems["host_ip"]
	itemstruct.Sftpbackup_Username = item.sftpitems["Username"]
	itemstruct.Sftpbackup_password = item.sftpitems["password"]
	itemstruct.Sftpbackup_Connecttimeout = item.sftpitems["Connect_timeout"]
	itemstruct.Sftpbackup_storagepath = item.sftpitems["storge_path"]
	itemstruct.Sftpbackup_Compress = item.sftpitems["Compress"]

	itemstruct.Encryptbackup_Compress = item.Encryptitems["Compress"]
	itemstruct.Encryptbackup_Encrypt = item.Encryptitems["Encrypt"]
	itemstruct.Encryptbackup_password = item.Encryptitems["Encrypt_password"]
	itemstruct.Encryptbackup_salt = item.Encryptitems["Encrypt_salt"]
	itemstruct.Encryptbackup_storagepath = item.Encryptitems["storge_path"]
	fmt.Println(itemstruct.Localbackup_Compress)

	t := template.Must(template.New("yaml").Parse(yaml))
	file, err := os.OpenFile("gobackup.yaml", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	err = t.Execute(file, itemstruct)
	if err != nil {
		fmt.Println("executing template:", err)
	}
}
