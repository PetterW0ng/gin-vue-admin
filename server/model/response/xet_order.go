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
	UserId       string `json:"user_id" bson:"userId"`
	PaymentType  int    `json:"payment_type" bson:"paymentType"`
	ResourceType int    `json:"resource_type" bson:"resourceType"`
	ProductId    string `json:"product_id" bson:"productId"`
	PurchaseName string `json:"purchase_name" bson:"purchaseName"`
	Price        int    `json:"price" bson:"price"`
	OrderState   int    `json:"order_state" bson:"orderState"`
	OutOrderId   string `json:"out_order_id" bson:"outOrderId"`
	WxAppType    int    `json:"wx_app_type" bson:"wxAppType"`
	Period       int    `json:"period" bson:"period"`
	PayTime      string `json:"pay_time" bson:"payTime"`
	CreatedAt    string `json:"created_at" bson:"createdAt"`
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
