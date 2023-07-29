package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"os"
	"source.golabs.io/scp/goto-profile-protos/gen/services"
	"time"
)

func main() {
	fmt.Println("PINGING Singapore service from taiwan......")
	_ = os.Setenv("GRPC_VERBOSITY", "DEBUG")
	_ = os.Setenv("GRPC_TRACE", "all")
	conn, err := grpc.Dial("10.225.96.9:80", grpc.WithInsecure())
	if err != nil {
		fmt.Println("[error] could not obtain gRPC connection")
		return
	}
	defer conn.Close()

	ctx := context.Background()
	client := services.NewProfilePingServiceClient(conn)
	if err != nil {
		fmt.Println("[error] could not get profileService client")
		return
	}
	startTime := time.Now()
	resp, err := client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("[error] from ping singapore PS")
		return
	}
	fmt.Printf("[response] timeTaken: [%v milliSeconds] from Singapore ProfileService: [%+v]", time.Since(startTime).Milliseconds(), resp)
}
