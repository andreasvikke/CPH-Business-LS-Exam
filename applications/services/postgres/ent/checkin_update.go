// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/postgres/ent/checkin"
	"github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/postgres/ent/predicate"
)

// CheckInUpdate is the builder for updating CheckIn entities.
type CheckInUpdate struct {
	config
	hooks    []Hook
	mutation *CheckInMutation
}

// Where appends a list predicates to the CheckInUpdate builder.
func (ciu *CheckInUpdate) Where(ps ...predicate.CheckIn) *CheckInUpdate {
	ciu.mutation.Where(ps...)
	return ciu
}

// SetAttendanceCode sets the "attendanceCode" field.
func (ciu *CheckInUpdate) SetAttendanceCode(i int64) *CheckInUpdate {
	ciu.mutation.ResetAttendanceCode()
	ciu.mutation.SetAttendanceCode(i)
	return ciu
}

// AddAttendanceCode adds i to the "attendanceCode" field.
func (ciu *CheckInUpdate) AddAttendanceCode(i int64) *CheckInUpdate {
	ciu.mutation.AddAttendanceCode(i)
	return ciu
}

// SetStudentId sets the "studentId" field.
func (ciu *CheckInUpdate) SetStudentId(s string) *CheckInUpdate {
	ciu.mutation.SetStudentId(s)
	return ciu
}

// SetStatus sets the "status" field.
func (ciu *CheckInUpdate) SetStatus(c checkin.Status) *CheckInUpdate {
	ciu.mutation.SetStatus(c)
	return ciu
}

// SetCheckinTime sets the "checkinTime" field.
func (ciu *CheckInUpdate) SetCheckinTime(i int64) *CheckInUpdate {
	ciu.mutation.ResetCheckinTime()
	ciu.mutation.SetCheckinTime(i)
	return ciu
}

// AddCheckinTime adds i to the "checkinTime" field.
func (ciu *CheckInUpdate) AddCheckinTime(i int64) *CheckInUpdate {
	ciu.mutation.AddCheckinTime(i)
	return ciu
}

// Mutation returns the CheckInMutation object of the builder.
func (ciu *CheckInUpdate) Mutation() *CheckInMutation {
	return ciu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ciu *CheckInUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ciu.hooks) == 0 {
		if err = ciu.check(); err != nil {
			return 0, err
		}
		affected, err = ciu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CheckInMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ciu.check(); err != nil {
				return 0, err
			}
			ciu.mutation = mutation
			affected, err = ciu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ciu.hooks) - 1; i >= 0; i-- {
			if ciu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ciu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ciu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ciu *CheckInUpdate) SaveX(ctx context.Context) int {
	affected, err := ciu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ciu *CheckInUpdate) Exec(ctx context.Context) error {
	_, err := ciu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciu *CheckInUpdate) ExecX(ctx context.Context) {
	if err := ciu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ciu *CheckInUpdate) check() error {
	if v, ok := ciu.mutation.AttendanceCode(); ok {
		if err := checkin.AttendanceCodeValidator(v); err != nil {
			return &ValidationError{Name: "attendanceCode", err: fmt.Errorf("ent: validator failed for field \"attendanceCode\": %w", err)}
		}
	}
	if v, ok := ciu.mutation.Status(); ok {
		if err := checkin.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (ciu *CheckInUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   checkin.Table,
			Columns: checkin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: checkin.FieldID,
			},
		},
	}
	if ps := ciu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciu.mutation.AttendanceCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: checkin.FieldAttendanceCode,
		})
	}
	if value, ok := ciu.mutation.AddedAttendanceCode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: checkin.FieldAttendanceCode,
		})
	}
	if value, ok := ciu.mutation.StudentId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: checkin.FieldStudentId,
		})
	}
	if value, ok := ciu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: checkin.FieldStatus,
		})
	}
	if value, ok := ciu.mutation.CheckinTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: checkin.FieldCheckinTime,
		})
	}
	if value, ok := ciu.mutation.AddedCheckinTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: checkin.FieldCheckinTime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ciu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{checkin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CheckInUpdateOne is the builder for updating a single CheckIn entity.
type CheckInUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CheckInMutation
}

