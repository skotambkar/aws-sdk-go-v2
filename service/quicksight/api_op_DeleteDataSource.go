// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteDataSourceInput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID of the data source. This ID is unique per AWS Region for each AWS
	// account.
	//
	// DataSourceId is a required field
	DataSourceId *string `location:"uri" locationName:"DataSourceId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDataSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDataSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDataSourceInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.DataSourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSourceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDataSourceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DataSourceId != nil {
		v := *s.DataSourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DataSourceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteDataSourceOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the data source that you deleted.
	Arn *string `type:"string"`

	// The ID of the data source. This ID is unique per AWS Region for each AWS
	// account.
	DataSourceId *string `type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s DeleteDataSourceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDataSourceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DataSourceId != nil {
		v := *s.DataSourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DataSourceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opDeleteDataSource = "DeleteDataSource"

// DeleteDataSourceRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Deletes the data source permanently. This action breaks all the datasets
// that reference the deleted data source.
//
//    // Example sending a request using DeleteDataSourceRequest.
//    req := client.DeleteDataSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DeleteDataSource
func (c *Client) DeleteDataSourceRequest(input *DeleteDataSourceInput) DeleteDataSourceRequest {
	op := &aws.Operation{
		Name:       opDeleteDataSource,
		HTTPMethod: "DELETE",
		HTTPPath:   "/accounts/{AwsAccountId}/data-sources/{DataSourceId}",
	}

	if input == nil {
		input = &DeleteDataSourceInput{}
	}

	req := c.newRequest(op, input, &DeleteDataSourceOutput{})
	return DeleteDataSourceRequest{Request: req, Input: input, Copy: c.DeleteDataSourceRequest}
}

// DeleteDataSourceRequest is the request type for the
// DeleteDataSource API operation.
type DeleteDataSourceRequest struct {
	*aws.Request
	Input *DeleteDataSourceInput
	Copy  func(*DeleteDataSourceInput) DeleteDataSourceRequest
}

// Send marshals and sends the DeleteDataSource API request.
func (r DeleteDataSourceRequest) Send(ctx context.Context) (*DeleteDataSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDataSourceResponse{
		DeleteDataSourceOutput: r.Request.Data.(*DeleteDataSourceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDataSourceResponse is the response type for the
// DeleteDataSource API operation.
type DeleteDataSourceResponse struct {
	*DeleteDataSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDataSource request.
func (r *DeleteDataSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
