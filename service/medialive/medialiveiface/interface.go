// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package medialiveiface provides an interface to enable mocking the AWS Elemental MediaLive service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package medialiveiface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/medialive"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
)

// ClientAPI provides an interface to enable mocking the
// medialive.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // MediaLive.
//    func myFunc(svc medialiveiface.ClientAPI) bool {
//        // Make svc.BatchUpdateSchedule request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := medialive.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        medialiveiface.ClientPI
//    }
//    func (m *mockClientClient) BatchUpdateSchedule(input *types.BatchUpdateScheduleInput) (*types.BatchUpdateScheduleOutput, error) {
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
	BatchUpdateScheduleRequest(*types.BatchUpdateScheduleInput) medialive.BatchUpdateScheduleRequest

	CreateChannelRequest(*types.CreateChannelInput) medialive.CreateChannelRequest

	CreateInputRequest(*types.CreateInputInput) medialive.CreateInputRequest

	CreateInputSecurityGroupRequest(*types.CreateInputSecurityGroupInput) medialive.CreateInputSecurityGroupRequest

	CreateTagsRequest(*types.CreateTagsInput) medialive.CreateTagsRequest

	DeleteChannelRequest(*types.DeleteChannelInput) medialive.DeleteChannelRequest

	DeleteInputRequest(*types.DeleteInputInput) medialive.DeleteInputRequest

	DeleteInputSecurityGroupRequest(*types.DeleteInputSecurityGroupInput) medialive.DeleteInputSecurityGroupRequest

	DeleteReservationRequest(*types.DeleteReservationInput) medialive.DeleteReservationRequest

	DeleteScheduleRequest(*types.DeleteScheduleInput) medialive.DeleteScheduleRequest

	DeleteTagsRequest(*types.DeleteTagsInput) medialive.DeleteTagsRequest

	DescribeChannelRequest(*types.DescribeChannelInput) medialive.DescribeChannelRequest

	DescribeInputRequest(*types.DescribeInputInput) medialive.DescribeInputRequest

	DescribeInputSecurityGroupRequest(*types.DescribeInputSecurityGroupInput) medialive.DescribeInputSecurityGroupRequest

	DescribeOfferingRequest(*types.DescribeOfferingInput) medialive.DescribeOfferingRequest

	DescribeReservationRequest(*types.DescribeReservationInput) medialive.DescribeReservationRequest

	DescribeScheduleRequest(*types.DescribeScheduleInput) medialive.DescribeScheduleRequest

	ListChannelsRequest(*types.ListChannelsInput) medialive.ListChannelsRequest

	ListInputSecurityGroupsRequest(*types.ListInputSecurityGroupsInput) medialive.ListInputSecurityGroupsRequest

	ListInputsRequest(*types.ListInputsInput) medialive.ListInputsRequest

	ListOfferingsRequest(*types.ListOfferingsInput) medialive.ListOfferingsRequest

	ListReservationsRequest(*types.ListReservationsInput) medialive.ListReservationsRequest

	ListTagsForResourceRequest(*types.ListTagsForResourceInput) medialive.ListTagsForResourceRequest

	PurchaseOfferingRequest(*types.PurchaseOfferingInput) medialive.PurchaseOfferingRequest

	StartChannelRequest(*types.StartChannelInput) medialive.StartChannelRequest

	StopChannelRequest(*types.StopChannelInput) medialive.StopChannelRequest

	UpdateChannelRequest(*types.UpdateChannelInput) medialive.UpdateChannelRequest

	UpdateChannelClassRequest(*types.UpdateChannelClassInput) medialive.UpdateChannelClassRequest

	UpdateInputRequest(*types.UpdateInputInput) medialive.UpdateInputRequest

	UpdateInputSecurityGroupRequest(*types.UpdateInputSecurityGroupInput) medialive.UpdateInputSecurityGroupRequest

	UpdateReservationRequest(*types.UpdateReservationInput) medialive.UpdateReservationRequest

	WaitUntilChannelCreated(context.Context, *types.DescribeChannelInput, ...aws.WaiterOption) error

	WaitUntilChannelDeleted(context.Context, *types.DescribeChannelInput, ...aws.WaiterOption) error

	WaitUntilChannelRunning(context.Context, *types.DescribeChannelInput, ...aws.WaiterOption) error

	WaitUntilChannelStopped(context.Context, *types.DescribeChannelInput, ...aws.WaiterOption) error
}

var _ ClientAPI = (*medialive.Client)(nil)
