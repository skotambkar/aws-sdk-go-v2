// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
)

const opBatchGrantPermissions = "BatchGrantPermissions"

// BatchGrantPermissionsRequest returns a request value for making API operation for
// AWS Lake Formation.
//
// Batch operation to grant permissions to the principal.
//
//    // Example sending a request using BatchGrantPermissionsRequest.
//    req := client.BatchGrantPermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lakeformation-2017-03-31/BatchGrantPermissions
func (c *Client) BatchGrantPermissionsRequest(input *types.BatchGrantPermissionsInput) BatchGrantPermissionsRequest {
	op := &aws.Operation{
		Name:       opBatchGrantPermissions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.BatchGrantPermissionsInput{}
	}

	req := c.newRequest(op, input, &types.BatchGrantPermissionsOutput{})
	return BatchGrantPermissionsRequest{Request: req, Input: input, Copy: c.BatchGrantPermissionsRequest}
}

// BatchGrantPermissionsRequest is the request type for the
// BatchGrantPermissions API operation.
type BatchGrantPermissionsRequest struct {
	*aws.Request
	Input *types.BatchGrantPermissionsInput
	Copy  func(*types.BatchGrantPermissionsInput) BatchGrantPermissionsRequest
}

// Send marshals and sends the BatchGrantPermissions API request.
func (r BatchGrantPermissionsRequest) Send(ctx context.Context) (*BatchGrantPermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchGrantPermissionsResponse{
		BatchGrantPermissionsOutput: r.Request.Data.(*types.BatchGrantPermissionsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchGrantPermissionsResponse is the response type for the
// BatchGrantPermissions API operation.
type BatchGrantPermissionsResponse struct {
	*types.BatchGrantPermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchGrantPermissions request.
func (r *BatchGrantPermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
