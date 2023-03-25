// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/zema1/watchvuln/ent/predicate"
	"github.com/zema1/watchvuln/ent/vulninformation"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeVulnInformation = "VulnInformation"
)

// VulnInformationMutation represents an operation that mutates the VulnInformation nodes in the graph.
type VulnInformationMutation struct {
	config
	op               Op
	typ              string
	id               *int
	key              *string
	title            *string
	description      *string
	severity         *string
	cve              *string
	disclosure       *string
	solutions        *string
	references       *[]string
	appendreferences []string
	tags             *[]string
	appendtags       []string
	from             *string
	clearedFields    map[string]struct{}
	done             bool
	oldValue         func(context.Context) (*VulnInformation, error)
	predicates       []predicate.VulnInformation
}

var _ ent.Mutation = (*VulnInformationMutation)(nil)

// vulninformationOption allows management of the mutation configuration using functional options.
type vulninformationOption func(*VulnInformationMutation)

// newVulnInformationMutation creates new mutation for the VulnInformation entity.
func newVulnInformationMutation(c config, op Op, opts ...vulninformationOption) *VulnInformationMutation {
	m := &VulnInformationMutation{
		config:        c,
		op:            op,
		typ:           TypeVulnInformation,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withVulnInformationID sets the ID field of the mutation.
func withVulnInformationID(id int) vulninformationOption {
	return func(m *VulnInformationMutation) {
		var (
			err   error
			once  sync.Once
			value *VulnInformation
		)
		m.oldValue = func(ctx context.Context) (*VulnInformation, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().VulnInformation.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withVulnInformation sets the old VulnInformation of the mutation.
func withVulnInformation(node *VulnInformation) vulninformationOption {
	return func(m *VulnInformationMutation) {
		m.oldValue = func(context.Context) (*VulnInformation, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m VulnInformationMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m VulnInformationMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *VulnInformationMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *VulnInformationMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().VulnInformation.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetKey sets the "key" field.
func (m *VulnInformationMutation) SetKey(s string) {
	m.key = &s
}

// Key returns the value of the "key" field in the mutation.
func (m *VulnInformationMutation) Key() (r string, exists bool) {
	v := m.key
	if v == nil {
		return
	}
	return *v, true
}

// OldKey returns the old "key" field's value of the VulnInformation entity.
// If the VulnInformation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VulnInformationMutation) OldKey(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldKey is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldKey requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldKey: %w", err)
	}
	return oldValue.Key, nil
}

// ResetKey resets all changes to the "key" field.
func (m *VulnInformationMutation) ResetKey() {
	m.key = nil
}

// SetTitle sets the "title" field.
func (m *VulnInformationMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *VulnInformationMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the VulnInformation entity.
// If the VulnInformation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VulnInformationMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *VulnInformationMutation) ResetTitle() {
	m.title = nil
}

// SetDescription sets the "description" field.
func (m *VulnInformationMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *VulnInformationMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the VulnInformation entity.
// If the VulnInformation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VulnInformationMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *VulnInformationMutation) ResetDescription() {
	m.description = nil
}

// SetSeverity sets the "severity" field.
func (m *VulnInformationMutation) SetSeverity(s string) {
	m.severity = &s
}

// Severity returns the value of the "severity" field in the mutation.
func (m *VulnInformationMutation) Severity() (r string, exists bool) {
	v := m.severity
	if v == nil {
		return
	}
	return *v, true
}

// OldSeverity returns the old "severity" field's value of the VulnInformation entity.
// If the VulnInformation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VulnInformationMutation) OldSeverity(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSeverity is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSeverity requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSeverity: %w", err)
	}
	return oldValue.Severity, nil
}

// ResetSeverity resets all changes to the "severity" field.
func (m *VulnInformationMutation) ResetSeverity() {
	m.severity = nil
}

// SetCve sets the "cve" field.
func (m *VulnInformationMutation) SetCve(s string) {
	m.cve = &s
}

// Cve returns the value of the "cve" field in the mutation.
func (m *VulnInformationMutation) Cve() (r string, exists bool) {
	v := m.cve
	if v == nil {
		return
	}
	return *v, true
}

// OldCve returns the old "cve" field's value of the VulnInformation entity.
// If the VulnInformation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VulnInformationMutation) OldCve(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCve is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCve requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCve: %w", err)
	}
	return oldValue.Cve, nil
}

// ResetCve resets all changes to the "cve" field.
func (m *VulnInformationMutation) ResetCve() {
	m.cve = nil
}

// SetDisclosure sets the "disclosure" field.
func (m *VulnInformationMutation) SetDisclosure(s string) {
	m.disclosure = &s
}

// Disclosure returns the value of the "disclosure" field in the mutation.
func (m *VulnInformationMutation) Disclosure() (r string, exists bool) {
	v := m.disclosure
	if v == nil {
		return
	}
	return *v, true
}

// OldDisclosure returns the old "disclosure" field's value of the VulnInformation entity.
// If the VulnInformation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VulnInformationMutation) OldDisclosure(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDisclosure is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDisclosure requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDisclosure: %w", err)
	}
	return oldValue.Disclosure, nil
}

