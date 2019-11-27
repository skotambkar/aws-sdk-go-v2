// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
)

const opStartStreamEncryption = "StartStreamEncryption"

// StartStreamEncryptionRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// Enables or updates server-side encryption using an AWS KMS key for a specified
// stream.
//
// Starting encryption is an asynchronous operation. Upon receiving the request,
// Kinesis Data Streams returns immediately and sets the status of the stream
// to UPDATING. After the update is complete, Kinesis Data Streams sets the
// status of the stream back to ACTIVE. Updating or applying encryption normally
// takes a few seconds to complete, but it can take minutes. You can continue
// to read and write data to your stream while its status is UPDATING. Once
// the status of the stream is ACTIVE, encryption begins for records written
// to the stream.
//
// API Limits: You can successfully apply a new AWS KMS key for server-side
// encryption 25 times in a rolling 24-hour period.
//
// Note: It can take up to 5 seconds after the stream is in an ACTIVE status
// before all records written to the stream are encrypted. After you enable
// encryption, you can verify that encryption is applied by inspecting the API
// response from PutRecord or PutRecords.
//
//    // Example sending a request using StartStreamEncryptionRequest.
//    req := client.StartStreamEncryptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/StartStreamEncryption
func (c *Client) StartStreamEncryptionRequest(input *types.StartStreamEncryptionInput) StartStreamEncryptionRequest {
	op := &aws.Operation{
		Name:       opStartStreamEncryption,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StartStreamEncryptionInput{}
	}

	req := c.newRequest(op, input, &types.StartStreamEncryptionOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return StartStreamEncryptionRequest{Request: req, Input: input, Copy: c.StartStreamEncryptionRequest}
}

// StartStreamEncryptionRequest is the request type for the
// StartStreamEncryption API operation.
type StartStreamEncryptionRequest struct {
	*aws.Request
	Input *types.StartStreamEncryptionInput
	Copy  func(*types.StartStreamEncryptionInput) StartStreamEncryptionRequest
}

// Send marshals and sends the StartStreamEncryption API request.
func (r StartStreamEncryptionRequest) Send(ctx context.Context) (*StartStreamEncryptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartStreamEncryptionResponse{
		StartStreamEncryptionOutput: r.Request.Data.(*types.StartStreamEncryptionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartStreamEncryptionResponse is the response type for the
// StartStreamEncryption API operation.
type StartStreamEncryptionResponse struct {
	*types.StartStreamEncryptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartStreamEncryption request.
func (r *StartStreamEncryptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
