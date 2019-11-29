// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package dataexchangeiface provides an interface to enable mocking the AWS Data Exchange service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package dataexchangeiface

import (
	"github.com/aws/aws-sdk-go-v2/service/dataexchange"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/types"
)

// ClientAPI provides an interface to enable mocking the
// dataexchange.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Data Exchange.
//    func myFunc(svc dataexchangeiface.ClientAPI) bool {
//        // Make svc.CancelJob request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := dataexchange.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        dataexchangeiface.ClientPI
//    }
//    func (m *mockClientClient) CancelJob(input *types.CancelJobInput) (*types.CancelJobOutput, error) {
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
	CancelJobRequest(*types.CancelJobInput) dataexchange.CancelJobRequest

	CreateDataSetRequest(*types.CreateDataSetInput) dataexchange.CreateDataSetRequest

	CreateJobRequest(*types.CreateJobInput) dataexchange.CreateJobRequest

	CreateRevisionRequest(*types.CreateRevisionInput) dataexchange.CreateRevisionRequest

	DeleteAssetRequest(*types.DeleteAssetInput) dataexchange.DeleteAssetRequest

	DeleteDataSetRequest(*types.DeleteDataSetInput) dataexchange.DeleteDataSetRequest

	DeleteRevisionRequest(*types.DeleteRevisionInput) dataexchange.DeleteRevisionRequest

	GetAssetRequest(*types.GetAssetInput) dataexchange.GetAssetRequest

	GetDataSetRequest(*types.GetDataSetInput) dataexchange.GetDataSetRequest

	GetJobRequest(*types.GetJobInput) dataexchange.GetJobRequest

	GetRevisionRequest(*types.GetRevisionInput) dataexchange.GetRevisionRequest

	ListDataSetRevisionsRequest(*types.ListDataSetRevisionsInput) dataexchange.ListDataSetRevisionsRequest

	ListDataSetsRequest(*types.ListDataSetsInput) dataexchange.ListDataSetsRequest

	ListJobsRequest(*types.ListJobsInput) dataexchange.ListJobsRequest

	ListRevisionAssetsRequest(*types.ListRevisionAssetsInput) dataexchange.ListRevisionAssetsRequest

	ListTagsForResourceRequest(*types.ListTagsForResourceInput) dataexchange.ListTagsForResourceRequest

	StartJobRequest(*types.StartJobInput) dataexchange.StartJobRequest

	TagResourceRequest(*types.TagResourceInput) dataexchange.TagResourceRequest

	UntagResourceRequest(*types.UntagResourceInput) dataexchange.UntagResourceRequest

	UpdateAssetRequest(*types.UpdateAssetInput) dataexchange.UpdateAssetRequest

	UpdateDataSetRequest(*types.UpdateDataSetInput) dataexchange.UpdateDataSetRequest

	UpdateRevisionRequest(*types.UpdateRevisionInput) dataexchange.UpdateRevisionRequest
}

var _ ClientAPI = (*dataexchange.Client)(nil)
