package mapper

import (
	"geo-search/internal/entities"
	"geo-search/internal/transports/grpc/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func EstateToGrpcEstateList(estates []entities.Estate) []*pb.Estate {
	res := make([]*pb.Estate, 0, len(estates))
	for _, estate := range estates {
		res = append(res, EstateToGrpcEstate(&estate))
	}

	return res
}

func EstateToGrpcEstate(estate *entities.Estate) *pb.Estate {
	res := &pb.Estate{
		Title: estate.Title,
		Location: &pb.Location{
			Lat: estate.Location.Lat,
			Lon: estate.Location.Lon,
		},
		CreatedAt: timestamppb.New(estate.CreatedAt),
		UpdatedAt: timestamppb.New(estate.UpdatedAt),
	}

	if estate.DeletedAt != nil {
		res.DeletedAt = timestamppb.New(*estate.DeletedAt)
	}

	return res
}
