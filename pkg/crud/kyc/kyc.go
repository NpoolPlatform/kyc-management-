package kyc

import (
	"context"

	"github.com/NpoolPlatform/kyc-management/message/npool"
	"github.com/NpoolPlatform/kyc-management/pkg/db"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent/kyc"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

const (
	WaitAudit = "auditing"
	PassAudit = "audited"
	FailAudit = "failed"
)

func DBRowToKyc(row *ent.Kyc) *npool.KycInfo {
	return &npool.KycInfo{
		ID:                  row.ID.String(),
		UserID:              row.UserID.String(),
		FirstName:           row.FirstName,
		LastName:            row.LastName,
		Region:              row.Region,
		CardType:            row.CardType,
		CardID:              row.CardID,
		FrontCardImg:        row.FrontCardImg,
		BackCardImg:         row.BackCardImg,
		UserHandlingCardImg: row.UserHandlingCardImg,
		ReviewStatus:        row.ReviewStatus,
		CreateAT:            row.CreateAt,
		UpdateAT:            row.UpdateAt,
	}
}

func Create(ctx context.Context, in *npool.CreateKycRecordRequest) (*npool.CreateKycRecordResponse, error) {
	userID, err := uuid.Parse(in.Info.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	appID, err := uuid.Parse(in.Info.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	_, err = cli.
		Kyc.
		Query().
		Where(
			kyc.And(
				kyc.AppID(appID),
				kyc.UserID(userID),
			),
		).Only(ctx)
	if err == nil {
		return nil, xerrors.Errorf("user kyc record has been existed: %v", err)
	}

	info, err := cli.
		Kyc.
		Create().
		SetUserID(userID).
		SetAppID(appID).
		SetFirstName(in.Info.FirstName).
		SetLastName(in.Info.LastName).
		SetRegion(in.Info.Region).
		SetCardType(in.Info.CardType).
		SetCardID(in.Info.CardID).
		SetFrontCardImg(in.Info.FrontCardImg).
		SetBackCardImg(in.Info.BackCardImg).
		SetUserHandlingCardImg(in.Info.UserHandlingCardImg).
		SetReviewStatus(WaitAudit).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to create user kyc: %v", err)
	}

	return &npool.CreateKycRecordResponse{
		Info: DBRowToKyc(info),
	}, nil
}

func parse2ID(userIDString, idString string) (uuid.UUID, uuid.UUID, error) { // nolint
	var userID, id uuid.UUID
	var err error
	if userIDString != "" {
		userID, err = uuid.Parse(userIDString)
		if err != nil {
			return uuid.UUID{}, uuid.UUID{}, xerrors.Errorf("invalid user id: %v", err)
		}
	}
	if idString != "" {
		id, err = uuid.Parse(idString)
		if err != nil {
			return uuid.UUID{}, uuid.UUID{}, xerrors.Errorf("invalid kyc id: %v", err)
		}
	}
	return userID, id, nil
}

func GetKycByUserIDAndAppID(ctx context.Context, appID, userID uuid.UUID) (*ent.Kyc, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	return cli.Kyc.Query().
		Where(
			kyc.And(
				kyc.AppID(appID),
				kyc.UserID(userID),
			),
		).Only(ctx)
}

func GetKycByID(ctx context.Context, kycID uuid.UUID) (*ent.Kyc, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	return cli.Kyc.Query().
		Where(
			kyc.ID(kycID),
		).Only(ctx)
}

func GetAll(ctx context.Context, in *npool.GetAllKycInfosRequest) (*npool.GetAllKycInfosResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	kycIDs := []uuid.UUID{}
	for _, id := range in.KycIDs {
		id, err := uuid.Parse(id)
		if err != nil {
			return nil, xerrors.Errorf("%v user kyc id invalid: %v", id, err)
		}
		kycIDs = append(kycIDs, id)
	}
	response := []*npool.KycInfo{}
	for _, kycid := range kycIDs {
		info, err := cli.
			Kyc.
			Query().
			Where(
				kyc.ID(kycid),
			).Only(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail to get %v kyc info: %v", kycid, err)
		}

		response = append(response, DBRowToKyc(info))
	}

	return &npool.GetAllKycInfosResponse{
		Infos: response,
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateKycRequest) (*npool.UpdateKycResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	userID, err := uuid.Parse(in.Info.UserID)
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	appID, err := uuid.Parse(in.Info.AppID)
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}
	id := ""
	if in.Info.ID == "" {
		info, err := cli.
			Kyc.
			Query().
			Where(
				kyc.And(
					kyc.UserID(userID),
					kyc.AppID(appID),
				),
			).Only(ctx)
		if err != nil {
			return nil, xerrors.Errorf("fail to query user kyc info: %v", err)
		}

		id = info.ID.String()
	} else {
		id = in.Info.ID
	}

	kycID, err := uuid.Parse(id)
	if err != nil {
		return nil, xerrors.Errorf("invalid kyc record id: %v", err)
	}

	info, err := cli.
		Kyc.
		UpdateOneID(kycID).
		SetUserID(userID).
		SetFirstName(in.Info.FirstName).
		SetLastName(in.Info.LastName).
		SetRegion(in.Info.Region).
		SetCardType(in.Info.CardType).
		SetCardID(in.Info.CardID).
		SetFrontCardImg(in.Info.FrontCardImg).
		SetBackCardImg(in.Info.BackCardImg).
		SetUserHandlingCardImg(in.Info.UserHandlingCardImg).
		SetReviewStatus(WaitAudit).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail to update user kyc: %v", err)
	}

	return &npool.UpdateKycResponse{
		Info: DBRowToKyc(info),
	}, nil
}

func UpdateReviewStatus(ctx context.Context, kycID uuid.UUID, status int32) (*ent.Kyc, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	var statusString string
	switch status {
	case 1:
		statusString = PassAudit
	case 2:
		statusString = FailAudit
	default:
		statusString = WaitAudit
	}

	resp, err := cli.
		Kyc.
		UpdateOneID(kycID).
		SetReviewStatus(statusString).
		Save(ctx)

	return resp, err
}
