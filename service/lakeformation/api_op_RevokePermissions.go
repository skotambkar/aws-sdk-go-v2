// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
)

const opRevokePermissions = "RevokePermissions"

// RevokePermissionsRequest returns a request value for making API operation for
// AWS Lake Formation.
//
// Revokes permissions to the principal to access metadata in the Data Catalog
// and data organized in underlying data storage such as Amazon S3.
//
//    // Example sending a request using RevokePermissionsRequest.
//    req := client.RevokePermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lakeformation-2017-03-31/RevokePermissions
func (c *Client) RevokePermissionsRequest(input *types.RevokePermissionsInput) RevokePermissionsRequest {
	op := &aws.Operation{
		Name:       opRevokePermissions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RevokePermissionsInput{}
	}

	req := c.newRequest(op, input, &types.RevokePermissionsOutput{})
	return RevokePermissionsRequest{Request: req, Input: input, Copy: c.RevokePermissionsRequest}
}

// RevokePermissionsRequest is the request type for the
// RevokePermissions API operation.
type RevokePermissionsRequest struct {
	*aws.Request
	Input *types.RevokePermissionsInput
	Copy  func(*types.RevokePermissionsInput) RevokePermissionsRequest
}

// Send marshals and sends the RevokePermissions API request.
func (r RevokePermissionsRequest) Send(ctx context.Context) (*RevokePermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RevokePermissionsResponse{
		RevokePermissionsOutput: r.Request.Data.(*types.RevokePermissionsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RevokePermissionsResponse is the response type for the
// RevokePermissions API operation.
type RevokePermissionsResponse struct {
	*types.RevokePermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RevokePermissions request.
func (r *RevokePermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
