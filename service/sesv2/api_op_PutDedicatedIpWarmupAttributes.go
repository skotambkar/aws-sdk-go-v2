// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sesv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

const opPutDedicatedIpWarmupAttributes = "PutDedicatedIpWarmupAttributes"

// PutDedicatedIpWarmupAttributesRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
//    // Example sending a request using PutDedicatedIpWarmupAttributesRequest.
//    req := client.PutDedicatedIpWarmupAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sesv2-2019-09-27/PutDedicatedIpWarmupAttributes
func (c *Client) PutDedicatedIpWarmupAttributesRequest(input *types.PutDedicatedIpWarmupAttributesInput) PutDedicatedIpWarmupAttributesRequest {
	op := &aws.Operation{
		Name:       opPutDedicatedIpWarmupAttributes,
		HTTPMethod: "PUT",
		HTTPPath:   "/v2/email/dedicated-ips/{IP}/warmup",
	}

	if input == nil {
		input = &types.PutDedicatedIpWarmupAttributesInput{}
	}

	req := c.newRequest(op, input, &types.PutDedicatedIpWarmupAttributesOutput{})
	return PutDedicatedIpWarmupAttributesRequest{Request: req, Input: input, Copy: c.PutDedicatedIpWarmupAttributesRequest}
}

// PutDedicatedIpWarmupAttributesRequest is the request type for the
// PutDedicatedIpWarmupAttributes API operation.
type PutDedicatedIpWarmupAttributesRequest struct {
	*aws.Request
	Input *types.PutDedicatedIpWarmupAttributesInput
	Copy  func(*types.PutDedicatedIpWarmupAttributesInput) PutDedicatedIpWarmupAttributesRequest
}

// Send marshals and sends the PutDedicatedIpWarmupAttributes API request.
func (r PutDedicatedIpWarmupAttributesRequest) Send(ctx context.Context) (*PutDedicatedIpWarmupAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutDedicatedIpWarmupAttributesResponse{
		PutDedicatedIpWarmupAttributesOutput: r.Request.Data.(*types.PutDedicatedIpWarmupAttributesOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutDedicatedIpWarmupAttributesResponse is the response type for the
// PutDedicatedIpWarmupAttributes API operation.
type PutDedicatedIpWarmupAttributesResponse struct {
	*types.PutDedicatedIpWarmupAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutDedicatedIpWarmupAttributes request.
func (r *PutDedicatedIpWarmupAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
