package change

import (
	"context"
)

type Action rune

const (
	NoOp   Action = 0
	Create Action = '+'
	Update Action = '~'
	Delete Action = '-'
)

type Change interface {
	String() string
	Apply(ctx context.Context, repoOwner string, repoName string) error
}

type ChangeSet []Change
