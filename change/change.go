package change

import (
	"context"
	"errors"

	"github.com/shotat/ghrc/status"
)

type Action rune

const (
	NoOp   Action = 0
	Create Action = '+'
	Update Action = '~'
	Delete Action = '-'
)

type ReposChange struct {
	Action Action
	Before *status.Repo
	After  *status.Repo
}

type LabelChange struct {
	Action Action
	Before *status.Label
	After  *status.Label
}

func (c *LabelChange) Apply(ctx context.Context) error {
	if c.Before == nil && c.After == nil {
		return errors.New("unexpected error")
	}
	if c.Before == nil && c.After != nil {
		if err := c.After.Create(ctx); err != nil {
			return err
		}
	}
	if c.Before != nil && c.After == nil {
		if err := c.Before.Destroy(ctx); err != nil {
			return err
		}
	}
	if c.Before != nil && c.After != nil {
		if err := c.After.Change(ctx); err != nil {
			return err
		}
	}

	return nil
}
