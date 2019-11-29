// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicequotas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
)

const opRequestServiceQuotaIncrease = "RequestServiceQuotaIncrease"

// RequestServiceQuotaIncreaseRequest returns a request value for making API operation for
// Service Quotas.
//
// Retrieves the details of a service quota increase request. The response to
// this command provides the details in the RequestedServiceQuotaChange object.
//
//    // Example sending a request using RequestServiceQuotaIncreaseRequest.
//    req := client.RequestServiceQuotaIncreaseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/RequestServiceQuotaIncrease
func (c *Client) RequestServiceQuotaIncreaseRequest(input *types.RequestServiceQuotaIncreaseInput) RequestServiceQuotaIncreaseRequest {
	op := &aws.Operation{
		Name:       opRequestServiceQuotaIncrease,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RequestServiceQuotaIncreaseInput{}
	}

	req := c.newRequest(op, input, &types.RequestServiceQuotaIncreaseOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.RequestServiceQuotaIncreaseMarshaler{Input: input}.GetNamedBuildHandler())

	return RequestServiceQuotaIncreaseRequest{Request: req, Input: input, Copy: c.RequestServiceQuotaIncreaseRequest}
}

// RequestServiceQuotaIncreaseRequest is the request type for the
// RequestServiceQuotaIncrease API operation.
type RequestServiceQuotaIncreaseRequest struct {
	*aws.Request
	Input *types.RequestServiceQuotaIncreaseInput
	Copy  func(*types.RequestServiceQuotaIncreaseInput) RequestServiceQuotaIncreaseRequest
}

// Send marshals and sends the RequestServiceQuotaIncrease API request.
func (r RequestServiceQuotaIncreaseRequest) Send(ctx context.Context) (*RequestServiceQuotaIncreaseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RequestServiceQuotaIncreaseResponse{
		RequestServiceQuotaIncreaseOutput: r.Request.Data.(*types.RequestServiceQuotaIncreaseOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RequestServiceQuotaIncreaseResponse is the response type for the
// RequestServiceQuotaIncrease API operation.
type RequestServiceQuotaIncreaseResponse struct {
	*types.RequestServiceQuotaIncreaseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RequestServiceQuotaIncrease request.
func (r *RequestServiceQuotaIncreaseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
