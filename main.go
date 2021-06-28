// package main

// import (
// 	"fmt"

// 	"github.com/litethink/huobiapi"
// )

// func main() {
// 	// 创建客户端实例
// 	market, err := huobiapi.NewMarket()
// 	if err != nil {
// 		panic(err)
// 	}
// 	// 订阅主题 market.ethbtc.kline.1min  market.yfiiusdt.trade.detail
// 	market.Subscribe("market.yfiiusdt.trade.detail", func(topic string, json *huobiapi.JSON) {
// 		// 收到数据更新时回调
// 		fmt.Println(topic, json)
// 	})
// 	// 请求数据
// 	//json, err := market.Request("market.ethbtc.kline.1min")
// 	//if err != nil {
// 		//panic(err)
// 	//} else {
// 	//	fmt.Println(json)
// 	//}
// 	// 进入阻塞等待，这样不会导致进程退出
// 	market.Loop()
// }


// package main

// import (
// 	//"reflect" //for TypeOf
// 	"fmt"
// 	// "time"
// 	"github.com/litethink/huobiapi"
// )


// func main() {
// 	// 创建客户端实例

// 	var dataArr [] []interface{}
// 	market1, err := huobiapi.NewMarket()
// 	if err != nil {
// 		panic(err)
// 	}
// 	// 订阅主题 market.ethbtc.kline.1min  market.yfiiusdt.trade.detail
// 	market1.Subscribe("market.btcusdt.trade.detail", func(topic string, json *huobiapi.JSON) {

        
//          //data was type of interface{}
//         data,err := (json.Get("tick").Get("data")).Array()
//         if err != nil {
//         	fmt.Println(err)
//         }

//         //data[0] was type of map[string]interface {}
//         var dataMap = make(map[string] interface{})
//         var arr []interface {}
//         for k1, v1 := range data[0].(map[string]interface{}){
//             dataMap[k1] = v1
//         }
//         arr =append(arr,dataMap["price"],dataMap["amount"],dataMap["direction"],dataMap["id"],dataMap["tradeId"],dataMap["ts"]) 
//         dataArr = append(dataArr,arr)
//         // fmt.Println(json)
//         // fmt.Println(json.Get("ts").MustInt())
//         // fmt.Println(json.Get("ch").MustString())
//         // fmt.Println(json.Get("tick").Get("data").Interface())
//         if len(dataArr) == 2 {
//         	fmt.Println(dataArr)
//         	dataArr = dataArr[2:]
//         } 
// 	})
	
// 	// 请求数据
// 	//json, err := market.Request("market.ethbtc.kline.1min")
// 	//if err != nil {
// 		//panic(err)
// 	//} else {
// 	//	fmt.Println(json)
// 	//}
// 	// 进入阻塞等待，这样不会导致进程退出
// 	// time.Sleep(time.Second*5)

// 	market1.Loop()
// }

package main

import (
	"fmt"
	"github.com/litethink/huobiapi"
)


func main() {

	var dataArr [] []interface{}
	market1, err := huobiapi.NewMarket()
	if err != nil {
		panic(err)
	}
	market1.Subscribe("market.btcusdt.trade.detail", func(topic string, json *huobiapi.JSON) {

        data,err := (json.Get("tick").Get("data")).Array()
        if err != nil {
        	fmt.Println(err)
        }
        var dataMap = make(map[string] interface{})
        var arr []interface {}
        for k1, v1 := range data[0].(map[string]interface{}){
            dataMap[k1] = v1
        }
        arr =append(arr,dataMap["price"],dataMap["amount"],dataMap["direction"],dataMap["id"],dataMap["tradeId"],dataMap["ts"]) 
        fmt.Println(arr[4],arr[0],arr[1])
        dataArr = append(dataArr,arr)
        // if len(dataArr) == 2 {
        	// fmt.Println(dataArr)
        // 	dataArr = dataArr[2:]
        // } 
	})
	market1.Loop()
}

// package main

// import (
// 	"fmt"

// 	"github.com/litethink/huobiapi"
// )

// func main() {
// 	// 创建客户端实例
// 	market, err := huobiapi.NewMarket()
// 	if err != nil {
// 		panic(err)
// 	}
// 	// 订阅主题 market.ethbtc.kline.1min  market.yfiiusdt.trade.detail
// 	market.Subscribe("market.btcusdt.trade.detail", func(topic string, json *huobiapi.JSON) {
// 		// 收到数据更新时回调
// 		fmt.Println(topic, json.Get("tick").Get("data"))
// 	})
// 	// 请求数据
// 	//json, err := market.Request("market.ethbtc.kline.1min")
// 	//if err != nil {
// 		//panic(err)
// 	//} else {
// 	//	fmt.Println(json)
// 	//}
// 	// 进入阻塞等待，这样不会导致进程退出
// 	market.Loop()
// }
