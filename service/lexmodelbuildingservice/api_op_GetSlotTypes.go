// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
)

const opGetSlotTypes = "GetSlotTypes"

// GetSlotTypesRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Returns slot type information as follows:
//
//    * If you specify the nameContains field, returns the $LATEST version of
//    all slot types that contain the specified string.
//
//    * If you don't specify the nameContains field, returns information about
//    the $LATEST version of all slot types.
//
// The operation requires permission for the lex:GetSlotTypes action.
//
//    // Example sending a request using GetSlotTypesRequest.
//    req := client.GetSlotTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetSlotTypes
func (c *Client) GetSlotTypesRequest(input *types.GetSlotTypesInput) GetSlotTypesRequest {
	op := &aws.Operation{
		Name:       opGetSlotTypes,
		HTTPMethod: "GET",
		HTTPPath:   "/slottypes/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.GetSlotTypesInput{}
	}

	req := c.newRequest(op, input, &types.GetSlotTypesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetSlotTypesMarshaler{Input: input}.GetNamedBuildHandler())

	return GetSlotTypesRequest{Request: req, Input: input, Copy: c.GetSlotTypesRequest}
}

// GetSlotTypesRequest is the request type for the
// GetSlotTypes API operation.
type GetSlotTypesRequest struct {
	*aws.Request
	Input *types.GetSlotTypesInput
	Copy  func(*types.GetSlotTypesInput) GetSlotTypesRequest
}

// Send marshals and sends the GetSlotTypes API request.
func (r GetSlotTypesRequest) Send(ctx context.Context) (*GetSlotTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSlotTypesResponse{
		GetSlotTypesOutput: r.Request.Data.(*types.GetSlotTypesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetSlotTypesRequestPaginator returns a paginator for GetSlotTypes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetSlotTypesRequest(input)
//   p := lexmodelbuildingservice.NewGetSlotTypesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetSlotTypesPaginator(req GetSlotTypesRequest) GetSlotTypesPaginator {
	return GetSlotTypesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.GetSlotTypesInput
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

// GetSlotTypesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetSlotTypesPaginator struct {
	aws.Pager
}

func (p *GetSlotTypesPaginator) CurrentPage() *types.GetSlotTypesOutput {
	return p.Pager.CurrentPage().(*types.GetSlotTypesOutput)
}

// GetSlotTypesResponse is the response type for the
// GetSlotTypes API operation.
type GetSlotTypesResponse struct {
	*types.GetSlotTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSlotTypes request.
func (r *GetSlotTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
