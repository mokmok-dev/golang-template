package marshaller

import (
	"fmt"

	domainmodel "github.com/mokmok-dev/golang-template/domain/model"
	recordmodel "github.com/mokmok-dev/golang-template/infra/postgres/model"
)

func RecordToUser(record *recordmodel.User) (*domainmodel.User, error) {
	id, err := domainmodel.ParseUserID(record.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert record to user: %w", err)
	}

	return &domainmodel.User{
		ID:        id,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}, nil
}

func UserToRecord(model *domainmodel.User) *recordmodel.User {
	return &recordmodel.User{
		ID:        model.ID.String(),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