// ResetDisclosure resets all changes to the "disclosure" field.
func (m *VulnInformationMutation) ResetDisclosure() {
	m.disclosure = nil
}

// SetSolutions sets the "solutions" field.
func (m *VulnInformationMutation) SetSolutions(s string) {
	m.solutions = &s
}

// Solutions returns the value of the "solutions" field in the mutation.
func (m *VulnInformationMutation) Solutions() (r string, exists bool) {
	v := m.solutions
	if v == nil {
		return
	}
	return *v, true
}

// OldSolutions returns the old "solutions" field's value of the VulnInformation entity.
// If the VulnInformation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VulnInformationMutation) OldSolutions(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSolutions is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSolutions requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSolutions: %w", err)
	}
	return oldValue.Solutions, nil
}

// ResetSolutions resets all changes to the "solutions" field.
func (m *VulnInformationMutation) ResetSolutions() {
	m.solutions = nil
}

// SetReferences sets the "references" field.
func (m *VulnInformationMutation) SetReferences(s []string) {
	m.references = &s
	m.appendreferences = nil
}

// References returns the value of the "references" field in the mutation.
func (m *VulnInformationMutation) References() (r []string, exists bool) {
	v := m.references
	if v == nil {
		return
	}
	return *v, true
}

// OldReferences returns the old "references" field's value of the VulnInformation entity.
// If the VulnInformation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VulnInformationMutation) OldReferences(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldReferences is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldReferences requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldReferences: %w", err)
	}
	return oldValue.References, nil
}

// AppendReferences adds s to the "references" field.
func (m *VulnInformationMutation) AppendReferences(s []string) {
	m.appendreferences = append(m.appendreferences, s...)
}

// AppendedReferences returns the list of values that were appended to the "references" field in this mutation.
func (m *VulnInformationMutation) AppendedReferences() ([]string, bool) {
	if len(m.appendreferences) == 0 {
		return nil, false
	}
	return m.appendreferences, true
}

// ClearReferences clears the value of the "references" field.
func (m *VulnInformationMutation) ClearReferences() {
	m.references = nil
	m.appendreferences = nil
	m.clearedFields[vulninformation.FieldReferences] = struct{}{}
}

// ReferencesCleared returns if the "references" field was cleared in this mutation.
func (m *VulnInformationMutation) ReferencesCleared() bool {
	_, ok := m.clearedFields[vulninformation.FieldReferences]
	return ok
}

// ResetReferences resets all changes to the "references" field.
func (m *VulnInformationMutation) ResetReferences() {
	m.references = nil
	m.appendreferences = nil
	delete(m.clearedFields, vulninformation.FieldReferences)
}

