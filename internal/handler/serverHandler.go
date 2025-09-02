package handler

import pb "gapi/proto/gen"

type Server struct {
	pb.UnimplementedTeachersServiceServer
	pb.UnimplementedStudentsServiceServer
	pb.UnimplementedExecsServiceServer
}
