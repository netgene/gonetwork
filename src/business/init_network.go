package business

import (
	logger "github.com/donnie4w/go-logger/logger"
	"network"
)

type DNetWork struct {
	Otc       network.Conn     //connect to otc
	Market    network.Conn     //connect to market
	Server    network.Server  //listen
}

/*
func InitNetwork(OtcConf network.ConnConf, MarketConf network.ConnConf, 
	ServiceConf network.ConnConf) (err error) {

	logger.info("InitNetWork")
	
	var network NetWork
	network.Otc.ConnConf = OtcConf
	network.Market.ConnConf = MarketConf
	network.Server.ServerConf = ServerConf

	for {
		//if err = network.Otc.Connect() ; err != nil { break }
		//if err = network.Market.Connect() ; err != nil { break }
		if err = network.Server.Listen() ; err != nil { break }
	}

	err == nil ? return : return err
}
*/

func InitNetwork() (err error) {

	logger.Info("InitNetWork")
	
	var n DNetWork
	var serviceConf network.ServerConf
	serviceConf.Addr = "127.0.0.1:9956"
	serviceConf.DealRecv = DealClientRecv
	n.Server.ServerConf = serviceConf

	err = n.Server.Listen()
	return	
}