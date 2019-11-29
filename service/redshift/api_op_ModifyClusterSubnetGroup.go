// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/redshift/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
)

const opModifyClusterSubnetGroup = "ModifyClusterSubnetGroup"

// ModifyClusterSubnetGroupRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Modifies a cluster subnet group to include the specified list of VPC subnets.
// The operation replaces the existing list of subnets with the new list of
// subnets.
//
//    // Example sending a request using ModifyClusterSubnetGroupRequest.
//    req := client.ModifyClusterSubnetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/ModifyClusterSubnetGroup
func (c *Client) ModifyClusterSubnetGroupRequest(input *types.ModifyClusterSubnetGroupInput) ModifyClusterSubnetGroupRequest {
	op := &aws.Operation{
		Name:       opModifyClusterSubnetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ModifyClusterSubnetGroupInput{}
	}

	req := c.newRequest(op, input, &types.ModifyClusterSubnetGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.ModifyClusterSubnetGroupMarshaler{Input: input}.GetNamedBuildHandler())

	return ModifyClusterSubnetGroupRequest{Request: req, Input: input, Copy: c.ModifyClusterSubnetGroupRequest}
}

// ModifyClusterSubnetGroupRequest is the request type for the
// ModifyClusterSubnetGroup API operation.
type ModifyClusterSubnetGroupRequest struct {
	*aws.Request
	Input *types.ModifyClusterSubnetGroupInput
	Copy  func(*types.ModifyClusterSubnetGroupInput) ModifyClusterSubnetGroupRequest
}

// Send marshals and sends the ModifyClusterSubnetGroup API request.
func (r ModifyClusterSubnetGroupRequest) Send(ctx context.Context) (*ModifyClusterSubnetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyClusterSubnetGroupResponse{
		ModifyClusterSubnetGroupOutput: r.Request.Data.(*types.ModifyClusterSubnetGroupOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyClusterSubnetGroupResponse is the response type for the
// ModifyClusterSubnetGroup API operation.
type ModifyClusterSubnetGroupResponse struct {
	*types.ModifyClusterSubnetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyClusterSubnetGroup request.
func (r *ModifyClusterSubnetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
