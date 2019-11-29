// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeImagesInput struct {
	_ struct{} `type:"structure"`

	// The filter key and value with which to filter your DescribeImages results.
	Filter *DescribeImagesFilter `locationName:"filter" type:"structure"`

	// The list of image IDs for the requested repository.
	ImageIds []ImageIdentifier `locationName:"imageIds" min:"1" type:"list"`

	// The maximum number of repository results returned by DescribeImages in paginated
	// output. When this parameter is used, DescribeImages only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another DescribeImages
	// request with the returned nextToken value. This value can be between 1 and
	// 1000. If this parameter is not used, then DescribeImages returns up to 100
	// results and a nextToken value, if applicable. This option cannot be used
	// when you specify images with imageIds.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The nextToken value returned from a previous paginated DescribeImages request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	// This option cannot be used when you specify images with imageIds.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The AWS account ID associated with the registry that contains the repository
	// in which to describe images. If you do not specify a registry, the default
	// registry is assumed.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The repository that contains the images to describe.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeImagesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeImagesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeImagesInput"}
	if s.ImageIds != nil && len(s.ImageIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ImageIds", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 2))
	}
	if s.ImageIds != nil {
		for i, v := range s.ImageIds {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ImageIds", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeImagesOutput struct {
	_ struct{} `type:"structure"`

	// A list of ImageDetail objects that contain data about the image.
	ImageDetails []ImageDetail `locationName:"imageDetails" type:"list"`

	// The nextToken value to include in a future DescribeImages request. When the
	// results of a DescribeImages request exceed maxResults, this value can be
	// used to retrieve the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeImagesOutput) String() string {
	return awsutil.Prettify(s)
}
