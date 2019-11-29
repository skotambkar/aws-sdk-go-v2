// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/sns/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

const opGetSMSAttributes = "GetSMSAttributes"

// GetSMSAttributesRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Returns the settings for sending SMS messages from your account.
//
// These settings are set with the SetSMSAttributes action.
//
//    // Example sending a request using GetSMSAttributesRequest.
//    req := client.GetSMSAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/GetSMSAttributes
func (c *Client) GetSMSAttributesRequest(input *types.GetSMSAttributesInput) GetSMSAttributesRequest {
	op := &aws.Operation{
		Name:       opGetSMSAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetSMSAttributesInput{}
	}

	req := c.newRequest(op, input, &types.GetSMSAttributesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.GetSMSAttributesMarshaler{Input: input}.GetNamedBuildHandler())

	return GetSMSAttributesRequest{Request: req, Input: input, Copy: c.GetSMSAttributesRequest}
}

// GetSMSAttributesRequest is the request type for the
// GetSMSAttributes API operation.
type GetSMSAttributesRequest struct {
	*aws.Request
	Input *types.GetSMSAttributesInput
	Copy  func(*types.GetSMSAttributesInput) GetSMSAttributesRequest
}

// Send marshals and sends the GetSMSAttributes API request.
func (r GetSMSAttributesRequest) Send(ctx context.Context) (*GetSMSAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSMSAttributesResponse{
		GetSMSAttributesOutput: r.Request.Data.(*types.GetSMSAttributesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSMSAttributesResponse is the response type for the
// GetSMSAttributes API operation.
type GetSMSAttributesResponse struct {
	*types.GetSMSAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSMSAttributes request.
func (r *GetSMSAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
