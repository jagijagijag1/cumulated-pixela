package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	pixela "github.com/gainings/pixela-go-client"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) error {
	// extract env var
	user := os.Getenv("PIXELA_USER")
	token := os.Getenv("PIXELA_TOKEN")
	graph := os.Getenv("PIXELA_GRAPH")

	// extract data from pixel on yesterday
	date, quantity := getPreviousPixel(user, token, graph)
	if date == "-1" || quantity == "-1" {
		return errors.New("Error in accessing pixela")
	}
	fmt.Printf("date: %s, quantity: %s\n", date, quantity)

	// record pixel
	perr := recordPixel(user, token, graph, date, quantity)
	if perr != nil {
		return errors.New("Error in accessing pixela")
	}

	return nil
}

func getPreviousPixel(user, token, graph string) (string, string) {
	c := pixela.NewClient(user, token)

	// set date
	yesterday := time.Now().AddDate(0, 0, -1).Format("20060102")
	today := time.Now().Format("20060102")

	// get yseterday's pixel
	q, err := c.GetPixelQuantity(graph, yesterday)
	if err != nil {
		fmt.Println(err)
		return "-1", "-1"
	}

	quantity := strconv.FormatFloat(q, 'f', 4, 64)

	return today, quantity
}

func recordPixel(user, token, graph, date, quantity string) error {
	c := pixela.NewClient(user, token)

	// try to record
	err := c.RegisterPixel(graph, date, quantity)
	if err == nil {
		fmt.Println("recorded")
		return err
	}

	// if fail, try to update
	err = c.UpdatePixelQuantity(graph, date, quantity)
	if err == nil {
		fmt.Println("updated")
	}

	return err
}

func main() {
	lambda.Start(Handler)
}
