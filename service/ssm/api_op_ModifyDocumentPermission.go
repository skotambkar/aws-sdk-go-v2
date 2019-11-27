// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opModifyDocumentPermission = "ModifyDocumentPermission"

// ModifyDocumentPermissionRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Shares a Systems Manager document publicly or privately. If you share a document
// privately, you must specify the AWS user account IDs for those people who
// can use the document. If you share a document publicly, you must specify
// All as the account ID.
//
//    // Example sending a request using ModifyDocumentPermissionRequest.
//    req := client.ModifyDocumentPermissionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/ModifyDocumentPermission
func (c *Client) ModifyDocumentPermissionRequest(input *types.ModifyDocumentPermissionInput) ModifyDocumentPermissionRequest {
	op := &aws.Operation{
		Name:       opModifyDocumentPermission,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ModifyDocumentPermissionInput{}
	}

	req := c.newRequest(op, input, &types.ModifyDocumentPermissionOutput{})
	return ModifyDocumentPermissionRequest{Request: req, Input: input, Copy: c.ModifyDocumentPermissionRequest}
}

// ModifyDocumentPermissionRequest is the request type for the
// ModifyDocumentPermission API operation.
type ModifyDocumentPermissionRequest struct {
	*aws.Request
	Input *types.ModifyDocumentPermissionInput
	Copy  func(*types.ModifyDocumentPermissionInput) ModifyDocumentPermissionRequest
}

// Send marshals and sends the ModifyDocumentPermission API request.
func (r ModifyDocumentPermissionRequest) Send(ctx context.Context) (*ModifyDocumentPermissionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyDocumentPermissionResponse{
		ModifyDocumentPermissionOutput: r.Request.Data.(*types.ModifyDocumentPermissionOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyDocumentPermissionResponse is the response type for the
// ModifyDocumentPermission API operation.
type ModifyDocumentPermissionResponse struct {
	*types.ModifyDocumentPermissionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyDocumentPermission request.
func (r *ModifyDocumentPermissionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
