// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/enums"
)

type PutFileInput struct {
	_ struct{} `type:"structure"`

	// The name of the branch where you want to add or update the file. If this
	// is an empty repository, this branch will be created.
	//
	// BranchName is a required field
	BranchName *string `locationName:"branchName" min:"1" type:"string" required:"true"`

	// A message about why this file was added or updated. While optional, adding
	// a message is strongly encouraged in order to provide a more useful commit
	// history for your repository.
	CommitMessage *string `locationName:"commitMessage" type:"string"`

	// An email address for the person adding or updating the file.
	Email *string `locationName:"email" type:"string"`

	// The content of the file, in binary object format.
	//
	// FileContent is automatically base64 encoded/decoded by the SDK.
	//
	// FileContent is a required field
	FileContent []byte `locationName:"fileContent" type:"blob" required:"true"`

	// The file mode permissions of the blob. Valid file mode permissions are listed
	// below.
	FileMode enums.FileModeTypeEnum `locationName:"fileMode" type:"string" enum:"true"`

	// The name of the file you want to add or update, including the relative path
	// to the file in the repository.
	//
	// If the path does not currently exist in the repository, the path will be
	// created as part of adding the file.
	//
	// FilePath is a required field
	FilePath *string `locationName:"filePath" type:"string" required:"true"`

	// The name of the person adding or updating the file. While optional, adding
	// a name is strongly encouraged in order to provide a more useful commit history
	// for your repository.
	Name *string `locationName:"name" type:"string"`

	// The full commit ID of the head commit in the branch where you want to add
	// or update the file. If this is an empty repository, no commit ID is required.
	// If this is not an empty repository, a commit ID is required.
	//
	// The commit ID must match the ID of the head commit at the time of the operation,
	// or an error will occur, and the file will not be added or updated.
	ParentCommitId *string `locationName:"parentCommitId" type:"string"`

	// The name of the repository where you want to add or update the file.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutFileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutFileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutFileInput"}

	if s.BranchName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BranchName"))
	}
	if s.BranchName != nil && len(*s.BranchName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BranchName", 1))
	}

	if s.FileContent == nil {
		invalidParams.Add(aws.NewErrParamRequired("FileContent"))
	}

	if s.FilePath == nil {
		invalidParams.Add(aws.NewErrParamRequired("FilePath"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutFileOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the blob, which is its SHA-1 pointer.
	//
	// BlobId is a required field
	BlobId *string `locationName:"blobId" type:"string" required:"true"`

	// The full SHA of the commit that contains this file change.
	//
	// CommitId is a required field
	CommitId *string `locationName:"commitId" type:"string" required:"true"`

	// The full SHA-1 pointer of the tree information for the commit that contains
	// this file change.
	//
	// TreeId is a required field
	TreeId *string `locationName:"treeId" type:"string" required:"true"`
}

// String returns the string representation
func (s PutFileOutput) String() string {
	return awsutil.Prettify(s)
}