// SetTags sets the "tags" field.
func (m *VulnInformationMutation) SetTags(s []string) {
	m.tags = &s
	m.appendtags = nil
}

// Tags returns the value of the "tags" field in the mutation.
func (m *VulnInformationMutation) Tags() (r []string, exists bool) {
	v := m.tags
	if v == nil {
		return
	}
	return *v, true
}

// OldTags returns the old "tags" field's value of the VulnInformation entity.
// If the VulnInformation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VulnInformationMutation) OldTags(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTags is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTags requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTags: %w", err)
	}
	return oldValue.Tags, nil
}

// AppendTags adds s to the "tags" field.
func (m *VulnInformationMutation) AppendTags(s []string) {
	m.appendtags = append(m.appendtags, s...)
}

// AppendedTags returns the list of values that were appended to the "tags" field in this mutation.
func (m *VulnInformationMutation) AppendedTags() ([]string, bool) {
	if len(m.appendtags) == 0 {
		return nil, false
	}
	return m.appendtags, true
}

// ClearTags clears the value of the "tags" field.
func (m *VulnInformationMutation) ClearTags() {
	m.tags = nil
	m.appendtags = nil
	m.clearedFields[vulninformation.FieldTags] = struct{}{}
}

// TagsCleared returns if the "tags" field was cleared in this mutation.
func (m *VulnInformationMutation) TagsCleared() bool {
	_, ok := m.clearedFields[vulninformation.FieldTags]
	return ok
}

// ResetTags resets all changes to the "tags" field.
func (m *VulnInformationMutation) ResetTags() {
	m.tags = nil
	m.appendtags = nil
	delete(m.clearedFields, vulninformation.FieldTags)
}

// SetFrom sets the "from" field.
func (m *VulnInformationMutation) SetFrom(s string) {
	m.from = &s
}

// From returns the value of the "from" field in the mutation.
func (m *VulnInformationMutation) From() (r string, exists bool) {
	v := m.from
	if v == nil {
		return
	}
	return *v, true
}

// OldFrom returns the old "from" field's value of the VulnInformation entity.
// If the VulnInformation object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VulnInformationMutation) OldFrom(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFrom is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFrom requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFrom: %w", err)
	}
	return oldValue.From, nil
}

// ResetFrom resets all changes to the "from" field.
func (m *VulnInformationMutation) ResetFrom() {
	m.from = nil
}

