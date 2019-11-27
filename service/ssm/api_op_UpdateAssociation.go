// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opUpdateAssociation = "UpdateAssociation"

// UpdateAssociationRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Updates an association. You can update the association name and version,
// the document version, schedule, parameters, and Amazon S3 output.
//
// In order to call this API action, your IAM user account, group, or role must
// be configured with permission to call the DescribeAssociation API action.
// If you don't have permission to call DescribeAssociation, then you receive
// the following error: An error occurred (AccessDeniedException) when calling
// the UpdateAssociation operation: User: <user_arn> is not authorized to perform:
// ssm:DescribeAssociation on resource: <resource_arn>
//
// When you update an association, the association immediately runs against
// the specified targets.
//
//    // Example sending a request using UpdateAssociationRequest.
//    req := client.UpdateAssociationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/UpdateAssociation
func (c *Client) UpdateAssociationRequest(input *types.UpdateAssociationInput) UpdateAssociationRequest {
	op := &aws.Operation{
		Name:       opUpdateAssociation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateAssociationInput{}
	}

	req := c.newRequest(op, input, &types.UpdateAssociationOutput{})
	return UpdateAssociationRequest{Request: req, Input: input, Copy: c.UpdateAssociationRequest}
}

// UpdateAssociationRequest is the request type for the
// UpdateAssociation API operation.
type UpdateAssociationRequest struct {
	*aws.Request
	Input *types.UpdateAssociationInput
	Copy  func(*types.UpdateAssociationInput) UpdateAssociationRequest
}

// Send marshals and sends the UpdateAssociation API request.
func (r UpdateAssociationRequest) Send(ctx context.Context) (*UpdateAssociationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAssociationResponse{
		UpdateAssociationOutput: r.Request.Data.(*types.UpdateAssociationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAssociationResponse is the response type for the
// UpdateAssociation API operation.
type UpdateAssociationResponse struct {
	*types.UpdateAssociationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAssociation request.
func (r *UpdateAssociationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
