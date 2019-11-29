// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opListCoreDefinitions = "ListCoreDefinitions"

// ListCoreDefinitionsRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves a list of core definitions.
//
//    // Example sending a request using ListCoreDefinitionsRequest.
//    req := client.ListCoreDefinitionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListCoreDefinitions
func (c *Client) ListCoreDefinitionsRequest(input *types.ListCoreDefinitionsInput) ListCoreDefinitionsRequest {
	op := &aws.Operation{
		Name:       opListCoreDefinitions,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/cores",
	}

	if input == nil {
		input = &types.ListCoreDefinitionsInput{}
	}

	req := c.newRequest(op, input, &types.ListCoreDefinitionsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListCoreDefinitionsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListCoreDefinitionsRequest{Request: req, Input: input, Copy: c.ListCoreDefinitionsRequest}
}

// ListCoreDefinitionsRequest is the request type for the
// ListCoreDefinitions API operation.
type ListCoreDefinitionsRequest struct {
	*aws.Request
	Input *types.ListCoreDefinitionsInput
	Copy  func(*types.ListCoreDefinitionsInput) ListCoreDefinitionsRequest
}

// Send marshals and sends the ListCoreDefinitions API request.
func (r ListCoreDefinitionsRequest) Send(ctx context.Context) (*ListCoreDefinitionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCoreDefinitionsResponse{
		ListCoreDefinitionsOutput: r.Request.Data.(*types.ListCoreDefinitionsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListCoreDefinitionsResponse is the response type for the
// ListCoreDefinitions API operation.
type ListCoreDefinitionsResponse struct {
	*types.ListCoreDefinitionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCoreDefinitions request.
func (r *ListCoreDefinitionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
