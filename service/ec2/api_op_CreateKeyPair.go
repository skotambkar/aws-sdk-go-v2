// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opCreateKeyPair = "CreateKeyPair"

// CreateKeyPairRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a 2048-bit RSA key pair with the specified name. Amazon EC2 stores
// the public key and displays the private key for you to save to a file. The
// private key is returned as an unencrypted PEM encoded PKCS#1 private key.
// If a key with the specified name already exists, Amazon EC2 returns an error.
//
// You can have up to five thousand key pairs per Region.
//
// The key pair returned to you is available only in the Region in which you
// create it. If you prefer, you can create your own key pair using a third-party
// tool and upload it to any Region using ImportKeyPair.
//
// For more information, see Key Pairs (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using CreateKeyPairRequest.
//    req := client.CreateKeyPairRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateKeyPair
func (c *Client) CreateKeyPairRequest(input *types.CreateKeyPairInput) CreateKeyPairRequest {
	op := &aws.Operation{
		Name:       opCreateKeyPair,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateKeyPairInput{}
	}

	req := c.newRequest(op, input, &types.CreateKeyPairOutput{})
	return CreateKeyPairRequest{Request: req, Input: input, Copy: c.CreateKeyPairRequest}
}

// CreateKeyPairRequest is the request type for the
// CreateKeyPair API operation.
type CreateKeyPairRequest struct {
	*aws.Request
	Input *types.CreateKeyPairInput
	Copy  func(*types.CreateKeyPairInput) CreateKeyPairRequest
}

// Send marshals and sends the CreateKeyPair API request.
func (r CreateKeyPairRequest) Send(ctx context.Context) (*CreateKeyPairResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateKeyPairResponse{
		CreateKeyPairOutput: r.Request.Data.(*types.CreateKeyPairOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateKeyPairResponse is the response type for the
// CreateKeyPair API operation.
type CreateKeyPairResponse struct {
	*types.CreateKeyPairOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateKeyPair request.
func (r *CreateKeyPairResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
