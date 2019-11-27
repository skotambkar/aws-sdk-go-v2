// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package personalizeiface provides an interface to enable mocking the Amazon Personalize service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package personalizeiface

import (
	"github.com/aws/aws-sdk-go-v2/service/personalize"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
)

// ClientAPI provides an interface to enable mocking the
// personalize.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Personalize.
//    func myFunc(svc personalizeiface.ClientAPI) bool {
//        // Make svc.CreateCampaign request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := personalize.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        personalizeiface.ClientPI
//    }
//    func (m *mockClientClient) CreateCampaign(input *types.CreateCampaignInput) (*types.CreateCampaignOutput, error) {
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
	CreateCampaignRequest(*types.CreateCampaignInput) personalize.CreateCampaignRequest

	CreateDatasetRequest(*types.CreateDatasetInput) personalize.CreateDatasetRequest

	CreateDatasetGroupRequest(*types.CreateDatasetGroupInput) personalize.CreateDatasetGroupRequest

	CreateDatasetImportJobRequest(*types.CreateDatasetImportJobInput) personalize.CreateDatasetImportJobRequest

	CreateEventTrackerRequest(*types.CreateEventTrackerInput) personalize.CreateEventTrackerRequest

	CreateSchemaRequest(*types.CreateSchemaInput) personalize.CreateSchemaRequest

	CreateSolutionRequest(*types.CreateSolutionInput) personalize.CreateSolutionRequest

	CreateSolutionVersionRequest(*types.CreateSolutionVersionInput) personalize.CreateSolutionVersionRequest

	DeleteCampaignRequest(*types.DeleteCampaignInput) personalize.DeleteCampaignRequest

	DeleteDatasetRequest(*types.DeleteDatasetInput) personalize.DeleteDatasetRequest

	DeleteDatasetGroupRequest(*types.DeleteDatasetGroupInput) personalize.DeleteDatasetGroupRequest

	DeleteEventTrackerRequest(*types.DeleteEventTrackerInput) personalize.DeleteEventTrackerRequest

	DeleteSchemaRequest(*types.DeleteSchemaInput) personalize.DeleteSchemaRequest

	DeleteSolutionRequest(*types.DeleteSolutionInput) personalize.DeleteSolutionRequest

	DescribeAlgorithmRequest(*types.DescribeAlgorithmInput) personalize.DescribeAlgorithmRequest

	DescribeCampaignRequest(*types.DescribeCampaignInput) personalize.DescribeCampaignRequest

	DescribeDatasetRequest(*types.DescribeDatasetInput) personalize.DescribeDatasetRequest

	DescribeDatasetGroupRequest(*types.DescribeDatasetGroupInput) personalize.DescribeDatasetGroupRequest

	DescribeDatasetImportJobRequest(*types.DescribeDatasetImportJobInput) personalize.DescribeDatasetImportJobRequest

	DescribeEventTrackerRequest(*types.DescribeEventTrackerInput) personalize.DescribeEventTrackerRequest

	DescribeFeatureTransformationRequest(*types.DescribeFeatureTransformationInput) personalize.DescribeFeatureTransformationRequest

	DescribeRecipeRequest(*types.DescribeRecipeInput) personalize.DescribeRecipeRequest

	DescribeSchemaRequest(*types.DescribeSchemaInput) personalize.DescribeSchemaRequest

	DescribeSolutionRequest(*types.DescribeSolutionInput) personalize.DescribeSolutionRequest

	DescribeSolutionVersionRequest(*types.DescribeSolutionVersionInput) personalize.DescribeSolutionVersionRequest

	GetSolutionMetricsRequest(*types.GetSolutionMetricsInput) personalize.GetSolutionMetricsRequest

	ListCampaignsRequest(*types.ListCampaignsInput) personalize.ListCampaignsRequest

	ListDatasetGroupsRequest(*types.ListDatasetGroupsInput) personalize.ListDatasetGroupsRequest

	ListDatasetImportJobsRequest(*types.ListDatasetImportJobsInput) personalize.ListDatasetImportJobsRequest

	ListDatasetsRequest(*types.ListDatasetsInput) personalize.ListDatasetsRequest

	ListEventTrackersRequest(*types.ListEventTrackersInput) personalize.ListEventTrackersRequest

	ListRecipesRequest(*types.ListRecipesInput) personalize.ListRecipesRequest

	ListSchemasRequest(*types.ListSchemasInput) personalize.ListSchemasRequest

	ListSolutionVersionsRequest(*types.ListSolutionVersionsInput) personalize.ListSolutionVersionsRequest

	ListSolutionsRequest(*types.ListSolutionsInput) personalize.ListSolutionsRequest

	UpdateCampaignRequest(*types.UpdateCampaignInput) personalize.UpdateCampaignRequest
}

var _ ClientAPI = (*personalize.Client)(nil)
