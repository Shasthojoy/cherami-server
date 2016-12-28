// Copyright (c) 2016 Uber Technologies, Inc.
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.


package metadata

import "github.com/uber/cherami-thrift/.generated/go/metadata"
import "github.com/uber/cherami-thrift/.generated/go/shared"
import "github.com/stretchr/testify/mock"

import "github.com/uber/tchannel-go/thrift"

// TChanMetadataServiceClient is an autogenerated mock type for the TChanMetadataServiceClient type
type TChanMetadataServiceClient struct {
	mock.Mock
}

// CreateConsumerGroup provides a mock function with given fields: ctx, registerRequest
func (_m *TChanMetadataServiceClient) CreateConsumerGroup(ctx thrift.Context, registerRequest *shared.CreateConsumerGroupRequest) (*shared.ConsumerGroupDescription, error) {
	ret := _m.Called(ctx, registerRequest)

	var r0 *shared.ConsumerGroupDescription
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.CreateConsumerGroupRequest) *shared.ConsumerGroupDescription); ok {
		r0 = rf(ctx, registerRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ConsumerGroupDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.CreateConsumerGroupRequest) error); ok {
		r1 = rf(ctx, registerRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadConsumerGroupByUUID provides a mock function with given fields: ctx, getRequest
func (_m *TChanMetadataServiceClient) ReadConsumerGroupByUUID(ctx thrift.Context, getRequest *metadata.ReadConsumerGroupRequest) (*shared.ConsumerGroupDescription, error) {
	ret := _m.Called(ctx, getRequest)

	var r0 *shared.ConsumerGroupDescription
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.ReadConsumerGroupRequest) *shared.ConsumerGroupDescription); ok {
		r0 = rf(ctx, getRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ConsumerGroupDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *metadata.ReadConsumerGroupRequest) error); ok {
		r1 = rf(ctx, getRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateConsumerGroupExtent provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) CreateConsumerGroupExtent(ctx thrift.Context, request *metadata.CreateConsumerGroupExtentRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.CreateConsumerGroupExtentRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateDestination provides a mock function with given fields: ctx, createRequest
func (_m *TChanMetadataServiceClient) CreateDestination(ctx thrift.Context, createRequest *shared.CreateDestinationRequest) (*shared.DestinationDescription, error) {
	ret := _m.Called(ctx, createRequest)

	var r0 *shared.DestinationDescription
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.CreateDestinationRequest) *shared.DestinationDescription); ok {
		r0 = rf(ctx, createRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.DestinationDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.CreateDestinationRequest) error); ok {
		r1 = rf(ctx, createRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDestinationUUID provides a mock function with given fields: ctx, createRequest
func (_m *TChanMetadataServiceClient) CreateDestinationUUID(ctx thrift.Context, createRequest *shared.CreateDestinationUUIDRequest) (*shared.DestinationDescription, error) {
	ret := _m.Called(ctx, createRequest)

	var r0 *shared.DestinationDescription
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.CreateDestinationUUIDRequest) *shared.DestinationDescription); ok {
		r0 = rf(ctx, createRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.DestinationDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.CreateDestinationUUIDRequest) error); ok {
		r1 = rf(ctx, createRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateExtent provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) CreateExtent(ctx thrift.Context, request *shared.CreateExtentRequest) (*shared.CreateExtentResult_, error) {
	ret := _m.Called(ctx, request)

	var r0 *shared.CreateExtentResult_
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.CreateExtentRequest) *shared.CreateExtentResult_); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.CreateExtentResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.CreateExtentRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteConsumerGroup provides a mock function with given fields: ctx, deleteRequest
func (_m *TChanMetadataServiceClient) DeleteConsumerGroup(ctx thrift.Context, deleteRequest *shared.DeleteConsumerGroupRequest) error {
	ret := _m.Called(ctx, deleteRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.DeleteConsumerGroupRequest) error); ok {
		r0 = rf(ctx, deleteRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDestination provides a mock function with given fields: ctx, deleteRequest
func (_m *TChanMetadataServiceClient) DeleteDestination(ctx thrift.Context, deleteRequest *shared.DeleteDestinationRequest) error {
	ret := _m.Called(ctx, deleteRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.DeleteDestinationRequest) error); ok {
		r0 = rf(ctx, deleteRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDestinationUUID provides a mock function with given fields: ctx, deleteRequest
func (_m *TChanMetadataServiceClient) DeleteDestinationUUID(ctx thrift.Context, deleteRequest *metadata.DeleteDestinationUUIDRequest) error {
	ret := _m.Called(ctx, deleteRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.DeleteDestinationUUIDRequest) error); ok {
		r0 = rf(ctx, deleteRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HostAddrToUUID provides a mock function with given fields: ctx, hostAddr
func (_m *TChanMetadataServiceClient) HostAddrToUUID(ctx thrift.Context, hostAddr string) (string, error) {
	ret := _m.Called(ctx, hostAddr)

	var r0 string
	if rf, ok := ret.Get(0).(func(thrift.Context, string) string); ok {
		r0 = rf(ctx, hostAddr)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, string) error); ok {
		r1 = rf(ctx, hostAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllConsumerGroups provides a mock function with given fields: ctx, listRequest
func (_m *TChanMetadataServiceClient) ListAllConsumerGroups(ctx thrift.Context, listRequest *metadata.ListConsumerGroupRequest) (*metadata.ListConsumerGroupResult_, error) {
	ret := _m.Called(ctx, listRequest)

	var r0 *metadata.ListConsumerGroupResult_
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.ListConsumerGroupRequest) *metadata.ListConsumerGroupResult_); ok {
		r0 = rf(ctx, listRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ListConsumerGroupResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *metadata.ListConsumerGroupRequest) error); ok {
		r1 = rf(ctx, listRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllConsumerGroups provides a mock function with given fields: ctx, listRequest
func (_m *TChanMetadataServiceClient) ListEntityOps(ctx thrift.Context, listRequest *metadata.ListEntityOpsRequest) (*metadata.ListEntityOpsResult_, error) {
	ret := _m.Called(ctx, listRequest)

	var r0 *metadata.ListEntityOpsResult_
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.ListEntityOpsRequest) *metadata.ListEntityOpsResult_); ok {
		r0 = rf(ctx, listRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ListEntityOpsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *metadata.ListEntityOpsRequest) error); ok {
		r1 = rf(ctx, listRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListConsumerGroups provides a mock function with given fields: ctx, listRequest
func (_m *TChanMetadataServiceClient) ListConsumerGroups(ctx thrift.Context, listRequest *metadata.ListConsumerGroupRequest) (*metadata.ListConsumerGroupResult_, error) {
	ret := _m.Called(ctx, listRequest)

	var r0 *metadata.ListConsumerGroupResult_
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.ListConsumerGroupRequest) *metadata.ListConsumerGroupResult_); ok {
		r0 = rf(ctx, listRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ListConsumerGroupResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *metadata.ListConsumerGroupRequest) error); ok {
		r1 = rf(ctx, listRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDestinations provides a mock function with given fields: ctx, listRequest
func (_m *TChanMetadataServiceClient) ListDestinations(ctx thrift.Context, listRequest *shared.ListDestinationsRequest) (*shared.ListDestinationsResult_, error) {
	ret := _m.Called(ctx, listRequest)

	var r0 *shared.ListDestinationsResult_
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.ListDestinationsRequest) *shared.ListDestinationsResult_); ok {
		r0 = rf(ctx, listRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ListDestinationsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.ListDestinationsRequest) error); ok {
		r1 = rf(ctx, listRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDestinationsByUUID provides a mock function with given fields: ctx, listRequest
func (_m *TChanMetadataServiceClient) ListDestinationsByUUID(ctx thrift.Context, listRequest *shared.ListDestinationsByUUIDRequest) (*shared.ListDestinationsResult_, error) {
	ret := _m.Called(ctx, listRequest)

	var r0 *shared.ListDestinationsResult_
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.ListDestinationsByUUIDRequest) *shared.ListDestinationsResult_); ok {
		r0 = rf(ctx, listRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ListDestinationsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.ListDestinationsByUUIDRequest) error); ok {
		r1 = rf(ctx, listRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListExtentsStats provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) ListExtentsStats(ctx thrift.Context, request *shared.ListExtentsStatsRequest) (*shared.ListExtentsStatsResult_, error) {
	ret := _m.Called(ctx, request)

	var r0 *shared.ListExtentsStatsResult_
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.ListExtentsStatsRequest) *shared.ListExtentsStatsResult_); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ListExtentsStatsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.ListExtentsStatsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListHosts provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) ListHosts(ctx thrift.Context, request *metadata.ListHostsRequest) (*metadata.ListHostsResult_, error) {
	ret := _m.Called(ctx, request)

	var r0 *metadata.ListHostsResult_
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.ListHostsRequest) *metadata.ListHostsResult_); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ListHostsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *metadata.ListHostsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListInputHostExtentsStats provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) ListInputHostExtentsStats(ctx thrift.Context, request *metadata.ListInputHostExtentsStatsRequest) (*metadata.ListInputHostExtentsStatsResult_, error) {
	ret := _m.Called(ctx, request)

	var r0 *metadata.ListInputHostExtentsStatsResult_
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.ListInputHostExtentsStatsRequest) *metadata.ListInputHostExtentsStatsResult_); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ListInputHostExtentsStatsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *metadata.ListInputHostExtentsStatsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListStoreExtentsStats provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) ListStoreExtentsStats(ctx thrift.Context, request *metadata.ListStoreExtentsStatsRequest) (*metadata.ListStoreExtentsStatsResult_, error) {
	ret := _m.Called(ctx, request)

	var r0 *metadata.ListStoreExtentsStatsResult_
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.ListStoreExtentsStatsRequest) *metadata.ListStoreExtentsStatsResult_); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ListStoreExtentsStatsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *metadata.ListStoreExtentsStatsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MoveExtent provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) MoveExtent(ctx thrift.Context, request *metadata.MoveExtentRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.MoveExtentRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadConsumerGroup provides a mock function with given fields: ctx, getRequest
func (_m *TChanMetadataServiceClient) ReadConsumerGroup(ctx thrift.Context, getRequest *metadata.ReadConsumerGroupRequest) (*shared.ConsumerGroupDescription, error) {
	ret := _m.Called(ctx, getRequest)

	var r0 *shared.ConsumerGroupDescription
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.ReadConsumerGroupRequest) *shared.ConsumerGroupDescription); ok {
		r0 = rf(ctx, getRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ConsumerGroupDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *metadata.ReadConsumerGroupRequest) error); ok {
		r1 = rf(ctx, getRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadConsumerGroupExtent provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) ReadConsumerGroupExtent(ctx thrift.Context, request *metadata.ReadConsumerGroupExtentRequest) (*metadata.ReadConsumerGroupExtentResult_, error) {
	ret := _m.Called(ctx, request)

	var r0 *metadata.ReadConsumerGroupExtentResult_
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.ReadConsumerGroupExtentRequest) *metadata.ReadConsumerGroupExtentResult_); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ReadConsumerGroupExtentResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *metadata.ReadConsumerGroupExtentRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadConsumerGroupExtents provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) ReadConsumerGroupExtents(ctx thrift.Context, request *metadata.ReadConsumerGroupExtentsRequest) (*metadata.ReadConsumerGroupExtentsResult_, error) {
	ret := _m.Called(ctx, request)

	var r0 *metadata.ReadConsumerGroupExtentsResult_
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.ReadConsumerGroupExtentsRequest) *metadata.ReadConsumerGroupExtentsResult_); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ReadConsumerGroupExtentsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *metadata.ReadConsumerGroupExtentsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadConsumerGroupExtentsByExtUUID provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) ReadConsumerGroupExtentsByExtUUID(ctx thrift.Context, request *metadata.ReadConsumerGroupExtentsByExtUUIDRequest) (*metadata.ReadConsumerGroupExtentsByExtUUIDResult_, error) {
	ret := _m.Called(ctx, request)

	var r0 *metadata.ReadConsumerGroupExtentsByExtUUIDResult_
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.ReadConsumerGroupExtentsByExtUUIDRequest) *metadata.ReadConsumerGroupExtentsByExtUUIDResult_); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ReadConsumerGroupExtentsByExtUUIDResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *metadata.ReadConsumerGroupExtentsByExtUUIDRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadDestination provides a mock function with given fields: ctx, getRequest
func (_m *TChanMetadataServiceClient) ReadDestination(ctx thrift.Context, getRequest *metadata.ReadDestinationRequest) (*shared.DestinationDescription, error) {
	ret := _m.Called(ctx, getRequest)

	var r0 *shared.DestinationDescription
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.ReadDestinationRequest) *shared.DestinationDescription); ok {
		r0 = rf(ctx, getRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.DestinationDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *metadata.ReadDestinationRequest) error); ok {
		r1 = rf(ctx, getRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadExtentStats provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) ReadExtentStats(ctx thrift.Context, request *metadata.ReadExtentStatsRequest) (*metadata.ReadExtentStatsResult_, error) {
	ret := _m.Called(ctx, request)

	var r0 *metadata.ReadExtentStatsResult_
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.ReadExtentStatsRequest) *metadata.ReadExtentStatsResult_); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ReadExtentStatsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *metadata.ReadExtentStatsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadStoreExtentReplicaStats provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) ReadStoreExtentReplicaStats(ctx thrift.Context, request *metadata.ReadStoreExtentReplicaStatsRequest) (*metadata.ReadStoreExtentReplicaStatsResult_, error) {
	ret := _m.Called(ctx, request)

	var r0 *metadata.ReadStoreExtentReplicaStatsResult_
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.ReadStoreExtentReplicaStatsRequest) *metadata.ReadStoreExtentReplicaStatsResult_); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ReadStoreExtentReplicaStatsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *metadata.ReadStoreExtentReplicaStatsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterHostUUID provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) RegisterHostUUID(ctx thrift.Context, request *metadata.RegisterHostUUIDRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.RegisterHostUUIDRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SealExtent provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) SealExtent(ctx thrift.Context, request *metadata.SealExtentRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.SealExtentRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetAckOffset provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) SetAckOffset(ctx thrift.Context, request *metadata.SetAckOffsetRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.SetAckOffsetRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetOutputHost provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) SetOutputHost(ctx thrift.Context, request *metadata.SetOutputHostRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.SetOutputHostRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UUIDToHostAddr provides a mock function with given fields: ctx, hostUUID
func (_m *TChanMetadataServiceClient) UUIDToHostAddr(ctx thrift.Context, hostUUID string) (string, error) {
	ret := _m.Called(ctx, hostUUID)

	var r0 string
	if rf, ok := ret.Get(0).(func(thrift.Context, string) string); ok {
		r0 = rf(ctx, hostUUID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, string) error); ok {
		r1 = rf(ctx, hostUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateConsumerGroup provides a mock function with given fields: ctx, updateRequest
func (_m *TChanMetadataServiceClient) UpdateConsumerGroup(ctx thrift.Context, updateRequest *shared.UpdateConsumerGroupRequest) (*shared.ConsumerGroupDescription, error) {
	ret := _m.Called(ctx, updateRequest)

	var r0 *shared.ConsumerGroupDescription
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.UpdateConsumerGroupRequest) *shared.ConsumerGroupDescription); ok {
		r0 = rf(ctx, updateRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ConsumerGroupDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.UpdateConsumerGroupRequest) error); ok {
		r1 = rf(ctx, updateRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateConsumerGroupExtentStatus provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) UpdateConsumerGroupExtentStatus(ctx thrift.Context, request *metadata.UpdateConsumerGroupExtentStatusRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.UpdateConsumerGroupExtentStatusRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDestination provides a mock function with given fields: ctx, updateRequest
func (_m *TChanMetadataServiceClient) UpdateDestination(ctx thrift.Context, updateRequest *shared.UpdateDestinationRequest) (*shared.DestinationDescription, error) {
	ret := _m.Called(ctx, updateRequest)

	var r0 *shared.DestinationDescription
	if rf, ok := ret.Get(0).(func(thrift.Context, *shared.UpdateDestinationRequest) *shared.DestinationDescription); ok {
		r0 = rf(ctx, updateRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.DestinationDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *shared.UpdateDestinationRequest) error); ok {
		r1 = rf(ctx, updateRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDestinationDLQCursors provides a mock function with given fields: ctx, updateRequest
func (_m *TChanMetadataServiceClient) UpdateDestinationDLQCursors(ctx thrift.Context, updateRequest *metadata.UpdateDestinationDLQCursorsRequest) (*shared.DestinationDescription, error) {
	ret := _m.Called(ctx, updateRequest)

	var r0 *shared.DestinationDescription
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.UpdateDestinationDLQCursorsRequest) *shared.DestinationDescription); ok {
		r0 = rf(ctx, updateRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.DestinationDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *metadata.UpdateDestinationDLQCursorsRequest) error); ok {
		r1 = rf(ctx, updateRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateExtentReplicaStats provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) UpdateExtentReplicaStats(ctx thrift.Context, request *metadata.UpdateExtentReplicaStatsRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.UpdateExtentReplicaStatsRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateExtentStats provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) UpdateExtentStats(ctx thrift.Context, request *metadata.UpdateExtentStatsRequest) (*metadata.UpdateExtentStatsResult_, error) {
	ret := _m.Called(ctx, request)

	var r0 *metadata.UpdateExtentStatsResult_
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.UpdateExtentStatsRequest) *metadata.UpdateExtentStatsResult_); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.UpdateExtentStatsResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *metadata.UpdateExtentStatsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStoreExtentReplicaStats provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) UpdateStoreExtentReplicaStats(ctx thrift.Context, request *metadata.UpdateStoreExtentReplicaStatsRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.UpdateStoreExtentReplicaStatsRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateHostInfo provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) CreateHostInfo(ctx thrift.Context, request *metadata.CreateHostInfoRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.CreateHostInfoRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteHostInfo provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) DeleteHostInfo(ctx thrift.Context, request *metadata.DeleteHostInfoRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.DeleteHostInfoRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadHostInfo provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) ReadHostInfo(ctx thrift.Context, request *metadata.ReadHostInfoRequest) (*metadata.ReadHostInfoResult_, error) {
	ret := _m.Called(ctx, request)

	var r0 *metadata.ReadHostInfoResult_
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.ReadHostInfoRequest) *metadata.ReadHostInfoResult_); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ReadHostInfoResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *metadata.ReadHostInfoRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateHostInfo provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) UpdateHostInfo(ctx thrift.Context, request *metadata.UpdateHostInfoRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.UpdateHostInfoRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateServiceConfig provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) CreateServiceConfig(ctx thrift.Context, request *metadata.CreateServiceConfigRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.CreateServiceConfigRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteServiceConfig provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) DeleteServiceConfig(ctx thrift.Context, request *metadata.DeleteServiceConfigRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.DeleteServiceConfigRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadServiceConfig provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) ReadServiceConfig(ctx thrift.Context, request *metadata.ReadServiceConfigRequest) (*metadata.ReadServiceConfigResult_, error) {
	ret := _m.Called(ctx, request)

	var r0 *metadata.ReadServiceConfigResult_
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.ReadServiceConfigRequest) *metadata.ReadServiceConfigResult_); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ReadServiceConfigResult_)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *metadata.ReadServiceConfigRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateServiceConfig provides a mock function with given fields: ctx, request
func (_m *TChanMetadataServiceClient) UpdateServiceConfig(ctx thrift.Context, request *metadata.UpdateServiceConfigRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *metadata.UpdateServiceConfigRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
