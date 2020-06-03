package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

//const configFile = "/etc/config/configmap-podsetlogger.yaml"
const configFileName = "podset-logger-config.yaml"
const configFilePath = "/etc/config/"

// Rad the config File
func readConfigFile() {
	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory

	// Read the config file
	viper.SetConfigName(configFileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configFilePath)
	viper.AddConfigPath(".")

	// Monitor the change to the config file
	viper.WatchConfig()
	//	viper.OnConfigChange(func(e fsnotify.Event) {
	//		log.Printf("Config file changed:", e.Name)
	//	})

	// Find and read the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}

// Here we read only string
func getAllKeys() []string {
	return viper.AllKeys()
}

// We only enable the reading of string
func getConfigVariable(key string) string {
	value, ok := viper.Get(key).(string)
	// If the type is a string then ok will be true
	if !ok && len(value) > 0 {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

// HttpHandler the handler struct
type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	readConfigFile()

	fmt.Fprintln(w, "This is the PodSet Logger application that read a config file & deplay the informations.")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Value read from the config file:")

	keys := getAllKeys()
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		fmt.Fprintln(w, "\t - "+strings.Title(key)+": "+getConfigVariable(key))
	}
}

func main() {

	log.Printf("Initializing the application")
	// Add this line to create the image that fail's at startup
	//	log.Fatalf("BAD VERSION POD CRASH SHOULD ROLE BACK")

	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	handler := HttpHandler{}

	http.ListenAndServe(":"+PORT, handler)
}
