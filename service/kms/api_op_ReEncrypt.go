// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kms/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

const opReEncrypt = "ReEncrypt"

// ReEncryptRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Encrypts data on the server side with a new customer master key (CMK) without
// exposing the plaintext of the data on the client side. The data is first
// decrypted and then reencrypted. You can also use this operation to change
// the encryption context of a ciphertext.
//
// You can reencrypt data using CMKs in different AWS accounts.
//
// Unlike other operations, ReEncrypt is authorized twice, once as ReEncryptFrom
// on the source CMK and once as ReEncryptTo on the destination CMK. We recommend
// that you include the "kms:ReEncrypt*" permission in your key policies (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html)
// to permit reencryption from or to the CMK. This permission is automatically
// included in the key policy when you create a CMK through the console. But
// you must include it manually when you create a CMK programmatically or when
// you set a key policy with the PutKeyPolicy operation.
//
// The result of this operation varies with the key state of the CMK. For details,
// see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using ReEncryptRequest.
//    req := client.ReEncryptRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/ReEncrypt
func (c *Client) ReEncryptRequest(input *types.ReEncryptInput) ReEncryptRequest {
	op := &aws.Operation{
		Name:       opReEncrypt,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ReEncryptInput{}
	}

	req := c.newRequest(op, input, &types.ReEncryptOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ReEncryptMarshaler{Input: input}.GetNamedBuildHandler())

	return ReEncryptRequest{Request: req, Input: input, Copy: c.ReEncryptRequest}
}

// ReEncryptRequest is the request type for the
// ReEncrypt API operation.
type ReEncryptRequest struct {
	*aws.Request
	Input *types.ReEncryptInput
	Copy  func(*types.ReEncryptInput) ReEncryptRequest
}

// Send marshals and sends the ReEncrypt API request.
func (r ReEncryptRequest) Send(ctx context.Context) (*ReEncryptResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReEncryptResponse{
		ReEncryptOutput: r.Request.Data.(*types.ReEncryptOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReEncryptResponse is the response type for the
// ReEncrypt API operation.
type ReEncryptResponse struct {
	*types.ReEncryptOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReEncrypt request.
func (r *ReEncryptResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
