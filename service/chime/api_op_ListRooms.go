// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opListRooms = "ListRooms"

// ListRoomsRequest returns a request value for making API operation for
// Amazon Chime.
//
// Lists the room details for the specified Amazon Chime account. Optionally,
// filter the results by a member ID (user ID or bot ID) to see a list of rooms
// that the member belongs to.
//
//    // Example sending a request using ListRoomsRequest.
//    req := client.ListRoomsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/ListRooms
func (c *Client) ListRoomsRequest(input *types.ListRoomsInput) ListRoomsRequest {
	op := &aws.Operation{
		Name:       opListRooms,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{accountId}/rooms",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListRoomsInput{}
	}

	req := c.newRequest(op, input, &types.ListRoomsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListRoomsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListRoomsRequest{Request: req, Input: input, Copy: c.ListRoomsRequest}
}

// ListRoomsRequest is the request type for the
// ListRooms API operation.
type ListRoomsRequest struct {
	*aws.Request
	Input *types.ListRoomsInput
	Copy  func(*types.ListRoomsInput) ListRoomsRequest
}

// Send marshals and sends the ListRooms API request.
func (r ListRoomsRequest) Send(ctx context.Context) (*ListRoomsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRoomsResponse{
		ListRoomsOutput: r.Request.Data.(*types.ListRoomsOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListRoomsRequestPaginator returns a paginator for ListRooms.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListRoomsRequest(input)
//   p := chime.NewListRoomsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListRoomsPaginator(req ListRoomsRequest) ListRoomsPaginator {
	return ListRoomsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListRoomsInput
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

// ListRoomsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListRoomsPaginator struct {
	aws.Pager
}

func (p *ListRoomsPaginator) CurrentPage() *types.ListRoomsOutput {
	return p.Pager.CurrentPage().(*types.ListRoomsOutput)
}

// ListRoomsResponse is the response type for the
// ListRooms API operation.
type ListRoomsResponse struct {
	*types.ListRoomsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRooms request.
func (r *ListRoomsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
