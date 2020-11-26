package cmd

import (
  "fmt"
  "os"
  "io/ioutil"
  "encoding/json" 
  "traktTV-cli/trakt"   

  "github.com/spf13/cobra"
)

func init() {
  RootCmd.AddCommand(showsCmd)
}

var showsCmd = &cobra.Command{
  Use:   "shows [OPTIONS]",
  Short: "returns information about shows",
  Long:  `returns information about shows allpopular`,
  Args: cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
        
    jsonFile, err := os.Open("./apidata.txt")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    defer jsonFile.Close()

    var tokenDat tokenData

    byteValue, _ := ioutil.ReadAll(jsonFile)
    json.Unmarshal(byteValue, &tokenDat)

    

    client := trakt.NewClient(
		""+client_id+"",
		trakt.TokenAuth{AccessToken: ""+tokenDat.AccessToken+""},
	)

    if args[0] == "allpopular"{
        shows, err := client.Shows().AllPopular()
            if err != nil {
                fmt.Println(err)
            }
	    fmt.Println(shows)
    }

  },
}
