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
	"github.com/joho/godotenv"
)

type Tracker struct {
	Count int
}

type Quote struct {
	Price float64
	MarketCap float64
	LastUpdated string
}

type EquityInfo struct {
	Name string
	Symbol string
	Quote Quote
}


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	key := os.Getenv("API_KEY")
	port := os.Getenv("PORT")

	client := &http.Client{}
	engine := html.New("./views", ".html")

	//req, err := http.NewRequest("GET","https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)

	req, err := http.NewRequest("GET","https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "5000")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
  req.Header.Add("X-CMC_PRO_API_KEY", key)
  req.URL.RawQuery = q.Encode()


  resp, err := client.Do(req);
  if err != nil {
    fmt.Println("Error sending request to server")
    os.Exit(1)
  }
  fmt.Println(resp.Status);
  respBody, _ := ioutil.ReadAll(resp.Body)

	var result interface{}

	err = json.Unmarshal(respBody, &result)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	data := result.(map[string]interface{})["data"].([]interface{})

	fmt.Println("api response(interface):")
	//fmt.Println(quote.(map[string]interface{})["price"])

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("static", "./static")

	app.Get("/", func(ctx *fiber.Ctx) error {
		info := []EquityInfo{}

		for key, _ := range data {
			data2 := data[key].(map[string]interface{})
			quote := data2["quote"].(map[string]interface{})["USD"].(map[string]interface{})
			if name, ok := data2["name"].(string); ok {
				if symbol, ok := data2["symbol"].(string); ok {
					if price, ok := quote["price"].(float64); ok {
						if marketCap, ok := quote["market_cap"].(float64); ok {
							if lastUpdate, ok := quote["last_updated"].(string); ok {
								info = append(info, EquityInfo{
									Name: name, 
									Symbol: symbol, 
									Quote: Quote{
										Price: price,
										MarketCap: marketCap,
										LastUpdated: lastUpdate,
									},
								})
							} else {
								fmt.Println("failed")
							}
						}
					}
				}
			}
		}

		return ctx.Render("index", fiber.Map{
			"Title": "Fiber",
			"Message": "Dynamic view",
			"BTC": info[0],
			"ETH": info[1],
			//"Quote": info[i].Quote,
		})
	})

	app.Get("/search", func(ctx *fiber.Ctx) error {
		input := ctx.FormValue("query")
		response := fmt.Sprintf("You typed: %s", input)
		fmt.Println(response)
		return ctx.SendString(response)
	})


	
	app.Listen(":" + port)

}
