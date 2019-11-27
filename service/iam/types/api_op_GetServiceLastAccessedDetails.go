// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/iam/enums"
)

type GetServiceLastAccessedDetailsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the request generated by the GenerateServiceLastAccessedDetails
	// operation.
	//
	// JobId is a required field
	JobId *string `min:"36" type:"string" required:"true"`

	// Use this parameter only when paginating results and only after you receive
	// a response indicating that the results are truncated. Set it to the value
	// of the Marker element in the response that you received to indicate where
	// the next call should start.
	Marker *string `min:"1" type:"string"`

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true.
	//
	// If you do not include this parameter, the number of items defaults to 100.
	// Note that IAM might return fewer results, even when there are more results
	// available. In that case, the IsTruncated response element returns true, and
	// Marker contains a value to include in the subsequent call that tells the
	// service where to continue from.
	MaxItems *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s GetServiceLastAccessedDetailsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetServiceLastAccessedDetailsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetServiceLastAccessedDetailsInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 36))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetServiceLastAccessedDetailsOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains details about the reason the operation failed.
	Error *ErrorDetails `type:"structure"`

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer
	// than the MaxItems number of results even when there are more results available.
	// We recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated *bool `type:"boolean"`

	// The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601),
	// when the generated report job was completed or failed.
	//
	// This field is null if the job is still in progress, as indicated by a job
	// status value of IN_PROGRESS.
	//
	// JobCompletionDate is a required field
	JobCompletionDate *time.Time `type:"timestamp" required:"true"`

	// The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601),
	// when the report job was created.
	//
	// JobCreationDate is a required field
	JobCreationDate *time.Time `type:"timestamp" required:"true"`

	// The status of the job.
	//
	// JobStatus is a required field
	JobStatus enums.JobStatusType `type:"string" required:"true" enum:"true"`

	// When IsTruncated is true, this element is present and contains the value
	// to use for the Marker parameter in a subsequent pagination request.
	Marker *string `type:"string"`

	// A ServiceLastAccessed object that contains details about the most recent
	// attempt to access the service.
	//
	// ServicesLastAccessed is a required field
	ServicesLastAccessed []ServiceLastAccessed `type:"list" required:"true"`
}

// String returns the string representation
func (s GetServiceLastAccessedDetailsOutput) String() string {
	return awsutil.Prettify(s)
}
