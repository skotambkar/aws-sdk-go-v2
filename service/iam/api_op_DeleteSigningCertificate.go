// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opDeleteSigningCertificate = "DeleteSigningCertificate"

// DeleteSigningCertificateRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Deletes a signing certificate associated with the specified IAM user.
//
// If you do not specify a user name, IAM determines the user name implicitly
// based on the AWS access key ID signing the request. This operation works
// for access keys under the AWS account. Consequently, you can use this operation
// to manage AWS account root user credentials even if the AWS account has no
// associated IAM users.
//
//    // Example sending a request using DeleteSigningCertificateRequest.
//    req := client.DeleteSigningCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteSigningCertificate
func (c *Client) DeleteSigningCertificateRequest(input *types.DeleteSigningCertificateInput) DeleteSigningCertificateRequest {
	op := &aws.Operation{
		Name:       opDeleteSigningCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteSigningCertificateInput{}
	}

	req := c.newRequest(op, input, &types.DeleteSigningCertificateOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteSigningCertificateRequest{Request: req, Input: input, Copy: c.DeleteSigningCertificateRequest}
}

// DeleteSigningCertificateRequest is the request type for the
// DeleteSigningCertificate API operation.
type DeleteSigningCertificateRequest struct {
	*aws.Request
	Input *types.DeleteSigningCertificateInput
	Copy  func(*types.DeleteSigningCertificateInput) DeleteSigningCertificateRequest
}

// Send marshals and sends the DeleteSigningCertificate API request.
func (r DeleteSigningCertificateRequest) Send(ctx context.Context) (*DeleteSigningCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSigningCertificateResponse{
		DeleteSigningCertificateOutput: r.Request.Data.(*types.DeleteSigningCertificateOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSigningCertificateResponse is the response type for the
// DeleteSigningCertificate API operation.
type DeleteSigningCertificateResponse struct {
	*types.DeleteSigningCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSigningCertificate request.
func (r *DeleteSigningCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
