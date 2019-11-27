// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
)

const opDescribeProvisioningParameters = "DescribeProvisioningParameters"

// DescribeProvisioningParametersRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Gets information about the configuration required to provision the specified
// product using the specified provisioning artifact.
//
// If the output contains a TagOption key with an empty list of values, there
// is a TagOption conflict for that key. The end user cannot take action to
// fix the conflict, and launch is not blocked. In subsequent calls to ProvisionProduct,
// do not include conflicted TagOption keys as tags, or this causes the error
// "Parameter validation failed: Missing required parameter in Tags[N]:Value".
// Tag the provisioned product with the value sc-tagoption-conflict-portfolioId-productId.
//
//    // Example sending a request using DescribeProvisioningParametersRequest.
//    req := client.DescribeProvisioningParametersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribeProvisioningParameters
func (c *Client) DescribeProvisioningParametersRequest(input *types.DescribeProvisioningParametersInput) DescribeProvisioningParametersRequest {
	op := &aws.Operation{
		Name:       opDescribeProvisioningParameters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeProvisioningParametersInput{}
	}

	req := c.newRequest(op, input, &types.DescribeProvisioningParametersOutput{})
	return DescribeProvisioningParametersRequest{Request: req, Input: input, Copy: c.DescribeProvisioningParametersRequest}
}

// DescribeProvisioningParametersRequest is the request type for the
// DescribeProvisioningParameters API operation.
type DescribeProvisioningParametersRequest struct {
	*aws.Request
	Input *types.DescribeProvisioningParametersInput
	Copy  func(*types.DescribeProvisioningParametersInput) DescribeProvisioningParametersRequest
}

// Send marshals and sends the DescribeProvisioningParameters API request.
func (r DescribeProvisioningParametersRequest) Send(ctx context.Context) (*DescribeProvisioningParametersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeProvisioningParametersResponse{
		DescribeProvisioningParametersOutput: r.Request.Data.(*types.DescribeProvisioningParametersOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeProvisioningParametersResponse is the response type for the
// DescribeProvisioningParameters API operation.
type DescribeProvisioningParametersResponse struct {
	*types.DescribeProvisioningParametersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeProvisioningParameters request.
func (r *DescribeProvisioningParametersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
