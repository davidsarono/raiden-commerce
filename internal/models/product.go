package models

import (
	"github.com/sev-2/raiden"
)

type Product struct {
	raiden.ModelBase
	Id          int32    `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('product_id_seq'::regclass)"`
	Name        *string  `json:"name,omitempty" column:"name:name;type:varchar;nullable"`
	Description *string  `json:"description,omitempty" column:"name:description;type:text;nullable"`
	Price       *float64 `json:"price,omitempty" column:"name:price;type:numeric;nullable"`
	CategoryId  *int32   `json:"category_id,omitempty" column:"name:category_id;type:integer;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Productcategory *Productcategory `json:"productcategory,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:category_id"`
	Orderitem       []*Orderitem     `json:"orderitem,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:product_id"`
	Order           []*Order         `json:"order,omitempty" join:"joinType:manyToMany;through:orderitem;sourcePrimaryKey:id;sourceForeignKey:product_id;targetPrimaryKey:id;targetForeign:product_id"`
}
