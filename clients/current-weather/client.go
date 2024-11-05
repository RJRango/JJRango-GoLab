package current_weather

import (
	"errors"
	"net/http"
	"net/url"
	"os"
)

type currentWeatherClient struct {
	baseUri *url.URL
	client  http.Client
}

func NewCurrentWeatherClient() (currentWeatherClient, error) {
	rawUri, ok := os.LookupEnv("URI_OPEN_WEATHER_CURRENT")
	if !ok {
		err := errors.New("cannot find URI_OPEN_WEATHER_CURRENT environment variable")
		return currentWeatherClient{}, err
	}

	base, err := url.Parse(rawUri)
	if err != nil {
		return currentWeatherClient{}, err
	}

	return currentWeatherClient{
		baseUri: base,
		client:  http.Client{},
	}, nil
}
