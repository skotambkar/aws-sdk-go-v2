// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Container for the parameters to the DeleteAnalysisScheme operation. Specifies
// the name of the domain you want to update and the analysis scheme you want
// to delete.
type DeleteAnalysisSchemeInput struct {
	_ struct{} `type:"structure"`

	// The name of the analysis scheme you want to delete.
	//
	// AnalysisSchemeName is a required field
	AnalysisSchemeName *string `min:"1" type:"string" required:"true"`

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start
	// with a letter or number and can contain the following characters: a-z (lowercase),
	// 0-9, and - (hyphen).
	//
	// DomainName is a required field
	DomainName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAnalysisSchemeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAnalysisSchemeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAnalysisSchemeInput"}

	if s.AnalysisSchemeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AnalysisSchemeName"))
	}
	if s.AnalysisSchemeName != nil && len(*s.AnalysisSchemeName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AnalysisSchemeName", 1))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a DeleteAnalysisScheme request. Contains the status of the
// deleted analysis scheme.
type DeleteAnalysisSchemeOutput struct {
	_ struct{} `type:"structure"`

	// The status of the analysis scheme being deleted.
	//
	// AnalysisScheme is a required field
	AnalysisScheme *AnalysisSchemeStatus `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteAnalysisSchemeOutput) String() string {
	return awsutil.Prettify(s)
}
