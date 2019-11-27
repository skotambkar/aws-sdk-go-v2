// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDeleteNetworkInterface = "DeleteNetworkInterface"

// DeleteNetworkInterfaceRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified network interface. You must detach the network interface
// before you can delete it.
//
//    // Example sending a request using DeleteNetworkInterfaceRequest.
//    req := client.DeleteNetworkInterfaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteNetworkInterface
func (c *Client) DeleteNetworkInterfaceRequest(input *types.DeleteNetworkInterfaceInput) DeleteNetworkInterfaceRequest {
	op := &aws.Operation{
		Name:       opDeleteNetworkInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteNetworkInterfaceInput{}
	}

	req := c.newRequest(op, input, &types.DeleteNetworkInterfaceOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteNetworkInterfaceRequest{Request: req, Input: input, Copy: c.DeleteNetworkInterfaceRequest}
}

// DeleteNetworkInterfaceRequest is the request type for the
// DeleteNetworkInterface API operation.
type DeleteNetworkInterfaceRequest struct {
	*aws.Request
	Input *types.DeleteNetworkInterfaceInput
	Copy  func(*types.DeleteNetworkInterfaceInput) DeleteNetworkInterfaceRequest
}

// Send marshals and sends the DeleteNetworkInterface API request.
func (r DeleteNetworkInterfaceRequest) Send(ctx context.Context) (*DeleteNetworkInterfaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteNetworkInterfaceResponse{
		DeleteNetworkInterfaceOutput: r.Request.Data.(*types.DeleteNetworkInterfaceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteNetworkInterfaceResponse is the response type for the
// DeleteNetworkInterface API operation.
type DeleteNetworkInterfaceResponse struct {
	*types.DeleteNetworkInterfaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteNetworkInterface request.
func (r *DeleteNetworkInterfaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
