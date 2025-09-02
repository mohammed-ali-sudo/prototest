package main

import (
	"fmt"
	"net"

	"gapi/internal/handler"
	pb "gapi/proto/gen"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// --- Init mgm (ODM) ---
	err := mgm.SetDefaultConfig(nil, "my_database",
		options.Client().ApplyURI("mongodb+srv://m56oo:1010204080@cluster0.mciy55k.mongodb.net/?retryWrites=true&w=majority"),
	)
	if err != nil {
		_ = handler.ErrorHandler(err, "❌ Failed to init mgm")
		return
	}
	fmt.Println("✅ Connected to MongoDB with mgm")

	// --- Setup gRPC server ---
	grpcServer := grpc.NewServer()

	// Register services (using same handler instance for all)
	svc := &handler.Server{}
	pb.RegisterExecsServiceServer(grpcServer, svc)
	pb.RegisterStudentsServiceServer(grpcServer, svc)
	pb.RegisterTeachersServiceServer(grpcServer, svc)

	// Enable reflection (debugging with evans / grpcurl)
	reflection.Register(grpcServer)

	// --- Start listening ---
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		_ = handler.ErrorHandler(err, "❌ Failed to listen on :8000")
		return
	}
	defer listener.Close()

	fmt.Println("🚀 gRPC server listening on :8000")

	if err := grpcServer.Serve(listener); err != nil {
		_ = handler.ErrorHandler(err, "❌ gRPC server failed")
	}
}