// Where appends a list predicates to the VulnInformationMutation builder.
func (m *VulnInformationMutation) Where(ps ...predicate.VulnInformation) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the VulnInformationMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *VulnInformationMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.VulnInformation, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *VulnInformationMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *VulnInformationMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (VulnInformation).
func (m *VulnInformationMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *VulnInformationMutation) Fields() []string {
	fields := make([]string, 0, 10)
	if m.key != nil {
		fields = append(fields, vulninformation.FieldKey)
	}
	if m.title != nil {
		fields = append(fields, vulninformation.FieldTitle)
	}
	if m.description != nil {
		fields = append(fields, vulninformation.FieldDescription)
	}
	if m.severity != nil {
		fields = append(fields, vulninformation.FieldSeverity)
	}
	if m.cve != nil {
		fields = append(fields, vulninformation.FieldCve)
	}
	if m.disclosure != nil {
		fields = append(fields, vulninformation.FieldDisclosure)
	}
	if m.solutions != nil {
		fields = append(fields, vulninformation.FieldSolutions)
	}
	if m.references != nil {
		fields = append(fields, vulninformation.FieldReferences)
	}
	if m.tags != nil {
		fields = append(fields, vulninformation.FieldTags)
	}
	if m.from != nil {
		fields = append(fields, vulninformation.FieldFrom)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *VulnInformationMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case vulninformation.FieldKey:
		return m.Key()
	case vulninformation.FieldTitle:
		return m.Title()
	case vulninformation.FieldDescription:
		return m.Description()
	case vulninformation.FieldSeverity:
		return m.Severity()
	case vulninformation.FieldCve:
		return m.Cve()
	case vulninformation.FieldDisclosure:
		return m.Disclosure()
	case vulninformation.FieldSolutions:
		return m.Solutions()
	case vulninformation.FieldReferences:
		return m.References()
	case vulninformation.FieldTags:
		return m.Tags()
	case vulninformation.FieldFrom:
		return m.From()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *VulnInformationMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case vulninformation.FieldKey:
		return m.OldKey(ctx)
	case vulninformation.FieldTitle:
		return m.OldTitle(ctx)
	case vulninformation.FieldDescription:
		return m.OldDescription(ctx)
	case vulninformation.FieldSeverity:
		return m.OldSeverity(ctx)
	case vulninformation.FieldCve:
		return m.OldCve(ctx)
	case vulninformation.FieldDisclosure:
		return m.OldDisclosure(ctx)
	case vulninformation.FieldSolutions:
		return m.OldSolutions(ctx)
	case vulninformation.FieldReferences:
		return m.OldReferences(ctx)
	case vulninformation.FieldTags:
		return m.OldTags(ctx)
	case vulninformation.FieldFrom:
		return m.OldFrom(ctx)
	}
	return nil, fmt.Errorf("unknown VulnInformation field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *VulnInformationMutation) SetField(name string, value ent.Value) error {
	switch name {
	case vulninformation.FieldKey:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetKey(v)
		return nil
	case vulninformation.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case vulninformation.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case vulninformation.FieldSeverity:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSeverity(v)
		return nil
	case vulninformation.FieldCve:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCve(v)
		return nil
	case vulninformation.FieldDisclosure:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDisclosure(v)
		return nil
	case vulninformation.FieldSolutions:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSolutions(v)
		return nil
	case vulninformation.FieldReferences:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetReferences(v)
		return nil
	case vulninformation.FieldTags:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTags(v)
		return nil
	case vulninformation.FieldFrom:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFrom(v)
		return nil
	}
	return fmt.Errorf("unknown VulnInformation field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *VulnInformationMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *VulnInformationMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *VulnInformationMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown VulnInformation numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *VulnInformationMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(vulninformation.FieldReferences) {
		fields = append(fields, vulninformation.FieldReferences)
	}
	if m.FieldCleared(vulninformation.FieldTags) {
		fields = append(fields, vulninformation.FieldTags)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *VulnInformationMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *VulnInformationMutation) ClearField(name string) error {
	switch name {
	case vulninformation.FieldReferences:
		m.ClearReferences()
		return nil
	case vulninformation.FieldTags:
		m.ClearTags()
		return nil
	}
	return fmt.Errorf("unknown VulnInformation nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *VulnInformationMutation) ResetField(name string) error {
	switch name {
	case vulninformation.FieldKey:
		m.ResetKey()
		return nil
	case vulninformation.FieldTitle:
		m.ResetTitle()
		return nil
	case vulninformation.FieldDescription:
		m.ResetDescription()
		return nil
	case vulninformation.FieldSeverity:
		m.ResetSeverity()
		return nil
	case vulninformation.FieldCve:
		m.ResetCve()
		return nil
	case vulninformation.FieldDisclosure:
		m.ResetDisclosure()
		return nil
	case vulninformation.FieldSolutions:
		m.ResetSolutions()
		return nil
	case vulninformation.FieldReferences:
		m.ResetReferences()
		return nil
	case vulninformation.FieldTags:
		m.ResetTags()
		return nil
	case vulninformation.FieldFrom:
		m.ResetFrom()
		return nil
	}
	return fmt.Errorf("unknown VulnInformation field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *VulnInformationMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *VulnInformationMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *VulnInformationMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *VulnInformationMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *VulnInformationMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *VulnInformationMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *VulnInformationMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown VulnInformation unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *VulnInformationMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown VulnInformation edge %s", name)
}
