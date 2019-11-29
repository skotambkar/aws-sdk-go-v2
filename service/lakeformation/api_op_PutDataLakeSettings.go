// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
)

const opPutDataLakeSettings = "PutDataLakeSettings"

// PutDataLakeSettingsRequest returns a request value for making API operation for
// AWS Lake Formation.
//
// The AWS Lake Formation principal.
//
//    // Example sending a request using PutDataLakeSettingsRequest.
//    req := client.PutDataLakeSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lakeformation-2017-03-31/PutDataLakeSettings
func (c *Client) PutDataLakeSettingsRequest(input *types.PutDataLakeSettingsInput) PutDataLakeSettingsRequest {
	op := &aws.Operation{
		Name:       opPutDataLakeSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutDataLakeSettingsInput{}
	}

	req := c.newRequest(op, input, &types.PutDataLakeSettingsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.PutDataLakeSettingsMarshaler{Input: input}.GetNamedBuildHandler())

	return PutDataLakeSettingsRequest{Request: req, Input: input, Copy: c.PutDataLakeSettingsRequest}
}

// PutDataLakeSettingsRequest is the request type for the
// PutDataLakeSettings API operation.
type PutDataLakeSettingsRequest struct {
	*aws.Request
	Input *types.PutDataLakeSettingsInput
	Copy  func(*types.PutDataLakeSettingsInput) PutDataLakeSettingsRequest
}

// Send marshals and sends the PutDataLakeSettings API request.
func (r PutDataLakeSettingsRequest) Send(ctx context.Context) (*PutDataLakeSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutDataLakeSettingsResponse{
		PutDataLakeSettingsOutput: r.Request.Data.(*types.PutDataLakeSettingsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutDataLakeSettingsResponse is the response type for the
// PutDataLakeSettings API operation.
type PutDataLakeSettingsResponse struct {
	*types.PutDataLakeSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutDataLakeSettings request.
func (r *PutDataLakeSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
