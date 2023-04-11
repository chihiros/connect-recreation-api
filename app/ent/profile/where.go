// Code generated by ent, DO NOT EDIT.

package profile

import (
	"app/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Profile {
	return predicate.Profile(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Profile {
	return predicate.Profile(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Profile {
	return predicate.Profile(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Profile {
	return predicate.Profile(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Profile {
	return predicate.Profile(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Profile {
	return predicate.Profile(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Profile {
	return predicate.Profile(sql.FieldLTE(FieldID, id))
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldNickname, v))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldUUID, v))
}

// IconURL applies equality check predicate on the "icon_url" field. It's identical to IconURLEQ.
func IconURL(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldIconURL, v))
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldNickname, v))
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.Profile {
	return predicate.Profile(sql.FieldNEQ(FieldNickname, v))
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.Profile {
	return predicate.Profile(sql.FieldIn(FieldNickname, vs...))
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.Profile {
	return predicate.Profile(sql.FieldNotIn(FieldNickname, vs...))
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.Profile {
	return predicate.Profile(sql.FieldGT(FieldNickname, v))
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.Profile {
	return predicate.Profile(sql.FieldGTE(FieldNickname, v))
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.Profile {
	return predicate.Profile(sql.FieldLT(FieldNickname, v))
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.Profile {
	return predicate.Profile(sql.FieldLTE(FieldNickname, v))
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.Profile {
	return predicate.Profile(sql.FieldContains(FieldNickname, v))
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.Profile {
	return predicate.Profile(sql.FieldHasPrefix(FieldNickname, v))
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.Profile {
	return predicate.Profile(sql.FieldHasSuffix(FieldNickname, v))
}

// NicknameEqualFold applies the EqualFold predicate on the "nickname" field.
func NicknameEqualFold(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEqualFold(FieldNickname, v))
}

// NicknameContainsFold applies the ContainsFold predicate on the "nickname" field.
func NicknameContainsFold(v string) predicate.Profile {
	return predicate.Profile(sql.FieldContainsFold(FieldNickname, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.Profile {
	return predicate.Profile(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.Profile {
	return predicate.Profile(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...string) predicate.Profile {
	return predicate.Profile(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v string) predicate.Profile {
	return predicate.Profile(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.Profile {
	return predicate.Profile(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.Profile {
	return predicate.Profile(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.Profile {
	return predicate.Profile(sql.FieldLTE(FieldUUID, v))
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.Profile {
	return predicate.Profile(sql.FieldContains(FieldUUID, v))
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.Profile {
	return predicate.Profile(sql.FieldHasPrefix(FieldUUID, v))
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.Profile {
	return predicate.Profile(sql.FieldHasSuffix(FieldUUID, v))
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEqualFold(FieldUUID, v))
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.Profile {
	return predicate.Profile(sql.FieldContainsFold(FieldUUID, v))
}

// IconURLEQ applies the EQ predicate on the "icon_url" field.
func IconURLEQ(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEQ(FieldIconURL, v))
}

// IconURLNEQ applies the NEQ predicate on the "icon_url" field.
func IconURLNEQ(v string) predicate.Profile {
	return predicate.Profile(sql.FieldNEQ(FieldIconURL, v))
}

// IconURLIn applies the In predicate on the "icon_url" field.
func IconURLIn(vs ...string) predicate.Profile {
	return predicate.Profile(sql.FieldIn(FieldIconURL, vs...))
}

// IconURLNotIn applies the NotIn predicate on the "icon_url" field.
func IconURLNotIn(vs ...string) predicate.Profile {
	return predicate.Profile(sql.FieldNotIn(FieldIconURL, vs...))
}

// IconURLGT applies the GT predicate on the "icon_url" field.
func IconURLGT(v string) predicate.Profile {
	return predicate.Profile(sql.FieldGT(FieldIconURL, v))
}

// IconURLGTE applies the GTE predicate on the "icon_url" field.
func IconURLGTE(v string) predicate.Profile {
	return predicate.Profile(sql.FieldGTE(FieldIconURL, v))
}

// IconURLLT applies the LT predicate on the "icon_url" field.
func IconURLLT(v string) predicate.Profile {
	return predicate.Profile(sql.FieldLT(FieldIconURL, v))
}

// IconURLLTE applies the LTE predicate on the "icon_url" field.
func IconURLLTE(v string) predicate.Profile {
	return predicate.Profile(sql.FieldLTE(FieldIconURL, v))
}

// IconURLContains applies the Contains predicate on the "icon_url" field.
func IconURLContains(v string) predicate.Profile {
	return predicate.Profile(sql.FieldContains(FieldIconURL, v))
}

// IconURLHasPrefix applies the HasPrefix predicate on the "icon_url" field.
func IconURLHasPrefix(v string) predicate.Profile {
	return predicate.Profile(sql.FieldHasPrefix(FieldIconURL, v))
}

// IconURLHasSuffix applies the HasSuffix predicate on the "icon_url" field.
func IconURLHasSuffix(v string) predicate.Profile {
	return predicate.Profile(sql.FieldHasSuffix(FieldIconURL, v))
}

// IconURLEqualFold applies the EqualFold predicate on the "icon_url" field.
func IconURLEqualFold(v string) predicate.Profile {
	return predicate.Profile(sql.FieldEqualFold(FieldIconURL, v))
}

// IconURLContainsFold applies the ContainsFold predicate on the "icon_url" field.
func IconURLContainsFold(v string) predicate.Profile {
	return predicate.Profile(sql.FieldContainsFold(FieldIconURL, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Profile) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Profile) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
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
func Not(p predicate.Profile) predicate.Profile {
	return predicate.Profile(func(s *sql.Selector) {
		p(s.Not())
	})
}
