// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/glue/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
)

const opGetCrawler = "GetCrawler"

// GetCrawlerRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves metadata for a specified crawler.
//
//    // Example sending a request using GetCrawlerRequest.
//    req := client.GetCrawlerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetCrawler
func (c *Client) GetCrawlerRequest(input *types.GetCrawlerInput) GetCrawlerRequest {
	op := &aws.Operation{
		Name:       opGetCrawler,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetCrawlerInput{}
	}

	req := c.newRequest(op, input, &types.GetCrawlerOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetCrawlerMarshaler{Input: input}.GetNamedBuildHandler())

	return GetCrawlerRequest{Request: req, Input: input, Copy: c.GetCrawlerRequest}
}

// GetCrawlerRequest is the request type for the
// GetCrawler API operation.
type GetCrawlerRequest struct {
	*aws.Request
	Input *types.GetCrawlerInput
	Copy  func(*types.GetCrawlerInput) GetCrawlerRequest
}

// Send marshals and sends the GetCrawler API request.
func (r GetCrawlerRequest) Send(ctx context.Context) (*GetCrawlerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCrawlerResponse{
		GetCrawlerOutput: r.Request.Data.(*types.GetCrawlerOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCrawlerResponse is the response type for the
// GetCrawler API operation.
type GetCrawlerResponse struct {
	*types.GetCrawlerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCrawler request.
func (r *GetCrawlerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
