package provider

import (
        "bookstore/pkg/eventing"
        "bookstore/pkg/model"
	"context"
	"github.com/jinzhu/gorm"
)



type GormStorageProvider struct {
	db *gorm.DB
}

func NewGormStorageProvider(db *gorm.DB) *GormStorageProvider {
	return &GormStorageProvider{
		db: db,
	}
}

func (provider *GormStorageProvider) SaveEvent(ctx context.Context, event eventing.Event) error {

	return nil
}

func (provider *GormStorageProvider) SaveProjection(ctx context.Context, projection eventing.Projection) error {

	projecData := projection.(*model.Projection)
	err := provider.db.Save(&projecData.Data).Error
	if err != nil {
		return err
	}
	return nil
}

func (provider *GormStorageProvider) GetProjection(ctx context.Context, entityID string, projection eventing.Projection) (eventing.Projection, error) {

	return nil, nil
}

func (provider *GormStorageProvider) GetLatestEventIDForEntityID(ctx context.Context, entityID string) (string, error) {

	return "", nil
}

func (provider *GormStorageProvider) GetSortedEventsForEntityID(ctx context.Context, entityID string) ([]eventing.Event, error) {

	return nil, nil
}
