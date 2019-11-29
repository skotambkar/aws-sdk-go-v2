// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/ses/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

const opCreateConfigurationSet = "CreateConfigurationSet"

// CreateConfigurationSetRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Creates a configuration set.
//
// Configuration sets enable you to publish email sending events. For information
// about using configuration sets, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using CreateConfigurationSetRequest.
//    req := client.CreateConfigurationSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/CreateConfigurationSet
func (c *Client) CreateConfigurationSetRequest(input *types.CreateConfigurationSetInput) CreateConfigurationSetRequest {
	op := &aws.Operation{
		Name:       opCreateConfigurationSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateConfigurationSetInput{}
	}

	req := c.newRequest(op, input, &types.CreateConfigurationSetOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.CreateConfigurationSetMarshaler{Input: input}.GetNamedBuildHandler())

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
