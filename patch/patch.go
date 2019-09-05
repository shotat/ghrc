package patch

import (
	"context"
	"errors"

	"github.com/shotat/ghrc/status"
)

type RepositoryPatch struct {
	Before *status.RepositoryStatus
	After  *status.RepositoryStatus
}

type LabelPatch struct {
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
