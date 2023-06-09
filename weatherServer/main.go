package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"

	"github.com/GurbanowGadam/grpc_weather/weather"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}

	srv := grpc.NewServer()
	weather.RegisterWeatherServiceServer(srv, &myWeatherService{})
	fmt.Println("Starting server...")
	panic(srv.Serve(lis))
}

type myWeatherService struct {
	weather.UnimplementedWeatherServiceServer
}

func (m *myWeatherService) ListCities(ctx context.Context, req *weather.ListCitiesRequest) (*weather.ListCitiesResponse, error) {
	return &weather.ListCitiesResponse{
		Items: []*weather.CityEntry{
			&weather.CityEntry{CityCode: "tkm_ag", CityName: "Ashgabat"},
			&weather.CityEntry{CityCode: "tkm_mr", CityName: "Mary"},
		},
	}, nil
}

func (m *myWeatherService) QueryWeather(req *weather.WeatherRequest, resp weather.WeatherService_QueryWeatherServer) error {
	for {
		err := resp.Send(&weather.WeatherResponse{Temperature: rand.Float32()*10 + 10})
		if err != nil {
			break
		}
	}
	return nil
}
