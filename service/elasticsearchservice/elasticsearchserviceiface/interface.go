// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package elasticsearchserviceiface provides an interface to enable mocking the Amazon Elasticsearch Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package elasticsearchserviceiface

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
)

// ClientAPI provides an interface to enable mocking the
// elasticsearchservice.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Elasticsearch Service.
//    func myFunc(svc elasticsearchserviceiface.ClientAPI) bool {
//        // Make svc.AddTags request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := elasticsearchservice.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        elasticsearchserviceiface.ClientPI
//    }
//    func (m *mockClientClient) AddTags(input *types.AddTagsInput) (*types.AddTagsOutput, error) {
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
	AddTagsRequest(*types.AddTagsInput) elasticsearchservice.AddTagsRequest

	CancelElasticsearchServiceSoftwareUpdateRequest(*types.CancelElasticsearchServiceSoftwareUpdateInput) elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateRequest

	CreateElasticsearchDomainRequest(*types.CreateElasticsearchDomainInput) elasticsearchservice.CreateElasticsearchDomainRequest

	DeleteElasticsearchDomainRequest(*types.DeleteElasticsearchDomainInput) elasticsearchservice.DeleteElasticsearchDomainRequest

	DeleteElasticsearchServiceRoleRequest(*types.DeleteElasticsearchServiceRoleInput) elasticsearchservice.DeleteElasticsearchServiceRoleRequest

	DescribeElasticsearchDomainRequest(*types.DescribeElasticsearchDomainInput) elasticsearchservice.DescribeElasticsearchDomainRequest

	DescribeElasticsearchDomainConfigRequest(*types.DescribeElasticsearchDomainConfigInput) elasticsearchservice.DescribeElasticsearchDomainConfigRequest

	DescribeElasticsearchDomainsRequest(*types.DescribeElasticsearchDomainsInput) elasticsearchservice.DescribeElasticsearchDomainsRequest

	DescribeElasticsearchInstanceTypeLimitsRequest(*types.DescribeElasticsearchInstanceTypeLimitsInput) elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsRequest

	DescribeReservedElasticsearchInstanceOfferingsRequest(*types.DescribeReservedElasticsearchInstanceOfferingsInput) elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsRequest

	DescribeReservedElasticsearchInstancesRequest(*types.DescribeReservedElasticsearchInstancesInput) elasticsearchservice.DescribeReservedElasticsearchInstancesRequest

	GetCompatibleElasticsearchVersionsRequest(*types.GetCompatibleElasticsearchVersionsInput) elasticsearchservice.GetCompatibleElasticsearchVersionsRequest

	GetUpgradeHistoryRequest(*types.GetUpgradeHistoryInput) elasticsearchservice.GetUpgradeHistoryRequest

	GetUpgradeStatusRequest(*types.GetUpgradeStatusInput) elasticsearchservice.GetUpgradeStatusRequest

	ListDomainNamesRequest(*types.ListDomainNamesInput) elasticsearchservice.ListDomainNamesRequest

	ListElasticsearchInstanceTypesRequest(*types.ListElasticsearchInstanceTypesInput) elasticsearchservice.ListElasticsearchInstanceTypesRequest

	ListElasticsearchVersionsRequest(*types.ListElasticsearchVersionsInput) elasticsearchservice.ListElasticsearchVersionsRequest

	ListTagsRequest(*types.ListTagsInput) elasticsearchservice.ListTagsRequest

	PurchaseReservedElasticsearchInstanceOfferingRequest(*types.PurchaseReservedElasticsearchInstanceOfferingInput) elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingRequest

	RemoveTagsRequest(*types.RemoveTagsInput) elasticsearchservice.RemoveTagsRequest

	StartElasticsearchServiceSoftwareUpdateRequest(*types.StartElasticsearchServiceSoftwareUpdateInput) elasticsearchservice.StartElasticsearchServiceSoftwareUpdateRequest

	UpdateElasticsearchDomainConfigRequest(*types.UpdateElasticsearchDomainConfigInput) elasticsearchservice.UpdateElasticsearchDomainConfigRequest

	UpgradeElasticsearchDomainRequest(*types.UpgradeElasticsearchDomainInput) elasticsearchservice.UpgradeElasticsearchDomainRequest
}

var _ ClientAPI = (*elasticsearchservice.Client)(nil)
