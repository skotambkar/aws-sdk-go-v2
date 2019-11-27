// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/mediastore/enums"
)

// This section describes operations that you can perform on an AWS Elemental
// MediaStore container.
type Container struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the container. The ARN has the following
	// format:
	//
	// arn:aws:<region>:<account that owns this container>:container/<name of container>
	//
	// For example: arn:aws:mediastore:us-west-2:111122223333:container/movies
	ARN *string `min:"1" type:"string"`

	// The state of access logging on the container. This value is false by default,
	// indicating that AWS Elemental MediaStore does not send access logs to Amazon
	// CloudWatch Logs. When you enable access logging on the container, MediaStore
	// changes this value to true, indicating that the service delivers access logs
	// for objects stored in that container to CloudWatch Logs.
	AccessLoggingEnabled *bool `type:"boolean"`

	// Unix timestamp.
	CreationTime *time.Time `type:"timestamp"`

	// The DNS endpoint of the container. Use the endpoint to identify the specific
	// container when sending requests to the data plane. The service assigns this
	// value when the container is created. Once the value has been assigned, it
	// does not change.
	Endpoint *string `min:"1" type:"string"`

	// The name of the container.
	Name *string `min:"1" type:"string"`

	// The status of container creation or deletion. The status is one of the following:
	// CREATING, ACTIVE, or DELETING. While the service is creating the container,
	// the status is CREATING. When the endpoint is available, the status changes
	// to ACTIVE.
	Status enums.ContainerStatus `min:"1" type:"string" enum:"true"`
}

// String returns the string representation
func (s Container) String() string {
	return awsutil.Prettify(s)
}

// A rule for a CORS policy. You can add up to 100 rules to a CORS policy. If
// more than one rule applies, the service uses the first applicable rule listed.
type CorsRule struct {
	_ struct{} `type:"structure"`

	// Specifies which headers are allowed in a preflight OPTIONS request through
	// the Access-Control-Request-Headers header. Each header name that is specified
	// in Access-Control-Request-Headers must have a corresponding entry in the
	// rule. Only the headers that were requested are sent back.
	//
	// This element can contain only one wildcard character (*).
	//
	// AllowedHeaders is a required field
	AllowedHeaders []string `type:"list" required:"true"`

	// Identifies an HTTP method that the origin that is specified in the rule is
	// allowed to execute.
	//
	// Each CORS rule must contain at least one AllowedMethods and one AllowedOrigins
	// element.
	AllowedMethods []enums.MethodName `min:"1" type:"list"`

	// One or more response headers that you want users to be able to access from
	// their applications (for example, from a JavaScript XMLHttpRequest object).
	//
	// Each CORS rule must have at least one AllowedOrigins element. The string
	// value can include only one wildcard character (*), for example, http://*.example.com.
	// Additionally, you can specify only one wildcard character to allow cross-origin
	// access for all origins.
	//
	// AllowedOrigins is a required field
	AllowedOrigins []string `min:"1" type:"list" required:"true"`

	// One or more headers in the response that you want users to be able to access
	// from their applications (for example, from a JavaScript XMLHttpRequest object).
	//
	// This element is optional for each rule.
	ExposeHeaders []string `type:"list"`

	// The time in seconds that your browser caches the preflight response for the
	// specified resource.
	//
	// A CORS rule can have only one MaxAgeSeconds element.
	MaxAgeSeconds *int64 `type:"integer"`
}

// String returns the string representation
func (s CorsRule) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CorsRule) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CorsRule"}

	if s.AllowedHeaders == nil {
		invalidParams.Add(aws.NewErrParamRequired("AllowedHeaders"))
	}
	if s.AllowedMethods != nil && len(s.AllowedMethods) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AllowedMethods", 1))
	}

	if s.AllowedOrigins == nil {
		invalidParams.Add(aws.NewErrParamRequired("AllowedOrigins"))
	}
	if s.AllowedOrigins != nil && len(s.AllowedOrigins) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AllowedOrigins", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A collection of tags associated with a container. Each tag consists of a
// key:value pair, which can be anything you define. Typically, the tag key
// represents a category (such as "environment") and the tag value represents
// a specific value within that category (such as "test," "development," or
// "production"). You can add up to 50 tags to each container. For more information
// about tagging, including naming and usage conventions, see Tagging Resources
// in MediaStore (https://aws.amazon.com/documentation/mediastore/tagging).
type Tag struct {
	_ struct{} `type:"structure"`

	// Part of the key:value pair that defines a tag. You can use a tag key to describe
	// a category of information, such as "customer." Tag keys are case-sensitive.
	Key *string `min:"1" type:"string"`

	// Part of the key:value pair that defines a tag. You can use a tag value to
	// describe a specific value within a category, such as "companyA" or "companyB."
	// Tag values are case-sensitive.
	Value *string `type:"string"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
