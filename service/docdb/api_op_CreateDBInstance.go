// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
)

const opCreateDBInstance = "CreateDBInstance"

// CreateDBInstanceRequest returns a request value for making API operation for
// Amazon DocumentDB with MongoDB compatibility.
//
// Creates a new DB instance.
//
//    // Example sending a request using CreateDBInstanceRequest.
//    req := client.CreateDBInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/CreateDBInstance
func (c *Client) CreateDBInstanceRequest(input *types.CreateDBInstanceInput) CreateDBInstanceRequest {
	op := &aws.Operation{
		Name:       opCreateDBInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateDBInstanceInput{}
	}

	req := c.newRequest(op, input, &types.CreateDBInstanceOutput{})
	return CreateDBInstanceRequest{Request: req, Input: input, Copy: c.CreateDBInstanceRequest}
}

// CreateDBInstanceRequest is the request type for the
// CreateDBInstance API operation.
type CreateDBInstanceRequest struct {
	*aws.Request
	Input *types.CreateDBInstanceInput
	Copy  func(*types.CreateDBInstanceInput) CreateDBInstanceRequest
}

// Send marshals and sends the CreateDBInstance API request.
func (r CreateDBInstanceRequest) Send(ctx context.Context) (*CreateDBInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDBInstanceResponse{
		CreateDBInstanceOutput: r.Request.Data.(*types.CreateDBInstanceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDBInstanceResponse is the response type for the
// CreateDBInstance API operation.
type CreateDBInstanceResponse struct {
	*types.CreateDBInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDBInstance request.
func (r *CreateDBInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
