// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transfer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type StopServerInput struct {
	_ struct{} `type:"structure"`

	// A system-assigned unique identifier for an SFTP server that you stopped.
	//
	// ServerId is a required field
	ServerId *string `min:"19" type:"string" required:"true"`
}

// String returns the string representation
func (s StopServerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopServerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopServerInput"}

	if s.ServerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerId"))
	}
	if s.ServerId != nil && len(*s.ServerId) < 19 {
		invalidParams.Add(aws.NewErrParamMinLen("ServerId", 19))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopServerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopServerOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopServer = "StopServer"

// StopServerRequest returns a request value for making API operation for
// AWS Transfer for SFTP.
//
// Changes the state of an SFTP server from ONLINE to OFFLINE. An OFFLINE server
// cannot accept and process file transfer jobs. Information tied to your server
// such as server and user properties are not affected by stopping your server.
// Stopping a server will not reduce or impact your Secure File Transfer Protocol
// (SFTP) endpoint billing.
//
// The state of STOPPING indicates that the server is in an intermediate state,
// either not fully able to respond, or not fully offline. The values of STOP_FAILED
// can indicate an error condition.
//
// No response is returned from this call.
//
//    // Example sending a request using StopServerRequest.
//    req := client.StopServerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/StopServer
func (c *Client) StopServerRequest(input *StopServerInput) StopServerRequest {
	op := &aws.Operation{
		Name:       opStopServer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopServerInput{}
	}

	req := c.newRequest(op, input, &StopServerOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return StopServerRequest{Request: req, Input: input, Copy: c.StopServerRequest}
}

// StopServerRequest is the request type for the
// StopServer API operation.
type StopServerRequest struct {
	*aws.Request
	Input *StopServerInput
	Copy  func(*StopServerInput) StopServerRequest
}

// Send marshals and sends the StopServer API request.
func (r StopServerRequest) Send(ctx context.Context) (*StopServerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopServerResponse{
		StopServerOutput: r.Request.Data.(*StopServerOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopServerResponse is the response type for the
// StopServer API operation.
type StopServerResponse struct {
	*StopServerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopServer request.
func (r *StopServerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
