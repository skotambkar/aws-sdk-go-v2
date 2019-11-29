// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package qldb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/qldb/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
)

const opDescribeLedger = "DescribeLedger"

// DescribeLedgerRequest returns a request value for making API operation for
// Amazon QLDB.
//
// Returns information about a ledger, including its state and when it was created.
//
//    // Example sending a request using DescribeLedgerRequest.
//    req := client.DescribeLedgerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/qldb-2019-01-02/DescribeLedger
func (c *Client) DescribeLedgerRequest(input *types.DescribeLedgerInput) DescribeLedgerRequest {
	op := &aws.Operation{
		Name:       opDescribeLedger,
		HTTPMethod: "GET",
		HTTPPath:   "/ledgers/{name}",
	}

	if input == nil {
		input = &types.DescribeLedgerInput{}
	}

	req := c.newRequest(op, input, &types.DescribeLedgerOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DescribeLedgerMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeLedgerRequest{Request: req, Input: input, Copy: c.DescribeLedgerRequest}
}

// DescribeLedgerRequest is the request type for the
// DescribeLedger API operation.
type DescribeLedgerRequest struct {
	*aws.Request
	Input *types.DescribeLedgerInput
	Copy  func(*types.DescribeLedgerInput) DescribeLedgerRequest
}

// Send marshals and sends the DescribeLedger API request.
func (r DescribeLedgerRequest) Send(ctx context.Context) (*DescribeLedgerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLedgerResponse{
		DescribeLedgerOutput: r.Request.Data.(*types.DescribeLedgerOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLedgerResponse is the response type for the
// DescribeLedger API operation.
type DescribeLedgerResponse struct {
	*types.DescribeLedgerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLedger request.
func (r *DescribeLedgerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
