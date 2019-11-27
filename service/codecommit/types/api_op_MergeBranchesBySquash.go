// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/enums"
)

type MergeBranchesBySquashInput struct {
	_ struct{} `type:"structure"`

	// The name of the author who created the commit. This information will be used
	// as both the author and committer for the commit.
	AuthorName *string `locationName:"authorName" type:"string"`

	// The commit message for the merge.
	CommitMessage *string `locationName:"commitMessage" type:"string"`

	// The level of conflict detail to use. If unspecified, the default FILE_LEVEL
	// is used, which will return a not mergeable result if the same file has differences
	// in both branches. If LINE_LEVEL is specified, a conflict will be considered
	// not mergeable if the same file in both branches has differences on the same
	// line.
	ConflictDetailLevel enums.ConflictDetailLevelTypeEnum `locationName:"conflictDetailLevel" type:"string" enum:"true"`

	// A list of inputs to use when resolving conflicts during a merge if AUTOMERGE
	// is chosen as the conflict resolution strategy.
	ConflictResolution *ConflictResolution `locationName:"conflictResolution" type:"structure"`

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

	// The email address of the person merging the branches. This information will
	// be used in the commit information for the merge.
	Email *string `locationName:"email" type:"string"`

	// If the commit contains deletions, whether to keep a folder or folder structure
	// if the changes leave the folders empty. If this is specified as true, a .gitkeep
	// file will be created for empty folders. The default is false.
	KeepEmptyFolders *bool `locationName:"keepEmptyFolders" type:"boolean"`

	// The name of the repository where you want to merge two branches.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`

	// The branch, tag, HEAD, or other fully qualified reference used to identify
	// a commit. For example, a branch name or a full commit ID.
	//
	// SourceCommitSpecifier is a required field
	SourceCommitSpecifier *string `locationName:"sourceCommitSpecifier" type:"string" required:"true"`

	// The branch where the merge will be applied.
	TargetBranch *string `locationName:"targetBranch" min:"1" type:"string"`
}

// String returns the string representation
func (s MergeBranchesBySquashInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MergeBranchesBySquashInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MergeBranchesBySquashInput"}

	if s.DestinationCommitSpecifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationCommitSpecifier"))
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
	if s.TargetBranch != nil && len(*s.TargetBranch) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetBranch", 1))
	}
	if s.ConflictResolution != nil {
		if err := s.ConflictResolution.Validate(); err != nil {
			invalidParams.AddNested("ConflictResolution", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type MergeBranchesBySquashOutput struct {
	_ struct{} `type:"structure"`

	// The commit ID of the merge in the destination or target branch.
	CommitId *string `locationName:"commitId" type:"string"`

	// The tree ID of the merge in the destination or target branch.
	TreeId *string `locationName:"treeId" type:"string"`
}

// String returns the string representation
func (s MergeBranchesBySquashOutput) String() string {
	return awsutil.Prettify(s)
}
