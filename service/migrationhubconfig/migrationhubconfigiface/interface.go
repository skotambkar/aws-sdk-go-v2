// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package migrationhubconfigiface provides an interface to enable mocking the AWS Migration Hub Config service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package migrationhubconfigiface

import (
	"github.com/aws/aws-sdk-go-v2/service/migrationhubconfig"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubconfig/types"
)

// ClientAPI provides an interface to enable mocking the
// migrationhubconfig.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Migration Hub Config.
//    func myFunc(svc migrationhubconfigiface.ClientAPI) bool {
//        // Make svc.CreateHomeRegionControl request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := migrationhubconfig.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        migrationhubconfigiface.ClientPI
//    }
//    func (m *mockClientClient) CreateHomeRegionControl(input *types.CreateHomeRegionControlInput) (*types.CreateHomeRegionControlOutput, error) {
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
	CreateHomeRegionControlRequest(*types.CreateHomeRegionControlInput) migrationhubconfig.CreateHomeRegionControlRequest

	DescribeHomeRegionControlsRequest(*types.DescribeHomeRegionControlsInput) migrationhubconfig.DescribeHomeRegionControlsRequest

	GetHomeRegionRequest(*types.GetHomeRegionInput) migrationhubconfig.GetHomeRegionRequest
}

var _ ClientAPI = (*migrationhubconfig.Client)(nil)
