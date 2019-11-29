// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dataexchange

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/types"
)

const opCancelJob = "CancelJob"

// CancelJobRequest returns a request value for making API operation for
// AWS Data Exchange.
//
// This operation cancels a job. Jobs can be cancelled only when they are in
// the WAITING state.
//
//    // Example sending a request using CancelJobRequest.
//    req := client.CancelJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dataexchange-2017-07-25/CancelJob
func (c *Client) CancelJobRequest(input *types.CancelJobInput) CancelJobRequest {
	op := &aws.Operation{
		Name:       opCancelJob,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/jobs/{JobId}",
	}

	if input == nil {
		input = &types.CancelJobInput{}
	}

	req := c.newRequest(op, input, &types.CancelJobOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
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
