package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sev-2/raiden"
	"github.com/sev-2/raiden/pkg/supabase/objects"
)

type Order struct {
	raiden.ModelBase
	Id        int32      `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('\"Order_id_seq\"'::regclass)"`
	OrderDate *time.Time `json:"order_date,omitempty" column:"name:order_date;type:timestamp;nullable;default:CURRENT_TIMESTAMP"`
	UserId    *uuid.UUID `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Users     *objects.User `json:"users,omitempty" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
	Orderitem []*Orderitem  `json:"orderitem,omitempty" join:"joinType:hasMany;primaryKey:id;foreignKey:order_id"`
	Product   []*Product    `json:"product,omitempty" join:"joinType:manyToMany;through:orderitem;sourcePrimaryKey:id;sourceForeignKey:order_id;targetPrimaryKey:id;targetForeign:order_id"`
}
