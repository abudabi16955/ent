package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SelfServicePackage holds the schema definition for the SelfServicePackage entity.
type SelfServicePackage struct {
	ent.Schema
}

// Fields of the SelfServicePackage.
func (SelfServicePackage) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
		field.String("category").Optional().Default("General"),
		field.String("version").Optional(),
		field.Enum("package_type").Values("msi", "exe", "zip", "winget"),
		field.Enum("source_type").Values("upload", "https"),
		field.String("source_url").Optional(),
		field.String("file_path").Optional(),
		field.String("checksum_sha256").Optional(),
		field.Enum("execution_context").Values("system", "user"),
		field.String("silent_args").Optional(),
		field.String("icon").Optional(),
		field.Bool("enabled").Default(true),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}
