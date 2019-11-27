// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
)

const opGetIPSet = "GetIPSet"

// GetIPSetRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Retrieves the IPSet specified by the IPSet ID.
//
//    // Example sending a request using GetIPSetRequest.
//    req := client.GetIPSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/GetIPSet
func (c *Client) GetIPSetRequest(input *types.GetIPSetInput) GetIPSetRequest {
	op := &aws.Operation{
		Name:       opGetIPSet,
		HTTPMethod: "GET",
		HTTPPath:   "/detector/{detectorId}/ipset/{ipSetId}",
	}

	if input == nil {
		input = &types.GetIPSetInput{}
	}

	req := c.newRequest(op, input, &types.GetIPSetOutput{})
	return GetIPSetRequest{Request: req, Input: input, Copy: c.GetIPSetRequest}
}

// GetIPSetRequest is the request type for the
// GetIPSet API operation.
type GetIPSetRequest struct {
	*aws.Request
	Input *types.GetIPSetInput
	Copy  func(*types.GetIPSetInput) GetIPSetRequest
}

// Send marshals and sends the GetIPSet API request.
func (r GetIPSetRequest) Send(ctx context.Context) (*GetIPSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetIPSetResponse{
		GetIPSetOutput: r.Request.Data.(*types.GetIPSetOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetIPSetResponse is the response type for the
// GetIPSet API operation.
type GetIPSetResponse struct {
	*types.GetIPSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetIPSet request.
func (r *GetIPSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