// SetAttendanceCode sets the "attendanceCode" field.
func (ciuo *CheckInUpdateOne) SetAttendanceCode(i int64) *CheckInUpdateOne {
	ciuo.mutation.ResetAttendanceCode()
	ciuo.mutation.SetAttendanceCode(i)
	return ciuo
}

// AddAttendanceCode adds i to the "attendanceCode" field.
func (ciuo *CheckInUpdateOne) AddAttendanceCode(i int64) *CheckInUpdateOne {
	ciuo.mutation.AddAttendanceCode(i)
	return ciuo
}

// SetStudentId sets the "studentId" field.
func (ciuo *CheckInUpdateOne) SetStudentId(s string) *CheckInUpdateOne {
	ciuo.mutation.SetStudentId(s)
	return ciuo
}

// SetStatus sets the "status" field.
func (ciuo *CheckInUpdateOne) SetStatus(c checkin.Status) *CheckInUpdateOne {
	ciuo.mutation.SetStatus(c)
	return ciuo
}

// SetCheckinTime sets the "checkinTime" field.
func (ciuo *CheckInUpdateOne) SetCheckinTime(i int64) *CheckInUpdateOne {
	ciuo.mutation.ResetCheckinTime()
	ciuo.mutation.SetCheckinTime(i)
	return ciuo
}

// AddCheckinTime adds i to the "checkinTime" field.
func (ciuo *CheckInUpdateOne) AddCheckinTime(i int64) *CheckInUpdateOne {
	ciuo.mutation.AddCheckinTime(i)
	return ciuo
}

// Mutation returns the CheckInMutation object of the builder.
func (ciuo *CheckInUpdateOne) Mutation() *CheckInMutation {
	return ciuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ciuo *CheckInUpdateOne) Select(field string, fields ...string) *CheckInUpdateOne {
	ciuo.fields = append([]string{field}, fields...)
	return ciuo
}

// Save executes the query and returns the updated CheckIn entity.
func (ciuo *CheckInUpdateOne) Save(ctx context.Context) (*CheckIn, error) {
	var (
		err  error
		node *CheckIn
	)
	if len(ciuo.hooks) == 0 {
		if err = ciuo.check(); err != nil {
			return nil, err
		}
		node, err = ciuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CheckInMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ciuo.check(); err != nil {
				return nil, err
			}
			ciuo.mutation = mutation
			node, err = ciuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ciuo.hooks) - 1; i >= 0; i-- {
			if ciuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ciuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ciuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ciuo *CheckInUpdateOne) SaveX(ctx context.Context) *CheckIn {
	node, err := ciuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ciuo *CheckInUpdateOne) Exec(ctx context.Context) error {
	_, err := ciuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciuo *CheckInUpdateOne) ExecX(ctx context.Context) {
	if err := ciuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ciuo *CheckInUpdateOne) check() error {
	if v, ok := ciuo.mutation.AttendanceCode(); ok {
		if err := checkin.AttendanceCodeValidator(v); err != nil {
			return &ValidationError{Name: "attendanceCode", err: fmt.Errorf("ent: validator failed for field \"attendanceCode\": %w", err)}
		}
	}
	if v, ok := ciuo.mutation.Status(); ok {
		if err := checkin.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (ciuo *CheckInUpdateOne) sqlSave(ctx context.Context) (_node *CheckIn, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   checkin.Table,
			Columns: checkin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: checkin.FieldID,
			},
		},
	}
	id, ok := ciuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CheckIn.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ciuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, checkin.FieldID)
		for _, f := range fields {
			if !checkin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != checkin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ciuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciuo.mutation.AttendanceCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: checkin.FieldAttendanceCode,
		})
	}
	if value, ok := ciuo.mutation.AddedAttendanceCode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: checkin.FieldAttendanceCode,
		})
	}
	if value, ok := ciuo.mutation.StudentId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: checkin.FieldStudentId,
		})
	}
	if value, ok := ciuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: checkin.FieldStatus,
		})
	}
	if value, ok := ciuo.mutation.CheckinTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: checkin.FieldCheckinTime,
		})
	}
	if value, ok := ciuo.mutation.AddedCheckinTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: checkin.FieldCheckinTime,
		})
	}
	_node = &CheckIn{config: ciuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ciuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{checkin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
