// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDescribePlacementGroups = "DescribePlacementGroups"

// DescribePlacementGroupsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified placement groups or all of your placement groups.
// For more information, see Placement Groups (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribePlacementGroupsRequest.
//    req := client.DescribePlacementGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribePlacementGroups
func (c *Client) DescribePlacementGroupsRequest(input *types.DescribePlacementGroupsInput) DescribePlacementGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribePlacementGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribePlacementGroupsInput{}
	}

	req := c.newRequest(op, input, &types.DescribePlacementGroupsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.DescribePlacementGroupsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribePlacementGroupsRequest{Request: req, Input: input, Copy: c.DescribePlacementGroupsRequest}
}

// DescribePlacementGroupsRequest is the request type for the
// DescribePlacementGroups API operation.
type DescribePlacementGroupsRequest struct {
	*aws.Request
	Input *types.DescribePlacementGroupsInput
	Copy  func(*types.DescribePlacementGroupsInput) DescribePlacementGroupsRequest
}

// Send marshals and sends the DescribePlacementGroups API request.
func (r DescribePlacementGroupsRequest) Send(ctx context.Context) (*DescribePlacementGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePlacementGroupsResponse{
		DescribePlacementGroupsOutput: r.Request.Data.(*types.DescribePlacementGroupsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribePlacementGroupsResponse is the response type for the
// DescribePlacementGroups API operation.
type DescribePlacementGroupsResponse struct {
	*types.DescribePlacementGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePlacementGroups request.
func (r *DescribePlacementGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
