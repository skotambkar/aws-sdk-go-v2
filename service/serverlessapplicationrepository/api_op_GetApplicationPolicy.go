// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/types"
)

const opGetApplicationPolicy = "GetApplicationPolicy"

// GetApplicationPolicyRequest returns a request value for making API operation for
// AWSServerlessApplicationRepository.
//
// Retrieves the policy for the application.
//
//    // Example sending a request using GetApplicationPolicyRequest.
//    req := client.GetApplicationPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/GetApplicationPolicy
func (c *Client) GetApplicationPolicyRequest(input *types.GetApplicationPolicyInput) GetApplicationPolicyRequest {
	op := &aws.Operation{
		Name:       opGetApplicationPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/applications/{applicationId}/policy",
	}

	if input == nil {
		input = &types.GetApplicationPolicyInput{}
	}

	req := c.newRequest(op, input, &types.GetApplicationPolicyOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetApplicationPolicyMarshaler{Input: input}.GetNamedBuildHandler())

	return GetApplicationPolicyRequest{Request: req, Input: input, Copy: c.GetApplicationPolicyRequest}
}

// GetApplicationPolicyRequest is the request type for the
// GetApplicationPolicy API operation.
type GetApplicationPolicyRequest struct {
	*aws.Request
	Input *types.GetApplicationPolicyInput
	Copy  func(*types.GetApplicationPolicyInput) GetApplicationPolicyRequest
}

// Send marshals and sends the GetApplicationPolicy API request.
func (r GetApplicationPolicyRequest) Send(ctx context.Context) (*GetApplicationPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetApplicationPolicyResponse{
		GetApplicationPolicyOutput: r.Request.Data.(*types.GetApplicationPolicyOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetApplicationPolicyResponse is the response type for the
// GetApplicationPolicy API operation.
type GetApplicationPolicyResponse struct {
	*types.GetApplicationPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetApplicationPolicy request.
func (r *GetApplicationPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
