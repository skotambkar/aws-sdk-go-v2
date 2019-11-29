// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/personalize/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
)

const opDescribeEventTracker = "DescribeEventTracker"

// DescribeEventTrackerRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Describes an event tracker. The response includes the trackingId and status
// of the event tracker. For more information on event trackers, see CreateEventTracker.
//
//    // Example sending a request using DescribeEventTrackerRequest.
//    req := client.DescribeEventTrackerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/DescribeEventTracker
func (c *Client) DescribeEventTrackerRequest(input *types.DescribeEventTrackerInput) DescribeEventTrackerRequest {
	op := &aws.Operation{
		Name:       opDescribeEventTracker,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeEventTrackerInput{}
	}

	req := c.newRequest(op, input, &types.DescribeEventTrackerOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeEventTrackerMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeEventTrackerRequest{Request: req, Input: input, Copy: c.DescribeEventTrackerRequest}
}

// DescribeEventTrackerRequest is the request type for the
// DescribeEventTracker API operation.
type DescribeEventTrackerRequest struct {
	*aws.Request
	Input *types.DescribeEventTrackerInput
	Copy  func(*types.DescribeEventTrackerInput) DescribeEventTrackerRequest
}

// Send marshals and sends the DescribeEventTracker API request.
func (r DescribeEventTrackerRequest) Send(ctx context.Context) (*DescribeEventTrackerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEventTrackerResponse{
		DescribeEventTrackerOutput: r.Request.Data.(*types.DescribeEventTrackerOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEventTrackerResponse is the response type for the
// DescribeEventTracker API operation.
type DescribeEventTrackerResponse struct {
	*types.DescribeEventTrackerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEventTracker request.
func (r *DescribeEventTrackerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
