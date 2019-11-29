// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sesv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

const opDeleteConfigurationSet = "DeleteConfigurationSet"

// DeleteConfigurationSetRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Delete an existing configuration set.
//
// Configuration sets are groups of rules that you can apply to the emails you
// send. You apply a configuration set to an email by including a reference
// to the configuration set in the headers of the email. When you apply a configuration
// set to an email, all of the rules in that configuration set are applied to
// the email.
//
//    // Example sending a request using DeleteConfigurationSetRequest.
//    req := client.DeleteConfigurationSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sesv2-2019-09-27/DeleteConfigurationSet
func (c *Client) DeleteConfigurationSetRequest(input *types.DeleteConfigurationSetInput) DeleteConfigurationSetRequest {
	op := &aws.Operation{
		Name:       opDeleteConfigurationSet,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v2/email/configuration-sets/{ConfigurationSetName}",
	}

	if input == nil {
		input = &types.DeleteConfigurationSetInput{}
	}

	req := c.newRequest(op, input, &types.DeleteConfigurationSetOutput{})
	return DeleteConfigurationSetRequest{Request: req, Input: input, Copy: c.DeleteConfigurationSetRequest}
}

// DeleteConfigurationSetRequest is the request type for the
// DeleteConfigurationSet API operation.
type DeleteConfigurationSetRequest struct {
	*aws.Request
	Input *types.DeleteConfigurationSetInput
	Copy  func(*types.DeleteConfigurationSetInput) DeleteConfigurationSetRequest
}

// Send marshals and sends the DeleteConfigurationSet API request.
func (r DeleteConfigurationSetRequest) Send(ctx context.Context) (*DeleteConfigurationSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConfigurationSetResponse{
		DeleteConfigurationSetOutput: r.Request.Data.(*types.DeleteConfigurationSetOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConfigurationSetResponse is the response type for the
// DeleteConfigurationSet API operation.
type DeleteConfigurationSetResponse struct {
	*types.DeleteConfigurationSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConfigurationSet request.
func (r *DeleteConfigurationSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
