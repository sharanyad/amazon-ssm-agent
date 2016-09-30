// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package service wraps SSM service
package service

import (
	"github.com/aws/amazon-ssm-agent/agent/association/model"
	"github.com/aws/amazon-ssm-agent/agent/contracts"
	"github.com/aws/amazon-ssm-agent/agent/log"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/stretchr/testify/mock"
)

// AssociationServiceMock stands for a mocked association service.
type AssociationServiceMock struct {
	mock.Mock
}

// NewMockDefault returns an instance of Mock with default expectations set.
func NewMockDefault() *AssociationServiceMock {
	return new(AssociationServiceMock)
}

// ListAssociations mocks implementation for ListAssociations
func (m *AssociationServiceMock) ListAssociations(log log.T, instanceID string) (*model.AssociationRawData, error) {
	args := m.Called(log, instanceID)
	return args.Get(0).(*model.AssociationRawData), args.Error(1)
}

// CreateNewServiceIfUnHealthy mocks implementation for CreateNewServiceIfUnHealthy
func (m *AssociationServiceMock) CreateNewServiceIfUnHealthy(log log.T) {
	m.Called(log)
}

// LoadAssociationDetail mocks implementation for LoadAssociationDetail
func (m *AssociationServiceMock) LoadAssociationDetail(log log.T, assoc *model.AssociationRawData) error {
	args := m.Called(log, assoc)
	return args.Error(0)
}

// UpdateAssociationStatus mocks implementation for UpdateAssociationStatus
func (m *AssociationServiceMock) UpdateAssociationStatus(
	log log.T,
	instanceID string,
	name string,
	status string,
	message string,
	agentInfo *contracts.AgentInfo) (*ssm.UpdateAssociationStatusOutput, error) {
	args := m.Called(log, instanceID, name, status, message, agentInfo)
	return args.Get(0).(*ssm.UpdateAssociationStatusOutput), args.Error(1)
}