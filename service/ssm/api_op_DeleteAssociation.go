// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opDeleteAssociation = "DeleteAssociation"

// DeleteAssociationRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Disassociates the specified Systems Manager document from the specified instance.
//
// When you disassociate a document from an instance, it does not change the
// configuration of the instance. To change the configuration state of an instance
// after you disassociate a document, you must create a new document with the
// desired configuration and associate it with the instance.
//
//    // Example sending a request using DeleteAssociationRequest.
//    req := client.DeleteAssociationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeleteAssociation
func (c *Client) DeleteAssociationRequest(input *types.DeleteAssociationInput) DeleteAssociationRequest {
	op := &aws.Operation{
		Name:       opDeleteAssociation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteAssociationInput{}
	}

	req := c.newRequest(op, input, &types.DeleteAssociationOutput{})
	return DeleteAssociationRequest{Request: req, Input: input, Copy: c.DeleteAssociationRequest}
}

// DeleteAssociationRequest is the request type for the
// DeleteAssociation API operation.
type DeleteAssociationRequest struct {
	*aws.Request
	Input *types.DeleteAssociationInput
	Copy  func(*types.DeleteAssociationInput) DeleteAssociationRequest
}

// Send marshals and sends the DeleteAssociation API request.
func (r DeleteAssociationRequest) Send(ctx context.Context) (*DeleteAssociationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAssociationResponse{
		DeleteAssociationOutput: r.Request.Data.(*types.DeleteAssociationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAssociationResponse is the response type for the
// DeleteAssociation API operation.
type DeleteAssociationResponse struct {
	*types.DeleteAssociationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAssociation request.
func (r *DeleteAssociationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
