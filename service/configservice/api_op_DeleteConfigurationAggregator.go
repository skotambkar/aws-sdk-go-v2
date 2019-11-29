// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
)

const opDeleteConfigurationAggregator = "DeleteConfigurationAggregator"

// DeleteConfigurationAggregatorRequest returns a request value for making API operation for
// AWS Config.
//
// Deletes the specified configuration aggregator and the aggregated data associated
// with the aggregator.
//
//    // Example sending a request using DeleteConfigurationAggregatorRequest.
//    req := client.DeleteConfigurationAggregatorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DeleteConfigurationAggregator
func (c *Client) DeleteConfigurationAggregatorRequest(input *types.DeleteConfigurationAggregatorInput) DeleteConfigurationAggregatorRequest {
	op := &aws.Operation{
		Name:       opDeleteConfigurationAggregator,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteConfigurationAggregatorInput{}
	}

	req := c.newRequest(op, input, &types.DeleteConfigurationAggregatorOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteConfigurationAggregatorRequest{Request: req, Input: input, Copy: c.DeleteConfigurationAggregatorRequest}
}

// DeleteConfigurationAggregatorRequest is the request type for the
// DeleteConfigurationAggregator API operation.
type DeleteConfigurationAggregatorRequest struct {
	*aws.Request
	Input *types.DeleteConfigurationAggregatorInput
	Copy  func(*types.DeleteConfigurationAggregatorInput) DeleteConfigurationAggregatorRequest
}

// Send marshals and sends the DeleteConfigurationAggregator API request.
func (r DeleteConfigurationAggregatorRequest) Send(ctx context.Context) (*DeleteConfigurationAggregatorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConfigurationAggregatorResponse{
		DeleteConfigurationAggregatorOutput: r.Request.Data.(*types.DeleteConfigurationAggregatorOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConfigurationAggregatorResponse is the response type for the
// DeleteConfigurationAggregator API operation.
type DeleteConfigurationAggregatorResponse struct {
	*types.DeleteConfigurationAggregatorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConfigurationAggregator request.
func (r *DeleteConfigurationAggregatorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
