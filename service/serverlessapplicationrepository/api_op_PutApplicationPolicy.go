// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/types"
)

const opPutApplicationPolicy = "PutApplicationPolicy"

// PutApplicationPolicyRequest returns a request value for making API operation for
// AWSServerlessApplicationRepository.
//
// Sets the permission policy for an application. For the list of actions supported
// for this operation, see Application Permissions (https://docs.aws.amazon.com/serverlessrepo/latest/devguide/access-control-resource-based.html#application-permissions) .
//
//    // Example sending a request using PutApplicationPolicyRequest.
//    req := client.PutApplicationPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/PutApplicationPolicy
func (c *Client) PutApplicationPolicyRequest(input *types.PutApplicationPolicyInput) PutApplicationPolicyRequest {
	op := &aws.Operation{
		Name:       opPutApplicationPolicy,
		HTTPMethod: "PUT",
		HTTPPath:   "/applications/{applicationId}/policy",
	}

	if input == nil {
		input = &types.PutApplicationPolicyInput{}
	}

	req := c.newRequest(op, input, &types.PutApplicationPolicyOutput{})
	return PutApplicationPolicyRequest{Request: req, Input: input, Copy: c.PutApplicationPolicyRequest}
}

// PutApplicationPolicyRequest is the request type for the
// PutApplicationPolicy API operation.
type PutApplicationPolicyRequest struct {
	*aws.Request
	Input *types.PutApplicationPolicyInput
	Copy  func(*types.PutApplicationPolicyInput) PutApplicationPolicyRequest
}

// Send marshals and sends the PutApplicationPolicy API request.
func (r PutApplicationPolicyRequest) Send(ctx context.Context) (*PutApplicationPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutApplicationPolicyResponse{
		PutApplicationPolicyOutput: r.Request.Data.(*types.PutApplicationPolicyOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutApplicationPolicyResponse is the response type for the
// PutApplicationPolicy API operation.
type PutApplicationPolicyResponse struct {
	*types.PutApplicationPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutApplicationPolicy request.
func (r *PutApplicationPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
