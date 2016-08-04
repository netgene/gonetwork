package main

import (
	//"fmt"
	logger "github.com/donnie4w/go-logger/logger"
    "time"
    "os"
    "os/exec"
    "path/filepath"
    "os/signal"
    "syscall"

    "business"
    "network"
)

var exit bool	

type etrconf struct {
	serverConf      network.ServerConf
	otcConf         network.ConnConf
	marketConf      network.ConnConf
}

func InitConfig() {
	logger.Info("InitConfig")
}

func daemonlize() (exit bool) {
	if os.Getppid() != 1 {
		filePath, _ := filepath.Abs(os.Args[0])
		cmd := exec.Command(filePath, os.Args[1:]...)
		cmd.Start()
		return true
	}
	return false
}

func main() {
	logger.SetRollingFile("log", "entrust.log", 10, 5, logger.KB)
	logger.SetRollingDaily("log", "entrust.log")
	logger.SetLevel(logger.INFO)

	//damemon
	//if daemonlize() {
	//	return
	//}
	logger.Info("program start.PID[", os.Getpid(), "]")

	//signals
	sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        sig := <- sigs;
        logger.Info("catch sig:", sig)
        exit = true
    }()

    //business
    InitConfig()
    business.InitDB()
	business.InitNetwork()

	//mygod! think goroutine close.
	//if main goroutine exit, all goroutines will exit.
	for !exit {
		time.Sleep(1 * time.Second)
	}

	logger.Info("program exit.")
	defer business.CloseDB()
	//defer logger.Flush()
}
