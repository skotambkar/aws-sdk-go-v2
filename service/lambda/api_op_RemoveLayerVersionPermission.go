// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

const opRemoveLayerVersionPermission = "RemoveLayerVersionPermission"

// RemoveLayerVersionPermissionRequest returns a request value for making API operation for
// AWS Lambda.
//
// Removes a statement from the permissions policy for a version of an AWS Lambda
// layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
// For more information, see AddLayerVersionPermission.
//
//    // Example sending a request using RemoveLayerVersionPermissionRequest.
//    req := client.RemoveLayerVersionPermissionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/RemoveLayerVersionPermission
func (c *Client) RemoveLayerVersionPermissionRequest(input *types.RemoveLayerVersionPermissionInput) RemoveLayerVersionPermissionRequest {
	op := &aws.Operation{
		Name:       opRemoveLayerVersionPermission,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2018-10-31/layers/{LayerName}/versions/{VersionNumber}/policy/{StatementId}",
	}

	if input == nil {
		input = &types.RemoveLayerVersionPermissionInput{}
	}

	req := c.newRequest(op, input, &types.RemoveLayerVersionPermissionOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RemoveLayerVersionPermissionRequest{Request: req, Input: input, Copy: c.RemoveLayerVersionPermissionRequest}
}

// RemoveLayerVersionPermissionRequest is the request type for the
// RemoveLayerVersionPermission API operation.
type RemoveLayerVersionPermissionRequest struct {
	*aws.Request
	Input *types.RemoveLayerVersionPermissionInput
	Copy  func(*types.RemoveLayerVersionPermissionInput) RemoveLayerVersionPermissionRequest
}

// Send marshals and sends the RemoveLayerVersionPermission API request.
func (r RemoveLayerVersionPermissionRequest) Send(ctx context.Context) (*RemoveLayerVersionPermissionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveLayerVersionPermissionResponse{
		RemoveLayerVersionPermissionOutput: r.Request.Data.(*types.RemoveLayerVersionPermissionOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveLayerVersionPermissionResponse is the response type for the
// RemoveLayerVersionPermission API operation.
type RemoveLayerVersionPermissionResponse struct {
	*types.RemoveLayerVersionPermissionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveLayerVersionPermission request.
func (r *RemoveLayerVersionPermissionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
