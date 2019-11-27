// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opAssociateRoleToGroup = "AssociateRoleToGroup"

// AssociateRoleToGroupRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Associates a role with a group. Your Greengrass core will use the role to
// access AWS cloud services. The role's permissions should allow Greengrass
// core Lambda functions to perform actions against the cloud.
//
//    // Example sending a request using AssociateRoleToGroupRequest.
//    req := client.AssociateRoleToGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/AssociateRoleToGroup
func (c *Client) AssociateRoleToGroupRequest(input *types.AssociateRoleToGroupInput) AssociateRoleToGroupRequest {
	op := &aws.Operation{
		Name:       opAssociateRoleToGroup,
		HTTPMethod: "PUT",
		HTTPPath:   "/greengrass/groups/{GroupId}/role",
	}

	if input == nil {
		input = &types.AssociateRoleToGroupInput{}
	}

	req := c.newRequest(op, input, &types.AssociateRoleToGroupOutput{})
	return AssociateRoleToGroupRequest{Request: req, Input: input, Copy: c.AssociateRoleToGroupRequest}
}

// AssociateRoleToGroupRequest is the request type for the
// AssociateRoleToGroup API operation.
type AssociateRoleToGroupRequest struct {
	*aws.Request
	Input *types.AssociateRoleToGroupInput
	Copy  func(*types.AssociateRoleToGroupInput) AssociateRoleToGroupRequest
}

// Send marshals and sends the AssociateRoleToGroup API request.
func (r AssociateRoleToGroupRequest) Send(ctx context.Context) (*AssociateRoleToGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateRoleToGroupResponse{
		AssociateRoleToGroupOutput: r.Request.Data.(*types.AssociateRoleToGroupOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateRoleToGroupResponse is the response type for the
// AssociateRoleToGroup API operation.
type AssociateRoleToGroupResponse struct {
	*types.AssociateRoleToGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateRoleToGroup request.
func (r *AssociateRoleToGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
