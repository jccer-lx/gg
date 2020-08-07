package helper

import (
	"encoding/json"
	"fmt"
	"testing"
)

type goodsUpdateParams struct {
	Name            string
	Price           json.Number
	Sort            json.Number
	Sales           json.Number
	Stock           json.Number
	GiveIntegral    json.Number
	StrPrice        string
	StrSort         string
	StrSales        string
	StrStock        string
	StrGiveIntegral string
}

type goods struct {
	Name         string
	Price        float64
	Sort         int
	Sales        int64
	Stock        uint
	GiveIntegral string

	StrPrice        float64
	StrSort         int
	StrSales        int64
	StrStock        uint
	StrGiveIntegral string
}

func TestReflectiveStructToStructWithJson(t *testing.T) {
	gup := new(goodsUpdateParams)
	gup.Name = "ABC"
	gup.Price = "25.5"
	gup.Sort = "60"
	gup.Sales = "50000"
	gup.Stock = "99999"
	gup.GiveIntegral = "3"
	gup.StrPrice = "26.5"
	gup.StrSort = "61"
	gup.StrSales = "50001"
	gup.StrStock = "999999"
	gup.StrGiveIntegral = "31"
	g := new(goods)
	ReflectiveStructToStructWithJson(g, gup)
	fmt.Println(g.Name)
	fmt.Println(g.Price)
	fmt.Println(g.Sort)
	fmt.Println(g.Sales)
	fmt.Println(g.Stock)
	fmt.Println(g.GiveIntegral)
	fmt.Println(g.StrPrice)
	fmt.Println(g.StrSort)
	fmt.Println(g.StrSales)
	fmt.Println(g.StrStock)
	fmt.Println(g.StrGiveIntegral)
}
