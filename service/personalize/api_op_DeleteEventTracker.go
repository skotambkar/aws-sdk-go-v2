// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/personalize/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
)

const opDeleteEventTracker = "DeleteEventTracker"

// DeleteEventTrackerRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Deletes the event tracker. Does not delete the event-interactions dataset
// from the associated dataset group. For more information on event trackers,
// see CreateEventTracker.
//
//    // Example sending a request using DeleteEventTrackerRequest.
//    req := client.DeleteEventTrackerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/DeleteEventTracker
func (c *Client) DeleteEventTrackerRequest(input *types.DeleteEventTrackerInput) DeleteEventTrackerRequest {
	op := &aws.Operation{
		Name:       opDeleteEventTracker,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteEventTrackerInput{}
	}

	req := c.newRequest(op, input, &types.DeleteEventTrackerOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteEventTrackerMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteEventTrackerRequest{Request: req, Input: input, Copy: c.DeleteEventTrackerRequest}
}

// DeleteEventTrackerRequest is the request type for the
// DeleteEventTracker API operation.
type DeleteEventTrackerRequest struct {
	*aws.Request
	Input *types.DeleteEventTrackerInput
	Copy  func(*types.DeleteEventTrackerInput) DeleteEventTrackerRequest
}

// Send marshals and sends the DeleteEventTracker API request.
func (r DeleteEventTrackerRequest) Send(ctx context.Context) (*DeleteEventTrackerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteEventTrackerResponse{
		DeleteEventTrackerOutput: r.Request.Data.(*types.DeleteEventTrackerOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteEventTrackerResponse is the response type for the
// DeleteEventTracker API operation.
type DeleteEventTrackerResponse struct {
	*types.DeleteEventTrackerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteEventTracker request.
func (r *DeleteEventTrackerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
