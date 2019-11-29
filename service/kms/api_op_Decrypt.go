// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

const opDecrypt = "Decrypt"

// DecryptRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Decrypts ciphertext. Ciphertext is plaintext that has been previously encrypted
// by using any of the following operations:
//
//    * GenerateDataKey
//
//    * GenerateDataKeyWithoutPlaintext
//
//    * Encrypt
//
// Whenever possible, use key policies to give users permission to call the
// Decrypt operation on the CMK, instead of IAM policies. Otherwise, you might
// create an IAM user policy that gives the user Decrypt permission on all CMKs.
// This user could decrypt ciphertext that was encrypted by CMKs in other accounts
// if the key policy for the cross-account CMK permits it. If you must use an
// IAM policy for Decrypt permissions, limit the user to particular CMKs or
// particular trusted accounts.
//
// The result of this operation varies with the key state of the CMK. For details,
// see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using DecryptRequest.
//    req := client.DecryptRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/Decrypt
func (c *Client) DecryptRequest(input *types.DecryptInput) DecryptRequest {
	op := &aws.Operation{
		Name:       opDecrypt,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DecryptInput{}
	}

	req := c.newRequest(op, input, &types.DecryptOutput{})
	return DecryptRequest{Request: req, Input: input, Copy: c.DecryptRequest}
}

// DecryptRequest is the request type for the
// Decrypt API operation.
type DecryptRequest struct {
	*aws.Request
	Input *types.DecryptInput
	Copy  func(*types.DecryptInput) DecryptRequest
}

// Send marshals and sends the Decrypt API request.
func (r DecryptRequest) Send(ctx context.Context) (*DecryptResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DecryptResponse{
		DecryptOutput: r.Request.Data.(*types.DecryptOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DecryptResponse is the response type for the
// Decrypt API operation.
type DecryptResponse struct {
	*types.DecryptOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// Decrypt request.
func (r *DecryptResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
