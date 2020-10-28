package s3shared

import (
	"fmt"
	"strings"

	smithyhttp "github.com/awslabs/smithy-go/transport/http"

	awsarn "github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/internal/s3shared/arn"
)

type ResourceRequestOptions struct {
	// RequestRegion is the region configured on the request config
	RequestRegion string

	// SigningRegion is the signing region resolved for the request
	SigningRegion string

	// PartitionID is the resolved partition id for the provided request region
	PartitionID string

	// UseARNRegion indicates if client should use the region provided in an ARN resource
	UseARNRegion bool

	// HasCustomEndpoint indicates if a custom endpoint is provided
	HasCustomEndpoint bool
}

type ResourceRequest struct {
	Resource arn.Resource
	Options  ResourceRequestOptions
	Request  *smithyhttp.Request
}

// ARN returns the resource ARN
func (r ResourceRequest) ARN() awsarn.ARN {
	return r.Resource.GetARN()
}

// UseFIPS returns true if request config region is FIPS region.
func (r ResourceRequest) UseFips() (bool, error) {
	return IsFIPS(r.Options.RequestRegion)
}

// ResourceConfiguredForFIPS returns true if resource ARNs region is FIPS
func (r ResourceRequest) ResourceConfiguredForFIPS() (bool, error) {
	return IsFIPS(r.ARN().Region)
}

// AllowCrossRegion returns a bool value to denote if S3UseARNRegion flag is set
func (r ResourceRequest) AllowCrossRegion() bool {
	return r.Options.UseARNRegion
}

// IsCrossPartition returns true if request is configured for region of another partition, than
// the partition that resource ARN region resolves to.
func (r ResourceRequest) IsCrossPartition() (bool, error) {
	// These error checks should never be triggered, unless validations are turned off
	rv := r.Options.PartitionID
	if len(rv) == 0 {
		return false, fmt.Errorf("partition id was not found for provided request region")
	}

	av := r.Resource.GetARN().Partition
	if len(av) == 0 {
		return false, fmt.Errorf("no partition id for provided ARN")
	}

	return !strings.EqualFold(rv, av), nil
}

// IsCrossRegion returns true if request signing region is not same as arn region
func (r ResourceRequest) IsCrossRegion() (bool, error) {
	v := r.Options.SigningRegion
	if len(v) == 0 {
		return false, fmt.Errorf("empty signing region provided for IsCrossRegion check")
	}
	return !strings.EqualFold(v, r.Resource.GetARN().Region), nil
}

// HasCustomEndpoint returns true if custom endpoint is provided
func (r ResourceRequest) HasCustomEndpoint() bool {
	return r.Options.HasCustomEndpoint
}

// TODO: should this be moved in aws endpoints package
// IsFIPS returns true if region is a fips region
func IsFIPS(clientRegion string) (bool, error) {
	if len(clientRegion) == 0 {
		return false, fmt.Errorf("empty region provided for isFIPS assertion")
	}

	return (strings.HasPrefix(clientRegion, "fips-") ||
		strings.HasSuffix(clientRegion, "-fips")), nil
}
