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

    
    switch com := args[0]; com {
   
    case "allpopular":
        shows, err := client.Shows().AllPopular()
            if err != nil {
                fmt.Println(err)
            }
	    fmt.Println(shows)
    case "search":
        if len(args) > 1 {
        showResults, err := client.Shows().Search(args[1])
            if err != nil {
                fmt.Println(err)
            }
        for _, showResult := range showResults {
            fmt.Println(showResult.Show)
        }
        }else
        {   
          fmt.Println("correct use: search \"name of the show\"")  
        }
    default:
        fmt.Println("commands")
        fmt.Println("allpopular")
        fmt.Println("search \"name of the show\"")
    }

  },
}
