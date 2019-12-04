// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateThreatIntelSetInput struct {
	_ struct{} `type:"structure"`

	// A boolean value that indicates whether GuardDuty is to start using the uploaded
	// ThreatIntelSet.
	//
	// Activate is a required field
	Activate *bool `locationName:"activate" type:"boolean" required:"true"`

	// The idempotency token for the create request.
	ClientToken *string `locationName:"clientToken" type:"string" idempotencyToken:"true"`

	// The unique ID of the detector of the GuardDuty account for which you want
	// to create a threatIntelSet.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// The format of the file that contains the ThreatIntelSet.
	//
	// Format is a required field
	Format ThreatIntelSetFormat `locationName:"format" min:"1" type:"string" required:"true" enum:"true"`

	// The URI of the file that contains the ThreatIntelSet. For example (https://s3.us-west-2.amazonaws.com/my-bucket/my-object-key).
	//
	// Location is a required field
	Location *string `locationName:"location" min:"1" type:"string" required:"true"`

	// A user-friendly ThreatIntelSet name that is displayed in all finding generated
	// by activity that involves IP addresses included in this ThreatIntelSet.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// The tags to be added to a new Threat List resource.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`
}

// String returns the string representation
func (s CreateThreatIntelSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateThreatIntelSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateThreatIntelSetInput"}

	if s.Activate == nil {
		invalidParams.Add(aws.NewErrParamRequired("Activate"))
	}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}
	if len(s.Format) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Format"))
	}

	if s.Location == nil {
		invalidParams.Add(aws.NewErrParamRequired("Location"))
	}
	if s.Location != nil && len(*s.Location) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Location", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateThreatIntelSetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Activate != nil {
		v := *s.Activate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "activate", protocol.BoolValue(v), metadata)
	}
	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Format) > 0 {
		v := s.Format

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "format", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "location", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateThreatIntelSetOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the ThreatIntelSet resource.
	//
	// ThreatIntelSetId is a required field
	ThreatIntelSetId *string `locationName:"threatIntelSetId" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateThreatIntelSetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateThreatIntelSetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ThreatIntelSetId != nil {
		v := *s.ThreatIntelSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "threatIntelSetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateThreatIntelSet = "CreateThreatIntelSet"

// CreateThreatIntelSetRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Create a new ThreatIntelSet. ThreatIntelSets consist of known malicious IP
// addresses. GuardDuty generates findings based on ThreatIntelSets. Only users
// of the master account can use this operation.
//
//    // Example sending a request using CreateThreatIntelSetRequest.
//    req := client.CreateThreatIntelSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/CreateThreatIntelSet
func (c *Client) CreateThreatIntelSetRequest(input *CreateThreatIntelSetInput) CreateThreatIntelSetRequest {
	op := &aws.Operation{
		Name:       opCreateThreatIntelSet,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/threatintelset",
	}

	if input == nil {
		input = &CreateThreatIntelSetInput{}
	}

	req := c.newRequest(op, input, &CreateThreatIntelSetOutput{})
	return CreateThreatIntelSetRequest{Request: req, Input: input, Copy: c.CreateThreatIntelSetRequest}
}

// CreateThreatIntelSetRequest is the request type for the
// CreateThreatIntelSet API operation.
type CreateThreatIntelSetRequest struct {
	*aws.Request
	Input *CreateThreatIntelSetInput
	Copy  func(*CreateThreatIntelSetInput) CreateThreatIntelSetRequest
}

// Send marshals and sends the CreateThreatIntelSet API request.
func (r CreateThreatIntelSetRequest) Send(ctx context.Context) (*CreateThreatIntelSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateThreatIntelSetResponse{
		CreateThreatIntelSetOutput: r.Request.Data.(*CreateThreatIntelSetOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateThreatIntelSetResponse is the response type for the
// CreateThreatIntelSet API operation.
type CreateThreatIntelSetResponse struct {
	*CreateThreatIntelSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateThreatIntelSet request.
func (r *CreateThreatIntelSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
