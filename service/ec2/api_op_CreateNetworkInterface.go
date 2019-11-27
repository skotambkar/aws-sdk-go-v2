// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opCreateNetworkInterface = "CreateNetworkInterface"

// CreateNetworkInterfaceRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a network interface in the specified subnet.
//
// For more information about network interfaces, see Elastic Network Interfaces
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html) in the
// Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using CreateNetworkInterfaceRequest.
//    req := client.CreateNetworkInterfaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateNetworkInterface
func (c *Client) CreateNetworkInterfaceRequest(input *types.CreateNetworkInterfaceInput) CreateNetworkInterfaceRequest {
	op := &aws.Operation{
		Name:       opCreateNetworkInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateNetworkInterfaceInput{}
	}

	req := c.newRequest(op, input, &types.CreateNetworkInterfaceOutput{})
	return CreateNetworkInterfaceRequest{Request: req, Input: input, Copy: c.CreateNetworkInterfaceRequest}
}

// CreateNetworkInterfaceRequest is the request type for the
// CreateNetworkInterface API operation.
type CreateNetworkInterfaceRequest struct {
	*aws.Request
	Input *types.CreateNetworkInterfaceInput
	Copy  func(*types.CreateNetworkInterfaceInput) CreateNetworkInterfaceRequest
}

// Send marshals and sends the CreateNetworkInterface API request.
func (r CreateNetworkInterfaceRequest) Send(ctx context.Context) (*CreateNetworkInterfaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateNetworkInterfaceResponse{
		CreateNetworkInterfaceOutput: r.Request.Data.(*types.CreateNetworkInterfaceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateNetworkInterfaceResponse is the response type for the
// CreateNetworkInterface API operation.
type CreateNetworkInterfaceResponse struct {
	*types.CreateNetworkInterfaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateNetworkInterface request.
func (r *CreateNetworkInterfaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
