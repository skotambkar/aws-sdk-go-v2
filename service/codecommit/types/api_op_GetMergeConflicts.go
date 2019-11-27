// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/enums"
)

type GetMergeConflictsInput struct {
	_ struct{} `type:"structure"`

	// The level of conflict detail to use. If unspecified, the default FILE_LEVEL
	// is used, which will return a not mergeable result if the same file has differences
	// in both branches. If LINE_LEVEL is specified, a conflict will be considered
	// not mergeable if the same file in both branches has differences on the same
	// line.
	ConflictDetailLevel enums.ConflictDetailLevelTypeEnum `locationName:"conflictDetailLevel" type:"string" enum:"true"`

	// Specifies which branch to use when resolving conflicts, or whether to attempt
	// automatically merging two versions of a file. The default is NONE, which
	// requires any conflicts to be resolved manually before the merge operation
	// will be successful.
	ConflictResolutionStrategy enums.ConflictResolutionStrategyTypeEnum `locationName:"conflictResolutionStrategy" type:"string" enum:"true"`

	// The branch, tag, HEAD, or other fully qualified reference used to identify
	// a commit. For example, a branch name or a full commit ID.
	//
	// DestinationCommitSpecifier is a required field
	DestinationCommitSpecifier *string `locationName:"destinationCommitSpecifier" type:"string" required:"true"`

	// The maximum number of files to include in the output.
	MaxConflictFiles *int64 `locationName:"maxConflictFiles" type:"integer"`

	// The merge option or strategy you want to use to merge the code.
	//
	// MergeOption is a required field
	MergeOption enums.MergeOptionTypeEnum `locationName:"mergeOption" type:"string" required:"true" enum:"true"`

	// An enumeration token that when provided in a request, returns the next batch
	// of the results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The name of the repository where the pull request was created.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`

	// The branch, tag, HEAD, or other fully qualified reference used to identify
	// a commit. For example, a branch name or a full commit ID.
	//
	// SourceCommitSpecifier is a required field
	SourceCommitSpecifier *string `locationName:"sourceCommitSpecifier" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMergeConflictsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMergeConflictsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMergeConflictsInput"}

	if s.DestinationCommitSpecifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationCommitSpecifier"))
	}
	if len(s.MergeOption) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("MergeOption"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if s.SourceCommitSpecifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceCommitSpecifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetMergeConflictsOutput struct {
	_ struct{} `type:"structure"`

	// The commit ID of the merge base.
	BaseCommitId *string `locationName:"baseCommitId" type:"string"`

	// A list of metadata for any conflicting files. If the specified merge strategy
	// is FAST_FORWARD_MERGE, this list will always be empty.
	//
	// ConflictMetadataList is a required field
	ConflictMetadataList []ConflictMetadata `locationName:"conflictMetadataList" type:"list" required:"true"`

	// The commit ID of the destination commit specifier that was used in the merge
	// evaluation.
	//
	// DestinationCommitId is a required field
	DestinationCommitId *string `locationName:"destinationCommitId" type:"string" required:"true"`

	// A Boolean value that indicates whether the code is mergeable by the specified
	// merge option.
	//
	// Mergeable is a required field
	Mergeable *bool `locationName:"mergeable" type:"boolean" required:"true"`

	// An enumeration token that can be used in a request to return the next batch
	// of the results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The commit ID of the source commit specifier that was used in the merge evaluation.
	//
	// SourceCommitId is a required field
	SourceCommitId *string `locationName:"sourceCommitId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMergeConflictsOutput) String() string {
	return awsutil.Prettify(s)
}
