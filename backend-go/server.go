package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 3080
// Hash 9811
// Hash 1931
// Hash 6323
// Hash 2637
// Hash 3073
// Hash 9162
// Hash 4612
// Hash 2492
// Hash 5698
// Hash 1420
// Hash 6554
// Hash 3219
// Hash 2073
// Hash 7920
// Hash 4181
// Hash 7702
// Hash 2897
// Hash 8275
// Hash 8925
// Hash 4486
// Hash 9307
// Hash 5518
// Hash 7187
// Hash 7770
// Hash 2057