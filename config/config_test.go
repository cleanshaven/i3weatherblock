package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func checkValues(t *testing.T) {
	if defaultLatitude != MyConfig.Location.Latitude {
		t.Fatalf("Latitude wanted %s got %s", defaultLatitude, MyConfig.Location.Latitude)
	}

	if defaultLongitude != MyConfig.Location.Longitude {
		t.Fatalf("Longitude wanted %s go %s", defaultLongitude, MyConfig.Location.Longitude)
	}

	if defaultBlockTemplate != MyConfig.Block.Template {
		t.Fatalf("Block Template wanted %s got %s", defaultBlockTemplate, MyConfig.Block.Template)
	}

	if defaultAlertTemplate != MyConfig.AlertPopup.Template {
		t.Fatalf("Alert Template wanted %s got %s", defaultAlertTemplate, MyConfig.AlertPopup.Template)
	}

	if defaultTimeToShow != MyConfig.AlertPopup.TimeToShow {
		t.Fatalf("Alert Timeout wanted %d got %d", defaultTimeToShow, MyConfig.AlertPopup.TimeToShow)
	}

	if defaultDetailTemplate != MyConfig.DetailPopup.Template {
		t.Fatalf("Detail Template wanted %s got %s", defaultDetailTemplate, MyConfig.DetailPopup.Template)
	}

	if defaultTimeToShow != MyConfig.DetailPopup.TimeToShow {
		t.Fatalf("Detail Timeout wanted %d got %d", defaultTimeToShow, MyConfig.DetailPopup.TimeToShow)
	}

}

func setEnvironmentVariables() {
	os.Setenv("LOCATION_LATITUDE", defaultLatitude)
	os.Setenv("LOCATION_LONGITUDE", defaultLongitude)
	os.Setenv("BLOCK_TEMPLATE", defaultBlockTemplate)
	os.Setenv("ALERTPOPUP_TEMPLATE", defaultAlertTemplate)
	os.Setenv("DETAILPOPUP_TEMPLATE", defaultDetailTemplate)
	os.Setenv("ALERTPOPUP_TIMETOSHOW", "30")
	os.Setenv("DETAILPOPUP_TIMETOSHOW", "30")
}

func TestDefaults(t *testing.T) {
	viper.Reset()
	setDefaults()
	err := viper.Unmarshal(&MyConfig)
	if err != nil {
		t.Fatalf(`Test Defaults failed to unmarshel config %v`, err)
	}

	checkValues(t)
}

func TestEnvironment(t *testing.T) {
	viper.Reset()
	setEnvironmentVariables()
	checkValues(t)
}

func TestConfigFile(t *testing.T) {
	viper.Reset()
	ConfigLocation = "."
	err := setConfigFile()
	if err != nil {
		t.Fatalf("Error reading config file %v", err)
	}

	err = viper.Unmarshal(&MyConfig)
	if err != nil {
		t.Fatalf("Error unmarshelling viper %v", err)
	}

	checkValues(t)
}
