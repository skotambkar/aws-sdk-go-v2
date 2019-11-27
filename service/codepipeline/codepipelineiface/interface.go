// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codepipelineiface provides an interface to enable mocking the AWS CodePipeline service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codepipelineiface

import (
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
)

// ClientAPI provides an interface to enable mocking the
// codepipeline.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // CodePipeline.
//    func myFunc(svc codepipelineiface.ClientAPI) bool {
//        // Make svc.AcknowledgeJob request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := codepipeline.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        codepipelineiface.ClientPI
//    }
//    func (m *mockClientClient) AcknowledgeJob(input *types.AcknowledgeJobInput) (*types.AcknowledgeJobOutput, error) {
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
	AcknowledgeJobRequest(*types.AcknowledgeJobInput) codepipeline.AcknowledgeJobRequest

	AcknowledgeThirdPartyJobRequest(*types.AcknowledgeThirdPartyJobInput) codepipeline.AcknowledgeThirdPartyJobRequest

	CreateCustomActionTypeRequest(*types.CreateCustomActionTypeInput) codepipeline.CreateCustomActionTypeRequest

	CreatePipelineRequest(*types.CreatePipelineInput) codepipeline.CreatePipelineRequest

	DeleteCustomActionTypeRequest(*types.DeleteCustomActionTypeInput) codepipeline.DeleteCustomActionTypeRequest

	DeletePipelineRequest(*types.DeletePipelineInput) codepipeline.DeletePipelineRequest

	DeleteWebhookRequest(*types.DeleteWebhookInput) codepipeline.DeleteWebhookRequest

	DeregisterWebhookWithThirdPartyRequest(*types.DeregisterWebhookWithThirdPartyInput) codepipeline.DeregisterWebhookWithThirdPartyRequest

	DisableStageTransitionRequest(*types.DisableStageTransitionInput) codepipeline.DisableStageTransitionRequest

	EnableStageTransitionRequest(*types.EnableStageTransitionInput) codepipeline.EnableStageTransitionRequest

	GetJobDetailsRequest(*types.GetJobDetailsInput) codepipeline.GetJobDetailsRequest

	GetPipelineRequest(*types.GetPipelineInput) codepipeline.GetPipelineRequest

	GetPipelineExecutionRequest(*types.GetPipelineExecutionInput) codepipeline.GetPipelineExecutionRequest

	GetPipelineStateRequest(*types.GetPipelineStateInput) codepipeline.GetPipelineStateRequest

	GetThirdPartyJobDetailsRequest(*types.GetThirdPartyJobDetailsInput) codepipeline.GetThirdPartyJobDetailsRequest

	ListActionExecutionsRequest(*types.ListActionExecutionsInput) codepipeline.ListActionExecutionsRequest

	ListActionTypesRequest(*types.ListActionTypesInput) codepipeline.ListActionTypesRequest

	ListPipelineExecutionsRequest(*types.ListPipelineExecutionsInput) codepipeline.ListPipelineExecutionsRequest

	ListPipelinesRequest(*types.ListPipelinesInput) codepipeline.ListPipelinesRequest

	ListTagsForResourceRequest(*types.ListTagsForResourceInput) codepipeline.ListTagsForResourceRequest

	ListWebhooksRequest(*types.ListWebhooksInput) codepipeline.ListWebhooksRequest

	PollForJobsRequest(*types.PollForJobsInput) codepipeline.PollForJobsRequest

	PollForThirdPartyJobsRequest(*types.PollForThirdPartyJobsInput) codepipeline.PollForThirdPartyJobsRequest

	PutActionRevisionRequest(*types.PutActionRevisionInput) codepipeline.PutActionRevisionRequest

	PutApprovalResultRequest(*types.PutApprovalResultInput) codepipeline.PutApprovalResultRequest

	PutJobFailureResultRequest(*types.PutJobFailureResultInput) codepipeline.PutJobFailureResultRequest

	PutJobSuccessResultRequest(*types.PutJobSuccessResultInput) codepipeline.PutJobSuccessResultRequest

	PutThirdPartyJobFailureResultRequest(*types.PutThirdPartyJobFailureResultInput) codepipeline.PutThirdPartyJobFailureResultRequest

	PutThirdPartyJobSuccessResultRequest(*types.PutThirdPartyJobSuccessResultInput) codepipeline.PutThirdPartyJobSuccessResultRequest

	PutWebhookRequest(*types.PutWebhookInput) codepipeline.PutWebhookRequest

	RegisterWebhookWithThirdPartyRequest(*types.RegisterWebhookWithThirdPartyInput) codepipeline.RegisterWebhookWithThirdPartyRequest

	RetryStageExecutionRequest(*types.RetryStageExecutionInput) codepipeline.RetryStageExecutionRequest

	StartPipelineExecutionRequest(*types.StartPipelineExecutionInput) codepipeline.StartPipelineExecutionRequest

	TagResourceRequest(*types.TagResourceInput) codepipeline.TagResourceRequest

	UntagResourceRequest(*types.UntagResourceInput) codepipeline.UntagResourceRequest

	UpdatePipelineRequest(*types.UpdatePipelineInput) codepipeline.UpdatePipelineRequest
}

var _ ClientAPI = (*codepipeline.Client)(nil)
