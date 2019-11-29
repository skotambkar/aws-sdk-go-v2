// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/sns/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

const opSetSMSAttributes = "SetSMSAttributes"

// SetSMSAttributesRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Use this request to set the default settings for sending SMS messages and
// receiving daily SMS usage reports.
//
// You can override some of these settings for a single message when you use
// the Publish action with the MessageAttributes.entry.N parameter. For more
// information, see Sending an SMS Message (https://docs.aws.amazon.com/sns/latest/dg/sms_publish-to-phone.html)
// in the Amazon SNS Developer Guide.
//
//    // Example sending a request using SetSMSAttributesRequest.
//    req := client.SetSMSAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/SetSMSAttributes
func (c *Client) SetSMSAttributesRequest(input *types.SetSMSAttributesInput) SetSMSAttributesRequest {
	op := &aws.Operation{
		Name:       opSetSMSAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.SetSMSAttributesInput{}
	}

	req := c.newRequest(op, input, &types.SetSMSAttributesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.SetSMSAttributesMarshaler{Input: input}.GetNamedBuildHandler())

	return SetSMSAttributesRequest{Request: req, Input: input, Copy: c.SetSMSAttributesRequest}
}

// SetSMSAttributesRequest is the request type for the
// SetSMSAttributes API operation.
type SetSMSAttributesRequest struct {
	*aws.Request
	Input *types.SetSMSAttributesInput
	Copy  func(*types.SetSMSAttributesInput) SetSMSAttributesRequest
}

// Send marshals and sends the SetSMSAttributes API request.
func (r SetSMSAttributesRequest) Send(ctx context.Context) (*SetSMSAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetSMSAttributesResponse{
		SetSMSAttributesOutput: r.Request.Data.(*types.SetSMSAttributesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetSMSAttributesResponse is the response type for the
// SetSMSAttributes API operation.
type SetSMSAttributesResponse struct {
	*types.SetSMSAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetSMSAttributes request.
func (r *SetSMSAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
