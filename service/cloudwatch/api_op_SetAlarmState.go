// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

const opSetAlarmState = "SetAlarmState"

// SetAlarmStateRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// Temporarily sets the state of an alarm for testing purposes. When the updated
// state differs from the previous value, the action configured for the appropriate
// state is invoked. For example, if your alarm is configured to send an Amazon
// SNS message when an alarm is triggered, temporarily changing the alarm state
// to ALARM sends an SNS message. The alarm returns to its actual state (often
// within seconds). Because the alarm state change happens quickly, it is typically
// only visible in the alarm's History tab in the Amazon CloudWatch console
// or through DescribeAlarmHistory.
//
//    // Example sending a request using SetAlarmStateRequest.
//    req := client.SetAlarmStateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/SetAlarmState
func (c *Client) SetAlarmStateRequest(input *types.SetAlarmStateInput) SetAlarmStateRequest {
	op := &aws.Operation{
		Name:       opSetAlarmState,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.SetAlarmStateInput{}
	}

	req := c.newRequest(op, input, &types.SetAlarmStateOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.SetAlarmStateMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetAlarmStateRequest{Request: req, Input: input, Copy: c.SetAlarmStateRequest}
}

// SetAlarmStateRequest is the request type for the
// SetAlarmState API operation.
type SetAlarmStateRequest struct {
	*aws.Request
	Input *types.SetAlarmStateInput
	Copy  func(*types.SetAlarmStateInput) SetAlarmStateRequest
}

// Send marshals and sends the SetAlarmState API request.
func (r SetAlarmStateRequest) Send(ctx context.Context) (*SetAlarmStateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetAlarmStateResponse{
		SetAlarmStateOutput: r.Request.Data.(*types.SetAlarmStateOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetAlarmStateResponse is the response type for the
// SetAlarmState API operation.
type SetAlarmStateResponse struct {
	*types.SetAlarmStateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetAlarmState request.
func (r *SetAlarmStateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
