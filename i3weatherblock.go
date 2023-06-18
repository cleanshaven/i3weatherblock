package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/esiqveland/notify"
	"github.com/godbus/dbus/v5"

	"bytes"

	"github.com/cleanshaven/i3weatherblock/config"
	"github.com/cleanshaven/noaaweather"
)

type ForecastToOutput struct {
	Forecast noaaweather.PeriodJson
	Alert    noaaweather.AlertFeatureJson
}

func main() {
	err := config.SetupConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	forecast, alerts, err := noaaweather.GetWeather(config.MyConfig.Location.Latitude,
		config.MyConfig.Location.Longitude)

	output := ForecastToOutput{}
	output.Forecast = forecast.Properties.Periods[0]
	if len(alerts.Features) > 0 {
		output.Alert = alerts.Features[0]
		//		displayAlert(output)
	}

	writeFormatedWeather(output)
	if config.IsButtonPressed() {
		displayDetail(output)
	}

}

func getIcon(url string) (filename string, err error) {
	filename = ""
	err = nil

	filename, err = noaaweather.GetIcon(url)
	return
}

func writeFormatedWeather(forecast ForecastToOutput) error {

	tmpl, err := template.New("weather").Parse(config.MyConfig.Block.Template)
	if err != nil {
		log.Println(err)
		return err
	}
	err = tmpl.Execute(os.Stdout, forecast)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("") // required for I3Blocks to clear the cache

	return err

}

func displayAlert(forecast ForecastToOutput) error {
	tmpl, err := template.New("alertWeather").Parse(config.MyConfig.AlertPopup.Template)
	if err != nil {
		return err
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, forecast)
	if err != nil {
		return err
	}
	result := tpl.String()
	alertNotification(config.MyConfig.AlertPopup.Title, result, "", config.MyConfig.AlertPopup.TimeToShow)
	return nil
}

func displayDetail(forecast ForecastToOutput) error {
	tmpl, err := template.New("detailedWeather").Parse(config.MyConfig.DetailPopup.Template)
	if err != nil {
		return err
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, forecast)
	if err != nil {
		return err
	}
	result := tpl.String()
	icon, err := getIcon(forecast.Forecast.Icon)
	alertNotification(config.MyConfig.DetailPopup.Title, result, icon, config.MyConfig.DetailPopup.TimeToShow)

	return nil
}

func alertNotification(summary, description, icon string, expireTime int) {
	conn, err := dbus.SessionBusPrivate()
	if err != nil {
		return
	}
	defer conn.Close()

	if err = conn.Auth(nil); err != nil {
		return
	}

	if err = conn.Hello(); err != nil {
		return
	}

	n := notify.Notification{
		AppName:       "I3Weather",
		ReplacesID:    uint32(0),
		Summary:       summary,
		Body:          description,
		ExpireTimeout: time.Second * time.Duration(expireTime),
		AppIcon:       icon,
	}

	notify.SendNotification(conn, n)
}
