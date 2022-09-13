package entitycache

import (
	"context"

	entitycachev1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/common/entitycache/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// EntityCacheService is the implementation of entitycachev1.EntityCacheService interface.
type EntityCacheService struct {
	entitycachev1.UnimplementedEntityCacheServiceServer
	entityCache *EntityCache
}

// RegisterEntityCacheService registers a service for entity cache.
func RegisterEntityCacheService(server *grpc.Server, cache *EntityCache) {
	svc := &EntityCacheService{
		entityCache: cache,
	}
	entitycachev1.RegisterEntityCacheServiceServer(server, svc)
}

// GetServicesList returns a list of services based on entities in cache.
func (c *EntityCacheService) GetServicesList(ctx context.Context, _ *emptypb.Empty) (*entitycachev1.ServicesList, error) {
	return c.entityCache.Services(), nil
}

// GetEntityByIP returns an entity found with given ip address.
func (c *EntityCacheService) GetEntityByIP(ctx context.Context, req *entitycachev1.GetEntityByIpRequest) (*entitycachev1.Entity, error) {
	return c.entityCache.GetByIP(req.IpAddress), nil
}

// GetEntityByName returns an entity found with given entity name.
func (c *EntityCacheService) GetEntityByName(ctx context.Context, req *entitycachev1.GetEntityByNameRequest) (*entitycachev1.Entity, error) {
	return c.entityCache.GetByName(req.EntityName), nil
}