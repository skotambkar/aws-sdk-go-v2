// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudtrail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
)

const opStopLogging = "StopLogging"

// StopLoggingRequest returns a request value for making API operation for
// AWS CloudTrail.
//
// Suspends the recording of AWS API calls and log file delivery for the specified
// trail. Under most circumstances, there is no need to use this action. You
// can update a trail without stopping it first. This action is the only way
// to stop recording. For a trail enabled in all regions, this operation must
// be called from the region in which the trail was created, or an InvalidHomeRegionException
// will occur. This operation cannot be called on the shadow trails (replicated
// trails in other regions) of a trail enabled in all regions.
//
//    // Example sending a request using StopLoggingRequest.
//    req := client.StopLoggingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudtrail-2013-11-01/StopLogging
func (c *Client) StopLoggingRequest(input *types.StopLoggingInput) StopLoggingRequest {
	op := &aws.Operation{
		Name:       opStopLogging,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StopLoggingInput{}
	}

	req := c.newRequest(op, input, &types.StopLoggingOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.StopLoggingMarshaler{Input: input}.GetNamedBuildHandler())

	return StopLoggingRequest{Request: req, Input: input, Copy: c.StopLoggingRequest}
}

// StopLoggingRequest is the request type for the
// StopLogging API operation.
type StopLoggingRequest struct {
	*aws.Request
	Input *types.StopLoggingInput
	Copy  func(*types.StopLoggingInput) StopLoggingRequest
}

// Send marshals and sends the StopLogging API request.
func (r StopLoggingRequest) Send(ctx context.Context) (*StopLoggingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopLoggingResponse{
		StopLoggingOutput: r.Request.Data.(*types.StopLoggingOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopLoggingResponse is the response type for the
// StopLogging API operation.
type StopLoggingResponse struct {
	*types.StopLoggingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopLogging request.
func (r *StopLoggingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
