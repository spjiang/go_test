package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	message "test/chatroom/common"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte // 传输时，使用的缓冲
}

// 接受数据，并解析成message.Message结构体
func (t *Transfer) ReadPkg() (mes message.Message, err error) {
	// 接受内容
	_, err = t.Conn.Read(t.Buf[:4])
	if err != nil {
		fmt.Println("conn.Read err", err)
		//err = errors.New("read pkg header err")
		return
	}

	// 根据buf[:4]转化成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(t.Buf[0:4])

	// 根据pkgLen长度，读取内容放到buf里面去
	n, err := t.Conn.Read(t.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read fail err=", err)
		//err = errors.New("read body header err")
		return
	}

	// 把pkgLen 反序列化-> message.Message
	err = json.Unmarshal(t.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	return
}

func (t *Transfer) WritePkg(data []byte) (err error) {
	// 1、先发送一个数据长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(t.Buf[:4], pkgLen)
	// 发送长度
	_, err = t.Conn.Write(t.Buf[:4])
	if err != nil {
		fmt.Println("conn.Write(buf[:4]) fail err=", err)
		return
	}
	// 发送数据包
	n, err := t.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) fail err=", err)
		return
	}
	return
}
