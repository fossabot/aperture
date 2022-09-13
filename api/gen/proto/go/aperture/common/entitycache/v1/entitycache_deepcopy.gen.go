// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package entitycachev1

import (
	proto "github.com/golang/protobuf/proto"
)

// DeepCopyInto supports using ServicesList within kubernetes types, where deepcopy-gen is used.
func (in *ServicesList) DeepCopyInto(out *ServicesList) {
	p := proto.Clone(in).(*ServicesList)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesList. Required by controller-gen.
func (in *ServicesList) DeepCopy() *ServicesList {
	if in == nil {
		return nil
	}
	out := new(ServicesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ServicesList. Required by controller-gen.
func (in *ServicesList) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Service within kubernetes types, where deepcopy-gen is used.
func (in *Service) DeepCopyInto(out *Service) {
	p := proto.Clone(in).(*Service)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service. Required by controller-gen.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Service. Required by controller-gen.
func (in *Service) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using OverlappingService within kubernetes types, where deepcopy-gen is used.
func (in *OverlappingService) DeepCopyInto(out *OverlappingService) {
	p := proto.Clone(in).(*OverlappingService)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OverlappingService. Required by controller-gen.
func (in *OverlappingService) DeepCopy() *OverlappingService {
	if in == nil {
		return nil
	}
	out := new(OverlappingService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new OverlappingService. Required by controller-gen.
func (in *OverlappingService) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GetEntityByIpRequest within kubernetes types, where deepcopy-gen is used.
func (in *GetEntityByIpRequest) DeepCopyInto(out *GetEntityByIpRequest) {
	p := proto.Clone(in).(*GetEntityByIpRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetEntityByIpRequest. Required by controller-gen.
func (in *GetEntityByIpRequest) DeepCopy() *GetEntityByIpRequest {
	if in == nil {
		return nil
	}
	out := new(GetEntityByIpRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetEntityByIpRequest. Required by controller-gen.
func (in *GetEntityByIpRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GetEntityByNameRequest within kubernetes types, where deepcopy-gen is used.
func (in *GetEntityByNameRequest) DeepCopyInto(out *GetEntityByNameRequest) {
	p := proto.Clone(in).(*GetEntityByNameRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetEntityByNameRequest. Required by controller-gen.
func (in *GetEntityByNameRequest) DeepCopy() *GetEntityByNameRequest {
	if in == nil {
		return nil
	}
	out := new(GetEntityByNameRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetEntityByNameRequest. Required by controller-gen.
func (in *GetEntityByNameRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Entities within kubernetes types, where deepcopy-gen is used.
func (in *Entities) DeepCopyInto(out *Entities) {
	p := proto.Clone(in).(*Entities)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Entities. Required by controller-gen.
func (in *Entities) DeepCopy() *Entities {
	if in == nil {
		return nil
	}
	out := new(Entities)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Entities. Required by controller-gen.
func (in *Entities) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Entity within kubernetes types, where deepcopy-gen is used.
func (in *Entity) DeepCopyInto(out *Entity) {
	p := proto.Clone(in).(*Entity)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Entity. Required by controller-gen.
func (in *Entity) DeepCopy() *Entity {
	if in == nil {
		return nil
	}
	out := new(Entity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Entity. Required by controller-gen.
func (in *Entity) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Entity_EntityID within kubernetes types, where deepcopy-gen is used.
func (in *Entity_EntityID) DeepCopyInto(out *Entity_EntityID) {
	p := proto.Clone(in).(*Entity_EntityID)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Entity_EntityID. Required by controller-gen.
func (in *Entity_EntityID) DeepCopy() *Entity_EntityID {
	if in == nil {
		return nil
	}
	out := new(Entity_EntityID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Entity_EntityID. Required by controller-gen.
func (in *Entity_EntityID) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
