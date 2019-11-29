// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sesv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

const opPutConfigurationSetDeliveryOptions = "PutConfigurationSetDeliveryOptions"

// PutConfigurationSetDeliveryOptionsRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Associate a configuration set with a dedicated IP pool. You can use dedicated
// IP pools to create groups of dedicated IP addresses for sending specific
// types of email.
//
//    // Example sending a request using PutConfigurationSetDeliveryOptionsRequest.
//    req := client.PutConfigurationSetDeliveryOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sesv2-2019-09-27/PutConfigurationSetDeliveryOptions
func (c *Client) PutConfigurationSetDeliveryOptionsRequest(input *types.PutConfigurationSetDeliveryOptionsInput) PutConfigurationSetDeliveryOptionsRequest {
	op := &aws.Operation{
		Name:       opPutConfigurationSetDeliveryOptions,
		HTTPMethod: "PUT",
		HTTPPath:   "/v2/email/configuration-sets/{ConfigurationSetName}/delivery-options",
	}

	if input == nil {
		input = &types.PutConfigurationSetDeliveryOptionsInput{}
	}

	req := c.newRequest(op, input, &types.PutConfigurationSetDeliveryOptionsOutput{})
	return PutConfigurationSetDeliveryOptionsRequest{Request: req, Input: input, Copy: c.PutConfigurationSetDeliveryOptionsRequest}
}

// PutConfigurationSetDeliveryOptionsRequest is the request type for the
// PutConfigurationSetDeliveryOptions API operation.
type PutConfigurationSetDeliveryOptionsRequest struct {
	*aws.Request
	Input *types.PutConfigurationSetDeliveryOptionsInput
	Copy  func(*types.PutConfigurationSetDeliveryOptionsInput) PutConfigurationSetDeliveryOptionsRequest
}

// Send marshals and sends the PutConfigurationSetDeliveryOptions API request.
func (r PutConfigurationSetDeliveryOptionsRequest) Send(ctx context.Context) (*PutConfigurationSetDeliveryOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutConfigurationSetDeliveryOptionsResponse{
		PutConfigurationSetDeliveryOptionsOutput: r.Request.Data.(*types.PutConfigurationSetDeliveryOptionsOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutConfigurationSetDeliveryOptionsResponse is the response type for the
// PutConfigurationSetDeliveryOptions API operation.
type PutConfigurationSetDeliveryOptionsResponse struct {
	*types.PutConfigurationSetDeliveryOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutConfigurationSetDeliveryOptions request.
func (r *PutConfigurationSetDeliveryOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
