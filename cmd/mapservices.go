/*
Copyright © 2023 Rudolf Farkas rudi.farkas@gmail.com
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const url = "https://api3.geo.admin.ch/rest/services/api/MapServer"

// mapservicesCmd represents the mapservices command
var mapservicesCmd = &cobra.Command{
	Use:   "mapservices",
	Short: "Returns a list of map services",
	Long: `Gets a list of map services from the Swiss Federal Office of Topography
` + `  (https://api3.geo.admin.ch/rest/services/api/MapServer)
` + `  as a JSON string and prints the list to stdout.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(os.Stderr, "=== mapservices response:")
		compact, _ := cmd.Flags().GetBool("compact")
		jsonStr := GetMapServerInfo(compact)
		fmt.Println(jsonStr)
	},
}

func init() {
	rootCmd.AddCommand(mapservicesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mapservicesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	mapservicesCmd.Flags().BoolP("compact", "c", false, "Print compact JSON")
}

// GetMapServerInfo returns a list of map services
func GetMapServerInfo(compact bool) string {

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if compact {
		return string(body)
	} else {
		ppBody, err := Prettyfmt(string(body))
		if err != nil {
			log.Fatal(err)
		}
		return ppBody
	}
}

// TODO: move this to a util package

func Prettyfmt(input string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(input), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
