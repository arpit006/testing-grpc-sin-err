package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"os"
	"source.golabs.io/scp/goto-profile-protos/gen/models"
	"source.golabs.io/scp/goto-profile-protos/gen/services"
	"time"
)

func main() {
	fmt.Println("PINGING Singapore service from taiwan......")
	_ = os.Setenv("GRPC_VERBOSITY", "DEBUG")
	_ = os.Setenv("GRPC_TRACE", "all")
	conn, err := grpc.Dial("10.225.96.9:80", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("[error] could not obtain gRPC connection. err: [%s]\n", err)
		return
	}
	defer conn.Close()

	ctx := context.Background()
	client := services.NewProfilePingServiceClient(conn)
	if err != nil {
		fmt.Printf("[error] could not get profileService client. err: [%s]\n", err)
		return
	}
	startTime := time.Now()
	resp, err := client.Ping(ctx, &models.NoParam{})
	if err != nil {
		fmt.Printf("[error] from ping singapore PS. err: [%s]\n", err)
		return
	}
	fmt.Printf("[response] timeTaken: [%v milliSeconds] from Singapore ProfileService: [%+v]", time.Since(startTime).Milliseconds(), resp)
}
