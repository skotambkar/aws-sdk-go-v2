// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
)

const opDescribeResourcePermissions = "DescribeResourcePermissions"

// DescribeResourcePermissionsRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Describes the permissions of a specified resource.
//
//    // Example sending a request using DescribeResourcePermissionsRequest.
//    req := client.DescribeResourcePermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DescribeResourcePermissions
func (c *Client) DescribeResourcePermissionsRequest(input *types.DescribeResourcePermissionsInput) DescribeResourcePermissionsRequest {
	op := &aws.Operation{
		Name:       opDescribeResourcePermissions,
		HTTPMethod: "GET",
		HTTPPath:   "/api/v1/resources/{ResourceId}/permissions",
	}

	if input == nil {
		input = &types.DescribeResourcePermissionsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeResourcePermissionsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DescribeResourcePermissionsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeResourcePermissionsRequest{Request: req, Input: input, Copy: c.DescribeResourcePermissionsRequest}
}

// DescribeResourcePermissionsRequest is the request type for the
// DescribeResourcePermissions API operation.
type DescribeResourcePermissionsRequest struct {
	*aws.Request
	Input *types.DescribeResourcePermissionsInput
	Copy  func(*types.DescribeResourcePermissionsInput) DescribeResourcePermissionsRequest
}

// Send marshals and sends the DescribeResourcePermissions API request.
func (r DescribeResourcePermissionsRequest) Send(ctx context.Context) (*DescribeResourcePermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeResourcePermissionsResponse{
		DescribeResourcePermissionsOutput: r.Request.Data.(*types.DescribeResourcePermissionsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeResourcePermissionsResponse is the response type for the
// DescribeResourcePermissions API operation.
type DescribeResourcePermissionsResponse struct {
	*types.DescribeResourcePermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeResourcePermissions request.
func (r *DescribeResourcePermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
