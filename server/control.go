package server

import (
	"goftp/log"
	"net"
	"sync"
)

type Control struct {
	log.Logger // add log
    listener      *net.TCPListener
	pipesTable    map[string]*Pipe
	sync.Mutex
	close         chan bool
	cwd           string
}

func NewControl(listener *net.TCPListener) *Control{
    ctl := &Control{
    	Logger:log.NewPrefixedLogger("Control"),
    	pipesTable:make(map[string]*Pipe),
        listener:listener,
        close:make(chan bool),
    	}
    go ctl.listenTcp()
    return ctl
}

func (ctl *Control) listenTcp(){
	defer func() {
		if r := recover() ; r != nil{
			ctl.Debug("Closed with panic :%v",r)
		}
	}()

	ctl.Debug("Listen tcp :%s",ctl.listener.Addr().String())
	for {
		tcpConn , err := ctl.listener.AcceptTCP()
		if err != nil{
			ctl.Debug("Listen tcp failed with error :%v",err)
		}//
		ctl.HandleConnection(tcpConn)
	}//for
}

func (ctl *Control) RegisterPipe(id string, pipe *Pipe){
	ctl.Lock()
	defer ctl.Unlock()
	ctl.Debug("Register pipe %s",id)
	ctl.pipesTable[id]  = pipe
}

func (ctl *Control) UnRegisterPipe(id string) {
	ctl.Lock()
	defer ctl.Unlock()
	if ctl.pipesTable[id] == nil{
		return
	}
	ctl.Debug("Unregister pipe %s",id)
	delete(ctl.pipesTable,id)
}

func (ctl *Control)HandleConnection(c net.Conn){
	ctl.Debug("New connection from %s",c.RemoteAddr())
	NewPipe(c,ctl)
}

func (ctl *Control) WaitClose(){
	<-ctl.close
}

func (ctl *Control) Close(){
	close(ctl.close)
}