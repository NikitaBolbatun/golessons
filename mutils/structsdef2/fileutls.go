package structsdef1

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	openFileError   = errors.New("Open file error")
	collectionEmpty = errors.New("Empty collection")
	temDirs         = `C:\MyGo\tmp\`
)

func gprep(colname string) *os.File {
	vcolname := fmt.Sprintf("%s.txt", colname)
	f, err := os.OpenFile(fmt.Sprintf("%s%s", temDirs, vcolname),
		os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {
		f = nil
		panic(openFileError)
	}
	return f
}

func gdef(f *os.File, colname string) {
	if r := recover(); r != nil {
		switch r {
		case openFileError:
			{
				fmt.Println(openFileError)
				f.Close()
			}
		case collectionEmpty:
			{
				fmt.Println(collectionEmpty)
			}
		default:
			{
				if f != nil {
					fmt.Println("неизвестная ошибка")
					f.Close()
				}
			}
		}
	} else {
		f.Close()
		fmt.Printf("%s Saved !!!\n", colname)
	}
}

//SaveAcountList - схраним пользователей
func SaveAcountList(vacountList map[string]*User) {
	if len(vacountList) == 0 {
		panic(collectionEmpty)
	}
	f := gprep("acountList")
	for id, item := range vacountList {
		str := fmt.Sprintf("%s:%f:%s:%d\n",
			id, item.Account, item.Email, item.UserType)
		_, err := f.WriteString(str)
		if err != nil {
			fmt.Println("Строка %s пропущена с ошибкой %v ", str, err)
			continue
		}
	}
	defer gdef(f, "acountList")
}

//SaveBillList - схраним историю счтов
func SaveBillList(vbillList map[string]map[int]float32) {
	if len(vbillList) == 0 {
		panic(collectionEmpty)
	}
	f := gprep("billList")
	for id, bills := range vbillList {
		strb := ""
		for idx, bill := range bills {
			strb = fmt.Sprintf("%s,%d:%f", strb, idx, bill)
		}
		str := fmt.Sprintf("%s{%s}\n", id, strb[1:]) // выкинул лишнюю запятую
		_, err := f.WriteString(str)
		if err != nil {
			fmt.Println("Строка %s пропущена с ошибкой %v ", str, err)
			continue
		}
	}
	defer gdef(f, "billList")
}

//SaveItemsPrice - схраним каталог товаров
func SaveItemsPrice(vitemsPrice map[string]*ItemPrice) {
	if len(vitemsPrice) == 0 {
		panic(collectionEmpty)
	}
	f := gprep("itemsPrice")
	for name, items := range vitemsPrice {
		strb := fmt.Sprintf("{ItemPrice:%f,ItemType:%d}",
			items.ItemPrice, items.ItemType)
		str := fmt.Sprintf("%s%s\n", name, strb)
		_, err := f.WriteString(str)
		if err != nil {
			fmt.Println("Строка %s пропущена с ошибкой %v ", str, err)
			continue
		}
	}
	defer gdef(f, "itemsPrice")
}

//SaveOrdersPrice - схраним заказы с ценами
func SaveOrdersPrice(vordersPrice []Order) {
	if len(vordersPrice) == 0 {
		panic(collectionEmpty)
	}
	f := gprep("ordersPrice")
	for _, elem := range vordersPrice {
		str := strings.Join(elem.Items, ",")
		str = fmt.Sprintf("{{%s},%f,%d}\n", str, elem.TotalSum, elem.OrderType)
		_, err := f.WriteString(str)
		if err != nil {
			fmt.Println("Строка %s пропущена с ошибкой %v ", str, err)
			continue
		}
	}

	defer gdef(f, "ordersPrice")
}