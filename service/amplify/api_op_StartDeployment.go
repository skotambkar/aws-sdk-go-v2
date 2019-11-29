// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
)

const opStartDeployment = "StartDeployment"

// StartDeploymentRequest returns a request value for making API operation for
// AWS Amplify.
//
// Start a deployment for manual deploy apps. (Apps are not connected to repository)
//
//    // Example sending a request using StartDeploymentRequest.
//    req := client.StartDeploymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/StartDeployment
func (c *Client) StartDeploymentRequest(input *types.StartDeploymentInput) StartDeploymentRequest {
	op := &aws.Operation{
		Name:       opStartDeployment,
		HTTPMethod: "POST",
		HTTPPath:   "/apps/{appId}/branches/{branchName}/deployments/start",
	}

	if input == nil {
		input = &types.StartDeploymentInput{}
	}

	req := c.newRequest(op, input, &types.StartDeploymentOutput{})
	return StartDeploymentRequest{Request: req, Input: input, Copy: c.StartDeploymentRequest}
}

// StartDeploymentRequest is the request type for the
// StartDeployment API operation.
type StartDeploymentRequest struct {
	*aws.Request
	Input *types.StartDeploymentInput
	Copy  func(*types.StartDeploymentInput) StartDeploymentRequest
}

// Send marshals and sends the StartDeployment API request.
func (r StartDeploymentRequest) Send(ctx context.Context) (*StartDeploymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartDeploymentResponse{
		StartDeploymentOutput: r.Request.Data.(*types.StartDeploymentOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartDeploymentResponse is the response type for the
// StartDeployment API operation.
type StartDeploymentResponse struct {
	*types.StartDeploymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartDeployment request.
func (r *StartDeploymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
