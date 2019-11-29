// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
)

const opPutConferencePreference = "PutConferencePreference"

// PutConferencePreferenceRequest returns a request value for making API operation for
// Alexa For Business.
//
// Sets the conference preferences on a specific conference provider at the
// account level.
//
//    // Example sending a request using PutConferencePreferenceRequest.
//    req := client.PutConferencePreferenceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/PutConferencePreference
func (c *Client) PutConferencePreferenceRequest(input *types.PutConferencePreferenceInput) PutConferencePreferenceRequest {
	op := &aws.Operation{
		Name:       opPutConferencePreference,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutConferencePreferenceInput{}
	}

	req := c.newRequest(op, input, &types.PutConferencePreferenceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.PutConferencePreferenceMarshaler{Input: input}.GetNamedBuildHandler())

	return PutConferencePreferenceRequest{Request: req, Input: input, Copy: c.PutConferencePreferenceRequest}
}

// PutConferencePreferenceRequest is the request type for the
// PutConferencePreference API operation.
type PutConferencePreferenceRequest struct {
	*aws.Request
	Input *types.PutConferencePreferenceInput
	Copy  func(*types.PutConferencePreferenceInput) PutConferencePreferenceRequest
}

// Send marshals and sends the PutConferencePreference API request.
func (r PutConferencePreferenceRequest) Send(ctx context.Context) (*PutConferencePreferenceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutConferencePreferenceResponse{
		PutConferencePreferenceOutput: r.Request.Data.(*types.PutConferencePreferenceOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutConferencePreferenceResponse is the response type for the
// PutConferencePreference API operation.
type PutConferencePreferenceResponse struct {
	*types.PutConferencePreferenceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutConferencePreference request.
func (r *PutConferencePreferenceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
