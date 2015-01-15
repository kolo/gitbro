package main

import "github.com/libgit2/git2go"

type Repository struct {
	repo *git.Repository
}

func OpenRepository(path string) (*Repository, error) {
	repo, err := git.OpenRepository(path)
	if err != nil {
		return nil, err
	}

	return &Repository{
		repo: repo,
	}, nil
}

func (r *Repository) Branches() ([]string, error) {
	branches := []string{}

	it, err := r.repo.NewBranchIterator(git.BranchLocal)
	if err != nil {
		return nil, err
	}

	for {
		b, _, err := it.Next()
		if err != nil {
			if git.IsErrorCode(err, git.ErrIterOver) {
				break
			} else {
				return nil, err
			}
		}

		name, err := b.Name()
		if err != nil {
			return nil, err
		}

		branches = append(branches, name)
	}

	return branches, nil
}

func (r *Repository) Log(ref string) ([]string, error) {
	walk, err := r.repo.Walk()
	if err != nil {
		return nil, err
	}

	if err = walk.PushRef(ref); err != nil {
		return nil, err
	}

	commits := []string{}

	var i int
	walkFn := func(c *git.Commit) bool {
		if i > 9 {
			return false
		}

		i += 1
		commits = append(commits, c.Id().String())
		return true
	}

	if err = walk.Iterate(walkFn); err != nil {
		return nil, err
	}

	return commits, nil
}
