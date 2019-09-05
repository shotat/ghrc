package status

import "context"

type Creatable interface {
	Create(ctx context.Context) error
}
type Changeable interface {
	Change(ctx context.Context) error
}
type Destroyable interface {
	Destroy(ctx context.Context) error
}

type Patch struct {
	Changes      []Changeable
	Creations    []Creatable
	Destructions []Destroyable
}

func (p *Patch) Apply(ctx context.Context) error {
	for _, a := range p.Destructions {
		if err := a.Destroy(ctx); err != nil {
			return err
		}
	}
	for _, a := range p.Changes {
		if err := a.Change(ctx); err != nil {
			return err
		}
	}
	for _, a := range p.Creations {
		if err := a.Create(ctx); err != nil {
			return err
		}
	}
	return nil
}
