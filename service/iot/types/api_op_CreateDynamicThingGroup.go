// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDynamicThingGroupInput struct {
	_ struct{} `type:"structure"`

	// The dynamic thing group index name.
	//
	// Currently one index is supported: "AWS_Things".
	IndexName *string `locationName:"indexName" min:"1" type:"string"`

	// The dynamic thing group search query string.
	//
	// See Query Syntax (https://docs.aws.amazon.com/iot/latest/developerguide/query-syntax.html)
	// for information about query string syntax.
	//
	// QueryString is a required field
	QueryString *string `locationName:"queryString" min:"1" type:"string" required:"true"`

	// The dynamic thing group query version.
	//
	// Currently one query version is supported: "2017-09-30". If not specified,
	// the query version defaults to this value.
	QueryVersion *string `locationName:"queryVersion" type:"string"`

	// Metadata which can be used to manage the dynamic thing group.
	Tags []Tag `locationName:"tags" type:"list"`

	// The dynamic thing group name to create.
	//
	// ThingGroupName is a required field
	ThingGroupName *string `location:"uri" locationName:"thingGroupName" min:"1" type:"string" required:"true"`

	// The dynamic thing group properties.
	ThingGroupProperties *ThingGroupProperties `locationName:"thingGroupProperties" type:"structure"`
}

// String returns the string representation
func (s CreateDynamicThingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDynamicThingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDynamicThingGroupInput"}
	if s.IndexName != nil && len(*s.IndexName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IndexName", 1))
	}

	if s.QueryString == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueryString"))
	}
	if s.QueryString != nil && len(*s.QueryString) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("QueryString", 1))
	}

	if s.ThingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingGroupName"))
	}
	if s.ThingGroupName != nil && len(*s.ThingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDynamicThingGroupOutput struct {
	_ struct{} `type:"structure"`

	// The dynamic thing group index name.
	IndexName *string `locationName:"indexName" min:"1" type:"string"`

	// The dynamic thing group search query string.
	QueryString *string `locationName:"queryString" min:"1" type:"string"`

	// The dynamic thing group query version.
	QueryVersion *string `locationName:"queryVersion" type:"string"`

	// The dynamic thing group ARN.
	ThingGroupArn *string `locationName:"thingGroupArn" type:"string"`

	// The dynamic thing group ID.
	ThingGroupId *string `locationName:"thingGroupId" min:"1" type:"string"`

	// The dynamic thing group name.
	ThingGroupName *string `locationName:"thingGroupName" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateDynamicThingGroupOutput) String() string {
	return awsutil.Prettify(s)
}