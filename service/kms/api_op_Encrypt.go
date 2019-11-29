// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kms/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

const opEncrypt = "Encrypt"

// EncryptRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Encrypts plaintext into ciphertext by using a customer master key (CMK).
// The Encrypt operation has two primary use cases:
//
//    * You can encrypt up to 4 kilobytes (4096 bytes) of arbitrary data such
//    as an RSA key, a database password, or other sensitive information.
//
//    * You can use the Encrypt operation to move encrypted data from one AWS
//    region to another. In the first region, generate a data key and use the
//    plaintext key to encrypt the data. Then, in the new region, call the Encrypt
//    method on same plaintext data key. Now, you can safely move the encrypted
//    data and encrypted data key to the new region, and decrypt in the new
//    region when necessary.
//
// You don't need use this operation to encrypt a data key within a region.
// The GenerateDataKey and GenerateDataKeyWithoutPlaintext operations return
// an encrypted data key.
//
// Also, you don't need to use this operation to encrypt data in your application.
// You can use the plaintext and encrypted data keys that the GenerateDataKey
// operation returns.
//
// The result of this operation varies with the key state of the CMK. For details,
// see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
// To perform this operation on a CMK in a different AWS account, specify the
// key ARN or alias ARN in the value of the KeyId parameter.
//
//    // Example sending a request using EncryptRequest.
//    req := client.EncryptRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/Encrypt
func (c *Client) EncryptRequest(input *types.EncryptInput) EncryptRequest {
	op := &aws.Operation{
		Name:       opEncrypt,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.EncryptInput{}
	}

	req := c.newRequest(op, input, &types.EncryptOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.EncryptMarshaler{Input: input}.GetNamedBuildHandler())

	return EncryptRequest{Request: req, Input: input, Copy: c.EncryptRequest}
}

// EncryptRequest is the request type for the
// Encrypt API operation.
type EncryptRequest struct {
	*aws.Request
	Input *types.EncryptInput
	Copy  func(*types.EncryptInput) EncryptRequest
}

// Send marshals and sends the Encrypt API request.
func (r EncryptRequest) Send(ctx context.Context) (*EncryptResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EncryptResponse{
		EncryptOutput: r.Request.Data.(*types.EncryptOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EncryptResponse is the response type for the
// Encrypt API operation.
type EncryptResponse struct {
	*types.EncryptOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// Encrypt request.
func (r *EncryptResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
