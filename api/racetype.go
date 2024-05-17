package api

import (
	"context"

	"alex.tse/logic"
	pb "alex.tse/proto/service"
)

type Server struct {
	pb.UnimplementedHupuServiceServer
}

func (s *Server) ListRaceTypes(ctx context.Context, req *pb.ListRaceTypesReq) (*pb.ListRaceTypesRsp, error) {
	return logic.ListRaceTypes(ctx, req)
}
