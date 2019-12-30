// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request for GetProfile operation.
type GetProfileInput struct {
	_ struct{} `type:"structure"`

	// The format of the profile to return. Supports application/json or application/x-amzn-ion.
	// Defaults to application/x-amzn-ion.
	Accept *string `location:"header" locationName:"Accept" type:"string"`

	// The end time of the profile to get. Either period or endTime must be specified.
	// Must be greater than start and the overall time range to be in the past and
	// not larger than a week.
	EndTime *time.Time `location:"querystring" locationName:"endTime" type:"timestamp" timestampFormat:"iso8601"`

	// Limit the max depth of the profile.
	MaxDepth *int64 `location:"querystring" locationName:"maxDepth" min:"1" type:"integer"`

	// The period of the profile to get. Exactly two of startTime, period and endTime
	// must be specified. Must be positive and the overall time range to be in the
	// past and not larger than a week.
	Period *string `location:"querystring" locationName:"period" min:"1" type:"string"`

	// The name of the profiling group.
	//
	// ProfilingGroupName is a required field
	ProfilingGroupName *string `location:"uri" locationName:"profilingGroupName" min:"1" type:"string" required:"true"`

	// The start time of the profile to get.
	StartTime *time.Time `location:"querystring" locationName:"startTime" type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s GetProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetProfileInput"}
	if s.MaxDepth != nil && *s.MaxDepth < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxDepth", 1))
	}
	if s.Period != nil && len(*s.Period) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Period", 1))
	}

	if s.ProfilingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProfilingGroupName"))
	}
	if s.ProfilingGroupName != nil && len(*s.ProfilingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProfilingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetProfileInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Accept != nil {
		v := *s.Accept

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Accept", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProfilingGroupName != nil {
		v := *s.ProfilingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "profilingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EndTime != nil {
		v := *s.EndTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "endTime",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: false}, metadata)
	}
	if s.MaxDepth != nil {
		v := *s.MaxDepth

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxDepth", protocol.Int64Value(v), metadata)
	}
	if s.Period != nil {
		v := *s.Period

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "period", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StartTime != nil {
		v := *s.StartTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "startTime",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: false}, metadata)
	}
	return nil
}

// Response for GetProfile operation.
type GetProfileOutput struct {
	_ struct{} `type:"structure" payload:"Profile"`

	// The content encoding of the profile in the payload.
	ContentEncoding *string `location:"header" locationName:"Content-Encoding" type:"string"`

	// The content type of the profile in the payload. Will be application/json
	// or application/x-amzn-ion based on Accept header in the request.
	//
	// ContentType is a required field
	ContentType *string `location:"header" locationName:"Content-Type" type:"string" required:"true"`

	// The profile representing the aggregation of agent profiles of the profiling
	// group for a time range.
	//
	// Profile is a required field
	Profile []byte `locationName:"profile" type:"blob" required:"true"`
}

// String returns the string representation
func (s GetProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContentEncoding != nil {
		v := *s.ContentEncoding

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Encoding", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Profile != nil {
		v := s.Profile

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "profile", protocol.BytesStream(v), metadata)
	}
	return nil
}

const opGetProfile = "GetProfile"

// GetProfileRequest returns a request value for making API operation for
// Amazon CodeGuru Profiler.
//
// Get the aggregated profile of a profiling group for the specified time range.
// If the requested time range does not align with the available aggregated
// profiles, it will be expanded to attain alignment. If aggregated profiles
// are available only for part of the period requested, the profile is returned
// from the earliest available to the latest within the requested time range.
// For instance, if the requested time range is from 00:00 to 00:20 and the
// available profiles are from 00:15 to 00:25, then the returned profile will
// be from 00:15 to 00:20.
//
//    // Example sending a request using GetProfileRequest.
//    req := client.GetProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeguruprofiler-2019-07-18/GetProfile
func (c *Client) GetProfileRequest(input *GetProfileInput) GetProfileRequest {
	op := &aws.Operation{
		Name:       opGetProfile,
		HTTPMethod: "GET",
		HTTPPath:   "/profilingGroups/{profilingGroupName}/profile",
	}

	if input == nil {
		input = &GetProfileInput{}
	}

	req := c.newRequest(op, input, &GetProfileOutput{})
	return GetProfileRequest{Request: req, Input: input, Copy: c.GetProfileRequest}
}

// GetProfileRequest is the request type for the
// GetProfile API operation.
type GetProfileRequest struct {
	*aws.Request
	Input *GetProfileInput
	Copy  func(*GetProfileInput) GetProfileRequest
}

// Send marshals and sends the GetProfile API request.
func (r GetProfileRequest) Send(ctx context.Context) (*GetProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetProfileResponse{
		GetProfileOutput: r.Request.Data.(*GetProfileOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetProfileResponse is the response type for the
// GetProfile API operation.
type GetProfileResponse struct {
	*GetProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetProfile request.
func (r *GetProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
