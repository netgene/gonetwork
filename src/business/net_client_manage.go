package business

import (
	logger "github.com/donnie4w/go-logger/logger"
	"network"
	"protobuf"
	"strconv"
)

const (
	MAX_DEAL_CLIENT_PACKAGE = 10
	MAX_HEART_INTERVAL      = 5 //重发心跳间隔
)

type ClientPackage struct {
	head   *protobuf.CmdHeader
	buf    []byte
	client *network.Client
}

//实际处理相应的数据包
func dealClientPackage(DealChan chan ClientPackage) {
	defer func() {
		logger.Info("dealClientPackage goroutine out")
	}()
	for {
		clientPackage, ok := <-DealChan
		if ok == false {
			logger.Info("get Package from DealChan failed")
			break
		}
		switch clientPackage.head.Command {
		case CodeHeartBeatReq:
			dealClientHeartBeatReq(&clientPackage)
		default:
			logger.Error("Recv undefined Command:", strconv.FormatInt(int64(clientPackage.head.Command), 16), " from client")
		}
	}
}

func DealClientRecv(recvInfoChan chan []byte, client *network.Client) {
	DealChan := make(chan ClientPackage, MAX_DEAL_CONN_PACKAGE)
	defer close(DealChan)

	go dealClientPackage(DealChan)
	var buf []byte
	for {
		logger.Info("get Package from recvInfoChan start...")
		recvbuf, ok := <-recvInfoChan
		buf = append(buf, recvbuf...)
		if ok == false {
			logger.Error("get Package from recvInfoChan failed")
			break
		}
		logger.Info("get Package from recvInfoChan ok")

		for {
			bodylen, err := protobuf.GetBodyLen(buf)
			if err != nil {
				logger.Error("getBodyLen failed:", err)
				break
			}
			if len(buf) < int(bodylen) {
				break
			}
			if bodylen < protobuf.HEAD_BYTE_LEN { //包长度小于包头时，说明数据已经乱了，继续执行将导致panice
				//client.KillClient()
				return
			}
			
			var clientPackage ClientPackage
			clientPackage.head, err = protobuf.NewHeader(buf)
			if err != nil {
				logger.Error("NewHeader failed:", err)
				break
			}
			clientPackage.buf = buf[protobuf.HEAD_BYTE_LEN:bodylen]
			clientPackage.client = client

			DealChan <- clientPackage
			buf = buf[bodylen:]
			logger.Info("push Package to DealChan ok")
		}
	}
}
