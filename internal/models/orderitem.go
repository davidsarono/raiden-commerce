package models

import (
	"github.com/sev-2/raiden"
)

type Orderitem struct {
	raiden.ModelBase
	OrderItemId int32 `json:"order_item_id,omitempty" column:"name:order_item_id;type:integer;primaryKey;nullable:false;default:nextval('orderitem_order_item_id_seq'::regclass)"`
	OrderId *int32 `json:"order_id,omitempty" column:"name:order_id;type:integer;nullable"`
	ProductId *int32 `json:"product_id,omitempty" column:"name:product_id;type:integer;nullable"`
	Quantity *int32 `json:"quantity,omitempty" column:"name:quantity;type:integer;nullable"`
	Price *float64 `json:"price,omitempty" column:"name:price;type:numeric;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Order *Order `json:"order,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:order_id"`
	Product *Product `json:"product,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:product_id"`
}
