package main

import (
	"context"
	"fmt"
	"io"

	"github.com/GurbanowGadam/grpc_weather/weather"
	"google.golang.org/grpc"
)

func main() {
	addr := "127.0.0.1:8080"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc.Dial => ", err)
	}
	client := weather.NewWeatherServiceClient(conn)

	ctx := context.Background()

	resp, err := client.ListCities(ctx, &weather.ListCitiesRequest{})
	if err != nil {
		fmt.Println("client.ListCities => ", err)
	}
	fmt.Println("Cities => ")
	for _, city := range resp.Items {
		fmt.Println(city.CityCode, "|", city.CityName)
	}

	stream, err := client.QueryWeather(ctx, &weather.WeatherRequest{CityCode: "tm_mr"})
	if err != nil {
		fmt.Println("clinet.QueryWeather => ", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("stream.Recv => ", err)
		}
		fmt.Println("Temperature => ", msg)
	}
	fmt.Println("Server Spopped")
}
