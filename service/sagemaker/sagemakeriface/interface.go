// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sagemakeriface provides an interface to enable mocking the Amazon SageMaker Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package sagemakeriface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
)

// ClientAPI provides an interface to enable mocking the
// sagemaker.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // SageMaker.
//    func myFunc(svc sagemakeriface.ClientAPI) bool {
//        // Make svc.AddTags request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := sagemaker.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        sagemakeriface.ClientPI
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
	AddTagsRequest(*types.AddTagsInput) sagemaker.AddTagsRequest

	CreateAlgorithmRequest(*types.CreateAlgorithmInput) sagemaker.CreateAlgorithmRequest

	CreateCodeRepositoryRequest(*types.CreateCodeRepositoryInput) sagemaker.CreateCodeRepositoryRequest

	CreateCompilationJobRequest(*types.CreateCompilationJobInput) sagemaker.CreateCompilationJobRequest

	CreateEndpointRequest(*types.CreateEndpointInput) sagemaker.CreateEndpointRequest

	CreateEndpointConfigRequest(*types.CreateEndpointConfigInput) sagemaker.CreateEndpointConfigRequest

	CreateHyperParameterTuningJobRequest(*types.CreateHyperParameterTuningJobInput) sagemaker.CreateHyperParameterTuningJobRequest

	CreateLabelingJobRequest(*types.CreateLabelingJobInput) sagemaker.CreateLabelingJobRequest

	CreateModelRequest(*types.CreateModelInput) sagemaker.CreateModelRequest

	CreateModelPackageRequest(*types.CreateModelPackageInput) sagemaker.CreateModelPackageRequest

	CreateNotebookInstanceRequest(*types.CreateNotebookInstanceInput) sagemaker.CreateNotebookInstanceRequest

	CreateNotebookInstanceLifecycleConfigRequest(*types.CreateNotebookInstanceLifecycleConfigInput) sagemaker.CreateNotebookInstanceLifecycleConfigRequest

	CreatePresignedNotebookInstanceUrlRequest(*types.CreatePresignedNotebookInstanceUrlInput) sagemaker.CreatePresignedNotebookInstanceUrlRequest

	CreateTrainingJobRequest(*types.CreateTrainingJobInput) sagemaker.CreateTrainingJobRequest

	CreateTransformJobRequest(*types.CreateTransformJobInput) sagemaker.CreateTransformJobRequest

	CreateWorkteamRequest(*types.CreateWorkteamInput) sagemaker.CreateWorkteamRequest

	DeleteAlgorithmRequest(*types.DeleteAlgorithmInput) sagemaker.DeleteAlgorithmRequest

	DeleteCodeRepositoryRequest(*types.DeleteCodeRepositoryInput) sagemaker.DeleteCodeRepositoryRequest

	DeleteEndpointRequest(*types.DeleteEndpointInput) sagemaker.DeleteEndpointRequest

	DeleteEndpointConfigRequest(*types.DeleteEndpointConfigInput) sagemaker.DeleteEndpointConfigRequest

	DeleteModelRequest(*types.DeleteModelInput) sagemaker.DeleteModelRequest

	DeleteModelPackageRequest(*types.DeleteModelPackageInput) sagemaker.DeleteModelPackageRequest

	DeleteNotebookInstanceRequest(*types.DeleteNotebookInstanceInput) sagemaker.DeleteNotebookInstanceRequest

	DeleteNotebookInstanceLifecycleConfigRequest(*types.DeleteNotebookInstanceLifecycleConfigInput) sagemaker.DeleteNotebookInstanceLifecycleConfigRequest

	DeleteTagsRequest(*types.DeleteTagsInput) sagemaker.DeleteTagsRequest

	DeleteWorkteamRequest(*types.DeleteWorkteamInput) sagemaker.DeleteWorkteamRequest

	DescribeAlgorithmRequest(*types.DescribeAlgorithmInput) sagemaker.DescribeAlgorithmRequest

	DescribeCodeRepositoryRequest(*types.DescribeCodeRepositoryInput) sagemaker.DescribeCodeRepositoryRequest

	DescribeCompilationJobRequest(*types.DescribeCompilationJobInput) sagemaker.DescribeCompilationJobRequest

	DescribeEndpointRequest(*types.DescribeEndpointInput) sagemaker.DescribeEndpointRequest

	DescribeEndpointConfigRequest(*types.DescribeEndpointConfigInput) sagemaker.DescribeEndpointConfigRequest

	DescribeHyperParameterTuningJobRequest(*types.DescribeHyperParameterTuningJobInput) sagemaker.DescribeHyperParameterTuningJobRequest

	DescribeLabelingJobRequest(*types.DescribeLabelingJobInput) sagemaker.DescribeLabelingJobRequest

	DescribeModelRequest(*types.DescribeModelInput) sagemaker.DescribeModelRequest

	DescribeModelPackageRequest(*types.DescribeModelPackageInput) sagemaker.DescribeModelPackageRequest

	DescribeNotebookInstanceRequest(*types.DescribeNotebookInstanceInput) sagemaker.DescribeNotebookInstanceRequest

	DescribeNotebookInstanceLifecycleConfigRequest(*types.DescribeNotebookInstanceLifecycleConfigInput) sagemaker.DescribeNotebookInstanceLifecycleConfigRequest

	DescribeSubscribedWorkteamRequest(*types.DescribeSubscribedWorkteamInput) sagemaker.DescribeSubscribedWorkteamRequest

	DescribeTrainingJobRequest(*types.DescribeTrainingJobInput) sagemaker.DescribeTrainingJobRequest

	DescribeTransformJobRequest(*types.DescribeTransformJobInput) sagemaker.DescribeTransformJobRequest

	DescribeWorkteamRequest(*types.DescribeWorkteamInput) sagemaker.DescribeWorkteamRequest

	GetSearchSuggestionsRequest(*types.GetSearchSuggestionsInput) sagemaker.GetSearchSuggestionsRequest

	ListAlgorithmsRequest(*types.ListAlgorithmsInput) sagemaker.ListAlgorithmsRequest

	ListCodeRepositoriesRequest(*types.ListCodeRepositoriesInput) sagemaker.ListCodeRepositoriesRequest

	ListCompilationJobsRequest(*types.ListCompilationJobsInput) sagemaker.ListCompilationJobsRequest

	ListEndpointConfigsRequest(*types.ListEndpointConfigsInput) sagemaker.ListEndpointConfigsRequest

	ListEndpointsRequest(*types.ListEndpointsInput) sagemaker.ListEndpointsRequest

	ListHyperParameterTuningJobsRequest(*types.ListHyperParameterTuningJobsInput) sagemaker.ListHyperParameterTuningJobsRequest

	ListLabelingJobsRequest(*types.ListLabelingJobsInput) sagemaker.ListLabelingJobsRequest

	ListLabelingJobsForWorkteamRequest(*types.ListLabelingJobsForWorkteamInput) sagemaker.ListLabelingJobsForWorkteamRequest

	ListModelPackagesRequest(*types.ListModelPackagesInput) sagemaker.ListModelPackagesRequest

	ListModelsRequest(*types.ListModelsInput) sagemaker.ListModelsRequest

	ListNotebookInstanceLifecycleConfigsRequest(*types.ListNotebookInstanceLifecycleConfigsInput) sagemaker.ListNotebookInstanceLifecycleConfigsRequest

	ListNotebookInstancesRequest(*types.ListNotebookInstancesInput) sagemaker.ListNotebookInstancesRequest

	ListSubscribedWorkteamsRequest(*types.ListSubscribedWorkteamsInput) sagemaker.ListSubscribedWorkteamsRequest

	ListTagsRequest(*types.ListTagsInput) sagemaker.ListTagsRequest

	ListTrainingJobsRequest(*types.ListTrainingJobsInput) sagemaker.ListTrainingJobsRequest

	ListTrainingJobsForHyperParameterTuningJobRequest(*types.ListTrainingJobsForHyperParameterTuningJobInput) sagemaker.ListTrainingJobsForHyperParameterTuningJobRequest

	ListTransformJobsRequest(*types.ListTransformJobsInput) sagemaker.ListTransformJobsRequest

	ListWorkteamsRequest(*types.ListWorkteamsInput) sagemaker.ListWorkteamsRequest

	RenderUiTemplateRequest(*types.RenderUiTemplateInput) sagemaker.RenderUiTemplateRequest

	SearchRequest(*types.SearchInput) sagemaker.SearchRequest

	StartNotebookInstanceRequest(*types.StartNotebookInstanceInput) sagemaker.StartNotebookInstanceRequest

	StopCompilationJobRequest(*types.StopCompilationJobInput) sagemaker.StopCompilationJobRequest

	StopHyperParameterTuningJobRequest(*types.StopHyperParameterTuningJobInput) sagemaker.StopHyperParameterTuningJobRequest

	StopLabelingJobRequest(*types.StopLabelingJobInput) sagemaker.StopLabelingJobRequest

	StopNotebookInstanceRequest(*types.StopNotebookInstanceInput) sagemaker.StopNotebookInstanceRequest

	StopTrainingJobRequest(*types.StopTrainingJobInput) sagemaker.StopTrainingJobRequest

	StopTransformJobRequest(*types.StopTransformJobInput) sagemaker.StopTransformJobRequest

	UpdateCodeRepositoryRequest(*types.UpdateCodeRepositoryInput) sagemaker.UpdateCodeRepositoryRequest

	UpdateEndpointRequest(*types.UpdateEndpointInput) sagemaker.UpdateEndpointRequest

	UpdateEndpointWeightsAndCapacitiesRequest(*types.UpdateEndpointWeightsAndCapacitiesInput) sagemaker.UpdateEndpointWeightsAndCapacitiesRequest

	UpdateNotebookInstanceRequest(*types.UpdateNotebookInstanceInput) sagemaker.UpdateNotebookInstanceRequest

	UpdateNotebookInstanceLifecycleConfigRequest(*types.UpdateNotebookInstanceLifecycleConfigInput) sagemaker.UpdateNotebookInstanceLifecycleConfigRequest

	UpdateWorkteamRequest(*types.UpdateWorkteamInput) sagemaker.UpdateWorkteamRequest

	WaitUntilEndpointDeleted(context.Context, *types.DescribeEndpointInput, ...aws.WaiterOption) error

	WaitUntilEndpointInService(context.Context, *types.DescribeEndpointInput, ...aws.WaiterOption) error

	WaitUntilNotebookInstanceDeleted(context.Context, *types.DescribeNotebookInstanceInput, ...aws.WaiterOption) error

	WaitUntilNotebookInstanceInService(context.Context, *types.DescribeNotebookInstanceInput, ...aws.WaiterOption) error

	WaitUntilNotebookInstanceStopped(context.Context, *types.DescribeNotebookInstanceInput, ...aws.WaiterOption) error

	WaitUntilTrainingJobCompletedOrStopped(context.Context, *types.DescribeTrainingJobInput, ...aws.WaiterOption) error

	WaitUntilTransformJobCompletedOrStopped(context.Context, *types.DescribeTransformJobInput, ...aws.WaiterOption) error
}

var _ ClientAPI = (*sagemaker.Client)(nil)
