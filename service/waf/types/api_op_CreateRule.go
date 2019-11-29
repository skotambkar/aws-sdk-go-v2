// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateRuleInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// A friendly name or description for the metrics for this Rule. The name can
	// contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length
	// 128 and minimum length one. It can't contain whitespace or metric names reserved
	// for AWS WAF, including "All" and "Default_Action." You can't change the name
	// of the metric after you create the Rule.
	//
	// MetricName is a required field
	MetricName *string `type:"string" required:"true"`

	// A friendly name or description of the Rule. You can't change the name of
	// a Rule after you create it.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	Tags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s CreateRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRuleInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.MetricName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MetricName"))
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
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateRuleOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the CreateRule request. You can also
	// use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`

	// The Rule returned in the CreateRule response.
	Rule *Rule `type:"structure"`
}

// String returns the string representation
func (s CreateRuleOutput) String() string {
	return awsutil.Prettify(s)
}