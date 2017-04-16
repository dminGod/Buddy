package main

import (
	"google.golang.org/grpc"
	"github.com/dminGod/HowRU/proto"
	"fmt"
	"net"
	"golang.org/x/net/context"
	"os/exec"
)

// GRPCServer registers
type HeyBuddyServer struct{}

// Do is used to perform simple RPC communication to store and query data from the endpoints
func (g *HeyBuddyServer) ExecBuddy(ctx context.Context, req *remote.RemoteRequest) (*remote.RemoteResponse, error) {

	resp := new(remote.RemoteResponse)


	cmdd := exec.Command("bash", "-c", req.RequestBody)

//	cmdd := exec.Command("cmd", "/C", "dir")

	fmt.Println("This is cmdd", cmdd)

	cmdOutput, _ := cmdd.Output()

	fmt.Println( "Response of DIR", string( cmdOutput ) )
	fmt.Println("got this request", req)


	resp.ResponseBody = string( cmdOutput )

	return resp, nil
}


func main(){

	fmt.Println("Starting Server")

	listener, _ := net.Listen("tcp", "127.0.0.1:9191")

	grpcServer := grpc.NewServer()
	remote.RegisterHeyBuddyServer(grpcServer, new(HeyBuddyServer))
	grpcServer.Serve(listener)
}




