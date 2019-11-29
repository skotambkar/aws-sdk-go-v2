// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

const opEnableAlarmActions = "EnableAlarmActions"

// EnableAlarmActionsRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// Enables the actions for the specified alarms.
//
//    // Example sending a request using EnableAlarmActionsRequest.
//    req := client.EnableAlarmActionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/EnableAlarmActions
func (c *Client) EnableAlarmActionsRequest(input *types.EnableAlarmActionsInput) EnableAlarmActionsRequest {
	op := &aws.Operation{
		Name:       opEnableAlarmActions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.EnableAlarmActionsInput{}
	}

	req := c.newRequest(op, input, &types.EnableAlarmActionsOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return EnableAlarmActionsRequest{Request: req, Input: input, Copy: c.EnableAlarmActionsRequest}
}

// EnableAlarmActionsRequest is the request type for the
// EnableAlarmActions API operation.
type EnableAlarmActionsRequest struct {
	*aws.Request
	Input *types.EnableAlarmActionsInput
	Copy  func(*types.EnableAlarmActionsInput) EnableAlarmActionsRequest
}

// Send marshals and sends the EnableAlarmActions API request.
func (r EnableAlarmActionsRequest) Send(ctx context.Context) (*EnableAlarmActionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableAlarmActionsResponse{
		EnableAlarmActionsOutput: r.Request.Data.(*types.EnableAlarmActionsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableAlarmActionsResponse is the response type for the
// EnableAlarmActions API operation.
type EnableAlarmActionsResponse struct {
	*types.EnableAlarmActionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableAlarmActions request.
func (r *EnableAlarmActionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
