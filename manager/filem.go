package manager

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"
	"runtime"
	"strings"
)

//工作目录管理

type WDirM struct {
	cwd           string   //当前目录
}

func (d *WDirM) Pwd() string{
	return d.cwd
}

func (d *WDirM) Cd(tPath string) string{
    return path.Join(d.cwd,tPath)
}

func (d *WDirM) Mkdir(tPath string){
    aPath := path.Join(d.cwd , tPath)
    os.MkdirAll(aPath,os.ModePerm)
}

func Home() (string , error){
	if usr , err := user.Current() ; err == nil{
		return usr.HomeDir , nil
	}//

	if runtime.GOOS == "windows"{
        return homeWindows()
	}
	return homeLinuxLike()
}

func homeWindows() (string , error){
    drive  := os.Getenv("HOMEDRIVE")
    _path  := os.Getenv("HOMEPATH")
    home   := drive + _path
    if drive == "" || _path == ""{
		home = os.Getenv("USERPROFILE")
	}
    if home == ""{
    	return "" , fmt.Errorf("Can't find home in HOMEDRIVE,HOMEPATH,USERPROFILE")
	}//if
	return home , nil
}

func homeLinuxLike() (string , error){
    if home := os.Getenv("HOME") ; home != ""{
    	return home , nil
	}//

	var stdout bytes.Buffer //从命令行中得到

	cmd := exec.Command("sh","-c","eval echo ~$USER")
	cmd.Stdout   = &stdout
	if err := cmd.Run() ; err != nil{
		return "", err
	}

	res  := strings.TrimSpace(stdout.String())
	if res == ""{
		return "" , fmt.Errorf("Can't find home")
	}
	return res , nil
}

