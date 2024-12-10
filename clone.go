package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// CloneRepoAndGetBlobSizes clones a GitHub repo and returns the sizes of all blobs.
func CloneRepoAndGetBlobSizes(repoURL string) ([]int64, error) {
	// Create a temporary directory for cloning the repo
	tempDir, err := os.MkdirTemp("", "repo-clone-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Clone the repository
	_, err = git.PlainClone(tempDir, true, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to clone repository: %w", err)
	}

	// Open the cloned repository
	repo, err := git.PlainOpen(tempDir)
	if err != nil {
		return nil, fmt.Errorf("failed to open repository: %w", err)
	}

	// Get the HEAD reference
	ref, err := repo.Head()
	if err != nil {
		return nil, fmt.Errorf("failed to get HEAD reference: %w", err)
	}

	// Get the commit object
	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get commit object: %w", err)
	}

	// Get the tree from the commit
	tree, err := commit.Tree()
	if err != nil {
		return nil, fmt.Errorf("failed to get tree: %w", err)
	}

	// Iterate through the blobs and collect their sizes
	var blobSizes []int64
	err = tree.Files().ForEach(func(file *object.File) error {
		blobSizes = append(blobSizes, file.Size)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to iterate blobs: %w", err)
	}

	return blobSizes, nil
}

func main() {
	repoURL := "https://github.com/user/repo.git" // Replace with the actual repo URL

	blobSizes, err := CloneRepoAndGetBlobSizes(repoURL)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Blob sizes:", blobSizes)
}
