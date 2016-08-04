package network

import (
	logger "github.com/donnie4w/go-logger/logger"
	"net"
)

type FClientDeal func(recvInfoChan chan []byte, client *Client)

type ServerConf struct {
	Addr               string  			//host:port
	HeartBeatTimeOut   int64
	DealRecv          FClientDeal
}

type Server struct {
	ServerConf
	listener             net.Listener
}

func (server *Server) Listen() (err error) {
	server.listener, err = net.Listen("tcp", server.Addr)
	if err != nil {
		logger.Error("listen address [", server.Addr, "] failed:", err)
		return err
	}
	logger.Info("listen address [", server.Addr, "]")

	go server.accept_goroutine()
	return nil
}

func (server *Server) accept_goroutine() {
	defer server.listener.Close()

	var err error
	for {
		client := Client {
			//HeartbeatTimeout:  server.HeartbeatTimeout,
			DealRecv: server.DealRecv,
		}
		client.Conn, err = server.listener.Accept()
		if err != nil {
			logger.Error("accept ", server.listener.Addr().String(), "failed")
			continue
		}
		logger.Info("client conn server [", server.listener.Addr().String(), "]")
		
		go client.client_goroutine()
	}
}
