// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opCreateRouteTable = "CreateRouteTable"

// CreateRouteTableRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a route table for the specified VPC. After you create a route table,
// you can add routes and associate the table with a subnet.
//
// For more information, see Route Tables (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using CreateRouteTableRequest.
//    req := client.CreateRouteTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateRouteTable
func (c *Client) CreateRouteTableRequest(input *types.CreateRouteTableInput) CreateRouteTableRequest {
	op := &aws.Operation{
		Name:       opCreateRouteTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateRouteTableInput{}
	}

	req := c.newRequest(op, input, &types.CreateRouteTableOutput{})
	return CreateRouteTableRequest{Request: req, Input: input, Copy: c.CreateRouteTableRequest}
}

// CreateRouteTableRequest is the request type for the
// CreateRouteTable API operation.
type CreateRouteTableRequest struct {
	*aws.Request
	Input *types.CreateRouteTableInput
	Copy  func(*types.CreateRouteTableInput) CreateRouteTableRequest
}

// Send marshals and sends the CreateRouteTable API request.
func (r CreateRouteTableRequest) Send(ctx context.Context) (*CreateRouteTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRouteTableResponse{
		CreateRouteTableOutput: r.Request.Data.(*types.CreateRouteTableOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRouteTableResponse is the response type for the
// CreateRouteTable API operation.
type CreateRouteTableResponse struct {
	*types.CreateRouteTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRouteTable request.
func (r *CreateRouteTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
