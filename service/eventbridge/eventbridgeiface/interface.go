// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package eventbridgeiface provides an interface to enable mocking the Amazon EventBridge service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package eventbridgeiface

import (
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
)

// ClientAPI provides an interface to enable mocking the
// eventbridge.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon EventBridge.
//    func myFunc(svc eventbridgeiface.ClientAPI) bool {
//        // Make svc.ActivateEventSource request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := eventbridge.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        eventbridgeiface.ClientPI
//    }
//    func (m *mockClientClient) ActivateEventSource(input *types.ActivateEventSourceInput) (*types.ActivateEventSourceOutput, error) {
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
	ActivateEventSourceRequest(*types.ActivateEventSourceInput) eventbridge.ActivateEventSourceRequest

	CreateEventBusRequest(*types.CreateEventBusInput) eventbridge.CreateEventBusRequest

	CreatePartnerEventSourceRequest(*types.CreatePartnerEventSourceInput) eventbridge.CreatePartnerEventSourceRequest

	DeactivateEventSourceRequest(*types.DeactivateEventSourceInput) eventbridge.DeactivateEventSourceRequest

	DeleteEventBusRequest(*types.DeleteEventBusInput) eventbridge.DeleteEventBusRequest

	DeletePartnerEventSourceRequest(*types.DeletePartnerEventSourceInput) eventbridge.DeletePartnerEventSourceRequest

	DeleteRuleRequest(*types.DeleteRuleInput) eventbridge.DeleteRuleRequest

	DescribeEventBusRequest(*types.DescribeEventBusInput) eventbridge.DescribeEventBusRequest

	DescribeEventSourceRequest(*types.DescribeEventSourceInput) eventbridge.DescribeEventSourceRequest

	DescribePartnerEventSourceRequest(*types.DescribePartnerEventSourceInput) eventbridge.DescribePartnerEventSourceRequest

	DescribeRuleRequest(*types.DescribeRuleInput) eventbridge.DescribeRuleRequest

	DisableRuleRequest(*types.DisableRuleInput) eventbridge.DisableRuleRequest

	EnableRuleRequest(*types.EnableRuleInput) eventbridge.EnableRuleRequest

	ListEventBusesRequest(*types.ListEventBusesInput) eventbridge.ListEventBusesRequest

	ListEventSourcesRequest(*types.ListEventSourcesInput) eventbridge.ListEventSourcesRequest

	ListPartnerEventSourceAccountsRequest(*types.ListPartnerEventSourceAccountsInput) eventbridge.ListPartnerEventSourceAccountsRequest

	ListPartnerEventSourcesRequest(*types.ListPartnerEventSourcesInput) eventbridge.ListPartnerEventSourcesRequest

	ListRuleNamesByTargetRequest(*types.ListRuleNamesByTargetInput) eventbridge.ListRuleNamesByTargetRequest

	ListRulesRequest(*types.ListRulesInput) eventbridge.ListRulesRequest

	ListTagsForResourceRequest(*types.ListTagsForResourceInput) eventbridge.ListTagsForResourceRequest

	ListTargetsByRuleRequest(*types.ListTargetsByRuleInput) eventbridge.ListTargetsByRuleRequest

	PutEventsRequest(*types.PutEventsInput) eventbridge.PutEventsRequest

	PutPartnerEventsRequest(*types.PutPartnerEventsInput) eventbridge.PutPartnerEventsRequest

	PutPermissionRequest(*types.PutPermissionInput) eventbridge.PutPermissionRequest

	PutRuleRequest(*types.PutRuleInput) eventbridge.PutRuleRequest

	PutTargetsRequest(*types.PutTargetsInput) eventbridge.PutTargetsRequest

	RemovePermissionRequest(*types.RemovePermissionInput) eventbridge.RemovePermissionRequest

	RemoveTargetsRequest(*types.RemoveTargetsInput) eventbridge.RemoveTargetsRequest

	TagResourceRequest(*types.TagResourceInput) eventbridge.TagResourceRequest

	TestEventPatternRequest(*types.TestEventPatternInput) eventbridge.TestEventPatternRequest

	UntagResourceRequest(*types.UntagResourceInput) eventbridge.UntagResourceRequest
}

var _ ClientAPI = (*eventbridge.Client)(nil)
