// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourcegroups

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
)

const opDeleteGroup = "DeleteGroup"

// DeleteGroupRequest returns a request value for making API operation for
// AWS Resource Groups.
//
// Deletes a specified resource group. Deleting a resource group does not delete
// resources that are members of the group; it only deletes the group structure.
//
//    // Example sending a request using DeleteGroupRequest.
//    req := client.DeleteGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/resource-groups-2017-11-27/DeleteGroup
func (c *Client) DeleteGroupRequest(input *types.DeleteGroupInput) DeleteGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteGroup,
		HTTPMethod: "DELETE",
		HTTPPath:   "/groups/{GroupName}",
	}

	if input == nil {
		input = &types.DeleteGroupInput{}
	}

	req := c.newRequest(op, input, &types.DeleteGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DeleteGroupMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteGroupRequest{Request: req, Input: input, Copy: c.DeleteGroupRequest}
}

// DeleteGroupRequest is the request type for the
// DeleteGroup API operation.
type DeleteGroupRequest struct {
	*aws.Request
	Input *types.DeleteGroupInput
	Copy  func(*types.DeleteGroupInput) DeleteGroupRequest
}

// Send marshals and sends the DeleteGroup API request.
func (r DeleteGroupRequest) Send(ctx context.Context) (*DeleteGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteGroupResponse{
		DeleteGroupOutput: r.Request.Data.(*types.DeleteGroupOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteGroupResponse is the response type for the
// DeleteGroup API operation.
type DeleteGroupResponse struct {
	*types.DeleteGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteGroup request.
func (r *DeleteGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
