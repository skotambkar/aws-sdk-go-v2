// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kendraiface provides an interface to enable mocking the AWSKendraFrontendService service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kendraiface

import (
	"github.com/aws/aws-sdk-go-v2/service/kendra"
)

// ClientAPI provides an interface to enable mocking the
// kendra.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // kendra.
//    func myFunc(svc kendraiface.ClientAPI) bool {
//        // Make svc.BatchDeleteDocument request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := kendra.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        kendraiface.ClientPI
//    }
//    func (m *mockClientClient) BatchDeleteDocument(input *kendra.BatchDeleteDocumentInput) (*kendra.BatchDeleteDocumentOutput, error) {
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
	BatchDeleteDocumentRequest(*kendra.BatchDeleteDocumentInput) kendra.BatchDeleteDocumentRequest

	BatchPutDocumentRequest(*kendra.BatchPutDocumentInput) kendra.BatchPutDocumentRequest

	CreateDataSourceRequest(*kendra.CreateDataSourceInput) kendra.CreateDataSourceRequest

	CreateFaqRequest(*kendra.CreateFaqInput) kendra.CreateFaqRequest

	CreateIndexRequest(*kendra.CreateIndexInput) kendra.CreateIndexRequest

	DeleteFaqRequest(*kendra.DeleteFaqInput) kendra.DeleteFaqRequest

	DeleteIndexRequest(*kendra.DeleteIndexInput) kendra.DeleteIndexRequest

	DescribeDataSourceRequest(*kendra.DescribeDataSourceInput) kendra.DescribeDataSourceRequest

	DescribeFaqRequest(*kendra.DescribeFaqInput) kendra.DescribeFaqRequest

	DescribeIndexRequest(*kendra.DescribeIndexInput) kendra.DescribeIndexRequest

	ListDataSourceSyncJobsRequest(*kendra.ListDataSourceSyncJobsInput) kendra.ListDataSourceSyncJobsRequest

	ListDataSourcesRequest(*kendra.ListDataSourcesInput) kendra.ListDataSourcesRequest

	ListFaqsRequest(*kendra.ListFaqsInput) kendra.ListFaqsRequest

	ListIndicesRequest(*kendra.ListIndicesInput) kendra.ListIndicesRequest

	QueryRequest(*kendra.QueryInput) kendra.QueryRequest

	StartDataSourceSyncJobRequest(*kendra.StartDataSourceSyncJobInput) kendra.StartDataSourceSyncJobRequest

	StopDataSourceSyncJobRequest(*kendra.StopDataSourceSyncJobInput) kendra.StopDataSourceSyncJobRequest

	SubmitFeedbackRequest(*kendra.SubmitFeedbackInput) kendra.SubmitFeedbackRequest

	UpdateDataSourceRequest(*kendra.UpdateDataSourceInput) kendra.UpdateDataSourceRequest

	UpdateIndexRequest(*kendra.UpdateIndexInput) kendra.UpdateIndexRequest
}

var _ ClientAPI = (*kendra.Client)(nil)
