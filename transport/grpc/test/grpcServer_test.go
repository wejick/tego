package test

import (
	"context"
	"net"
	"testing"

	"github.com/wejick/tego/endpoint"
	"github.com/wejick/tego/transport/grpc"
	"github.com/wejick/tego/transport/grpc/test/pb"

	goGRPC "google.golang.org/grpc"
)

func makeTestHaloEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if request.(string) == "" {
			return "halo", nil
		}
		return request, nil
	}
}

func requestDecoder(ctx context.Context, req interface{}) (request interface{}, err error) {
	var r *pb.HaloRequest
	if req != nil {
		r = req.(*pb.HaloRequest)
	}
	if r.Message != "" {
		request = r.Message
	} else {
		request = ""
	}
	return
}

func responseEncoder(ctx context.Context, response interface{}) (grpcResponse interface{}, err error) {
	grpcResponse = pb.HaloResponse{
		Message: response.(string),
	}
	return
}

// here we implement the binding from proto interface with tego server
type serverBinder struct {
	halo grpc.Handler
}

func (s *serverBinder) Halo(ctx context.Context, req *pb.HaloRequest) (*pb.HaloResponse, error) {
	_, resp, err := s.halo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	haloResponse := resp.(pb.HaloResponse)
	return &haloResponse, nil
}

func TestServer(t *testing.T) {
	//spawning the server
	haloService := grpc.New(makeTestHaloEndpoint(), requestDecoder, responseEncoder)
	haloBinding := &serverBinder{
		halo: haloService,
	}
	server := goGRPC.NewServer()

	sc, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		t.Fatalf("unable to listen: %+v", err)
	}
	defer server.GracefulStop()

	go func() {
		pb.RegisterHaloServer(server, haloBinding)
		_ = server.Serve(sc)
	}()

	//creating client
	cc, err := goGRPC.Dial("localhost:8000", goGRPC.WithInsecure())
	if err != nil {
		t.Fatalf("unable to Dial: %+v", err)
	}
	client := pb.NewHaloClient(cc)

	//testing client
	resp, err := client.Halo(context.Background(), &pb.HaloRequest{Message: ""})
	if err != nil {
		t.Fatalf("unable to call: %+v", err)
	}
	if resp.Message != "halo" {
		t.Error("expecting resp.Message == halo, got", resp.Message)
	}
	resp, err = client.Halo(context.Background(), &pb.HaloRequest{Message: "hai"})
	if err != nil {
		t.Fatalf("unable to call: %+v", err)
	}
	if resp.Message != "hai" {
		t.Error("expecting resp.Message == hai, got", resp.Message)
	}
}
