// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconvert

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
)

const opCancelJob = "CancelJob"

// CancelJobRequest returns a request value for making API operation for
// AWS Elemental MediaConvert.
//
// Permanently cancel a job. Once you have canceled a job, you can't start it
// again.
//
//    // Example sending a request using CancelJobRequest.
//    req := client.CancelJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/CancelJob
func (c *Client) CancelJobRequest(input *types.CancelJobInput) CancelJobRequest {
	op := &aws.Operation{
		Name:       opCancelJob,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2017-08-29/jobs/{id}",
	}

	if input == nil {
		input = &types.CancelJobInput{}
	}

	req := c.newRequest(op, input, &types.CancelJobOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CancelJobMarshaler{Input: input}.GetNamedBuildHandler())

	return CancelJobRequest{Request: req, Input: input, Copy: c.CancelJobRequest}
}

// CancelJobRequest is the request type for the
// CancelJob API operation.
type CancelJobRequest struct {
	*aws.Request
	Input *types.CancelJobInput
	Copy  func(*types.CancelJobInput) CancelJobRequest
}

// Send marshals and sends the CancelJob API request.
func (r CancelJobRequest) Send(ctx context.Context) (*CancelJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelJobResponse{
		CancelJobOutput: r.Request.Data.(*types.CancelJobOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelJobResponse is the response type for the
// CancelJob API operation.
type CancelJobResponse struct {
	*types.CancelJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelJob request.
func (r *CancelJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
