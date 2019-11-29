// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sesv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

const opPutDedicatedIpInPool = "PutDedicatedIpInPool"

// PutDedicatedIpInPoolRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Move a dedicated IP address to an existing dedicated IP pool.
//
// The dedicated IP address that you specify must already exist, and must be
// associated with your AWS account.
//
// The dedicated IP pool you specify must already exist. You can create a new
// pool by using the CreateDedicatedIpPool operation.
//
//    // Example sending a request using PutDedicatedIpInPoolRequest.
//    req := client.PutDedicatedIpInPoolRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sesv2-2019-09-27/PutDedicatedIpInPool
func (c *Client) PutDedicatedIpInPoolRequest(input *types.PutDedicatedIpInPoolInput) PutDedicatedIpInPoolRequest {
	op := &aws.Operation{
		Name:       opPutDedicatedIpInPool,
		HTTPMethod: "PUT",
		HTTPPath:   "/v2/email/dedicated-ips/{IP}/pool",
	}

	if input == nil {
		input = &types.PutDedicatedIpInPoolInput{}
	}

	req := c.newRequest(op, input, &types.PutDedicatedIpInPoolOutput{})
	return PutDedicatedIpInPoolRequest{Request: req, Input: input, Copy: c.PutDedicatedIpInPoolRequest}
}

// PutDedicatedIpInPoolRequest is the request type for the
// PutDedicatedIpInPool API operation.
type PutDedicatedIpInPoolRequest struct {
	*aws.Request
	Input *types.PutDedicatedIpInPoolInput
	Copy  func(*types.PutDedicatedIpInPoolInput) PutDedicatedIpInPoolRequest
}

// Send marshals and sends the PutDedicatedIpInPool API request.
func (r PutDedicatedIpInPoolRequest) Send(ctx context.Context) (*PutDedicatedIpInPoolResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutDedicatedIpInPoolResponse{
		PutDedicatedIpInPoolOutput: r.Request.Data.(*types.PutDedicatedIpInPoolOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutDedicatedIpInPoolResponse is the response type for the
// PutDedicatedIpInPool API operation.
type PutDedicatedIpInPoolResponse struct {
	*types.PutDedicatedIpInPoolOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutDedicatedIpInPool request.
func (r *PutDedicatedIpInPoolResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
