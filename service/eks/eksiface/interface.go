// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package eksiface provides an interface to enable mocking the Amazon Elastic Kubernetes Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package eksiface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

// ClientAPI provides an interface to enable mocking the
// eks.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon EKS.
//    func myFunc(svc eksiface.ClientAPI) bool {
//        // Make svc.CreateCluster request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := eks.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        eksiface.ClientPI
//    }
//    func (m *mockClientClient) CreateCluster(input *eks.CreateClusterInput) (*eks.CreateClusterOutput, error) {
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
	CreateClusterRequest(*eks.CreateClusterInput) eks.CreateClusterRequest

	CreateNodegroupRequest(*eks.CreateNodegroupInput) eks.CreateNodegroupRequest

	DeleteClusterRequest(*eks.DeleteClusterInput) eks.DeleteClusterRequest

	DeleteNodegroupRequest(*eks.DeleteNodegroupInput) eks.DeleteNodegroupRequest

	DescribeClusterRequest(*eks.DescribeClusterInput) eks.DescribeClusterRequest

	DescribeNodegroupRequest(*eks.DescribeNodegroupInput) eks.DescribeNodegroupRequest

	DescribeUpdateRequest(*eks.DescribeUpdateInput) eks.DescribeUpdateRequest

	ListClustersRequest(*eks.ListClustersInput) eks.ListClustersRequest

	ListNodegroupsRequest(*eks.ListNodegroupsInput) eks.ListNodegroupsRequest

	ListTagsForResourceRequest(*eks.ListTagsForResourceInput) eks.ListTagsForResourceRequest

	ListUpdatesRequest(*eks.ListUpdatesInput) eks.ListUpdatesRequest

	TagResourceRequest(*eks.TagResourceInput) eks.TagResourceRequest

	UntagResourceRequest(*eks.UntagResourceInput) eks.UntagResourceRequest

	UpdateClusterConfigRequest(*eks.UpdateClusterConfigInput) eks.UpdateClusterConfigRequest

	UpdateClusterVersionRequest(*eks.UpdateClusterVersionInput) eks.UpdateClusterVersionRequest

	UpdateNodegroupConfigRequest(*eks.UpdateNodegroupConfigInput) eks.UpdateNodegroupConfigRequest

	UpdateNodegroupVersionRequest(*eks.UpdateNodegroupVersionInput) eks.UpdateNodegroupVersionRequest

	WaitUntilClusterActive(context.Context, *eks.DescribeClusterInput, ...aws.WaiterOption) error

	WaitUntilClusterDeleted(context.Context, *eks.DescribeClusterInput, ...aws.WaiterOption) error

	WaitUntilNodegroupActive(context.Context, *eks.DescribeNodegroupInput, ...aws.WaiterOption) error

	WaitUntilNodegroupDeleted(context.Context, *eks.DescribeNodegroupInput, ...aws.WaiterOption) error
}

var _ ClientAPI = (*eks.Client)(nil)
