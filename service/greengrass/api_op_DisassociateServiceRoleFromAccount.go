// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opDisassociateServiceRoleFromAccount = "DisassociateServiceRoleFromAccount"

// DisassociateServiceRoleFromAccountRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Disassociates the service role from your account. Without a service role,
// deployments will not work.
//
//    // Example sending a request using DisassociateServiceRoleFromAccountRequest.
//    req := client.DisassociateServiceRoleFromAccountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/DisassociateServiceRoleFromAccount
func (c *Client) DisassociateServiceRoleFromAccountRequest(input *types.DisassociateServiceRoleFromAccountInput) DisassociateServiceRoleFromAccountRequest {
	op := &aws.Operation{
		Name:       opDisassociateServiceRoleFromAccount,
		HTTPMethod: "DELETE",
		HTTPPath:   "/greengrass/servicerole",
	}

	if input == nil {
		input = &types.DisassociateServiceRoleFromAccountInput{}
	}

	req := c.newRequest(op, input, &types.DisassociateServiceRoleFromAccountOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DisassociateServiceRoleFromAccountMarshaler{Input: input}.GetNamedBuildHandler())

	return DisassociateServiceRoleFromAccountRequest{Request: req, Input: input, Copy: c.DisassociateServiceRoleFromAccountRequest}
}

// DisassociateServiceRoleFromAccountRequest is the request type for the
// DisassociateServiceRoleFromAccount API operation.
type DisassociateServiceRoleFromAccountRequest struct {
	*aws.Request
	Input *types.DisassociateServiceRoleFromAccountInput
	Copy  func(*types.DisassociateServiceRoleFromAccountInput) DisassociateServiceRoleFromAccountRequest
}

// Send marshals and sends the DisassociateServiceRoleFromAccount API request.
func (r DisassociateServiceRoleFromAccountRequest) Send(ctx context.Context) (*DisassociateServiceRoleFromAccountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateServiceRoleFromAccountResponse{
		DisassociateServiceRoleFromAccountOutput: r.Request.Data.(*types.DisassociateServiceRoleFromAccountOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateServiceRoleFromAccountResponse is the response type for the
// DisassociateServiceRoleFromAccount API operation.
type DisassociateServiceRoleFromAccountResponse struct {
	*types.DisassociateServiceRoleFromAccountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateServiceRoleFromAccount request.
func (r *DisassociateServiceRoleFromAccountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
