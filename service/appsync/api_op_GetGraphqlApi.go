// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
)

const opGetGraphqlApi = "GetGraphqlApi"

// GetGraphqlApiRequest returns a request value for making API operation for
// AWS AppSync.
//
// Retrieves a GraphqlApi object.
//
//    // Example sending a request using GetGraphqlApiRequest.
//    req := client.GetGraphqlApiRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/GetGraphqlApi
func (c *Client) GetGraphqlApiRequest(input *types.GetGraphqlApiInput) GetGraphqlApiRequest {
	op := &aws.Operation{
		Name:       opGetGraphqlApi,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apis/{apiId}",
	}

	if input == nil {
		input = &types.GetGraphqlApiInput{}
	}

	req := c.newRequest(op, input, &types.GetGraphqlApiOutput{})
	return GetGraphqlApiRequest{Request: req, Input: input, Copy: c.GetGraphqlApiRequest}
}

// GetGraphqlApiRequest is the request type for the
// GetGraphqlApi API operation.
type GetGraphqlApiRequest struct {
	*aws.Request
	Input *types.GetGraphqlApiInput
	Copy  func(*types.GetGraphqlApiInput) GetGraphqlApiRequest
}

// Send marshals and sends the GetGraphqlApi API request.
func (r GetGraphqlApiRequest) Send(ctx context.Context) (*GetGraphqlApiResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetGraphqlApiResponse{
		GetGraphqlApiOutput: r.Request.Data.(*types.GetGraphqlApiOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetGraphqlApiResponse is the response type for the
// GetGraphqlApi API operation.
type GetGraphqlApiResponse struct {
	*types.GetGraphqlApiOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetGraphqlApi request.
func (r *GetGraphqlApiResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
