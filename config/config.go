package config

import (
	"github.com/spf13/viper"
)

const (
	defaultLatitude       = "41.26"
	defaultLongitude      = "-81.86"
	defaultBlockTemplate  = `{{.Forecast.Temperature}}°F {{.Forecast.ShortForecast}} {{.Forecast.ProbabilityOfPrecipitation.Value}}% <span foreground="red">{{.Alert.Properties.Event}}</span>`
	defaultTimeToShow     = 30
	defaultAlertTemplate  = `{{.Alert.Properties.Event}}`
	defaultDetailTemplate = `{{.Alert.Properties.Description}}`
	defaultAlertTitle     = "Weather Alert"
	defaultDetailTitle    = "Weather Details"

	latitudeVariable              = "location.latitude"
	longitudeVariable             = "location.longitude"
	blockTemplateVariable         = "block.template"
	alertPopupTemplateVariable    = "alertPopup.template"
	alertPopupTimeToShowVariable  = "alertPopup.timeToShow"
	detailPopupTemplateVariable   = "detailPopup.template"
	detailPopupTimeToShowVariable = "detailPopup.timeToShow"
	buttonVariable                = "BLOCK_BUTTON"
	defaultBlockButton            = "-1"
	alertTitleVariable            = "alertPopup.title"
	detailTitleVariable           = "detailPopup.title"

	configName = "config.yml"
)

type Location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Block struct {
	Template string `json:"template"`
	Button   int    `json:"button"`
}

type Popup struct {
	Template   string `json:"template"`
	TimeToShow int    `json:"timeToShow"`
	Title      string `json:"title"`
}

type Config struct {
	Location    Location `json:"location"`
	Block       Block    `json:"block"`
	AlertPopup  Popup    `json:"alertPopup"`
	DetailPopup Popup    `json:"detailPopup"`
}

var MyConfig Config

var ConfigLocation string = `$HOME/.config/i3weatherblock`

func SetupConfig() error {
	setDefaults()
	setEnv()

	err := setConfigFile()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&MyConfig)
	return err
}

func setDefaults() {
	viper.SetDefault(latitudeVariable, defaultLatitude)
	viper.SetDefault(longitudeVariable, defaultLongitude)
	viper.SetDefault(blockTemplateVariable, defaultBlockTemplate)
	viper.SetDefault(alertPopupTemplateVariable, defaultAlertTemplate)
	viper.SetDefault(alertPopupTimeToShowVariable, defaultTimeToShow)
	viper.SetDefault(detailPopupTemplateVariable, defaultDetailTemplate)
	viper.SetDefault(detailPopupTimeToShowVariable, defaultTimeToShow)
	viper.SetDefault(buttonVariable, defaultBlockButton)
	viper.SetDefault(alertTitleVariable, defaultAlertTitle)
	viper.SetDefault(detailTitleVariable, defaultDetailTitle)
}

func setEnv() {
	viper.AutomaticEnv()
}

func setConfigFile() error {
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(ConfigLocation)
	err := viper.ReadInConfig()
	return err
}

func IsButtonPressed() bool {
	return viper.GetString(buttonVariable) != "-1"
}
