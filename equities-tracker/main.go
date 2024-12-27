package main

import (
	"net/http"
	"os"
	"fmt"
	"log"
	"io/ioutil"
	"net/url"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Tracker struct {
	Count int
}

type EquityInfo struct {
	Name string
	Symbol string
}


func main() {
	client := &http.Client{}
	engine := html.New("./views", ".html")

	req, err := http.NewRequest("GET","https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "5000")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
  req.Header.Add("X-CMC_PRO_API_KEY", "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c")
  req.URL.RawQuery = q.Encode()


  resp, err := client.Do(req);
  if err != nil {
    fmt.Println("Error sending request to server")
    os.Exit(1)
  }
  fmt.Println(resp.Status);
  respBody, _ := ioutil.ReadAll(resp.Body)
  //fmt.Println(string(respBody))

	var result interface{}

	err = json.Unmarshal(respBody, &result)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	data := result.(map[string]interface{})["data"].([]interface{})

	fmt.Println("api response(interface):")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("static", "./static")

	app.Get("/", func(ctx *fiber.Ctx) error {
		var tracker Tracker
		tracker.Count = 0

		//fmt.Println(data2)

		/*

		multiMap := map[string]map[string]int{
        "fruits": {
            "apple":  5,
            "banana": 10,
            "cherry": 15,
        },
        "vegetables": {
            "carrot": 3,
            "potato": 7,
        },
    }
		*/

		fmt.Println()
		fmt.Println()
		
		//fmt.Println(multiMap["fruits"]["apple"])

		/*
		var count int
		*/
		//var name []string
		//var symbol []string

		info := []EquityInfo{}

		
		for key, _ := range data {
			data2 := data[key].(map[string]interface{})
			/*
			*/
			if data3, ok := data2["name"].(string); ok {
				if symbol, ok := data2["symbol"].(string); ok {
					info = append(info, EquityInfo{Name: data3, Symbol: symbol})
				} else {
					fmt.Println("failed")
				}
				//fmt.Println(data3)
			} else {
				fmt.Println("failed")
			}
			
			//info = append(info, EquityInfo{Name: data2["name"], Symbol: data2["symbol"]})
			//symbol =  append(symbol, data2["symbol"])
			//count += 1
			//fmt.Printf("Name:%s, symbol:%s\n", data2["name"], data2["symbol"])
			//fmt.Printf("Name:%s, symbol:%s\n", data3, data3)
			//fmt.Println(count)
		}

		for i := 0; i < 5; i++ {
			tracker.Count += 1
		}
		fmt.Println(info[1].Name)
		fmt.Println(info[1].Symbol)

		return ctx.Render("index", fiber.Map{
			"Title": "Fiber",
			"Message": "Dynamic view",
			"Count": tracker.Count,
			"Name": info[1].Name,
			"Symbol": info[1].Symbol,
		})
	})

	
	app.Listen(":3000")

}
