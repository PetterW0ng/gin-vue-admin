package response

import (
	"strconv"
)

type XetOrder struct {
	UserId       string `json:"userId" bson:"userId"`
	PaymentType  string `json:"paymentType" bson:"paymentType"`
	ResourceType string `json:"resourceType" bson:"resourceType"`
	ProductId    string `json:"productId" bson:"productId"`
	PurchaseName string `json:"purchaseName" bson:"purchaseName"`
	Price        int    `json:"price" bson:"price"`
	OrderState   string `json:"orderState" bson:"orderState"`
	WxAppType    string `json:"wxAppType" bson:"wxAppType"`
	Period       string `json:"period" bson:"period"`
	PayTime      string `json:"payTime" bson:"payTime"`
	CreatedAt    string `json:"createdAt" bson:"createdAt"`
}

type XetOrderOuter struct {
	UserId       string `mapstructure:"user_id" bson:"userId"`
	PaymentType  int    `mapstructure:"payment_type" bson:"paymentType"`
	ResourceType int    `mapstructure:"resource_type" bson:"resourceType"`
	ProductId    string `mapstructure:"product_id" bson:"productId"`
	PurchaseName string `mapstructure:"purchase_name" bson:"purchaseName"`
	Price        int    `mapstructure:"price" bson:"price"`
	OrderState   int    `mapstructure:"order_state" bson:"orderState"`
	OutOrderId   string `mapstructure:"out_order_id" bson:"outOrderId"`
	WxAppType    int    `mapstructure:"wx_app_type" bson:"wxAppType"`
	Period       int    `mapstructure:"period" bson:"period"`
	PayTime      string `mapstructure:"pay_time" bson:"payTime"`
	CreatedAt    string `mapstructure:"created_at" bson:"createdAt"`
}

func (xetOrder *XetOrder) SetValues(outer XetOrderOuter) {
	xetOrder.UserId = outer.UserId
	xetOrder.ProductId = outer.ProductId
	xetOrder.PurchaseName = outer.PurchaseName
	xetOrder.Price = outer.Price
	xetOrder.Period = strconv.Itoa(outer.Period)
	xetOrder.PayTime = outer.PayTime
	xetOrder.CreatedAt = outer.CreatedAt
}
