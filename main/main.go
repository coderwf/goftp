package main

import (
	"goftp/log"
	"goftp/server"
	"net"
)

func main(){
	log.LogTo(log.STDOUT,log.DEBUG)
	Ip  := net.ParseIP("0.0.0.0")
	listener  , err := net.ListenTCP("tcp",&net.TCPAddr{IP:Ip,Port:9999})
	if err != nil{
		log.Warn("Listen with error:%v",err)
		return
	}//
	ctl := server.NewControl(listener)
	ctl.WaitClose()
}
