// Code generated by entc, DO NOT EDIT.

package checkin

import (
	"entgo.io/ent/dialect/sql"
	"github.com/AndreasVikke-School/CPH-Bussines-LS-Exam/applications/services/postgres/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AttendanceCode applies equality check predicate on the "attendanceCode" field. It's identical to AttendanceCodeEQ.
func AttendanceCode(v int64) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAttendanceCode), v))
	})
}

// StudentId applies equality check predicate on the "studentId" field. It's identical to StudentIdEQ.
func StudentId(v string) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStudentId), v))
	})
}

// CheckinTime applies equality check predicate on the "checkinTime" field. It's identical to CheckinTimeEQ.
func CheckinTime(v int64) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCheckinTime), v))
	})
}

// AttendanceCodeEQ applies the EQ predicate on the "attendanceCode" field.
func AttendanceCodeEQ(v int64) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAttendanceCode), v))
	})
}

// AttendanceCodeNEQ applies the NEQ predicate on the "attendanceCode" field.
func AttendanceCodeNEQ(v int64) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAttendanceCode), v))
	})
}

// AttendanceCodeIn applies the In predicate on the "attendanceCode" field.
func AttendanceCodeIn(vs ...int64) predicate.CheckIn {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckIn(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAttendanceCode), v...))
	})
}

// AttendanceCodeNotIn applies the NotIn predicate on the "attendanceCode" field.
func AttendanceCodeNotIn(vs ...int64) predicate.CheckIn {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckIn(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAttendanceCode), v...))
	})
}

// AttendanceCodeGT applies the GT predicate on the "attendanceCode" field.
func AttendanceCodeGT(v int64) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAttendanceCode), v))
	})
}

// AttendanceCodeGTE applies the GTE predicate on the "attendanceCode" field.
func AttendanceCodeGTE(v int64) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAttendanceCode), v))
	})
}

// AttendanceCodeLT applies the LT predicate on the "attendanceCode" field.
func AttendanceCodeLT(v int64) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAttendanceCode), v))
	})
}

// AttendanceCodeLTE applies the LTE predicate on the "attendanceCode" field.
func AttendanceCodeLTE(v int64) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAttendanceCode), v))
	})
}

// StudentIdEQ applies the EQ predicate on the "studentId" field.
func StudentIdEQ(v string) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStudentId), v))
	})
}

// StudentIdNEQ applies the NEQ predicate on the "studentId" field.
func StudentIdNEQ(v string) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStudentId), v))
	})
}

// StudentIdIn applies the In predicate on the "studentId" field.
func StudentIdIn(vs ...string) predicate.CheckIn {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckIn(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStudentId), v...))
	})
}

// StudentIdNotIn applies the NotIn predicate on the "studentId" field.
func StudentIdNotIn(vs ...string) predicate.CheckIn {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckIn(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStudentId), v...))
	})
}

// StudentIdGT applies the GT predicate on the "studentId" field.
func StudentIdGT(v string) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStudentId), v))
	})
}

// StudentIdGTE applies the GTE predicate on the "studentId" field.
func StudentIdGTE(v string) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStudentId), v))
	})
}

// StudentIdLT applies the LT predicate on the "studentId" field.
func StudentIdLT(v string) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStudentId), v))
	})
}

// StudentIdLTE applies the LTE predicate on the "studentId" field.
func StudentIdLTE(v string) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStudentId), v))
	})
}

// StudentIdContains applies the Contains predicate on the "studentId" field.
func StudentIdContains(v string) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStudentId), v))
	})
}

// StudentIdHasPrefix applies the HasPrefix predicate on the "studentId" field.
func StudentIdHasPrefix(v string) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStudentId), v))
	})
}

// StudentIdHasSuffix applies the HasSuffix predicate on the "studentId" field.
func StudentIdHasSuffix(v string) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStudentId), v))
	})
}

// StudentIdEqualFold applies the EqualFold predicate on the "studentId" field.
func StudentIdEqualFold(v string) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStudentId), v))
	})
}

// StudentIdContainsFold applies the ContainsFold predicate on the "studentId" field.
func StudentIdContainsFold(v string) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStudentId), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.CheckIn {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckIn(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.CheckIn {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckIn(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// CheckinTimeEQ applies the EQ predicate on the "checkinTime" field.
func CheckinTimeEQ(v int64) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCheckinTime), v))
	})
}

// CheckinTimeNEQ applies the NEQ predicate on the "checkinTime" field.
func CheckinTimeNEQ(v int64) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCheckinTime), v))
	})
}

// CheckinTimeIn applies the In predicate on the "checkinTime" field.
func CheckinTimeIn(vs ...int64) predicate.CheckIn {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckIn(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCheckinTime), v...))
	})
}

// CheckinTimeNotIn applies the NotIn predicate on the "checkinTime" field.
func CheckinTimeNotIn(vs ...int64) predicate.CheckIn {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CheckIn(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCheckinTime), v...))
	})
}

// CheckinTimeGT applies the GT predicate on the "checkinTime" field.
func CheckinTimeGT(v int64) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCheckinTime), v))
	})
}

// CheckinTimeGTE applies the GTE predicate on the "checkinTime" field.
func CheckinTimeGTE(v int64) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCheckinTime), v))
	})
}

// CheckinTimeLT applies the LT predicate on the "checkinTime" field.
func CheckinTimeLT(v int64) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCheckinTime), v))
	})
}

// CheckinTimeLTE applies the LTE predicate on the "checkinTime" field.
func CheckinTimeLTE(v int64) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCheckinTime), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CheckIn) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CheckIn) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CheckIn) predicate.CheckIn {
	return predicate.CheckIn(func(s *sql.Selector) {
		p(s.Not())
	})
}
