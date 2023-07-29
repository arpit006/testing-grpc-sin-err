package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"os"
	"source.golabs.io/scp/goto-profile-protos/gen/models"
	"source.golabs.io/scp/goto-profile-protos/gen/services"
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
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr))

	conn, err := grpc.Dial("10.225.96.9:80", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("[error] could not obtain gRPC connection. err: [%s]", err)
	}
	defer conn.Close()

	ctx := context.Background()
	client := services.NewProfilePingServiceClient(conn)
	if err != nil {
		log.Fatalf("[error] could not get profileService client. err: [%s]", err)
		return
	}
	startTime := time.Now()
	resp, err := client.Ping(ctx, &models.NoParam{})
	if err != nil {
		log.Fatalf("[error] from ping singapore PS. err: [%s]", err)
		return
	}
	log.Fatalf("[response] timeTaken: [%v milliSeconds] from Singapore ProfileService: [%+v]", time.Since(startTime).Milliseconds(), resp)
}
