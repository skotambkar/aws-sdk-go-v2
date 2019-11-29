// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opDeleteParameter = "DeleteParameter"

// DeleteParameterRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Delete a parameter from the system.
//
//    // Example sending a request using DeleteParameterRequest.
//    req := client.DeleteParameterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeleteParameter
func (c *Client) DeleteParameterRequest(input *types.DeleteParameterInput) DeleteParameterRequest {
	op := &aws.Operation{
		Name:       opDeleteParameter,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteParameterInput{}
	}

	req := c.newRequest(op, input, &types.DeleteParameterOutput{})
	return DeleteParameterRequest{Request: req, Input: input, Copy: c.DeleteParameterRequest}
}

// DeleteParameterRequest is the request type for the
// DeleteParameter API operation.
type DeleteParameterRequest struct {
	*aws.Request
	Input *types.DeleteParameterInput
	Copy  func(*types.DeleteParameterInput) DeleteParameterRequest
}

// Send marshals and sends the DeleteParameter API request.
func (r DeleteParameterRequest) Send(ctx context.Context) (*DeleteParameterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteParameterResponse{
		DeleteParameterOutput: r.Request.Data.(*types.DeleteParameterOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteParameterResponse is the response type for the
// DeleteParameter API operation.
type DeleteParameterResponse struct {
	*types.DeleteParameterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteParameter request.
func (r *DeleteParameterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
