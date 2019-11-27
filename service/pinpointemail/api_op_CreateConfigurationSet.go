// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
)

const opCreateConfigurationSet = "CreateConfigurationSet"

// CreateConfigurationSetRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Create a configuration set. Configuration sets are groups of rules that you
// can apply to the emails you send using Amazon Pinpoint. You apply a configuration
// set to an email by including a reference to the configuration set in the
// headers of the email. When you apply a configuration set to an email, all
// of the rules in that configuration set are applied to the email.
//
//    // Example sending a request using CreateConfigurationSetRequest.
//    req := client.CreateConfigurationSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/CreateConfigurationSet
func (c *Client) CreateConfigurationSetRequest(input *types.CreateConfigurationSetInput) CreateConfigurationSetRequest {
	op := &aws.Operation{
		Name:       opCreateConfigurationSet,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/email/configuration-sets",
	}

	if input == nil {
		input = &types.CreateConfigurationSetInput{}
	}

	req := c.newRequest(op, input, &types.CreateConfigurationSetOutput{})
	return CreateConfigurationSetRequest{Request: req, Input: input, Copy: c.CreateConfigurationSetRequest}
}

// CreateConfigurationSetRequest is the request type for the
// CreateConfigurationSet API operation.
type CreateConfigurationSetRequest struct {
	*aws.Request
	Input *types.CreateConfigurationSetInput
	Copy  func(*types.CreateConfigurationSetInput) CreateConfigurationSetRequest
}

// Send marshals and sends the CreateConfigurationSet API request.
func (r CreateConfigurationSetRequest) Send(ctx context.Context) (*CreateConfigurationSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateConfigurationSetResponse{
		CreateConfigurationSetOutput: r.Request.Data.(*types.CreateConfigurationSetOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateConfigurationSetResponse is the response type for the
// CreateConfigurationSet API operation.
type CreateConfigurationSetResponse struct {
	*types.CreateConfigurationSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateConfigurationSet request.
func (r *CreateConfigurationSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
