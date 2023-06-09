// Code generated by mockery v2.27.1. DO NOT EDIT.

package inventorymocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	inventory "api-gateway/proto/inventory"

	mock "github.com/stretchr/testify/mock"
)

// InventoryServiceClient is an autogenerated mock type for the InventoryServiceClient type
type InventoryServiceClient struct {
	mock.Mock
}

// AddItem provides a mock function with given fields: ctx, in, opts
func (_m *InventoryServiceClient) AddItem(ctx context.Context, in *inventory.AddItemRequest, opts ...grpc.CallOption) (*inventory.AddItemResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *inventory.AddItemResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *inventory.AddItemRequest, ...grpc.CallOption) (*inventory.AddItemResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *inventory.AddItemRequest, ...grpc.CallOption) *inventory.AddItemResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*inventory.AddItemResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *inventory.AddItemRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddQuantity provides a mock function with given fields: ctx, in, opts
func (_m *InventoryServiceClient) AddQuantity(ctx context.Context, in *inventory.AddQuantityRequest, opts ...grpc.CallOption) (*inventory.AddQuantityResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *inventory.AddQuantityResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *inventory.AddQuantityRequest, ...grpc.CallOption) (*inventory.AddQuantityResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *inventory.AddQuantityRequest, ...grpc.CallOption) *inventory.AddQuantityResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*inventory.AddQuantityResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *inventory.AddQuantityRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteItem provides a mock function with given fields: ctx, in, opts
func (_m *InventoryServiceClient) DeleteItem(ctx context.Context, in *inventory.DeleteItemRequest, opts ...grpc.CallOption) (*inventory.DeleteItemResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *inventory.DeleteItemResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *inventory.DeleteItemRequest, ...grpc.CallOption) (*inventory.DeleteItemResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *inventory.DeleteItemRequest, ...grpc.CallOption) *inventory.DeleteItemResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*inventory.DeleteItemResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *inventory.DeleteItemRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllItems provides a mock function with given fields: ctx, in, opts
func (_m *InventoryServiceClient) GetAllItems(ctx context.Context, in *inventory.GetAllItemsRequest, opts ...grpc.CallOption) (*inventory.GetAllItemsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *inventory.GetAllItemsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *inventory.GetAllItemsRequest, ...grpc.CallOption) (*inventory.GetAllItemsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *inventory.GetAllItemsRequest, ...grpc.CallOption) *inventory.GetAllItemsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*inventory.GetAllItemsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *inventory.GetAllItemsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetItem provides a mock function with given fields: ctx, in, opts
func (_m *InventoryServiceClient) GetItem(ctx context.Context, in *inventory.GetItemRequest, opts ...grpc.CallOption) (*inventory.GetItemResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *inventory.GetItemResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *inventory.GetItemRequest, ...grpc.CallOption) (*inventory.GetItemResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *inventory.GetItemRequest, ...grpc.CallOption) *inventory.GetItemResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*inventory.GetItemResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *inventory.GetItemRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LowerQuantity provides a mock function with given fields: ctx, in, opts
func (_m *InventoryServiceClient) LowerQuantity(ctx context.Context, in *inventory.LowerQuantityRequest, opts ...grpc.CallOption) (*inventory.LowerQuantityResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *inventory.LowerQuantityResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *inventory.LowerQuantityRequest, ...grpc.CallOption) (*inventory.LowerQuantityResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *inventory.LowerQuantityRequest, ...grpc.CallOption) *inventory.LowerQuantityResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*inventory.LowerQuantityResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *inventory.LowerQuantityRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewInventoryServiceClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewInventoryServiceClient creates a new instance of InventoryServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInventoryServiceClient(t mockConstructorTestingTNewInventoryServiceClient) *InventoryServiceClient {
	mock := &InventoryServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
