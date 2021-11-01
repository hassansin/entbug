// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"entgo.io/bug/ent/order"
	"entgo.io/bug/ent/predicate"
	"entgo.io/bug/ent/user"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeOrder = "Order"
	TypeUser  = "User"
)

// OrderMutation represents an operation that mutates the Order nodes in the graph.
type OrderMutation struct {
	config
	op            Op
	typ           string
	id            *uint
	clearedFields map[string]struct{}
	user          *int
	cleareduser   bool
	done          bool
	oldValue      func(context.Context) (*Order, error)
	predicates    []predicate.Order
}

var _ ent.Mutation = (*OrderMutation)(nil)

// orderOption allows management of the mutation configuration using functional options.
type orderOption func(*OrderMutation)

// newOrderMutation creates new mutation for the Order entity.
func newOrderMutation(c config, op Op, opts ...orderOption) *OrderMutation {
	m := &OrderMutation{
		config:        c,
		op:            op,
		typ:           TypeOrder,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withOrderID sets the ID field of the mutation.
func withOrderID(id uint) orderOption {
	return func(m *OrderMutation) {
		var (
			err   error
			once  sync.Once
			value *Order
		)
		m.oldValue = func(ctx context.Context) (*Order, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Order.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withOrder sets the old Order of the mutation.
func withOrder(node *Order) orderOption {
	return func(m *OrderMutation) {
		m.oldValue = func(context.Context) (*Order, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m OrderMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m OrderMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Order entities.
func (m *OrderMutation) SetID(id uint) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *OrderMutation) ID() (id uint, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetUserID sets the "user" edge to the User entity by id.
func (m *OrderMutation) SetUserID(id int) {
	m.user = &id
}

// ClearUser clears the "user" edge to the User entity.
func (m *OrderMutation) ClearUser() {
	m.cleareduser = true
}

// UserCleared reports if the "user" edge to the User entity was cleared.
func (m *OrderMutation) UserCleared() bool {
	return m.cleareduser
}

// UserID returns the "user" edge ID in the mutation.
func (m *OrderMutation) UserID() (id int, exists bool) {
	if m.user != nil {
		return *m.user, true
	}
	return
}

// UserIDs returns the "user" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// UserID instead. It exists only for internal usage by the builders.
func (m *OrderMutation) UserIDs() (ids []int) {
	if id := m.user; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetUser resets all changes to the "user" edge.
func (m *OrderMutation) ResetUser() {
	m.user = nil
	m.cleareduser = false
}

// Where appends a list predicates to the OrderMutation builder.
func (m *OrderMutation) Where(ps ...predicate.Order) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *OrderMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Order).
func (m *OrderMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *OrderMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *OrderMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *OrderMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown Order field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OrderMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Order field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *OrderMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *OrderMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OrderMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown Order numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *OrderMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *OrderMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *OrderMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Order nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *OrderMutation) ResetField(name string) error {
	return fmt.Errorf("unknown Order field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *OrderMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.user != nil {
		edges = append(edges, order.EdgeUser)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *OrderMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case order.EdgeUser:
		if id := m.user; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *OrderMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *OrderMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *OrderMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareduser {
		edges = append(edges, order.EdgeUser)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *OrderMutation) EdgeCleared(name string) bool {
	switch name {
	case order.EdgeUser:
		return m.cleareduser
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *OrderMutation) ClearEdge(name string) error {
	switch name {
	case order.EdgeUser:
		m.ClearUser()
		return nil
	}
	return fmt.Errorf("unknown Order unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *OrderMutation) ResetEdge(name string) error {
	switch name {
	case order.EdgeUser:
		m.ResetUser()
		return nil
	}
	return fmt.Errorf("unknown Order edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	clearedFields map[string]struct{}
	orders        map[uint]struct{}
	removedorders map[uint]struct{}
	clearedorders bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id int) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// AddOrderIDs adds the "orders" edge to the Order entity by ids.
func (m *UserMutation) AddOrderIDs(ids ...uint) {
	if m.orders == nil {
		m.orders = make(map[uint]struct{})
	}
	for i := range ids {
		m.orders[ids[i]] = struct{}{}
	}
}

// ClearOrders clears the "orders" edge to the Order entity.
func (m *UserMutation) ClearOrders() {
	m.clearedorders = true
}

// OrdersCleared reports if the "orders" edge to the Order entity was cleared.
func (m *UserMutation) OrdersCleared() bool {
	return m.clearedorders
}

// RemoveOrderIDs removes the "orders" edge to the Order entity by IDs.
func (m *UserMutation) RemoveOrderIDs(ids ...uint) {
	if m.removedorders == nil {
		m.removedorders = make(map[uint]struct{})
	}
	for i := range ids {
		delete(m.orders, ids[i])
		m.removedorders[ids[i]] = struct{}{}
	}
}

// RemovedOrders returns the removed IDs of the "orders" edge to the Order entity.
func (m *UserMutation) RemovedOrdersIDs() (ids []uint) {
	for id := range m.removedorders {
		ids = append(ids, id)
	}
	return
}

// OrdersIDs returns the "orders" edge IDs in the mutation.
func (m *UserMutation) OrdersIDs() (ids []uint) {
	for id := range m.orders {
		ids = append(ids, id)
	}
	return
}

// ResetOrders resets all changes to the "orders" edge.
func (m *UserMutation) ResetOrders() {
	m.orders = nil
	m.clearedorders = false
	m.removedorders = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.orders != nil {
		edges = append(edges, user.EdgeOrders)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeOrders:
		ids := make([]ent.Value, 0, len(m.orders))
		for id := range m.orders {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedorders != nil {
		edges = append(edges, user.EdgeOrders)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeOrders:
		ids := make([]ent.Value, 0, len(m.removedorders))
		for id := range m.removedorders {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedorders {
		edges = append(edges, user.EdgeOrders)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeOrders:
		return m.clearedorders
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeOrders:
		m.ResetOrders()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
