// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListSchemaVersionsInput struct {
	_ struct{} `type:"structure"`

	Limit *int64 `location:"querystring" locationName:"limit" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// RegistryName is a required field
	RegistryName *string `location:"uri" locationName:"registryName" type:"string" required:"true"`

	// SchemaName is a required field
	SchemaName *string `location:"uri" locationName:"schemaName" type:"string" required:"true"`
}

// String returns the string representation
func (s ListSchemaVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSchemaVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSchemaVersionsInput"}

	if s.RegistryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegistryName"))
	}

	if s.SchemaName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListSchemaVersionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RegistryName != nil {
		v := *s.RegistryName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "registryName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaName != nil {
		v := *s.SchemaName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "schemaName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListSchemaVersionsOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `type:"string"`

	SchemaVersions []SchemaVersionSummary `type:"list"`
}

// String returns the string representation
func (s ListSchemaVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListSchemaVersionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaVersions != nil {
		v := s.SchemaVersions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "SchemaVersions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListSchemaVersions = "ListSchemaVersions"

// ListSchemaVersionsRequest returns a request value for making API operation for
// Schemas.
//
// Provides a list of the schema versions and related information.
//
//    // Example sending a request using ListSchemaVersionsRequest.
//    req := client.ListSchemaVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/schemas-2019-12-02/ListSchemaVersions
func (c *Client) ListSchemaVersionsRequest(input *ListSchemaVersionsInput) ListSchemaVersionsRequest {
	op := &aws.Operation{
		Name:       opListSchemaVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/registries/name/{registryName}/schemas/name/{schemaName}/versions",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListSchemaVersionsInput{}
	}

	req := c.newRequest(op, input, &ListSchemaVersionsOutput{})
	return ListSchemaVersionsRequest{Request: req, Input: input, Copy: c.ListSchemaVersionsRequest}
}

// ListSchemaVersionsRequest is the request type for the
// ListSchemaVersions API operation.
type ListSchemaVersionsRequest struct {
	*aws.Request
	Input *ListSchemaVersionsInput
	Copy  func(*ListSchemaVersionsInput) ListSchemaVersionsRequest
}

// Send marshals and sends the ListSchemaVersions API request.
func (r ListSchemaVersionsRequest) Send(ctx context.Context) (*ListSchemaVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSchemaVersionsResponse{
		ListSchemaVersionsOutput: r.Request.Data.(*ListSchemaVersionsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListSchemaVersionsRequestPaginator returns a paginator for ListSchemaVersions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListSchemaVersionsRequest(input)
//   p := schemas.NewListSchemaVersionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListSchemaVersionsPaginator(req ListSchemaVersionsRequest) ListSchemaVersionsPaginator {
	return ListSchemaVersionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListSchemaVersionsInput
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

// ListSchemaVersionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListSchemaVersionsPaginator struct {
	aws.Pager
}

func (p *ListSchemaVersionsPaginator) CurrentPage() *ListSchemaVersionsOutput {
	return p.Pager.CurrentPage().(*ListSchemaVersionsOutput)
}

// ListSchemaVersionsResponse is the response type for the
// ListSchemaVersions API operation.
type ListSchemaVersionsResponse struct {
	*ListSchemaVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSchemaVersions request.
func (r *ListSchemaVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
