// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
)

const opCreateConnection = "CreateConnection"

// CreateConnectionRequest returns a request value for making API operation for
// AWS Glue.
//
// Creates a connection definition in the Data Catalog.
//
//    // Example sending a request using CreateConnectionRequest.
//    req := client.CreateConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/CreateConnection
func (c *Client) CreateConnectionRequest(input *types.CreateConnectionInput) CreateConnectionRequest {
	op := &aws.Operation{
		Name:       opCreateConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateConnectionInput{}
	}

	req := c.newRequest(op, input, &types.CreateConnectionOutput{})
	return CreateConnectionRequest{Request: req, Input: input, Copy: c.CreateConnectionRequest}
}

// CreateConnectionRequest is the request type for the
// CreateConnection API operation.
type CreateConnectionRequest struct {
	*aws.Request
	Input *types.CreateConnectionInput
	Copy  func(*types.CreateConnectionInput) CreateConnectionRequest
}

// Send marshals and sends the CreateConnection API request.
func (r CreateConnectionRequest) Send(ctx context.Context) (*CreateConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateConnectionResponse{
		CreateConnectionOutput: r.Request.Data.(*types.CreateConnectionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateConnectionResponse is the response type for the
// CreateConnection API operation.
type CreateConnectionResponse struct {
	*types.CreateConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateConnection request.
func (r *CreateConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
