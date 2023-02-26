package grpc

import (
	"context"
	"geo-search/internal/entities"
	"geo-search/internal/interactors"
	"geo-search/internal/transports/grpc/mapper"
	"geo-search/internal/transports/grpc/pb"

	"google.golang.org/grpc"
)

type EstateController struct {
	pb.UnimplementedEstateServiceServer

	estateInteractor *interactors.EstateInteractor
}

var _ BaseController = &EstateController{}
var _ pb.EstateServiceServer = &EstateController{}

func WireEstateController(estateInteractor *interactors.EstateInteractor) *EstateController {
	return &EstateController{
		estateInteractor: estateInteractor,
	}
}

func (ctrl *EstateController) Register(s grpc.ServiceRegistrar) {
	pb.RegisterEstateServiceServer(s, ctrl)
}

func (ctrl *EstateController) CreateEstate(ctx context.Context, request *pb.CreateEstateRequest) (*pb.CreateEstateResponse, error) {
	estate, err := ctrl.estateInteractor.Create(ctx, &entities.Estate{
		Title: request.Title,
		Location: entities.Location{
			Lat: request.Location.Lat,
			Lon: request.Location.Lon,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateEstateResponse{
		Estate: mapper.EstateToGrpcEstate(estate),
	}, nil
}

func (ctrl *EstateController) GetEstateByBound(ctx context.Context, request *pb.GetEstatesByBoundBoxRequest) (*pb.GetEstatesByBoundBoxResponse, error) {
	estates, hits, err := ctrl.estateInteractor.FindByBoundBox(
		ctx,
		&entities.Location{
			Lat: request.TopLeft.Lat, Lon: request.TopLeft.Lon,
		},
		&entities.Location{
			Lat: request.BottomRight.Lat, Lon: request.BottomRight.Lon,
		},
	)
	if err != nil {
		return nil, err
	}

	return &pb.GetEstatesByBoundBoxResponse{
		Estates: mapper.EstateToGrpcEstateList(estates),
		Hits:    int32(hits),
	}, nil
}
