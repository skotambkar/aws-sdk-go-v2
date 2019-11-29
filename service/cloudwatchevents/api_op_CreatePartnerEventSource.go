// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchevents

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
)

const opCreatePartnerEventSource = "CreatePartnerEventSource"

// CreatePartnerEventSourceRequest returns a request value for making API operation for
// Amazon CloudWatch Events.
//
// Called by an SaaS partner to create a partner event source.
//
// This operation is not used by AWS customers.
//
// Each partner event source can be used by one AWS account to create a matching
// partner event bus in that AWS account. A SaaS partner must create one partner
// event source for each AWS account that wants to receive those event types.
//
// A partner event source creates events based on resources in the SaaS partner's
// service or application.
//
// An AWS account that creates a partner event bus that matches the partner
// event source can use that event bus to receive events from the partner, and
// then process them using AWS Events rules and targets.
//
// Partner event source names follow this format:
//
// aws.partner/partner_name/event_namespace/event_name
//
//    * partner_name is determined during partner registration and identifies
//    the partner to AWS customers.
//
//    * For event_namespace, we recommend that partners use a string that identifies
//    the AWS customer within the partner's system. This should not be the customer's
//    AWS account ID.
//
//    * event_name is determined by the partner, and should uniquely identify
//    an event-generating resource within the partner system. This should help
//    AWS customers decide whether to create an event bus to receive these events.
//
//    // Example sending a request using CreatePartnerEventSourceRequest.
//    req := client.CreatePartnerEventSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/events-2015-10-07/CreatePartnerEventSource
func (c *Client) CreatePartnerEventSourceRequest(input *types.CreatePartnerEventSourceInput) CreatePartnerEventSourceRequest {
	op := &aws.Operation{
		Name:       opCreatePartnerEventSource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreatePartnerEventSourceInput{}
	}

	req := c.newRequest(op, input, &types.CreatePartnerEventSourceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreatePartnerEventSourceMarshaler{Input: input}.GetNamedBuildHandler())

	return CreatePartnerEventSourceRequest{Request: req, Input: input, Copy: c.CreatePartnerEventSourceRequest}
}

// CreatePartnerEventSourceRequest is the request type for the
// CreatePartnerEventSource API operation.
type CreatePartnerEventSourceRequest struct {
	*aws.Request
	Input *types.CreatePartnerEventSourceInput
	Copy  func(*types.CreatePartnerEventSourceInput) CreatePartnerEventSourceRequest
}

// Send marshals and sends the CreatePartnerEventSource API request.
func (r CreatePartnerEventSourceRequest) Send(ctx context.Context) (*CreatePartnerEventSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePartnerEventSourceResponse{
		CreatePartnerEventSourceOutput: r.Request.Data.(*types.CreatePartnerEventSourceOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePartnerEventSourceResponse is the response type for the
// CreatePartnerEventSource API operation.
type CreatePartnerEventSourceResponse struct {
	*types.CreatePartnerEventSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePartnerEventSource request.
func (r *CreatePartnerEventSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
