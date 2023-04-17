package main

import (
        "context"
        "fmt"
        "math/rand"
        "net"
        "weather/api"

        "google.golang.org/grpc"
)

func main() {
        lis, err := net.Listen("tcp", "127.0.0.1:8080")
        if err != nil {
                fmt.Println(err)
        }

        srv := grps.NewServer()
        api.RegisterWeatherServiceServer(srv, &myWeatherService{})
        fmt.Println("Starting server...")
        panic(srv.Serve(lis))
}

type myWeatherService struct {
        api.UnimplementedWeatherServiceServer
}

func (m *myWeatherService) ListCities(ctx context.Context, req *api.ListCitiesRequest) (*api.ListCitiesResponse, error) {
        return &api.ListCitiesResponse{
                Items: []*api.CityEntry{
                        &api.CityEntry{CityCode: "tkm_ag", CityName: "Ashgabat"},
                        &api.CityEntry{CityCode: "tkm_mr", CityName: "Mary"}
                },
        }, nil
}

func (m *myWeatherService) QueryWeather(req *api.WeatherRequest, resp api.WeatherService_QueryWeatherService) error { 
        for {
                err := resp.Send(&api.WeatherResponse{Temperature: rand.Float32()*10 + 10})
                if err != nil {
                        break
                }
        }
        return nil
}
