// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/types"
)

const opPutResourceAttributes = "PutResourceAttributes"

// PutResourceAttributesRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Provides identifying details of the resource being migrated so that it can
// be associated in the Application Discovery Service (ADS)'s repository. This
// association occurs asynchronously after PutResourceAttributes returns.
//
//    * Keep in mind that subsequent calls to PutResourceAttributes will override
//    previously stored attributes. For example, if it is first called with
//    a MAC address, but later, it is desired to add an IP address, it will
//    then be required to call it with both the IP and MAC addresses to prevent
//    overiding the MAC address.
//
//    * Note the instructions regarding the special use case of the ResourceAttributeList
//    (https://docs.aws.amazon.com/migrationhub/latest/ug/API_PutResourceAttributes.html#migrationhub-PutResourceAttributes-request-ResourceAttributeList)
//    parameter when specifying any "VM" related value.
//
// Because this is an asynchronous call, it will always return 200, whether
// an association occurs or not. To confirm if an association was found based
// on the provided details, call ListDiscoveredResources.
//
//    // Example sending a request using PutResourceAttributesRequest.
//    req := client.PutResourceAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/PutResourceAttributes
func (c *Client) PutResourceAttributesRequest(input *types.PutResourceAttributesInput) PutResourceAttributesRequest {
	op := &aws.Operation{
		Name:       opPutResourceAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutResourceAttributesInput{}
	}

	req := c.newRequest(op, input, &types.PutResourceAttributesOutput{})
	return PutResourceAttributesRequest{Request: req, Input: input, Copy: c.PutResourceAttributesRequest}
}

// PutResourceAttributesRequest is the request type for the
// PutResourceAttributes API operation.
type PutResourceAttributesRequest struct {
	*aws.Request
	Input *types.PutResourceAttributesInput
	Copy  func(*types.PutResourceAttributesInput) PutResourceAttributesRequest
}

// Send marshals and sends the PutResourceAttributes API request.
func (r PutResourceAttributesRequest) Send(ctx context.Context) (*PutResourceAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutResourceAttributesResponse{
		PutResourceAttributesOutput: r.Request.Data.(*types.PutResourceAttributesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutResourceAttributesResponse is the response type for the
// PutResourceAttributes API operation.
type PutResourceAttributesResponse struct {
	*types.PutResourceAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutResourceAttributes request.
func (r *PutResourceAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
