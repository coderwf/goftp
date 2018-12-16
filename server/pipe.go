package server

import (
	"fmt"
	"goftp/auth"
	"goftp/log"
	"goftp/msg"
	"goftp/util"
	"net"
)

type Pipe struct {
	//log
	log.Logger
	//id
	id              string
	//远程连接的地址
	Url             string
	//认证信息
	auth            *auth.Auth

	conn            net.Conn

	ctl             *Control //ctl

	closed          bool

}

func NewPipe(c net.Conn , ctl *Control) *Pipe{
	id    := util.RandId(3)
	url   := "tcp:" + c.RemoteAddr().String()
	pipe  := &Pipe{id:id,Url:url,conn:c,ctl:ctl,Logger:log.NewPrefixedLogger("Pipe")}
	pipe.AddPrefix(pipe.Id())
	pipe.Debug("Open")
	pipe.ctl.RegisterPipe(pipe.Id(),pipe)
	go pipe.manager()
	go pipe.processor()
	return pipe
}

func (p *Pipe) manager(){
    //fmt.Println("manager")
}

func (p *Pipe) processor(){
	var code    int16
	var param   string
	var err     error
    for {
    	if code , param , err = msg.ReadMsg(p.conn) ; err != nil{
    		p.Debug("Close with error %v",err)
    		p.ctl.UnRegisterPipe(p.Id())
			return
		}//if
		fmt.Println(code,param)
	}//
}

func (p *Pipe) Id() string{
	return p.id + "-" + p.Url
}

