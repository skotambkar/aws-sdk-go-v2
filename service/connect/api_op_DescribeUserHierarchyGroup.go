// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
)

const opDescribeUserHierarchyGroup = "DescribeUserHierarchyGroup"

// DescribeUserHierarchyGroupRequest returns a request value for making API operation for
// Amazon Connect Service.
//
// Returns a HierarchyGroup object that includes information about a hierarchy
// group in your instance.
//
//    // Example sending a request using DescribeUserHierarchyGroupRequest.
//    req := client.DescribeUserHierarchyGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/DescribeUserHierarchyGroup
func (c *Client) DescribeUserHierarchyGroupRequest(input *types.DescribeUserHierarchyGroupInput) DescribeUserHierarchyGroupRequest {
	op := &aws.Operation{
		Name:       opDescribeUserHierarchyGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/user-hierarchy-groups/{InstanceId}/{HierarchyGroupId}",
	}

	if input == nil {
		input = &types.DescribeUserHierarchyGroupInput{}
	}

	req := c.newRequest(op, input, &types.DescribeUserHierarchyGroupOutput{})
	return DescribeUserHierarchyGroupRequest{Request: req, Input: input, Copy: c.DescribeUserHierarchyGroupRequest}
}

// DescribeUserHierarchyGroupRequest is the request type for the
// DescribeUserHierarchyGroup API operation.
type DescribeUserHierarchyGroupRequest struct {
	*aws.Request
	Input *types.DescribeUserHierarchyGroupInput
	Copy  func(*types.DescribeUserHierarchyGroupInput) DescribeUserHierarchyGroupRequest
}

// Send marshals and sends the DescribeUserHierarchyGroup API request.
func (r DescribeUserHierarchyGroupRequest) Send(ctx context.Context) (*DescribeUserHierarchyGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeUserHierarchyGroupResponse{
		DescribeUserHierarchyGroupOutput: r.Request.Data.(*types.DescribeUserHierarchyGroupOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeUserHierarchyGroupResponse is the response type for the
// DescribeUserHierarchyGroup API operation.
type DescribeUserHierarchyGroupResponse struct {
	*types.DescribeUserHierarchyGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeUserHierarchyGroup request.
func (r *DescribeUserHierarchyGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
