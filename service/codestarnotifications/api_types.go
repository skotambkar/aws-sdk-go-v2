// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestarnotifications

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// Returns information about an event that has triggered a notification rule.
type EventTypeSummary struct {
	_ struct{} `type:"structure"`

	// The system-generated ID of the event.
	EventTypeId *string `min:"1" type:"string"`

	// The name of the event.
	EventTypeName *string `type:"string"`

	// The resource type of the event.
	ResourceType *string `min:"1" type:"string"`

	// The name of the service for which the event applies.
	ServiceName *string `type:"string"`
}

// String returns the string representation
func (s EventTypeSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EventTypeSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.EventTypeId != nil {
		v := *s.EventTypeId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EventTypeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EventTypeName != nil {
		v := *s.EventTypeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EventTypeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceType != nil {
		v := *s.ResourceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ResourceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ServiceName != nil {
		v := *s.ServiceName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ServiceName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about a filter to apply to the list of returned event types.
// You can filter by resource type or service name.
type ListEventTypesFilter struct {
	_ struct{} `type:"structure"`

	// The system-generated name of the filter type you want to filter by.
	//
	// Name is a required field
	Name ListEventTypesFilterName `type:"string" required:"true" enum:"true"`

	// The name of the resource type (for example, pipeline) or service name (for
	// example, CodePipeline) that you want to filter by.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListEventTypesFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListEventTypesFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListEventTypesFilter"}
	if len(s.Name) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListEventTypesFilter) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Name) > 0 {
		v := s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Value != nil {
		v := *s.Value

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Value", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about a filter to apply to the list of returned notification
// rules. You can filter by event type, owner, resource, or target.
type ListNotificationRulesFilter struct {
	_ struct{} `type:"structure"`

	// The name of the attribute you want to use to filter the returned notification
	// rules.
	//
	// Name is a required field
	Name ListNotificationRulesFilterName `type:"string" required:"true" enum:"true"`

	// The value of the attribute you want to use to filter the returned notification
	// rules. For example, if you specify filtering by RESOURCE in Name, you might
	// specify the ARN of a pipeline in AWS CodePipeline for the value.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListNotificationRulesFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListNotificationRulesFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListNotificationRulesFilter"}
	if len(s.Name) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListNotificationRulesFilter) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Name) > 0 {
		v := s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Value != nil {
		v := *s.Value

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Value", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about a filter to apply to the list of returned targets. You
// can filter by target type, address, or status. For example, to filter results
// to notification rules that have active Amazon SNS topics as targets, you
// could specify a ListTargetsFilter Name as TargetType and a Value of SNS,
// and a Name of TARGET_STATUS and a Value of ACTIVE.
type ListTargetsFilter struct {
	_ struct{} `type:"structure"`

	// The name of the attribute you want to use to filter the returned targets.
	//
	// Name is a required field
	Name ListTargetsFilterName `type:"string" required:"true" enum:"true"`

	// The value of the attribute you want to use to filter the returned targets.
	// For example, if you specify SNS for the Target type, you could specify an
	// Amazon Resource Name (ARN) for a topic as the value.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListTargetsFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTargetsFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTargetsFilter"}
	if len(s.Name) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListTargetsFilter) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Name) > 0 {
		v := s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Value != nil {
		v := *s.Value

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Value", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about a specified notification rule.
type NotificationRuleSummary struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the notification rule.
	Arn *string `type:"string"`

	// The unique ID of the notification rule.
	Id *string `min:"1" type:"string"`
}

// String returns the string representation
func (s NotificationRuleSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s NotificationRuleSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about the SNS topics associated with a notification rule.
type Target struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the SNS topic.
	TargetAddress *string `min:"1" type:"string" sensitive:"true"`

	// The target type. Can be an Amazon SNS topic.
	TargetType *string `type:"string"`
}

// String returns the string representation
func (s Target) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Target) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Target"}
	if s.TargetAddress != nil && len(*s.TargetAddress) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetAddress", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Target) MarshalFields(e protocol.FieldEncoder) error {
	if s.TargetAddress != nil {
		v := *s.TargetAddress

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TargetAddress", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TargetType != nil {
		v := *s.TargetType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TargetType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about the targets specified for a notification rule.
type TargetSummary struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the SNS topic.
	TargetAddress *string `min:"1" type:"string" sensitive:"true"`

	// The status of the target.
	TargetStatus TargetStatus `type:"string" enum:"true"`

	// The type of the target (for example, SNS).
	TargetType *string `type:"string"`
}

// String returns the string representation
func (s TargetSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TargetSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.TargetAddress != nil {
		v := *s.TargetAddress

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TargetAddress", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.TargetStatus) > 0 {
		v := s.TargetStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TargetStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.TargetType != nil {
		v := *s.TargetType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TargetType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}
