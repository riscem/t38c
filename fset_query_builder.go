package t38c

import (
	"context"
)

// FSetQueryBuilder struct
type FSetQueryBuilder struct {
	client   tile38Client
	key      string
	objectID string
	fields   []field
	xx       bool
}

func newFSetQueryBuilder(client tile38Client, key, objectID string) FSetQueryBuilder {
	return FSetQueryBuilder{
		client:   client,
		key:      key,
		objectID: objectID,
	}
}

func (query FSetQueryBuilder) toCmd() cmd {
	args := []string{query.key, query.objectID}
	if query.xx {
		args = append(args, "XX")
	}

	for _, field := range query.fields {
		args = append(args, field.Name, field.Value)
	}
	return newCmd("FSET", args...)
}

// Do cmd
func (query FSetQueryBuilder) Do(ctx context.Context) error {
	cmd := query.toCmd()
	return query.client.jExecute(ctx, nil, cmd.Name, cmd.Args...)
}

// Field sets the object field
func (query FSetQueryBuilder) Field(name string, value string) FSetQueryBuilder {
	query.fields = append(query.fields, field{name, value})
	return query
}

// IfExists only set the object if it already exist
func (query FSetQueryBuilder) IfExists() FSetQueryBuilder {
	query.xx = true
	return query
}
