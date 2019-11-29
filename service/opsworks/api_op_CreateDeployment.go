// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
)

const opCreateDeployment = "CreateDeployment"

// CreateDeploymentRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Runs deployment or stack commands. For more information, see Deploying Apps
// (https://docs.aws.amazon.com/opsworks/latest/userguide/workingapps-deploying.html)
// and Run Stack Commands (https://docs.aws.amazon.com/opsworks/latest/userguide/workingstacks-commands.html).
//
// Required Permissions: To use this action, an IAM user must have a Deploy
// or Manage permissions level for the stack, or an attached policy that explicitly
// grants permissions. For more information on user permissions, see Managing
// User Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using CreateDeploymentRequest.
//    req := client.CreateDeploymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/CreateDeployment
func (c *Client) CreateDeploymentRequest(input *types.CreateDeploymentInput) CreateDeploymentRequest {
	op := &aws.Operation{
		Name:       opCreateDeployment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateDeploymentInput{}
	}

	req := c.newRequest(op, input, &types.CreateDeploymentOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateDeploymentMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateDeploymentRequest{Request: req, Input: input, Copy: c.CreateDeploymentRequest}
}

// CreateDeploymentRequest is the request type for the
// CreateDeployment API operation.
type CreateDeploymentRequest struct {
	*aws.Request
	Input *types.CreateDeploymentInput
	Copy  func(*types.CreateDeploymentInput) CreateDeploymentRequest
}

// Send marshals and sends the CreateDeployment API request.
func (r CreateDeploymentRequest) Send(ctx context.Context) (*CreateDeploymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDeploymentResponse{
		CreateDeploymentOutput: r.Request.Data.(*types.CreateDeploymentOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDeploymentResponse is the response type for the
// CreateDeployment API operation.
type CreateDeploymentResponse struct {
	*types.CreateDeploymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDeployment request.
func (r *CreateDeploymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
