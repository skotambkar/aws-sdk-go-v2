// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeIngestionInput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID of the dataset used in the ingestion.
	//
	// DataSetId is a required field
	DataSetId *string `location:"uri" locationName:"DataSetId" type:"string" required:"true"`

	// An ID for the ingestion.
	//
	// IngestionId is a required field
	IngestionId *string `location:"uri" locationName:"IngestionId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeIngestionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeIngestionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeIngestionInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.DataSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSetId"))
	}

	if s.IngestionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IngestionId"))
	}
	if s.IngestionId != nil && len(*s.IngestionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IngestionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeIngestionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DataSetId != nil {
		v := *s.DataSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DataSetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IngestionId != nil {
		v := *s.IngestionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "IngestionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeIngestionOutput struct {
	_ struct{} `type:"structure"`

	// Information about the ingestion.
	Ingestion *Ingestion `type:"structure"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s DescribeIngestionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeIngestionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Ingestion != nil {
		v := s.Ingestion

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Ingestion", v, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opDescribeIngestion = "DescribeIngestion"

// DescribeIngestionRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Describes a SPICE ingestion.
//
//    // Example sending a request using DescribeIngestionRequest.
//    req := client.DescribeIngestionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DescribeIngestion
func (c *Client) DescribeIngestionRequest(input *DescribeIngestionInput) DescribeIngestionRequest {
	op := &aws.Operation{
		Name:       opDescribeIngestion,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/data-sets/{DataSetId}/ingestions/{IngestionId}",
	}

	if input == nil {
		input = &DescribeIngestionInput{}
	}

	req := c.newRequest(op, input, &DescribeIngestionOutput{})
	return DescribeIngestionRequest{Request: req, Input: input, Copy: c.DescribeIngestionRequest}
}

// DescribeIngestionRequest is the request type for the
// DescribeIngestion API operation.
type DescribeIngestionRequest struct {
	*aws.Request
	Input *DescribeIngestionInput
	Copy  func(*DescribeIngestionInput) DescribeIngestionRequest
}

// Send marshals and sends the DescribeIngestion API request.
func (r DescribeIngestionRequest) Send(ctx context.Context) (*DescribeIngestionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeIngestionResponse{
		DescribeIngestionOutput: r.Request.Data.(*DescribeIngestionOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeIngestionResponse is the response type for the
// DescribeIngestion API operation.
type DescribeIngestionResponse struct {
	*DescribeIngestionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeIngestion request.
func (r *DescribeIngestionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
