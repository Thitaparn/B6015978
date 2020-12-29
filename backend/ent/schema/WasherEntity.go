package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Washer holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

// Fields of the Washer.
func (Washer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("count").NotEmpty(),
		field.Time("time"),
	}
}

// Edges of the Washer.
func (Washer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("coin", coin.Type).Ref("Washers").Unique(),          
		edge.From("medicalcare", MedicalCare.Type).Ref("Washers").Unique(), 
		edge.From("employee", Employee.Type).Ref("Washers").Unique(),     
		edge.From("disease", Disease.Type).Ref("Washers").Unique(),   
	}