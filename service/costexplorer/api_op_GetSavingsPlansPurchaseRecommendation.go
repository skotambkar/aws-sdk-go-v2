// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

const opGetSavingsPlansPurchaseRecommendation = "GetSavingsPlansPurchaseRecommendation"

// GetSavingsPlansPurchaseRecommendationRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// Retrieves your request parameters, Savings Plan Recommendations Summary and
// Details.
//
//    // Example sending a request using GetSavingsPlansPurchaseRecommendationRequest.
//    req := client.GetSavingsPlansPurchaseRecommendationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetSavingsPlansPurchaseRecommendation
func (c *Client) GetSavingsPlansPurchaseRecommendationRequest(input *types.GetSavingsPlansPurchaseRecommendationInput) GetSavingsPlansPurchaseRecommendationRequest {
	op := &aws.Operation{
		Name:       opGetSavingsPlansPurchaseRecommendation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetSavingsPlansPurchaseRecommendationInput{}
	}

	req := c.newRequest(op, input, &types.GetSavingsPlansPurchaseRecommendationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetSavingsPlansPurchaseRecommendationMarshaler{Input: input}.GetNamedBuildHandler())

	return GetSavingsPlansPurchaseRecommendationRequest{Request: req, Input: input, Copy: c.GetSavingsPlansPurchaseRecommendationRequest}
}

// GetSavingsPlansPurchaseRecommendationRequest is the request type for the
// GetSavingsPlansPurchaseRecommendation API operation.
type GetSavingsPlansPurchaseRecommendationRequest struct {
	*aws.Request
	Input *types.GetSavingsPlansPurchaseRecommendationInput
	Copy  func(*types.GetSavingsPlansPurchaseRecommendationInput) GetSavingsPlansPurchaseRecommendationRequest
}

// Send marshals and sends the GetSavingsPlansPurchaseRecommendation API request.
func (r GetSavingsPlansPurchaseRecommendationRequest) Send(ctx context.Context) (*GetSavingsPlansPurchaseRecommendationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSavingsPlansPurchaseRecommendationResponse{
		GetSavingsPlansPurchaseRecommendationOutput: r.Request.Data.(*types.GetSavingsPlansPurchaseRecommendationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSavingsPlansPurchaseRecommendationResponse is the response type for the
// GetSavingsPlansPurchaseRecommendation API operation.
type GetSavingsPlansPurchaseRecommendationResponse struct {
	*types.GetSavingsPlansPurchaseRecommendationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSavingsPlansPurchaseRecommendation request.
func (r *GetSavingsPlansPurchaseRecommendationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
