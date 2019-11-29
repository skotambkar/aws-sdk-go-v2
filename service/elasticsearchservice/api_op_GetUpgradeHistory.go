// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
)

const opGetUpgradeHistory = "GetUpgradeHistory"

// GetUpgradeHistoryRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// Retrieves the complete history of the last 10 upgrades that were performed
// on the domain.
//
//    // Example sending a request using GetUpgradeHistoryRequest.
//    req := client.GetUpgradeHistoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetUpgradeHistoryRequest(input *types.GetUpgradeHistoryInput) GetUpgradeHistoryRequest {
	op := &aws.Operation{
		Name:       opGetUpgradeHistory,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-01-01/es/upgradeDomain/{DomainName}/history",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.GetUpgradeHistoryInput{}
	}

	req := c.newRequest(op, input, &types.GetUpgradeHistoryOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetUpgradeHistoryMarshaler{Input: input}.GetNamedBuildHandler())

	return GetUpgradeHistoryRequest{Request: req, Input: input, Copy: c.GetUpgradeHistoryRequest}
}

// GetUpgradeHistoryRequest is the request type for the
// GetUpgradeHistory API operation.
type GetUpgradeHistoryRequest struct {
	*aws.Request
	Input *types.GetUpgradeHistoryInput
	Copy  func(*types.GetUpgradeHistoryInput) GetUpgradeHistoryRequest
}

// Send marshals and sends the GetUpgradeHistory API request.
func (r GetUpgradeHistoryRequest) Send(ctx context.Context) (*GetUpgradeHistoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetUpgradeHistoryResponse{
		GetUpgradeHistoryOutput: r.Request.Data.(*types.GetUpgradeHistoryOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetUpgradeHistoryRequestPaginator returns a paginator for GetUpgradeHistory.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetUpgradeHistoryRequest(input)
//   p := elasticsearchservice.NewGetUpgradeHistoryRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetUpgradeHistoryPaginator(req GetUpgradeHistoryRequest) GetUpgradeHistoryPaginator {
	return GetUpgradeHistoryPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.GetUpgradeHistoryInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetUpgradeHistoryPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetUpgradeHistoryPaginator struct {
	aws.Pager
}

func (p *GetUpgradeHistoryPaginator) CurrentPage() *types.GetUpgradeHistoryOutput {
	return p.Pager.CurrentPage().(*types.GetUpgradeHistoryOutput)
}

// GetUpgradeHistoryResponse is the response type for the
// GetUpgradeHistory API operation.
type GetUpgradeHistoryResponse struct {
	*types.GetUpgradeHistoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetUpgradeHistory request.
func (r *GetUpgradeHistoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
