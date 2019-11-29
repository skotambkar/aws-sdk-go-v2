// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ecs/enums"
)

type DescribeTaskDefinitionInput struct {
	_ struct{} `type:"structure"`

	// Specifies whether to see the resource tags for the task definition. If TAGS
	// is specified, the tags are included in the response. If this field is omitted,
	// tags are not included in the response.
	Include []enums.TaskDefinitionField `locationName:"include" type:"list"`

	// The family for the latest ACTIVE revision, family and revision (family:revision)
	// for a specific revision in the family, or full Amazon Resource Name (ARN)
	// of the task definition to describe.
	//
	// TaskDefinition is a required field
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTaskDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTaskDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTaskDefinitionInput"}

	if s.TaskDefinition == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskDefinition"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeTaskDefinitionOutput struct {
	_ struct{} `type:"structure"`

	// The metadata that is applied to the task definition to help you categorize
	// and organize them. Each tag consists of a key and an optional value, both
	// of which you define.
	//
	// The following basic restrictions apply to tags:
	//
	//    * Maximum number of tags per resource - 50
	//
	//    * For each resource, each tag key must be unique, and each tag key can
	//    have only one value.
	//
	//    * Maximum key length - 128 Unicode characters in UTF-8
	//
	//    * Maximum value length - 256 Unicode characters in UTF-8
	//
	//    * If your tagging schema is used across multiple services and resources,
	//    remember that other services may have restrictions on allowed characters.
	//    Generally allowed characters are: letters, numbers, and spaces representable
	//    in UTF-8, and the following characters: + - = . _ : / @.
	//
	//    * Tag keys and values are case-sensitive.
	//
	//    * Do not use aws:, AWS:, or any upper or lowercase combination of such
	//    as a prefix for either keys or values as it is reserved for AWS use. You
	//    cannot edit or delete tag keys or values with this prefix. Tags with this
	//    prefix do not count against your tags per resource limit.
	Tags []Tag `locationName:"tags" type:"list"`

	// The full task definition description.
	TaskDefinition *TaskDefinition `locationName:"taskDefinition" type:"structure"`
}

// String returns the string representation
func (s DescribeTaskDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}
