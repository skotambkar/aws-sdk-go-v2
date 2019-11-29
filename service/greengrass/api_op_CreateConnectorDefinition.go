// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opCreateConnectorDefinition = "CreateConnectorDefinition"

// CreateConnectorDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Creates a connector definition. You may provide the initial version of the
// connector definition now or use ''CreateConnectorDefinitionVersion'' at a
// later time.
//
//    // Example sending a request using CreateConnectorDefinitionRequest.
//    req := client.CreateConnectorDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateConnectorDefinition
func (c *Client) CreateConnectorDefinitionRequest(input *types.CreateConnectorDefinitionInput) CreateConnectorDefinitionRequest {
	op := &aws.Operation{
		Name:       opCreateConnectorDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/greengrass/definition/connectors",
	}

	if input == nil {
		input = &types.CreateConnectorDefinitionInput{}
	}

	req := c.newRequest(op, input, &types.CreateConnectorDefinitionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreateConnectorDefinitionMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateConnectorDefinitionRequest{Request: req, Input: input, Copy: c.CreateConnectorDefinitionRequest}
}

// CreateConnectorDefinitionRequest is the request type for the
// CreateConnectorDefinition API operation.
type CreateConnectorDefinitionRequest struct {
	*aws.Request
	Input *types.CreateConnectorDefinitionInput
	Copy  func(*types.CreateConnectorDefinitionInput) CreateConnectorDefinitionRequest
}

// Send marshals and sends the CreateConnectorDefinition API request.
func (r CreateConnectorDefinitionRequest) Send(ctx context.Context) (*CreateConnectorDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateConnectorDefinitionResponse{
		CreateConnectorDefinitionOutput: r.Request.Data.(*types.CreateConnectorDefinitionOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateConnectorDefinitionResponse is the response type for the
// CreateConnectorDefinition API operation.
type CreateConnectorDefinitionResponse struct {
	*types.CreateConnectorDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateConnectorDefinition request.
func (r *CreateConnectorDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
