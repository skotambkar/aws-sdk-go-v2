// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ecsiface provides an interface to enable mocking the Amazon EC2 Container Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package ecsiface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

// ClientAPI provides an interface to enable mocking the
// ecs.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon ECS.
//    func myFunc(svc ecsiface.ClientAPI) bool {
//        // Make svc.CreateCluster request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := ecs.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        ecsiface.ClientPI
//    }
//    func (m *mockClientClient) CreateCluster(input *types.CreateClusterInput) (*types.CreateClusterOutput, error) {
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
	CreateClusterRequest(*types.CreateClusterInput) ecs.CreateClusterRequest

	CreateServiceRequest(*types.CreateServiceInput) ecs.CreateServiceRequest

	CreateTaskSetRequest(*types.CreateTaskSetInput) ecs.CreateTaskSetRequest

	DeleteAccountSettingRequest(*types.DeleteAccountSettingInput) ecs.DeleteAccountSettingRequest

	DeleteAttributesRequest(*types.DeleteAttributesInput) ecs.DeleteAttributesRequest

	DeleteClusterRequest(*types.DeleteClusterInput) ecs.DeleteClusterRequest

	DeleteServiceRequest(*types.DeleteServiceInput) ecs.DeleteServiceRequest

	DeleteTaskSetRequest(*types.DeleteTaskSetInput) ecs.DeleteTaskSetRequest

	DeregisterContainerInstanceRequest(*types.DeregisterContainerInstanceInput) ecs.DeregisterContainerInstanceRequest

	DeregisterTaskDefinitionRequest(*types.DeregisterTaskDefinitionInput) ecs.DeregisterTaskDefinitionRequest

	DescribeClustersRequest(*types.DescribeClustersInput) ecs.DescribeClustersRequest

	DescribeContainerInstancesRequest(*types.DescribeContainerInstancesInput) ecs.DescribeContainerInstancesRequest

	DescribeServicesRequest(*types.DescribeServicesInput) ecs.DescribeServicesRequest

	DescribeTaskDefinitionRequest(*types.DescribeTaskDefinitionInput) ecs.DescribeTaskDefinitionRequest

	DescribeTaskSetsRequest(*types.DescribeTaskSetsInput) ecs.DescribeTaskSetsRequest

	DescribeTasksRequest(*types.DescribeTasksInput) ecs.DescribeTasksRequest

	DiscoverPollEndpointRequest(*types.DiscoverPollEndpointInput) ecs.DiscoverPollEndpointRequest

	ListAccountSettingsRequest(*types.ListAccountSettingsInput) ecs.ListAccountSettingsRequest

	ListAttributesRequest(*types.ListAttributesInput) ecs.ListAttributesRequest

	ListClustersRequest(*types.ListClustersInput) ecs.ListClustersRequest

	ListContainerInstancesRequest(*types.ListContainerInstancesInput) ecs.ListContainerInstancesRequest

	ListServicesRequest(*types.ListServicesInput) ecs.ListServicesRequest

	ListTagsForResourceRequest(*types.ListTagsForResourceInput) ecs.ListTagsForResourceRequest

	ListTaskDefinitionFamiliesRequest(*types.ListTaskDefinitionFamiliesInput) ecs.ListTaskDefinitionFamiliesRequest

	ListTaskDefinitionsRequest(*types.ListTaskDefinitionsInput) ecs.ListTaskDefinitionsRequest

	ListTasksRequest(*types.ListTasksInput) ecs.ListTasksRequest

	PutAccountSettingRequest(*types.PutAccountSettingInput) ecs.PutAccountSettingRequest

	PutAccountSettingDefaultRequest(*types.PutAccountSettingDefaultInput) ecs.PutAccountSettingDefaultRequest

	PutAttributesRequest(*types.PutAttributesInput) ecs.PutAttributesRequest

	RegisterContainerInstanceRequest(*types.RegisterContainerInstanceInput) ecs.RegisterContainerInstanceRequest

	RegisterTaskDefinitionRequest(*types.RegisterTaskDefinitionInput) ecs.RegisterTaskDefinitionRequest

	RunTaskRequest(*types.RunTaskInput) ecs.RunTaskRequest

	StartTaskRequest(*types.StartTaskInput) ecs.StartTaskRequest

	StopTaskRequest(*types.StopTaskInput) ecs.StopTaskRequest

	SubmitAttachmentStateChangesRequest(*types.SubmitAttachmentStateChangesInput) ecs.SubmitAttachmentStateChangesRequest

	SubmitContainerStateChangeRequest(*types.SubmitContainerStateChangeInput) ecs.SubmitContainerStateChangeRequest

	SubmitTaskStateChangeRequest(*types.SubmitTaskStateChangeInput) ecs.SubmitTaskStateChangeRequest

	TagResourceRequest(*types.TagResourceInput) ecs.TagResourceRequest

	UntagResourceRequest(*types.UntagResourceInput) ecs.UntagResourceRequest

	UpdateClusterSettingsRequest(*types.UpdateClusterSettingsInput) ecs.UpdateClusterSettingsRequest

	UpdateContainerAgentRequest(*types.UpdateContainerAgentInput) ecs.UpdateContainerAgentRequest

	UpdateContainerInstancesStateRequest(*types.UpdateContainerInstancesStateInput) ecs.UpdateContainerInstancesStateRequest

	UpdateServiceRequest(*types.UpdateServiceInput) ecs.UpdateServiceRequest

	UpdateServicePrimaryTaskSetRequest(*types.UpdateServicePrimaryTaskSetInput) ecs.UpdateServicePrimaryTaskSetRequest

	UpdateTaskSetRequest(*types.UpdateTaskSetInput) ecs.UpdateTaskSetRequest

	WaitUntilServicesInactive(context.Context, *types.DescribeServicesInput, ...aws.WaiterOption) error

	WaitUntilServicesStable(context.Context, *types.DescribeServicesInput, ...aws.WaiterOption) error

	WaitUntilTasksRunning(context.Context, *types.DescribeTasksInput, ...aws.WaiterOption) error

	WaitUntilTasksStopped(context.Context, *types.DescribeTasksInput, ...aws.WaiterOption) error
}

var _ ClientAPI = (*ecs.Client)(nil)
