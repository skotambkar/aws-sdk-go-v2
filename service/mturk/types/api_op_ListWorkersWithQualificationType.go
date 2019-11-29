// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/mturk/enums"
)

type ListWorkersWithQualificationTypeInput struct {
	_ struct{} `type:"structure"`

	// Limit the number of results returned.
	MaxResults *int64 `min:"1" type:"integer"`

	// Pagination Token
	NextToken *string `min:"1" type:"string"`

	// The ID of the Qualification type of the Qualifications to return.
	//
	// QualificationTypeId is a required field
	QualificationTypeId *string `min:"1" type:"string" required:"true"`

	// The status of the Qualifications to return. Can be Granted | Revoked.
	Status enums.QualificationStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListWorkersWithQualificationTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListWorkersWithQualificationTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListWorkersWithQualificationTypeInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.QualificationTypeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("QualificationTypeId"))
	}
	if s.QualificationTypeId != nil && len(*s.QualificationTypeId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("QualificationTypeId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListWorkersWithQualificationTypeOutput struct {
	_ struct{} `type:"structure"`

	// If the previous response was incomplete (because there is more data to retrieve),
	// Amazon Mechanical Turk returns a pagination token in the response. You can
	// use this pagination token to retrieve the next set of results.
	NextToken *string `min:"1" type:"string"`

	// The number of Qualifications on this page in the filtered results list, equivalent
	// to the number of Qualifications being returned by this call.
	NumResults *int64 `type:"integer"`

	// The list of Qualification elements returned by this call.
	Qualifications []Qualification `type:"list"`
}

// String returns the string representation
func (s ListWorkersWithQualificationTypeOutput) String() string {
	return awsutil.Prettify(s)
}