package network

import (
	//logger "github.com/donnie4w/go-logger/logger"
	"time"
	"net"
	//"errors"
)

type Client struct {
	Conn              net.Conn                       
	heartbeatTimeout  int64              
	lastRecvTime      int64              
	lastSendTime      int64              
	bConnState        bool               
	DealRecv          FClientDeal            
}

//向链接写入数据
func (client *Client) Write(buf []byte) (err error) {
	/*
	client.writeLock.Lock()
	defer client.writeLock.Unlock()

	if client.bConnState == false {
		return errors.New("Connect server failed")
	} else {
		for n := 0; n != len(buf); {
			i, err := client.Conn.Write(buf[n:])
			if err != nil {
				return err
			}
			n += i
		}
		client.lastSendTime = time.Now().Unix()
	}
	log.Infor("write:", buf)
	*/
	return nil
}


func (client *Client) client_goroutine() {
	recvInfoChan := make(chan []byte, CHAN_BUF_COUNT)
	defer close(recvInfoChan)

	if client.DealRecv != nil {
		go client.DealRecv(recvInfoChan, client)
	}
	for {
		recvbuf := make([]byte, RECV_BUF_SIZE)
		client.Conn.SetReadDeadline(time.Now().Add(READ_TIME * time.Millisecond))
		n, err := client.Conn.Read(recvbuf)
		if err != nil {
		} else {
			recvInfoChan <- recvbuf[:n]
		}
	}
}
