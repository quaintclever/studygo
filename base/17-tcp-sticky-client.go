package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	conn, e := net.Dial("tcp", ":8080")
	if e != nil {
		fmt.Println("遇到错误")
		return
	}
	defer conn.Close()
	for i := 0; i < 200; i++ {
		//msg := "Hello quaint! "
		//conn.Write([]byte(msg))
		msg := "Hello quaint! "
		data, err := Encode(msg)
		if err != nil {
			fmt.Println("err", err)
			return
		}
		conn.Write(data)
	}
}

func Encode(message string) ([]byte, error) {
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}
