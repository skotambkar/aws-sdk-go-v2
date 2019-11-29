// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
)

const opForgetSmartHomeAppliances = "ForgetSmartHomeAppliances"

// ForgetSmartHomeAppliancesRequest returns a request value for making API operation for
// Alexa For Business.
//
// Forgets smart home appliances associated to a room.
//
//    // Example sending a request using ForgetSmartHomeAppliancesRequest.
//    req := client.ForgetSmartHomeAppliancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/ForgetSmartHomeAppliances
func (c *Client) ForgetSmartHomeAppliancesRequest(input *types.ForgetSmartHomeAppliancesInput) ForgetSmartHomeAppliancesRequest {
	op := &aws.Operation{
		Name:       opForgetSmartHomeAppliances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ForgetSmartHomeAppliancesInput{}
	}

	req := c.newRequest(op, input, &types.ForgetSmartHomeAppliancesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ForgetSmartHomeAppliancesMarshaler{Input: input}.GetNamedBuildHandler())

	return ForgetSmartHomeAppliancesRequest{Request: req, Input: input, Copy: c.ForgetSmartHomeAppliancesRequest}
}

// ForgetSmartHomeAppliancesRequest is the request type for the
// ForgetSmartHomeAppliances API operation.
type ForgetSmartHomeAppliancesRequest struct {
	*aws.Request
	Input *types.ForgetSmartHomeAppliancesInput
	Copy  func(*types.ForgetSmartHomeAppliancesInput) ForgetSmartHomeAppliancesRequest
}

// Send marshals and sends the ForgetSmartHomeAppliances API request.
func (r ForgetSmartHomeAppliancesRequest) Send(ctx context.Context) (*ForgetSmartHomeAppliancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ForgetSmartHomeAppliancesResponse{
		ForgetSmartHomeAppliancesOutput: r.Request.Data.(*types.ForgetSmartHomeAppliancesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ForgetSmartHomeAppliancesResponse is the response type for the
// ForgetSmartHomeAppliances API operation.
type ForgetSmartHomeAppliancesResponse struct {
	*types.ForgetSmartHomeAppliancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ForgetSmartHomeAppliances request.
func (r *ForgetSmartHomeAppliancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
