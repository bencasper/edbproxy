package mock

import (
	"encoding/xml"
	"time"
)

type EdbDelivery struct {
	XMLName xml.Name `xml:"items"`
	TotalResults int32 `xml:"totalResults"`
	TotalResultsAll int32 `xml:"totalResultsAll"`
	Rows ResultRow `xml:"Rows"`
}

type ResultRow struct {
	XMLName xml.Name `xml:"Rows"`
	Tid string `xml:"tid"`
	OutTid string `xml:"out_tid"`
	Status string `xml:"status"`
	ExpressNo string `xml:"express_no"`
	Express string `xml:"express"`
	ExpressCoding string `xml:"express_coding"`
	DeliveryTime string `xml:"delivery_time"`
	DeliveryStatus string `xml:"delivery_status"`
	Items []Item `xml:"tid_item>Item"`
}

type Item struct {
	XMLName xml.Name `xml:"Item"`
	ProductName string `xml:"product_name"`
}

func DeliveryXml(eorderNo string) *EdbDelivery {

	item := &Item{ProductName: "demo",}

	current := time.Now().Format("20060102150405")
	current1 := time.Now().Format("2006-01-02 15:04:05")
	row := &ResultRow{Tid: "S" + current,
		OutTid: eorderNo,
		Status: "已确认",
		ExpressNo: "ex" + current,
		DeliveryTime: current1,
		DeliveryStatus: "已发货",
		Items: []Item{*item},}
	delivery := &EdbDelivery{TotalResults: 1, TotalResultsAll: 1, Rows: *row,}

	//out, _ := xml.MarshalIndent(delivery, " ", "  ")

	//xml.Header + string(out)
	return delivery

}