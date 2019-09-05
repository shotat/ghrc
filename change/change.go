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

type StringChange struct {
	Action Action
	Before *string
	After  *string
}

type ReposChange struct {
	DescriptionChange StringChange
	PrivateChange     StringChange
	LabelChanges      []LabelChange
}

type LabelChange struct {
	Action Action
	Before *status.Label
	After  *status.Label
}

func (p *LabelPatch) Apply(ctx context.Context) error {
	if p.Before == nil && p.After == nil {
		return errors.New("unexpected error")
	}
	if p.Before == nil && p.After != nil {
		if err := p.After.Create(ctx); err != nil {
			return err
		}
	}
	if p.Before != nil && p.After == nil {
		if err := p.Before.Destroy(ctx); err != nil {
			return err
		}
	}
	if p.Before != nil && p.After != nil {
		if err := p.After.Change(ctx); err != nil {
			return err
		}
	}

	return nil
}
