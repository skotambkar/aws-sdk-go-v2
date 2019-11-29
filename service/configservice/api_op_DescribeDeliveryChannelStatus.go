// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/configservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
)

const opDescribeDeliveryChannelStatus = "DescribeDeliveryChannelStatus"

// DescribeDeliveryChannelStatusRequest returns a request value for making API operation for
// AWS Config.
//
// Returns the current status of the specified delivery channel. If a delivery
// channel is not specified, this action returns the current status of all delivery
// channels associated with the account.
//
// Currently, you can specify only one delivery channel per region in your account.
//
//    // Example sending a request using DescribeDeliveryChannelStatusRequest.
//    req := client.DescribeDeliveryChannelStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DescribeDeliveryChannelStatus
func (c *Client) DescribeDeliveryChannelStatusRequest(input *types.DescribeDeliveryChannelStatusInput) DescribeDeliveryChannelStatusRequest {
	op := &aws.Operation{
		Name:       opDescribeDeliveryChannelStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeDeliveryChannelStatusInput{}
	}

	req := c.newRequest(op, input, &types.DescribeDeliveryChannelStatusOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeDeliveryChannelStatusMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeDeliveryChannelStatusRequest{Request: req, Input: input, Copy: c.DescribeDeliveryChannelStatusRequest}
}

// DescribeDeliveryChannelStatusRequest is the request type for the
// DescribeDeliveryChannelStatus API operation.
type DescribeDeliveryChannelStatusRequest struct {
	*aws.Request
	Input *types.DescribeDeliveryChannelStatusInput
	Copy  func(*types.DescribeDeliveryChannelStatusInput) DescribeDeliveryChannelStatusRequest
}

// Send marshals and sends the DescribeDeliveryChannelStatus API request.
func (r DescribeDeliveryChannelStatusRequest) Send(ctx context.Context) (*DescribeDeliveryChannelStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDeliveryChannelStatusResponse{
		DescribeDeliveryChannelStatusOutput: r.Request.Data.(*types.DescribeDeliveryChannelStatusOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDeliveryChannelStatusResponse is the response type for the
// DescribeDeliveryChannelStatus API operation.
type DescribeDeliveryChannelStatusResponse struct {
	*types.DescribeDeliveryChannelStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDeliveryChannelStatus request.
func (r *DescribeDeliveryChannelStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
