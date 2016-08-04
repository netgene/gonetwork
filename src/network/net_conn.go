package network

import (
	//logger "github.com/donnie4w/go-logger/logger"
	"net"
)

const (
	SERVER_NETWORK    = "tcp"      //链接协议类型
	CONN_TIMEOUT_MILL = 5000       //链接超时时间，单位为ms
	RECV_BUF_SIZE     = 1024 * 100 //网络读取数据buf大小
	CHAN_BUF_COUNT    = 10         //将收到的数据发送给分包携程的切片数
	RECONN_TIME       = 5          //连接服务器断开时，重连间隔，单位为秒
	READ_TIME         = 500        //Read最长时间间隔，单位为ms
)

type ConnConf struct {
	addr               string  			//host:port
	addrBak            string  			//host:port
	connType           int    			//client 1, service 2
	HeartBeatTimeOut   int64
}

type Conn struct {
	//link info
	conconf            ConnConf
	conn               net.Conn
	bStatus            bool
	lastRecvTime       int64
	lastSendTime       int64
	lastHeartBeatTime  int64
	bHeartBeatTimeOut  bool
	recvChan           chan []byte
	sendChan           chan []byte
	//link action
	//RecvMsg            ConnRecvMsg
}

func (conn *Conn) clientConn_goroutine() {
}