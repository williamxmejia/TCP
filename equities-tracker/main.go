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
	data2 := data[1].(map[string]interface{})

	fmt.Println("api response(interface):")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("static", "./static")

	app.Get("/", func(ctx *fiber.Ctx) error {
		var tracker Tracker
		tracker.Count = 0

		//var keyItems []string
		//var valueItems []string

		fmt.Println(data2["cmc_rank"])
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

		fmt.Println()
		fmt.Println()
		
		fmt.Println(multiMap["fruits"]["apple"])

		/*
		var count int
		
		for key, value := range data {
			//keyItems = append(keyItems, key) 
			//valueItems =  append(valueItems, value)
			//count += 1
			fmt.Printf("%d: %v\n", key, value)
			//fmt.Println(count)
		}
		*/

		for i := 0; i < 5; i++ {
			tracker.Count += 1
		}

		return ctx.Render("index", fiber.Map{
			"Title": "Fiber",
			"Message": "Dynamic view",
			"Count": tracker.Count,
			//"respBody": valueItems,
		})
	})

	
	app.Listen(":3000")

}
