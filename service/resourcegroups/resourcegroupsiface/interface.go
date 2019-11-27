// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package resourcegroupsiface provides an interface to enable mocking the AWS Resource Groups service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package resourcegroupsiface

import (
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
)

// ClientAPI provides an interface to enable mocking the
// resourcegroups.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Resource Groups.
//    func myFunc(svc resourcegroupsiface.ClientAPI) bool {
//        // Make svc.CreateGroup request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := resourcegroups.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        resourcegroupsiface.ClientPI
//    }
//    func (m *mockClientClient) CreateGroup(input *types.CreateGroupInput) (*types.CreateGroupOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	CreateGroupRequest(*types.CreateGroupInput) resourcegroups.CreateGroupRequest

	DeleteGroupRequest(*types.DeleteGroupInput) resourcegroups.DeleteGroupRequest

	GetGroupRequest(*types.GetGroupInput) resourcegroups.GetGroupRequest

	GetGroupQueryRequest(*types.GetGroupQueryInput) resourcegroups.GetGroupQueryRequest

	GetTagsRequest(*types.GetTagsInput) resourcegroups.GetTagsRequest

	ListGroupResourcesRequest(*types.ListGroupResourcesInput) resourcegroups.ListGroupResourcesRequest

	ListGroupsRequest(*types.ListGroupsInput) resourcegroups.ListGroupsRequest

	SearchResourcesRequest(*types.SearchResourcesInput) resourcegroups.SearchResourcesRequest

	TagRequest(*types.TagInput) resourcegroups.TagRequest

	UntagRequest(*types.UntagInput) resourcegroups.UntagRequest

	UpdateGroupRequest(*types.UpdateGroupInput) resourcegroups.UpdateGroupRequest

	UpdateGroupQueryRequest(*types.UpdateGroupQueryInput) resourcegroups.UpdateGroupQueryRequest
}

var _ ClientAPI = (*resourcegroups.Client)(nil)
