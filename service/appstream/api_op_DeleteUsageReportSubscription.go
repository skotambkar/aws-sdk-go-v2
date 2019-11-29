// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/appstream/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
)

const opDeleteUsageReportSubscription = "DeleteUsageReportSubscription"

// DeleteUsageReportSubscriptionRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Disables usage report generation.
//
//    // Example sending a request using DeleteUsageReportSubscriptionRequest.
//    req := client.DeleteUsageReportSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DeleteUsageReportSubscription
func (c *Client) DeleteUsageReportSubscriptionRequest(input *types.DeleteUsageReportSubscriptionInput) DeleteUsageReportSubscriptionRequest {
	op := &aws.Operation{
		Name:       opDeleteUsageReportSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteUsageReportSubscriptionInput{}
	}

	req := c.newRequest(op, input, &types.DeleteUsageReportSubscriptionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteUsageReportSubscriptionMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteUsageReportSubscriptionRequest{Request: req, Input: input, Copy: c.DeleteUsageReportSubscriptionRequest}
}

// DeleteUsageReportSubscriptionRequest is the request type for the
// DeleteUsageReportSubscription API operation.
type DeleteUsageReportSubscriptionRequest struct {
	*aws.Request
	Input *types.DeleteUsageReportSubscriptionInput
	Copy  func(*types.DeleteUsageReportSubscriptionInput) DeleteUsageReportSubscriptionRequest
}

// Send marshals and sends the DeleteUsageReportSubscription API request.
func (r DeleteUsageReportSubscriptionRequest) Send(ctx context.Context) (*DeleteUsageReportSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteUsageReportSubscriptionResponse{
		DeleteUsageReportSubscriptionOutput: r.Request.Data.(*types.DeleteUsageReportSubscriptionOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteUsageReportSubscriptionResponse is the response type for the
// DeleteUsageReportSubscription API operation.
type DeleteUsageReportSubscriptionResponse struct {
	*types.DeleteUsageReportSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteUsageReportSubscription request.
func (r *DeleteUsageReportSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
