// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cognitosynciface provides an interface to enable mocking the Amazon Cognito Sync service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cognitosynciface

import (
	"github.com/aws/aws-sdk-go-v2/service/cognitosync"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
)

// ClientAPI provides an interface to enable mocking the
// cognitosync.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Cognito Sync.
//    func myFunc(svc cognitosynciface.ClientAPI) bool {
//        // Make svc.BulkPublish request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := cognitosync.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        cognitosynciface.ClientPI
//    }
//    func (m *mockClientClient) BulkPublish(input *types.BulkPublishInput) (*types.BulkPublishOutput, error) {
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
	BulkPublishRequest(*types.BulkPublishInput) cognitosync.BulkPublishRequest

	DeleteDatasetRequest(*types.DeleteDatasetInput) cognitosync.DeleteDatasetRequest

	DescribeDatasetRequest(*types.DescribeDatasetInput) cognitosync.DescribeDatasetRequest

	DescribeIdentityPoolUsageRequest(*types.DescribeIdentityPoolUsageInput) cognitosync.DescribeIdentityPoolUsageRequest

	DescribeIdentityUsageRequest(*types.DescribeIdentityUsageInput) cognitosync.DescribeIdentityUsageRequest

	GetBulkPublishDetailsRequest(*types.GetBulkPublishDetailsInput) cognitosync.GetBulkPublishDetailsRequest

	GetCognitoEventsRequest(*types.GetCognitoEventsInput) cognitosync.GetCognitoEventsRequest

	GetIdentityPoolConfigurationRequest(*types.GetIdentityPoolConfigurationInput) cognitosync.GetIdentityPoolConfigurationRequest

	ListDatasetsRequest(*types.ListDatasetsInput) cognitosync.ListDatasetsRequest

	ListIdentityPoolUsageRequest(*types.ListIdentityPoolUsageInput) cognitosync.ListIdentityPoolUsageRequest

	ListRecordsRequest(*types.ListRecordsInput) cognitosync.ListRecordsRequest

	RegisterDeviceRequest(*types.RegisterDeviceInput) cognitosync.RegisterDeviceRequest

	SetCognitoEventsRequest(*types.SetCognitoEventsInput) cognitosync.SetCognitoEventsRequest

	SetIdentityPoolConfigurationRequest(*types.SetIdentityPoolConfigurationInput) cognitosync.SetIdentityPoolConfigurationRequest

	SubscribeToDatasetRequest(*types.SubscribeToDatasetInput) cognitosync.SubscribeToDatasetRequest

	UnsubscribeFromDatasetRequest(*types.UnsubscribeFromDatasetInput) cognitosync.UnsubscribeFromDatasetRequest

	UpdateRecordsRequest(*types.UpdateRecordsInput) cognitosync.UpdateRecordsRequest
}

var _ ClientAPI = (*cognitosync.Client)(nil)
