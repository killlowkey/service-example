package main

import (
	"fmt"
	"github.com/kardianos/service"
	"os"
)

type SystemService struct{}

func (ss *SystemService) Start(s service.Service) error {
	fmt.Println("coming Start.......")
	go ss.run()
	return nil
}

func (ss *SystemService) run() {
	fmt.Println("coming run.......")
}

func (ss *SystemService) Stop(s service.Service) error {
	fmt.Println("coming Stop.......")
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
	fmt.Println("service.Interactive()---->", service.Interactive())
	svcConfig := &service.Config{
		Name:        "custom-service",
		DisplayName: "custom service",
		Description: "this is github.com/kardianos/service test case",
	}

	ss := &SystemService{}
	s, err := service.New(ss, svcConfig)
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
