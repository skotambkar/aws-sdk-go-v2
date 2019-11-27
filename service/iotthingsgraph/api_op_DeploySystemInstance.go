// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
)

const opDeploySystemInstance = "DeploySystemInstance"

// DeploySystemInstanceRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Greengrass and Cloud Deployments
//
// Deploys the system instance to the target specified in CreateSystemInstance.
//
// Greengrass Deployments
//
// If the system or any workflows and entities have been updated before this
// action is called, then the deployment will create a new Amazon Simple Storage
// Service resource file and then deploy it.
//
// Since this action creates a Greengrass deployment on the caller's behalf,
// the calling identity must have write permissions to the specified Greengrass
// group. Otherwise, the call will fail with an authorization error.
//
// For information about the artifacts that get added to your Greengrass core
// device when you use this API, see AWS IoT Things Graph and AWS IoT Greengrass
// (https://docs.aws.amazon.com/thingsgraph/latest/ug/iot-tg-greengrass.html).
//
//    // Example sending a request using DeploySystemInstanceRequest.
//    req := client.DeploySystemInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/DeploySystemInstance
func (c *Client) DeploySystemInstanceRequest(input *types.DeploySystemInstanceInput) DeploySystemInstanceRequest {
	op := &aws.Operation{
		Name:       opDeploySystemInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeploySystemInstanceInput{}
	}

	req := c.newRequest(op, input, &types.DeploySystemInstanceOutput{})
	return DeploySystemInstanceRequest{Request: req, Input: input, Copy: c.DeploySystemInstanceRequest}
}

// DeploySystemInstanceRequest is the request type for the
// DeploySystemInstance API operation.
type DeploySystemInstanceRequest struct {
	*aws.Request
	Input *types.DeploySystemInstanceInput
	Copy  func(*types.DeploySystemInstanceInput) DeploySystemInstanceRequest
}

// Send marshals and sends the DeploySystemInstance API request.
func (r DeploySystemInstanceRequest) Send(ctx context.Context) (*DeploySystemInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeploySystemInstanceResponse{
		DeploySystemInstanceOutput: r.Request.Data.(*types.DeploySystemInstanceOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeploySystemInstanceResponse is the response type for the
// DeploySystemInstance API operation.
type DeploySystemInstanceResponse struct {
	*types.DeploySystemInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeploySystemInstance request.
func (r *DeploySystemInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
