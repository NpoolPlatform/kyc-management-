package api

import (
	"context"
	"unicode/utf8"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/kyc-management/message/npool"
	"github.com/NpoolPlatform/kyc-management/pkg/crud/kyc"
	mygrpc "github.com/NpoolPlatform/kyc-management/pkg/grpc"
	myconst "github.com/NpoolPlatform/kyc-management/pkg/message/const"
	mkyc "github.com/NpoolPlatform/kyc-management/pkg/middleware/kyc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateKyc(ctx context.Context, in *npool.CreateKycRequest) (*npool.CreateKycResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		logger.Sugar().Errorf("CreateKyc error: app id is invalid: %v", err)
		return nil, status.Error(codes.InvalidArgument, "app id is invalid")
	}

	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		logger.Sugar().Errorf("CreateKyc error: user id is invalid: %v", err)
		return nil, status.Error(codes.InvalidArgument, "user id is invalid")
	}

	if in.GetCardID() == "" {
		logger.Sugar().Error("CreateKyc error: user card id can not be empty")
		return nil, status.Error(codes.InvalidArgument, "user card id can not be empty")
	}

	if len(in.GetCardID()) > 30 {
		logger.Sugar().Error("CreateKyc error: user card id is not invalid")
		return nil, status.Error(codes.InvalidArgument, "user card id is not invalid")
	}

	if in.GetCardType() == "" {
		logger.Sugar().Error("CreateKyc error: user card type can not be empty")
		return nil, status.Error(codes.InvalidArgument, "user card type can not be empty")
	}

	if in.GetFirstName() == "" {
		logger.Sugar().Error("CreateKyc error: user first name can not be empty")
		return nil, status.Error(codes.InvalidArgument, "user first name can not be empty")
	}

	if utf8.RuneCountInString(in.GetFirstName()) > 50 {
		logger.Sugar().Error("CreateKyc error: user first name is not invalid")
		return nil, status.Error(codes.InvalidArgument, "user first name is not invalid")
	}

	if in.GetLastName() == "" {
		logger.Sugar().Error("CreateKyc error: user last name can not be empty")
		return nil, status.Error(codes.InvalidArgument, "user last name can not be empty")
	}

	if utf8.RuneCountInString(in.GetLastName()) > 50 {
		logger.Sugar().Error("CreateKyc error: user last name is not invalid")
		return nil, status.Error(codes.InvalidArgument, "user last name is not invalid")
	}

	if in.GetRegion() == "" {
		logger.Sugar().Error("CreateKyc error: user region can not be empty")
		return nil, status.Error(codes.InvalidArgument, "user region can not be empty")
	}

	if in.GetFrontCardImg() == "" {
		logger.Sugar().Error("CreateKyc error: user front card image can not be empty")
		return nil, status.Error(codes.InvalidArgument, "user front card image can not be empty")
	}

	if in.GetUserHandlingCardImg() == "" {
		logger.Sugar().Error("CreateKyc error: user handling card image can not be empty")
		return nil, status.Error(codes.InvalidArgument, "user front card image can not be empty")
	}

	if in.GetBackCardImg() == "" {
		logger.Sugar().Error("CreateKyc error: user back card image can not be empty")
		return nil, status.Error(codes.InvalidArgument, "user back card image can not be empty")
	}

	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	if exist, err := kyc.ExistCradTypeCardIDInApp(ctx, in.GetCardType(), in.GetCardID(), appID); err != nil {
		logger.Sugar().Errorf("CreateKyc error: internal server error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	} else if exist {
		logger.Sugar().Error("CreayeKyc error: this card type card id has been existed in this app")
		return nil, status.Error(codes.AlreadyExists, "this card type card id has been existed in this app")
	}

	if exist, err := kyc.ExistKycByUserIDAndAppID(ctx, appID, userID); err != nil {
		logger.Sugar().Errorf("CreateKyc error: internal server error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	} else if exist {
		logger.Sugar().Error("CreateKyc error: user has already created a kyc record")
		return nil, status.Error(codes.AlreadyExists, "user has already created a kyc record")
	}

	resp, err := kyc.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("CreateKyc error: internal server error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	_, err = mygrpc.CreateKycReview(ctx, resp.GetInfo().GetID(), resp.GetInfo().GetAppID())
	if err != nil {
		logger.Sugar().Errorf("CreateKyc call CreateReview error: %v", err)

		err := kyc.DeleteUserKycRecordByKycID(ctx, uuid.MustParse(resp.GetInfo().GetID()))
		if err != nil {
			logger.Sugar().Errorf("CreateKyc call CreateReview and CreateReview call DeleteUserKycRecordByKycID error: %v", err)
			return nil, status.Errorf(codes.Internal, "internal server error")
		}
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return resp, nil
}

func (s *Server) GetAllKycInfos(ctx context.Context, in *npool.GetAllKycInfosRequest) (*npool.GetAllKycInfosResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, err := mkyc.GetKycInfo(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to get kyc record: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) UpdateKyc(ctx context.Context, in *npool.UpdateKycRequest) (*npool.UpdateKycResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, err := kyc.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to update kyc record: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return resp, nil
}
