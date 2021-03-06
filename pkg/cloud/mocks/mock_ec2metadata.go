// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/cloud (interfaces: EC2Metadata)

package mocks

import (
	ec2metadata "github.com/aws/aws-sdk-go/aws/ec2metadata"
	gomock "github.com/golang/mock/gomock"
)

// Mock of EC2Metadata interface
type MockEC2Metadata struct {
	ctrl     *gomock.Controller
	recorder *_MockEC2MetadataRecorder
}

// Recorder for MockEC2Metadata (not exported)
type _MockEC2MetadataRecorder struct {
	mock *MockEC2Metadata
}

func NewMockEC2Metadata(ctrl *gomock.Controller) *MockEC2Metadata {
	mock := &MockEC2Metadata{ctrl: ctrl}
	mock.recorder = &_MockEC2MetadataRecorder{mock}
	return mock
}

func (_m *MockEC2Metadata) EXPECT() *_MockEC2MetadataRecorder {
	return _m.recorder
}

func (_m *MockEC2Metadata) Available() bool {
	ret := _m.ctrl.Call(_m, "Available")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockEC2MetadataRecorder) Available() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Available")
}

func (_m *MockEC2Metadata) GetInstanceIdentityDocument() (ec2metadata.EC2InstanceIdentityDocument, error) {
	ret := _m.ctrl.Call(_m, "GetInstanceIdentityDocument")
	ret0, _ := ret[0].(ec2metadata.EC2InstanceIdentityDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEC2MetadataRecorder) GetInstanceIdentityDocument() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetInstanceIdentityDocument")
}
