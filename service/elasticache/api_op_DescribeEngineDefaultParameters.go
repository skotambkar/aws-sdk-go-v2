// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
)

const opDescribeEngineDefaultParameters = "DescribeEngineDefaultParameters"

// DescribeEngineDefaultParametersRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Returns the default engine and system parameter information for the specified
// cache engine.
//
//    // Example sending a request using DescribeEngineDefaultParametersRequest.
//    req := client.DescribeEngineDefaultParametersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DescribeEngineDefaultParameters
func (c *Client) DescribeEngineDefaultParametersRequest(input *types.DescribeEngineDefaultParametersInput) DescribeEngineDefaultParametersRequest {
	op := &aws.Operation{
		Name:       opDescribeEngineDefaultParameters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"EngineDefaults.Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.DescribeEngineDefaultParametersInput{}
	}

	req := c.newRequest(op, input, &types.DescribeEngineDefaultParametersOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DescribeEngineDefaultParametersMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeEngineDefaultParametersRequest{Request: req, Input: input, Copy: c.DescribeEngineDefaultParametersRequest}
}

// DescribeEngineDefaultParametersRequest is the request type for the
// DescribeEngineDefaultParameters API operation.
type DescribeEngineDefaultParametersRequest struct {
	*aws.Request
	Input *types.DescribeEngineDefaultParametersInput
	Copy  func(*types.DescribeEngineDefaultParametersInput) DescribeEngineDefaultParametersRequest
}

// Send marshals and sends the DescribeEngineDefaultParameters API request.
func (r DescribeEngineDefaultParametersRequest) Send(ctx context.Context) (*DescribeEngineDefaultParametersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEngineDefaultParametersResponse{
		DescribeEngineDefaultParametersOutput: r.Request.Data.(*types.DescribeEngineDefaultParametersOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeEngineDefaultParametersRequestPaginator returns a paginator for DescribeEngineDefaultParameters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeEngineDefaultParametersRequest(input)
//   p := elasticache.NewDescribeEngineDefaultParametersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeEngineDefaultParametersPaginator(req DescribeEngineDefaultParametersRequest) DescribeEngineDefaultParametersPaginator {
	return DescribeEngineDefaultParametersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeEngineDefaultParametersInput
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

// DescribeEngineDefaultParametersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeEngineDefaultParametersPaginator struct {
	aws.Pager
}

func (p *DescribeEngineDefaultParametersPaginator) CurrentPage() *types.DescribeEngineDefaultParametersOutput {
	return p.Pager.CurrentPage().(*types.DescribeEngineDefaultParametersOutput)
}

// DescribeEngineDefaultParametersResponse is the response type for the
// DescribeEngineDefaultParameters API operation.
type DescribeEngineDefaultParametersResponse struct {
	*types.DescribeEngineDefaultParametersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEngineDefaultParameters request.
func (r *DescribeEngineDefaultParametersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
