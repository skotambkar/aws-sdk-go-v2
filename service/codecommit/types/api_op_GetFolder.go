// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetFolderInput struct {
	_ struct{} `type:"structure"`

	// A fully-qualified reference used to identify a commit that contains the version
	// of the folder's content to return. A fully-qualified reference can be a commit
	// ID, branch name, tag, or reference such as HEAD. If no specifier is provided,
	// the folder content will be returned as it exists in the HEAD commit.
	CommitSpecifier *string `locationName:"commitSpecifier" type:"string"`

	// The fully-qualified path to the folder whose contents will be returned, including
	// the folder name. For example, /examples is a fully-qualified path to a folder
	// named examples that was created off of the root directory (/) of a repository.
	//
	// FolderPath is a required field
	FolderPath *string `locationName:"folderPath" type:"string" required:"true"`

	// The name of the repository.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetFolderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFolderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFolderInput"}

	if s.FolderPath == nil {
		invalidParams.Add(aws.NewErrParamRequired("FolderPath"))
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

type GetFolderOutput struct {
	_ struct{} `type:"structure"`

	// The full commit ID used as a reference for which version of the folder content
	// is returned.
	//
	// CommitId is a required field
	CommitId *string `locationName:"commitId" type:"string" required:"true"`

	// The list of files that exist in the specified folder, if any.
	Files []File `locationName:"files" type:"list"`

	// The fully-qualified path of the folder whose contents are returned.
	//
	// FolderPath is a required field
	FolderPath *string `locationName:"folderPath" type:"string" required:"true"`

	// The list of folders that exist beneath the specified folder, if any.
	SubFolders []Folder `locationName:"subFolders" type:"list"`

	// The list of submodules that exist in the specified folder, if any.
	SubModules []SubModule `locationName:"subModules" type:"list"`

	// The list of symbolic links to other files and folders that exist in the specified
	// folder, if any.
	SymbolicLinks []SymbolicLink `locationName:"symbolicLinks" type:"list"`

	// The full SHA-1 pointer of the tree information for the commit that contains
	// the folder.
	TreeId *string `locationName:"treeId" type:"string"`
}

// String returns the string representation
func (s GetFolderOutput) String() string {
	return awsutil.Prettify(s)
}
