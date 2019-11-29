// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mqiface provides an interface to enable mocking the AmazonMQ service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mqiface

import (
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
)

// ClientAPI provides an interface to enable mocking the
// mq.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AmazonMQ.
//    func myFunc(svc mqiface.ClientAPI) bool {
//        // Make svc.CreateBroker request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := mq.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        mqiface.ClientPI
//    }
//    func (m *mockClientClient) CreateBroker(input *types.CreateBrokerInput) (*types.CreateBrokerOutput, error) {
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
	CreateBrokerRequest(*types.CreateBrokerInput) mq.CreateBrokerRequest

	CreateConfigurationRequest(*types.CreateConfigurationInput) mq.CreateConfigurationRequest

	CreateTagsRequest(*types.CreateTagsInput) mq.CreateTagsRequest

	CreateUserRequest(*types.CreateUserInput) mq.CreateUserRequest

	DeleteBrokerRequest(*types.DeleteBrokerInput) mq.DeleteBrokerRequest

	DeleteTagsRequest(*types.DeleteTagsInput) mq.DeleteTagsRequest

	DeleteUserRequest(*types.DeleteUserInput) mq.DeleteUserRequest

	DescribeBrokerRequest(*types.DescribeBrokerInput) mq.DescribeBrokerRequest

	DescribeBrokerEngineTypesRequest(*types.DescribeBrokerEngineTypesInput) mq.DescribeBrokerEngineTypesRequest

	DescribeBrokerInstanceOptionsRequest(*types.DescribeBrokerInstanceOptionsInput) mq.DescribeBrokerInstanceOptionsRequest

	DescribeConfigurationRequest(*types.DescribeConfigurationInput) mq.DescribeConfigurationRequest

	DescribeConfigurationRevisionRequest(*types.DescribeConfigurationRevisionInput) mq.DescribeConfigurationRevisionRequest

	DescribeUserRequest(*types.DescribeUserInput) mq.DescribeUserRequest

	ListBrokersRequest(*types.ListBrokersInput) mq.ListBrokersRequest

	ListConfigurationRevisionsRequest(*types.ListConfigurationRevisionsInput) mq.ListConfigurationRevisionsRequest

	ListConfigurationsRequest(*types.ListConfigurationsInput) mq.ListConfigurationsRequest

	ListTagsRequest(*types.ListTagsInput) mq.ListTagsRequest

	ListUsersRequest(*types.ListUsersInput) mq.ListUsersRequest

	RebootBrokerRequest(*types.RebootBrokerInput) mq.RebootBrokerRequest

	UpdateBrokerRequest(*types.UpdateBrokerInput) mq.UpdateBrokerRequest

	UpdateConfigurationRequest(*types.UpdateConfigurationInput) mq.UpdateConfigurationRequest

	UpdateUserRequest(*types.UpdateUserInput) mq.UpdateUserRequest
}

var _ ClientAPI = (*mq.Client)(nil)
