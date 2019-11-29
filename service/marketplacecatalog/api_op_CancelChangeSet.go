// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplacecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog/types"
)

const opCancelChangeSet = "CancelChangeSet"

// CancelChangeSetRequest returns a request value for making API operation for
// AWS Marketplace Catalog Service.
//
// Used to cancel an open change request. Must be sent before the status of
// the request changes to APPLYING, the final stage of completing your change
// request. You can describe a change during the 60-day request history retention
// period for API calls.
//
//    // Example sending a request using CancelChangeSetRequest.
//    req := client.CancelChangeSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/marketplace-catalog-2018-09-17/CancelChangeSet
func (c *Client) CancelChangeSetRequest(input *types.CancelChangeSetInput) CancelChangeSetRequest {
	op := &aws.Operation{
		Name:       opCancelChangeSet,
		HTTPMethod: "PATCH",
		HTTPPath:   "/CancelChangeSet",
	}

	if input == nil {
		input = &types.CancelChangeSetInput{}
	}

	req := c.newRequest(op, input, &types.CancelChangeSetOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CancelChangeSetMarshaler{Input: input}.GetNamedBuildHandler())

	return CancelChangeSetRequest{Request: req, Input: input, Copy: c.CancelChangeSetRequest}
}

// CancelChangeSetRequest is the request type for the
// CancelChangeSet API operation.
type CancelChangeSetRequest struct {
	*aws.Request
	Input *types.CancelChangeSetInput
	Copy  func(*types.CancelChangeSetInput) CancelChangeSetRequest
}

// Send marshals and sends the CancelChangeSet API request.
func (r CancelChangeSetRequest) Send(ctx context.Context) (*CancelChangeSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelChangeSetResponse{
		CancelChangeSetOutput: r.Request.Data.(*types.CancelChangeSetOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelChangeSetResponse is the response type for the
// CancelChangeSet API operation.
type CancelChangeSetResponse struct {
	*types.CancelChangeSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelChangeSet request.
func (r *CancelChangeSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
