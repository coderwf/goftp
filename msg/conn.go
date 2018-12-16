package msg

import (
	"encoding/binary"
	"fmt"
	"net"
	"unicode/utf8"
)

//长度只是param的长度
func ReadMsg(c net.Conn)(code int16,param string,err error){
    var sz   int32
    if err = binary.Read(c,binary.BigEndian,&sz) ; err != nil{
    	return
	}//

	if err = binary.Read(c,binary.BigEndian,&code) ; err != nil{
		return
	}
	buffer := make([]byte,sz)
	var     n  int
	if n , err = c.Read(buffer) ; err != nil{
		return
	}//
	if n != int(sz){
		err = fmt.Errorf("Expected to read %d bytes,but only read %d bytes",sz,n)
		return
	}
	//编码为utf-8
	if !utf8.Valid(buffer) {
		err  = fmt.Errorf("Request UTF8 encoding")
		return
	}
    param   = string(buffer)
    return
}

func WriteMsg(c net.Conn , code int16 , param string)(err error){
	buffer        := []byte(param)
	if !utf8.Valid(buffer){
		err  = fmt.Errorf("Request UTF8 encoding")
		return
	}
	var sz   = int32(len(buffer))
    if err = binary.Write(c,binary.BigEndian,sz) ; err != nil{
    	return
	}//if
	if err = binary.Write(c,binary.BigEndian,code) ; err != nil{
		return
	}
	var n int
	if n , err = c.Write(buffer) ; err != nil{
		return
	}
	if n != int(sz){
		err = fmt.Errorf("Expected to write %d bytes , but only write %d bytes",sz,n)
		return
	}//if
	return
}

//按照格式写入msg
func WriteMsgF(c net.Conn , code int16 ,format string,args ...interface{}) error{
	param  := fmt.Sprintf(format,args)
	return WriteMsg(c,code,param)
}
//没有此用户
func ReplyNoUser(c net.Conn , user string) error{
	return WriteMsgF(c,NoUser,"No User named %s",user)
}

func ReplyAuthFailed(c net.Conn) error {
	return WriteMsg(c,AuthFail,"User or Password error")
}

func ReplyNoFound(c net.Conn , name string) error{
    return WriteMsgF(c,NotFound,"Directory or File %s Not found",name)
}

func ReplyAlreadyExist(c net.Conn,name string) error{
	return WriteMsgF(c,AlreadyExist,"Directory or File %s already exists")
}


