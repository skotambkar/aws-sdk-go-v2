// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opDeleteThing = "DeleteThing"

// DeleteThingRequest returns a request value for making API operation for
// AWS IoT.
//
// Deletes the specified thing. Returns successfully with no error if the deletion
// is successful or you specify a thing that doesn't exist.
//
//    // Example sending a request using DeleteThingRequest.
//    req := client.DeleteThingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteThingRequest(input *types.DeleteThingInput) DeleteThingRequest {
	op := &aws.Operation{
		Name:       opDeleteThing,
		HTTPMethod: "DELETE",
		HTTPPath:   "/things/{thingName}",
	}

	if input == nil {
		input = &types.DeleteThingInput{}
	}

	req := c.newRequest(op, input, &types.DeleteThingOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DeleteThingMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteThingRequest{Request: req, Input: input, Copy: c.DeleteThingRequest}
}

// DeleteThingRequest is the request type for the
// DeleteThing API operation.
type DeleteThingRequest struct {
	*aws.Request
	Input *types.DeleteThingInput
	Copy  func(*types.DeleteThingInput) DeleteThingRequest
}

// Send marshals and sends the DeleteThing API request.
func (r DeleteThingRequest) Send(ctx context.Context) (*DeleteThingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteThingResponse{
		DeleteThingOutput: r.Request.Data.(*types.DeleteThingOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteThingResponse is the response type for the
// DeleteThing API operation.
type DeleteThingResponse struct {
	*types.DeleteThingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteThing request.
func (r *DeleteThingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
