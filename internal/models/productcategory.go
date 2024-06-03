package models

import (
	"github.com/sev-2/raiden"
)

type Productcategory struct {
	raiden.ModelBase
	Id   int32   `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('productcategory_id_seq'::regclass)"`
	Name *string `json:"name,omitempty" column:"name:name;type:varchar;nullable;unique"`

	// Table information
	Metadata string `json:"-" schema:"public" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Product []*Product `json:"product,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:category_id"`
}
