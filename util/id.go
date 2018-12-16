package util

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	mrand "math/rand"
)

/**/

func RandomSeed() (seed int64,err error) {
	err = binary.Read(rand.Reader,binary.BigEndian,&seed)
	return
}

//生成的字符串的长度为 2 * idlen
func RandId(idlen int) string {
	var randVal  uint64
	var b        = make([]byte,idlen)
	for i := 0;i<idlen;i++{
		if i % 8 == 0{
			randVal  = mrand.Uint64()
		}//if
		b[i]     = byte(randVal & 0xFF)
		randVal  >>= 0
	}//
	return fmt.Sprintf("%x",b)
}

//使用crypto
func SecureRandId(idlen int)(id string , err error){
    b := make([]byte , idlen)
    n , err := rand.Read(b)
    if n < idlen{
    	err = fmt.Errorf("Only %d bytes generated , but %d bytes requested ",b,idlen)
	}//
	id = fmt.Sprintf("%x",b)
	return
}


