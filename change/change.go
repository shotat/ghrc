package change

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/shotat/ghrc/config"
	"github.com/shotat/ghrc/state"
)

type Action rune

type Change interface {
	String() string
	Apply(ctx context.Context, repoOwner string, repoName string) error
}

type ChangeSet []Change

const (
	NoOp   Action = 0
	Create Action = '+'
	Update Action = '~'
	Delete Action = '-'
)

func CalculateChangeSet(ctx context.Context, c *config.Config) (ChangeSet, error) {
	repo, err := state.FindRepo(ctx, c.Metadata.Owner, c.Metadata.Name)
	if err != nil {
		return nil, err
	}
	labels, err := state.FindLabels(ctx, c.Metadata.Owner, c.Metadata.Name)
	if err != nil {
		return nil, err
	}

	changeSet := make(ChangeSet, 0)
	changeSet = append(changeSet, GetRepoChange(repo, c.Spec.Repo))
	for _, labelChange := range GetLabelsChangeSet(labels, c.Spec.Labels) {
		changeSet = append(changeSet, labelChange)
	}
	return changeSet, nil
}

// Plan shows the expected changes without changing actual states.
func Plan(ctx context.Context, c *config.Config) error {
	cs, err := CalculateChangeSet(ctx, c)
	if err != nil {
		return err
	}

	for _, ch := range cs {
		fmt.Println(ch)
	}
	return nil
}

// Apply changes the remote configurations based on this Config.
func Apply(ctx context.Context, c *config.Config) error {
	cs, err := CalculateChangeSet(ctx, c)
	if err != nil {
		return err
	}

	for _, ch := range cs {
		if err := ch.Apply(ctx, c.Metadata.Owner, c.Metadata.Name); err != nil {
			return err
		}
	}
	return nil
}

type ReposChange struct {
	Action Action
	Before *state.Repo
	After  *state.Repo
}

func (c *ReposChange) Apply(ctx context.Context, repoOwner string, repoName string) error {
	return c.After.Update(ctx, repoOwner, repoName)
}

func (c *ReposChange) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("%s Repo\n", string(c.Action)))
	switch c.Action {
	case Update:
		buf.WriteString(fmt.Sprintf("\tdescription\t%v\n", *c.After.Description))
		buf.WriteString(fmt.Sprintf("\thomepage\t%v\n", *c.After.Homepage))
		buf.WriteString(fmt.Sprintf("\tprivate\t%v\n", *c.After.Private))
		buf.WriteString(fmt.Sprintf("\tallowSquashMerge\t%v\n", *c.After.AllowSquashMerge))
		buf.WriteString(fmt.Sprintf("\tallowMergeCommit\t%v\n", *c.After.AllowMergeCommit))
		buf.WriteString(fmt.Sprintf("\tallowRebaseMerge\t%v\n", *c.After.AllowRebaseMerge))
		buf.WriteString(fmt.Sprintf("\ttopics\t%v\n", c.After.Topics))
	}
	return buf.String()
}

type LabelChange struct {
	Action Action
	Before *state.Label
	After  *state.Label
}

func (c *LabelChange) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("%s Label\n", string(c.Action)))
	switch c.Action {
	case Create:
		buf.WriteString(fmt.Sprintf("\tname\t%v\n", c.After.Name))
		buf.WriteString(fmt.Sprintf("\tcolor\t%v\n", c.After.Color))
	case Update:
		buf.WriteString(fmt.Sprintf("\tname\t%v\n", c.After.Name))
		buf.WriteString(fmt.Sprintf("\tcolor\t%v\n", c.After.Color))
	case Delete:
		buf.WriteString(fmt.Sprintf("\tname\t%v\n", c.Before.Name))
		buf.WriteString(fmt.Sprintf("\tcolor\t%v\n", c.Before.Color))
	}
	return buf.String()
}

func (c *LabelChange) Apply(ctx context.Context, repoOwner string, repoName string) error {
	if c.Before == nil && c.After == nil {
		return errors.New("unexpected error")
	}
	if c.Before == nil && c.After != nil {
		if err := c.After.Create(ctx, repoOwner, repoName); err != nil {
			return err
		}
	}
	if c.Before != nil && c.After == nil {
		if err := c.Before.Destroy(ctx, repoOwner, repoName); err != nil {
			return err
		}
	}
	if c.Before != nil && c.After != nil {
		if err := c.After.Update(ctx, repoOwner, repoName); err != nil {
			return err
		}
	}

	return nil
}
