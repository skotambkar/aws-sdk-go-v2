// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
)

const opAssociateElasticIp = "AssociateElasticIp"

// AssociateElasticIpRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Associates one of the stack's registered Elastic IP addresses with a specified
// instance. The address must first be registered with the stack by calling
// RegisterElasticIp. For more information, see Resource Management (https://docs.aws.amazon.com/opsworks/latest/userguide/resources.html).
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using AssociateElasticIpRequest.
//    req := client.AssociateElasticIpRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/AssociateElasticIp
func (c *Client) AssociateElasticIpRequest(input *types.AssociateElasticIpInput) AssociateElasticIpRequest {
	op := &aws.Operation{
		Name:       opAssociateElasticIp,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AssociateElasticIpInput{}
	}

	req := c.newRequest(op, input, &types.AssociateElasticIpOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.AssociateElasticIpMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AssociateElasticIpRequest{Request: req, Input: input, Copy: c.AssociateElasticIpRequest}
}

// AssociateElasticIpRequest is the request type for the
// AssociateElasticIp API operation.
type AssociateElasticIpRequest struct {
	*aws.Request
	Input *types.AssociateElasticIpInput
	Copy  func(*types.AssociateElasticIpInput) AssociateElasticIpRequest
}

// Send marshals and sends the AssociateElasticIp API request.
func (r AssociateElasticIpRequest) Send(ctx context.Context) (*AssociateElasticIpResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateElasticIpResponse{
		AssociateElasticIpOutput: r.Request.Data.(*types.AssociateElasticIpOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateElasticIpResponse is the response type for the
// AssociateElasticIp API operation.
type AssociateElasticIpResponse struct {
	*types.AssociateElasticIpOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateElasticIp request.
func (r *AssociateElasticIpResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
