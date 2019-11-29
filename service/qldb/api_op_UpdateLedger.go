// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package qldb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/qldb/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
)

const opUpdateLedger = "UpdateLedger"

// UpdateLedgerRequest returns a request value for making API operation for
// Amazon QLDB.
//
// Updates properties on a ledger.
//
//    // Example sending a request using UpdateLedgerRequest.
//    req := client.UpdateLedgerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/qldb-2019-01-02/UpdateLedger
func (c *Client) UpdateLedgerRequest(input *types.UpdateLedgerInput) UpdateLedgerRequest {
	op := &aws.Operation{
		Name:       opUpdateLedger,
		HTTPMethod: "PATCH",
		HTTPPath:   "/ledgers/{name}",
	}

	if input == nil {
		input = &types.UpdateLedgerInput{}
	}

	req := c.newRequest(op, input, &types.UpdateLedgerOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateLedgerMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateLedgerRequest{Request: req, Input: input, Copy: c.UpdateLedgerRequest}
}

// UpdateLedgerRequest is the request type for the
// UpdateLedger API operation.
type UpdateLedgerRequest struct {
	*aws.Request
	Input *types.UpdateLedgerInput
	Copy  func(*types.UpdateLedgerInput) UpdateLedgerRequest
}

// Send marshals and sends the UpdateLedger API request.
func (r UpdateLedgerRequest) Send(ctx context.Context) (*UpdateLedgerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateLedgerResponse{
		UpdateLedgerOutput: r.Request.Data.(*types.UpdateLedgerOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateLedgerResponse is the response type for the
// UpdateLedger API operation.
type UpdateLedgerResponse struct {
	*types.UpdateLedgerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateLedger request.
func (r *UpdateLedgerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
