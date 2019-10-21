package configs

import (
	"os"

	"github.com/lbrulet/APINIT-GO/src/models"
	"github.com/lbrulet/gonfig"
)

// Config is a global variable to get informations from the Configuration structure
var Config models.Configuration

// InitConfig is a function that fetch information from a file / environnement
func InitConfig() {
	if pwd, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		if os.Getenv("ENVIRONMENT") == "LOCAL" {
			if err := gonfig.GetConf(pwd+"/src/configs/local/config.json", &Config); err != nil {
				panic(err)
			}
		} else {
			if err := gonfig.GetConf(pwd+"/src/configs/dev/config.json", &Config); err != nil {
				panic(err)
			}
		}
	}
}
