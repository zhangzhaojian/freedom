//Package po generated by 'freedom new-po'
package po

import (
	"github.com/jinzhu/gorm"
	"time"
)

// OrderDetail .
type OrderDetail struct {
	changes   map[string]interface{}
	ID        int       `gorm:"primary_key;column:id"`
	OrderNo   string    `gorm:"column:order_no"`   // 订单id
	GoodsID   int       `gorm:"column:goods_id"`   // 商品id
	Num       int       `gorm:"column:num"`        // 数量
	GoodsName string    `gorm:"column:goods_name"` // 商品名称
	Created   time.Time `gorm:"column:created"`
	Updated   time.Time `gorm:"column:updated"`
}

// TableName .
func (obj *OrderDetail) TableName() string {
	return "order_detail"
}

// TakeChanges .
func (obj *OrderDetail) TakeChanges() map[string]interface{} {
	if obj.changes == nil {
		return nil
	}
	result := make(map[string]interface{})
	for k, v := range obj.changes {
		result[k] = v
	}
	obj.changes = nil
	return result
}

// updateChanges .
func (obj *OrderDetail) setChanges(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}

// SetOrderNo .
func (obj *OrderDetail) SetOrderNo(orderNo string) {
	obj.OrderNo = orderNo
	obj.setChanges("order_no", orderNo)
}

// SetGoodsID .
func (obj *OrderDetail) SetGoodsID(goodsID int) {
	obj.GoodsID = goodsID
	obj.setChanges("goods_id", goodsID)
}

// SetNum .
func (obj *OrderDetail) SetNum(num int) {
	obj.Num = num
	obj.setChanges("num", num)
}

// SetGoodsName .
func (obj *OrderDetail) SetGoodsName(goodsName string) {
	obj.GoodsName = goodsName
	obj.setChanges("goods_name", goodsName)
}

// SetCreated .
func (obj *OrderDetail) SetCreated(created time.Time) {
	obj.Created = created
	obj.setChanges("created", created)
}

// SetUpdated .
func (obj *OrderDetail) SetUpdated(updated time.Time) {
	obj.Updated = updated
	obj.setChanges("updated", updated)
}

// AddGoodsID .
func (obj *OrderDetail) AddGoodsID(goodsID int) {
	obj.GoodsID += goodsID
	obj.setChanges("goods_id", gorm.Expr("goods_id + ?", goodsID))
}

// AddNum .
func (obj *OrderDetail) AddNum(num int) {
	obj.Num += num
	obj.setChanges("num", gorm.Expr("num + ?", num))
}
