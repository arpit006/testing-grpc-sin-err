package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"source.golabs.io/scp/goto-profile-protos/gen/models"
	"source.golabs.io/scp/goto-profile-protos/gen/services"
	"sync"
	"time"
)

func main() {
	fmt.Println("PINGING Singapore service from taiwan......")
	//_ = os.Setenv("GRPC_VERBOSITY", "DEBUG")
	//_ = os.Setenv("GRPC_TRACE", "all")

	// Set GRPC_GO_LOG_VERBOSITY environment variable to 'debug'
	_ = os.Setenv("GRPC_GO_LOG_VERBOSITY", "debug")

	// Set GRPC_GO_LOG_SEVERITY_LEVEL environment variable to 'debug'
	_ = os.Setenv("GRPC_GO_LOG_SEVERITY_LEVEL", "debug")

	// Initialize the custom logger with debug verbosity
	//grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr))

	startTime := time.Now()
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(5)

	//conn, err := grpc.Dial("10.225.139.211:80", grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("[error] could not obtain gRPC connection. err: [%s]", err)
	//}
	//defer conn.Close()

	go initConnectionAndCall(getConn(), "FIRST", waitGroup)
	go initConnectionAndCall(getConn(), "SECOND", waitGroup)
	go initConnectionAndCall(getConn(),"THIRD", waitGroup)
	go initConnectionAndCall(getConn(), "FOURTH", waitGroup)
	go initConnectionAndCall(getConn(), "FIFTH", waitGroup)

	waitGroup.Wait()

	log.Fatalf("[successful] All batch calls completed succesffully from Profile service in: [%v milliSeconds]", time.Since(startTime).Milliseconds())
}

func getConn() *grpc.ClientConn {
	conn, err := grpc.Dial("10.225.96.9:80", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("[error] could not obtain gRPC connection. err: [%s]", err)
	}
	//defer conn.Close()
	return conn
}

func initConnectionAndCall(conn *grpc.ClientConn, leader string, waitGroup *sync.WaitGroup) {


	pingInLoop(conn, leader)
	waitGroup.Done()
}

func pingInLoop(conn *grpc.ClientConn, leader string) {
	for i := 1; i <= 30; i++ {
		ctx := context.Background()
		client := services.NewProfilePingServiceClient(conn)
		startTime := time.Now()
		resp, err := client.Ping(ctx, &models.NoParam{})
		if err != nil {
			log.Printf("[%s] [error] from ping singapore PS. err: [%s]", leader, err)
		}
		log.Printf("[%s] [response] timeTaken: [%v milliSeconds] from Singapore ProfileService: [%+v]", leader, time.Since(startTime).Milliseconds(), resp)
	}
}
