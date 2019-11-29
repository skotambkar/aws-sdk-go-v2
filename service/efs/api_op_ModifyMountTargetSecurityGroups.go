// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package efs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/efs/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
)

const opModifyMountTargetSecurityGroups = "ModifyMountTargetSecurityGroups"

// ModifyMountTargetSecurityGroupsRequest returns a request value for making API operation for
// Amazon Elastic File System.
//
// Modifies the set of security groups in effect for a mount target.
//
// When you create a mount target, Amazon EFS also creates a new network interface.
// For more information, see CreateMountTarget. This operation replaces the
// security groups in effect for the network interface associated with a mount
// target, with the SecurityGroups provided in the request. This operation requires
// that the network interface of the mount target has been created and the lifecycle
// state of the mount target is not deleted.
//
// The operation requires permissions for the following actions:
//
//    * elasticfilesystem:ModifyMountTargetSecurityGroups action on the mount
//    target's file system.
//
//    * ec2:ModifyNetworkInterfaceAttribute action on the mount target's network
//    interface.
//
//    // Example sending a request using ModifyMountTargetSecurityGroupsRequest.
//    req := client.ModifyMountTargetSecurityGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/ModifyMountTargetSecurityGroups
func (c *Client) ModifyMountTargetSecurityGroupsRequest(input *types.ModifyMountTargetSecurityGroupsInput) ModifyMountTargetSecurityGroupsRequest {
	op := &aws.Operation{
		Name:       opModifyMountTargetSecurityGroups,
		HTTPMethod: "PUT",
		HTTPPath:   "/2015-02-01/mount-targets/{MountTargetId}/security-groups",
	}

	if input == nil {
		input = &types.ModifyMountTargetSecurityGroupsInput{}
	}

	req := c.newRequest(op, input, &types.ModifyMountTargetSecurityGroupsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ModifyMountTargetSecurityGroupsMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return ModifyMountTargetSecurityGroupsRequest{Request: req, Input: input, Copy: c.ModifyMountTargetSecurityGroupsRequest}
}

// ModifyMountTargetSecurityGroupsRequest is the request type for the
// ModifyMountTargetSecurityGroups API operation.
type ModifyMountTargetSecurityGroupsRequest struct {
	*aws.Request
	Input *types.ModifyMountTargetSecurityGroupsInput
	Copy  func(*types.ModifyMountTargetSecurityGroupsInput) ModifyMountTargetSecurityGroupsRequest
}

// Send marshals and sends the ModifyMountTargetSecurityGroups API request.
func (r ModifyMountTargetSecurityGroupsRequest) Send(ctx context.Context) (*ModifyMountTargetSecurityGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyMountTargetSecurityGroupsResponse{
		ModifyMountTargetSecurityGroupsOutput: r.Request.Data.(*types.ModifyMountTargetSecurityGroupsOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyMountTargetSecurityGroupsResponse is the response type for the
// ModifyMountTargetSecurityGroups API operation.
type ModifyMountTargetSecurityGroupsResponse struct {
	*types.ModifyMountTargetSecurityGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyMountTargetSecurityGroups request.
func (r *ModifyMountTargetSecurityGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
