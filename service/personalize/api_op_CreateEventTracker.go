// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
)

const opCreateEventTracker = "CreateEventTracker"

// CreateEventTrackerRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Creates an event tracker that you use when sending event data to the specified
// dataset group using the PutEvents (https://docs.aws.amazon.com/personalize/latest/dg/API_UBS_PutEvents.html)
// API.
//
// When Amazon Personalize creates an event tracker, it also creates an event-interactions
// dataset in the dataset group associated with the event tracker. The event-interactions
// dataset stores the event data from the PutEvents call. The contents of this
// dataset are not available to the user.
//
// Only one event tracker can be associated with a dataset group. You will get
// an error if you call CreateEventTracker using the same dataset group as an
// existing event tracker.
//
// When you send event data you include your tracking ID. The tracking ID identifies
// the customer and authorizes the customer to send the data.
//
// The event tracker can be in one of the following states:
//
//    * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
//    * DELETE PENDING > DELETE IN_PROGRESS
//
// To get the status of the event tracker, call DescribeEventTracker.
//
// The event tracker must be in the ACTIVE state before using the tracking ID.
//
// Related APIs
//
//    * ListEventTrackers
//
//    * DescribeEventTracker
//
//    * DeleteEventTracker
//
//    // Example sending a request using CreateEventTrackerRequest.
//    req := client.CreateEventTrackerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/CreateEventTracker
func (c *Client) CreateEventTrackerRequest(input *types.CreateEventTrackerInput) CreateEventTrackerRequest {
	op := &aws.Operation{
		Name:       opCreateEventTracker,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateEventTrackerInput{}
	}

	req := c.newRequest(op, input, &types.CreateEventTrackerOutput{})
	return CreateEventTrackerRequest{Request: req, Input: input, Copy: c.CreateEventTrackerRequest}
}

// CreateEventTrackerRequest is the request type for the
// CreateEventTracker API operation.
type CreateEventTrackerRequest struct {
	*aws.Request
	Input *types.CreateEventTrackerInput
	Copy  func(*types.CreateEventTrackerInput) CreateEventTrackerRequest
}

// Send marshals and sends the CreateEventTracker API request.
func (r CreateEventTrackerRequest) Send(ctx context.Context) (*CreateEventTrackerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateEventTrackerResponse{
		CreateEventTrackerOutput: r.Request.Data.(*types.CreateEventTrackerOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateEventTrackerResponse is the response type for the
// CreateEventTracker API operation.
type CreateEventTrackerResponse struct {
	*types.CreateEventTrackerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateEventTracker request.
func (r *CreateEventTrackerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
