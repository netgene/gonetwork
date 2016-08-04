package business

import (
	logger "github.com/donnie4w/go-logger/logger"
	"protobuf"
	"protobuf/protomsg"
	//"network"
)

//处理客户端发来的心跳
func dealClientHeartBeatReq(clientPackage *ClientPackage) {
	//解包
	HeartReq := &protomsg.HeartReq{}
	err := protobuf.Nunpack(clientPackage.buf, HeartReq)
	if err != nil {
		logger.Error("Unpack HeartBeatReq failed:", err)
		return
	}
	//打包
	HeartRes := &protomsg.HeartRes{SID: HeartReq.SID}
	sendBuf, err := protobuf.Npack(HeartRes, CodeHeartBeatRes)
	if err != nil {
		logger.Error("pack HeartBeatRes failed:", err)
		return
	}
	//发送
	err = clientPackage.client.Write(sendBuf)
	if err != nil {
		logger.Error("Write HeartBeatRes failed:", err)
		return
	}
	logger.Info("recv HeartBeatReq[", HeartReq.GetSID(), "]")
	return
}
