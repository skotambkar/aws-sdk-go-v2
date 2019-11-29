// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
)

const opCreateDeploymentConfig = "CreateDeploymentConfig"

// CreateDeploymentConfigRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Creates a deployment configuration.
//
//    // Example sending a request using CreateDeploymentConfigRequest.
//    req := client.CreateDeploymentConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/CreateDeploymentConfig
func (c *Client) CreateDeploymentConfigRequest(input *types.CreateDeploymentConfigInput) CreateDeploymentConfigRequest {
	op := &aws.Operation{
		Name:       opCreateDeploymentConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateDeploymentConfigInput{}
	}

	req := c.newRequest(op, input, &types.CreateDeploymentConfigOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateDeploymentConfigMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateDeploymentConfigRequest{Request: req, Input: input, Copy: c.CreateDeploymentConfigRequest}
}

// CreateDeploymentConfigRequest is the request type for the
// CreateDeploymentConfig API operation.
type CreateDeploymentConfigRequest struct {
	*aws.Request
	Input *types.CreateDeploymentConfigInput
	Copy  func(*types.CreateDeploymentConfigInput) CreateDeploymentConfigRequest
}

// Send marshals and sends the CreateDeploymentConfig API request.
func (r CreateDeploymentConfigRequest) Send(ctx context.Context) (*CreateDeploymentConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDeploymentConfigResponse{
		CreateDeploymentConfigOutput: r.Request.Data.(*types.CreateDeploymentConfigOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDeploymentConfigResponse is the response type for the
// CreateDeploymentConfig API operation.
type CreateDeploymentConfigResponse struct {
	*types.CreateDeploymentConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDeploymentConfig request.
func (r *CreateDeploymentConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
