// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

const opCreateGrant = "CreateGrant"

// CreateGrantRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Adds a grant to a customer master key (CMK). The grant allows the grantee
// principal to use the CMK when the conditions specified in the grant are met.
// When setting permissions, grants are an alternative to key policies.
//
// To create a grant that allows a cryptographic operation only when the encryption
// context in the operation request matches or includes a specified encryption
// context, use the Constraints parameter. For details, see GrantConstraints.
//
// To perform this operation on a CMK in a different AWS account, specify the
// key ARN in the value of the KeyId parameter. For more information about grants,
// see Grants (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html)
// in the AWS Key Management Service Developer Guide .
//
// The result of this operation varies with the key state of the CMK. For details,
// see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using CreateGrantRequest.
//    req := client.CreateGrantRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/CreateGrant
func (c *Client) CreateGrantRequest(input *types.CreateGrantInput) CreateGrantRequest {
	op := &aws.Operation{
		Name:       opCreateGrant,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateGrantInput{}
	}

	req := c.newRequest(op, input, &types.CreateGrantOutput{})
	return CreateGrantRequest{Request: req, Input: input, Copy: c.CreateGrantRequest}
}

// CreateGrantRequest is the request type for the
// CreateGrant API operation.
type CreateGrantRequest struct {
	*aws.Request
	Input *types.CreateGrantInput
	Copy  func(*types.CreateGrantInput) CreateGrantRequest
}

// Send marshals and sends the CreateGrant API request.
func (r CreateGrantRequest) Send(ctx context.Context) (*CreateGrantResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateGrantResponse{
		CreateGrantOutput: r.Request.Data.(*types.CreateGrantOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateGrantResponse is the response type for the
// CreateGrant API operation.
type CreateGrantResponse struct {
	*types.CreateGrantOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateGrant request.
func (r *CreateGrantResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
