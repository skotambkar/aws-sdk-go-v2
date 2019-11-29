// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
)

const opCreateDeploymentJob = "CreateDeploymentJob"

// CreateDeploymentJobRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Deploys a specific version of a robot application to robots in a fleet.
//
// The robot application must have a numbered applicationVersion for consistency
// reasons. To create a new version, use CreateRobotApplicationVersion or see
// Creating a Robot Application Version (https://docs.aws.amazon.com/robomaker/latest/dg/create-robot-application-version.html).
//
// After 90 days, deployment jobs expire and will be deleted. They will no longer
// be accessible.
//
//    // Example sending a request using CreateDeploymentJobRequest.
//    req := client.CreateDeploymentJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/CreateDeploymentJob
func (c *Client) CreateDeploymentJobRequest(input *types.CreateDeploymentJobInput) CreateDeploymentJobRequest {
	op := &aws.Operation{
		Name:       opCreateDeploymentJob,
		HTTPMethod: "POST",
		HTTPPath:   "/createDeploymentJob",
	}

	if input == nil {
		input = &types.CreateDeploymentJobInput{}
	}

	req := c.newRequest(op, input, &types.CreateDeploymentJobOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreateDeploymentJobMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateDeploymentJobRequest{Request: req, Input: input, Copy: c.CreateDeploymentJobRequest}
}

// CreateDeploymentJobRequest is the request type for the
// CreateDeploymentJob API operation.
type CreateDeploymentJobRequest struct {
	*aws.Request
	Input *types.CreateDeploymentJobInput
	Copy  func(*types.CreateDeploymentJobInput) CreateDeploymentJobRequest
}

// Send marshals and sends the CreateDeploymentJob API request.
func (r CreateDeploymentJobRequest) Send(ctx context.Context) (*CreateDeploymentJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDeploymentJobResponse{
		CreateDeploymentJobOutput: r.Request.Data.(*types.CreateDeploymentJobOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDeploymentJobResponse is the response type for the
// CreateDeploymentJob API operation.
type CreateDeploymentJobResponse struct {
	*types.CreateDeploymentJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDeploymentJob request.
func (r *CreateDeploymentJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
