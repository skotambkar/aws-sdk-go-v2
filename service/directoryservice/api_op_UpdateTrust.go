// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
)

const opUpdateTrust = "UpdateTrust"

// UpdateTrustRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Updates the trust that has been set up between your AWS Managed Microsoft
// AD directory and an on-premises Active Directory.
//
//    // Example sending a request using UpdateTrustRequest.
//    req := client.UpdateTrustRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/UpdateTrust
func (c *Client) UpdateTrustRequest(input *types.UpdateTrustInput) UpdateTrustRequest {
	op := &aws.Operation{
		Name:       opUpdateTrust,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateTrustInput{}
	}

	req := c.newRequest(op, input, &types.UpdateTrustOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateTrustMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateTrustRequest{Request: req, Input: input, Copy: c.UpdateTrustRequest}
}

// UpdateTrustRequest is the request type for the
// UpdateTrust API operation.
type UpdateTrustRequest struct {
	*aws.Request
	Input *types.UpdateTrustInput
	Copy  func(*types.UpdateTrustInput) UpdateTrustRequest
}

// Send marshals and sends the UpdateTrust API request.
func (r UpdateTrustRequest) Send(ctx context.Context) (*UpdateTrustResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTrustResponse{
		UpdateTrustOutput: r.Request.Data.(*types.UpdateTrustOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTrustResponse is the response type for the
// UpdateTrust API operation.
type UpdateTrustResponse struct {
	*types.UpdateTrustOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTrust request.
func (r *UpdateTrustResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
