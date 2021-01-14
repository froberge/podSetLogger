package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
)

const configFileName = "podset-logger-config.yaml"
const configFilePath = "/etc/config/"

// Read the config File
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

	fmt.Fprintln(w, "PodSet Logger V2")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "This is the PodSet Logger application that read a config file & deplay the informations found according to what we are looking for.")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Value read from the config file")
	fmt.Fprintln(w, "")

	// Great the table
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader([]string{"Key", "Value"})
	table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	table.SetCenterSeparator("|")
	table.SetRowLine(true)

	keys := getAllKeys()
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		table.Append([]string{strings.Title(key), getConfigVariable(key)})
	}

	table.Render()

	fmt.Fprintln(w, tableString.String())
}

func main() {

	log.Printf("Initializing the application")

	log.Fatalf("Simulate bad version and code crash.")

	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	handler := HttpHandler{}

	http.ListenAndServe(":"+PORT, handler)
}
