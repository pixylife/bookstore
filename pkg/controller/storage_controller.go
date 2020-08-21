package controller

import (
	"context"
            event "bookstore/pkg/eventing"
)

type BuildProjectionFunc func(ctx context.Context, events []event.Event) (event.Projection, error)

type StorageProvider interface {
	SaveEvent(ctx context.Context, event event.Event) error
	SaveProjection(ctx context.Context, projection event.Projection) error
	GetProjection(ctx context.Context, entityID string, projection event.Projection) (event.Projection, error)
	GetLatestEventIDForEntityID(ctx context.Context, entityID string) (string, error)
	GetSortedEventsForEntityID(ctx context.Context, entityID string) ([]event.Event, error)
}

// Controllers are the primary way to save events and build projections based upon these.
type StorageCtrl struct {
	provider        StorageProvider
	projection      event.Projection
	buildProjection BuildProjectionFunc
}

// NewController creates a new instance of an controller by collection all necessary data.
func NewController(projection event.Projection, buildProjection BuildProjectionFunc, provider StorageProvider) *StorageCtrl {
	return &StorageCtrl{
		provider:        provider,
		projection:      projection,
		buildProjection: buildProjection,
	}
}

// GetLatestProjection retrieves the latest projection the managed projection-entity.
func (controller *StorageCtrl) GetLatestProjection(ctx context.Context, entityID string) (event.Projection, error) {
	latestProjection, err := controller.provider.GetProjection(ctx, entityID, controller.projection)
	if err != nil {
		return nil, err
	}
	return latestProjection, nil
}

// SaveEvent stores an event and builds an up-to-date projection.
func (controller *StorageCtrl) SaveEvent(ctx context.Context, event event.Event) error {
	if err := controller.provider.SaveEvent(ctx, event); err != nil {
		return err
	}
	return nil
}

func (controller *StorageCtrl) ProjectEvent(ctx context.Context, event event.Event) (event.Projection,error) {

	events, err := controller.provider.GetSortedEventsForEntityID(ctx, event.GetEntityID())
	if err != nil {
		return nil,err
	}

	return controller.buildProjection(ctx, events)
}

func (controller *StorageCtrl) ProjectSave(ctx context.Context,proj event.Projection) error {

	return controller.provider.SaveProjection(ctx, proj)
}