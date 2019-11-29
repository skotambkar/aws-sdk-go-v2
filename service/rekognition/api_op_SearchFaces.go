// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

const opSearchFaces = "SearchFaces"

// SearchFacesRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// For a given input face ID, searches for matching faces in the collection
// the face belongs to. You get a face ID when you add a face to the collection
// using the IndexFaces operation. The operation compares the features of the
// input face with faces in the specified collection.
//
// You can also search faces without indexing faces by using the SearchFacesByImage
// operation.
//
// The operation response returns an array of faces that match, ordered by similarity
// score with the highest similarity first. More specifically, it is an array
// of metadata for each face match that is found. Along with the metadata, the
// response also includes a confidence value for each face match, indicating
// the confidence that the specific face matches the input face.
//
// For an example, see Searching for a Face Using Its Face ID in the Amazon
// Rekognition Developer Guide.
//
// This operation requires permissions to perform the rekognition:SearchFaces
// action.
//
//    // Example sending a request using SearchFacesRequest.
//    req := client.SearchFacesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) SearchFacesRequest(input *types.SearchFacesInput) SearchFacesRequest {
	op := &aws.Operation{
		Name:       opSearchFaces,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.SearchFacesInput{}
	}

	req := c.newRequest(op, input, &types.SearchFacesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.SearchFacesMarshaler{Input: input}.GetNamedBuildHandler())

	return SearchFacesRequest{Request: req, Input: input, Copy: c.SearchFacesRequest}
}

// SearchFacesRequest is the request type for the
// SearchFaces API operation.
type SearchFacesRequest struct {
	*aws.Request
	Input *types.SearchFacesInput
	Copy  func(*types.SearchFacesInput) SearchFacesRequest
}

// Send marshals and sends the SearchFaces API request.
func (r SearchFacesRequest) Send(ctx context.Context) (*SearchFacesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SearchFacesResponse{
		SearchFacesOutput: r.Request.Data.(*types.SearchFacesOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SearchFacesResponse is the response type for the
// SearchFaces API operation.
type SearchFacesResponse struct {
	*types.SearchFacesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SearchFaces request.
func (r *SearchFacesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
