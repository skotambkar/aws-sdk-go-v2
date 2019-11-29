// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
)

const opCreateSlotTypeVersion = "CreateSlotTypeVersion"

// CreateSlotTypeVersionRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Creates a new version of a slot type based on the $LATEST version of the
// specified slot type. If the $LATEST version of this resource has not changed
// since the last version that you created, Amazon Lex doesn't create a new
// version. It returns the last version that you created.
//
// You can update only the $LATEST version of a slot type. You can't update
// the numbered versions that you create with the CreateSlotTypeVersion operation.
//
// When you create a version of a slot type, Amazon Lex sets the version to
// 1. Subsequent versions increment by 1. For more information, see versioning-intro.
//
// This operation requires permissions for the lex:CreateSlotTypeVersion action.
//
//    // Example sending a request using CreateSlotTypeVersionRequest.
//    req := client.CreateSlotTypeVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/CreateSlotTypeVersion
func (c *Client) CreateSlotTypeVersionRequest(input *types.CreateSlotTypeVersionInput) CreateSlotTypeVersionRequest {
	op := &aws.Operation{
		Name:       opCreateSlotTypeVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/slottypes/{name}/versions",
	}

	if input == nil {
		input = &types.CreateSlotTypeVersionInput{}
	}

	req := c.newRequest(op, input, &types.CreateSlotTypeVersionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreateSlotTypeVersionMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateSlotTypeVersionRequest{Request: req, Input: input, Copy: c.CreateSlotTypeVersionRequest}
}

// CreateSlotTypeVersionRequest is the request type for the
// CreateSlotTypeVersion API operation.
type CreateSlotTypeVersionRequest struct {
	*aws.Request
	Input *types.CreateSlotTypeVersionInput
	Copy  func(*types.CreateSlotTypeVersionInput) CreateSlotTypeVersionRequest
}

// Send marshals and sends the CreateSlotTypeVersion API request.
func (r CreateSlotTypeVersionRequest) Send(ctx context.Context) (*CreateSlotTypeVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSlotTypeVersionResponse{
		CreateSlotTypeVersionOutput: r.Request.Data.(*types.CreateSlotTypeVersionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSlotTypeVersionResponse is the response type for the
// CreateSlotTypeVersion API operation.
type CreateSlotTypeVersionResponse struct {
	*types.CreateSlotTypeVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSlotTypeVersion request.
func (r *CreateSlotTypeVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
