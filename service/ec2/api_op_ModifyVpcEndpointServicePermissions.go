// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opModifyVpcEndpointServicePermissions = "ModifyVpcEndpointServicePermissions"

// ModifyVpcEndpointServicePermissionsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the permissions for your VPC endpoint service (https://docs.aws.amazon.com/vpc/latest/userguide/endpoint-service.html).
// You can add or remove permissions for service consumers (IAM users, IAM roles,
// and AWS accounts) to connect to your endpoint service.
//
// If you grant permissions to all principals, the service is public. Any users
// who know the name of a public service can send a request to attach an endpoint.
// If the service does not require manual approval, attachments are automatically
// approved.
//
//    // Example sending a request using ModifyVpcEndpointServicePermissionsRequest.
//    req := client.ModifyVpcEndpointServicePermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyVpcEndpointServicePermissions
func (c *Client) ModifyVpcEndpointServicePermissionsRequest(input *types.ModifyVpcEndpointServicePermissionsInput) ModifyVpcEndpointServicePermissionsRequest {
	op := &aws.Operation{
		Name:       opModifyVpcEndpointServicePermissions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ModifyVpcEndpointServicePermissionsInput{}
	}

	req := c.newRequest(op, input, &types.ModifyVpcEndpointServicePermissionsOutput{})
	return ModifyVpcEndpointServicePermissionsRequest{Request: req, Input: input, Copy: c.ModifyVpcEndpointServicePermissionsRequest}
}

// ModifyVpcEndpointServicePermissionsRequest is the request type for the
// ModifyVpcEndpointServicePermissions API operation.
type ModifyVpcEndpointServicePermissionsRequest struct {
	*aws.Request
	Input *types.ModifyVpcEndpointServicePermissionsInput
	Copy  func(*types.ModifyVpcEndpointServicePermissionsInput) ModifyVpcEndpointServicePermissionsRequest
}

// Send marshals and sends the ModifyVpcEndpointServicePermissions API request.
func (r ModifyVpcEndpointServicePermissionsRequest) Send(ctx context.Context) (*ModifyVpcEndpointServicePermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyVpcEndpointServicePermissionsResponse{
		ModifyVpcEndpointServicePermissionsOutput: r.Request.Data.(*types.ModifyVpcEndpointServicePermissionsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyVpcEndpointServicePermissionsResponse is the response type for the
// ModifyVpcEndpointServicePermissions API operation.
type ModifyVpcEndpointServicePermissionsResponse struct {
	*types.ModifyVpcEndpointServicePermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyVpcEndpointServicePermissions request.
func (r *ModifyVpcEndpointServicePermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
