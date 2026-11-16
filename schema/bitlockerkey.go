package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// BitlockerKey holds the schema definition for the BitlockerKey entity.
type BitlockerKey struct {
	ent.Schema
}

// Fields of the BitlockerKey.
func (BitlockerKey) Fields() []ent.Field {
	return []ent.Field{
		field.String("agent_id"),
		field.String("drive_letter").Default("C:"),
		field.String("key_id"),
		field.String("encrypted_password"),
		field.Bool("is_active").Default(true),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Indexes of the BitlockerKey.
func (BitlockerKey) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("agent_id"),
		index.Fields("key_id"),
	}
}
