package logic

import (
	"context"

	"alex.tse/dao"
	pb "alex.tse/proto/service"
)

func ListRaceTypes(ctx context.Context, req *pb.ListRaceTypesReq) (*pb.ListRaceTypesRsp, error) {
	raceTypes, err := dao.ListRaceTypes()
	if err != nil {
		return nil, err
	}
	pRaceTypes := make([]*pb.RaceType, len(raceTypes))
	for i, rt := range raceTypes {
		pRaceTypes[i] = &pb.RaceType{
			Id:   int32(rt.ID),
			Name: rt.Name,
		}
	}
	return &pb.ListRaceTypesRsp{
		Data: pRaceTypes,
	}, nil
}
