// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dataexchange

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/types"
)

const opStartJob = "StartJob"

// StartJobRequest returns a request value for making API operation for
// AWS Data Exchange.
//
// This operation starts a job.
//
//    // Example sending a request using StartJobRequest.
//    req := client.StartJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dataexchange-2017-07-25/StartJob
func (c *Client) StartJobRequest(input *types.StartJobInput) StartJobRequest {
	op := &aws.Operation{
		Name:       opStartJob,
		HTTPMethod: "PATCH",
		HTTPPath:   "/v1/jobs/{JobId}",
	}

	if input == nil {
		input = &types.StartJobInput{}
	}

	req := c.newRequest(op, input, &types.StartJobOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.StartJobMarshaler{Input: input}.GetNamedBuildHandler())

	return StartJobRequest{Request: req, Input: input, Copy: c.StartJobRequest}
}

// StartJobRequest is the request type for the
// StartJob API operation.
type StartJobRequest struct {
	*aws.Request
	Input *types.StartJobInput
	Copy  func(*types.StartJobInput) StartJobRequest
}

// Send marshals and sends the StartJob API request.
func (r StartJobRequest) Send(ctx context.Context) (*StartJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartJobResponse{
		StartJobOutput: r.Request.Data.(*types.StartJobOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartJobResponse is the response type for the
// StartJob API operation.
type StartJobResponse struct {
	*types.StartJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartJob request.
func (r *StartJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
