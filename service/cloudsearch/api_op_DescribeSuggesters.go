// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
)

const opDescribeSuggesters = "DescribeSuggesters"

// DescribeSuggestersRequest returns a request value for making API operation for
// Amazon CloudSearch.
//
// Gets the suggesters configured for a domain. A suggester enables you to display
// possible matches before users finish typing their queries. Can be limited
// to specific suggesters by name. By default, shows all suggesters and includes
// any pending changes to the configuration. Set the Deployed option to true
// to show the active configuration and exclude pending changes. For more information,
// see Getting Search Suggestions (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-suggestions.html)
// in the Amazon CloudSearch Developer Guide.
//
//    // Example sending a request using DescribeSuggestersRequest.
//    req := client.DescribeSuggestersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeSuggestersRequest(input *types.DescribeSuggestersInput) DescribeSuggestersRequest {
	op := &aws.Operation{
		Name:       opDescribeSuggesters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeSuggestersInput{}
	}

	req := c.newRequest(op, input, &types.DescribeSuggestersOutput{})
	return DescribeSuggestersRequest{Request: req, Input: input, Copy: c.DescribeSuggestersRequest}
}

// DescribeSuggestersRequest is the request type for the
// DescribeSuggesters API operation.
type DescribeSuggestersRequest struct {
	*aws.Request
	Input *types.DescribeSuggestersInput
	Copy  func(*types.DescribeSuggestersInput) DescribeSuggestersRequest
}

// Send marshals and sends the DescribeSuggesters API request.
func (r DescribeSuggestersRequest) Send(ctx context.Context) (*DescribeSuggestersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSuggestersResponse{
		DescribeSuggestersOutput: r.Request.Data.(*types.DescribeSuggestersOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSuggestersResponse is the response type for the
// DescribeSuggesters API operation.
type DescribeSuggestersResponse struct {
	*types.DescribeSuggestersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSuggesters request.
func (r *DescribeSuggestersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
