// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
)

const opRegisterElasticIp = "RegisterElasticIp"

// RegisterElasticIpRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Registers an Elastic IP address with a specified stack. An address can be
// registered with only one stack at a time. If the address is already registered,
// you must first deregister it by calling DeregisterElasticIp. For more information,
// see Resource Management (https://docs.aws.amazon.com/opsworks/latest/userguide/resources.html).
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using RegisterElasticIpRequest.
//    req := client.RegisterElasticIpRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/RegisterElasticIp
func (c *Client) RegisterElasticIpRequest(input *types.RegisterElasticIpInput) RegisterElasticIpRequest {
	op := &aws.Operation{
		Name:       opRegisterElasticIp,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RegisterElasticIpInput{}
	}

	req := c.newRequest(op, input, &types.RegisterElasticIpOutput{})
	return RegisterElasticIpRequest{Request: req, Input: input, Copy: c.RegisterElasticIpRequest}
}

// RegisterElasticIpRequest is the request type for the
// RegisterElasticIp API operation.
type RegisterElasticIpRequest struct {
	*aws.Request
	Input *types.RegisterElasticIpInput
	Copy  func(*types.RegisterElasticIpInput) RegisterElasticIpRequest
}

// Send marshals and sends the RegisterElasticIp API request.
func (r RegisterElasticIpRequest) Send(ctx context.Context) (*RegisterElasticIpResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterElasticIpResponse{
		RegisterElasticIpOutput: r.Request.Data.(*types.RegisterElasticIpOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterElasticIpResponse is the response type for the
// RegisterElasticIp API operation.
type RegisterElasticIpResponse struct {
	*types.RegisterElasticIpOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterElasticIp request.
func (r *RegisterElasticIpResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
