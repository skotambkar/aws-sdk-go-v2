// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transfer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/transfer/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
)

const opStartServer = "StartServer"

// StartServerRequest returns a request value for making API operation for
// AWS Transfer for SFTP.
//
// Changes the state of a Secure File Transfer Protocol (SFTP) server from OFFLINE
// to ONLINE. It has no impact on an SFTP server that is already ONLINE. An
// ONLINE server can accept and process file transfer jobs.
//
// The state of STARTING indicates that the server is in an intermediate state,
// either not fully able to respond, or not fully online. The values of START_FAILED
// can indicate an error condition.
//
// No response is returned from this call.
//
//    // Example sending a request using StartServerRequest.
//    req := client.StartServerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/StartServer
func (c *Client) StartServerRequest(input *types.StartServerInput) StartServerRequest {
	op := &aws.Operation{
		Name:       opStartServer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StartServerInput{}
	}

	req := c.newRequest(op, input, &types.StartServerOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.StartServerMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return StartServerRequest{Request: req, Input: input, Copy: c.StartServerRequest}
}

// StartServerRequest is the request type for the
// StartServer API operation.
type StartServerRequest struct {
	*aws.Request
	Input *types.StartServerInput
	Copy  func(*types.StartServerInput) StartServerRequest
}

// Send marshals and sends the StartServer API request.
func (r StartServerRequest) Send(ctx context.Context) (*StartServerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartServerResponse{
		StartServerOutput: r.Request.Data.(*types.StartServerOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartServerResponse is the response type for the
// StartServer API operation.
type StartServerResponse struct {
	*types.StartServerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartServer request.
func (r *StartServerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
