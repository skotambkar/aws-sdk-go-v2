// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
)

const opGetUtterancesView = "GetUtterancesView"

// GetUtterancesViewRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Use the GetUtterancesView operation to get information about the utterances
// that your users have made to your bot. You can use this list to tune the
// utterances that your bot responds to.
//
// For example, say that you have created a bot to order flowers. After your
// users have used your bot for a while, use the GetUtterancesView operation
// to see the requests that they have made and whether they have been successful.
// You might find that the utterance "I want flowers" is not being recognized.
// You could add this utterance to the OrderFlowers intent so that your bot
// recognizes that utterance.
//
// After you publish a new version of a bot, you can get information about the
// old version and the new so that you can compare the performance across the
// two versions.
//
// Utterance statistics are generated once a day. Data is available for the
// last 15 days. You can request information for up to 5 versions in each request.
// The response contains information about a maximum of 100 utterances for each
// version.
//
// This operation requires permissions for the lex:GetUtterancesView action.
//
//    // Example sending a request using GetUtterancesViewRequest.
//    req := client.GetUtterancesViewRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetUtterancesView
func (c *Client) GetUtterancesViewRequest(input *types.GetUtterancesViewInput) GetUtterancesViewRequest {
	op := &aws.Operation{
		Name:       opGetUtterancesView,
		HTTPMethod: "GET",
		HTTPPath:   "/bots/{botname}/utterances?view=aggregation",
	}

	if input == nil {
		input = &types.GetUtterancesViewInput{}
	}

	req := c.newRequest(op, input, &types.GetUtterancesViewOutput{})
	return GetUtterancesViewRequest{Request: req, Input: input, Copy: c.GetUtterancesViewRequest}
}

// GetUtterancesViewRequest is the request type for the
// GetUtterancesView API operation.
type GetUtterancesViewRequest struct {
	*aws.Request
	Input *types.GetUtterancesViewInput
	Copy  func(*types.GetUtterancesViewInput) GetUtterancesViewRequest
}

// Send marshals and sends the GetUtterancesView API request.
func (r GetUtterancesViewRequest) Send(ctx context.Context) (*GetUtterancesViewResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetUtterancesViewResponse{
		GetUtterancesViewOutput: r.Request.Data.(*types.GetUtterancesViewOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetUtterancesViewResponse is the response type for the
// GetUtterancesView API operation.
type GetUtterancesViewResponse struct {
	*types.GetUtterancesViewOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetUtterancesView request.
func (r *GetUtterancesViewResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
