// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package qldb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
)

const opCreateLedger = "CreateLedger"

// CreateLedgerRequest returns a request value for making API operation for
// Amazon QLDB.
//
// Creates a new ledger in your AWS account.
//
//    // Example sending a request using CreateLedgerRequest.
//    req := client.CreateLedgerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/qldb-2019-01-02/CreateLedger
func (c *Client) CreateLedgerRequest(input *types.CreateLedgerInput) CreateLedgerRequest {
	op := &aws.Operation{
		Name:       opCreateLedger,
		HTTPMethod: "POST",
		HTTPPath:   "/ledgers",
	}

	if input == nil {
		input = &types.CreateLedgerInput{}
	}

	req := c.newRequest(op, input, &types.CreateLedgerOutput{})
	return CreateLedgerRequest{Request: req, Input: input, Copy: c.CreateLedgerRequest}
}

// CreateLedgerRequest is the request type for the
// CreateLedger API operation.
type CreateLedgerRequest struct {
	*aws.Request
	Input *types.CreateLedgerInput
	Copy  func(*types.CreateLedgerInput) CreateLedgerRequest
}

// Send marshals and sends the CreateLedger API request.
func (r CreateLedgerRequest) Send(ctx context.Context) (*CreateLedgerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLedgerResponse{
		CreateLedgerOutput: r.Request.Data.(*types.CreateLedgerOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLedgerResponse is the response type for the
// CreateLedger API operation.
type CreateLedgerResponse struct {
	*types.CreateLedgerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLedger request.
func (r *CreateLedgerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
