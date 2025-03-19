package main

import (
	"fmt"
	"github.com/kardianos/service"
	"os"
	"time"
)

type SystemService struct {
	exit chan struct{}
}

func NewSystemService() service.Interface {
	return &SystemService{exit: make(chan struct{}, 1)}
}

func (ss *SystemService) Start(s service.Service) error {
	fmt.Println("coming Start.......")
	go ss.run()
	return nil
}

func (ss *SystemService) run() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Printf("%s: running\n", time.Now().Format(time.DateTime))
		case <-ss.exit:
			return
		}
	}
}

func (ss *SystemService) Stop(s service.Service) error {
	fmt.Println("coming Stop.......")
	ss.exit <- struct{}{}
	return nil
}

// go build -o main main.go
// ./main install
// ./main start
// ./main restart
// ./main stop
// ./main uninstall
// systemctl status custom-service
// systemctl restart custom-service
// systemctl start custom-service
// systemctl stop custom-service
func main() {
	svcConfig := &service.Config{
		Name:        "custom-service",
		DisplayName: "custom service",
		Description: "this is github.com/kardianos/service test case",
	}

	s, err := service.New(NewSystemService(), svcConfig)
	if err != nil {
		fmt.Printf("service New failed, err: %v\n", err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		if err = service.Control(s, os.Args[1]); err != nil {
			fmt.Printf("service Control failed, err: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// 默认 运行 Run
	if err = s.Run(); err != nil {
		fmt.Printf("service Control 222 failed, err: %v\n", err)
		os.Exit(1)
	}
}
