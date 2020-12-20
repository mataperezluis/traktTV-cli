//Package cmd ...
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"traktTV-cli/trakt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(calendarCmd)
	calendarCmd.SetHelpTemplate("use: calendars [OPTIONS]\n"+
	"\navailable options:\n"+

	"  \nmyshows [start date] [days]"+
	"  \nmynewshows [start date] [days]"+
	"  \nmyseasonpremiere [start date] [days]"+
	"  \nmymovies [start date] [days]"+
	"  \nmydvd [start date] [days]"+
	"  \nallshows [start date] [days]"+
	"  \nallnewshows [start date] [days]"+
	"  \nallpremieres [start date] [days]"+
	"  \nallmovies [start date] [days]"+
	"  \nalldvd [start date] [days]"+
	"  \n\n")

}

var calendarCmd = &cobra.Command{
	Use:   "calendars [OPTIONS]",
	Short: "returns information about all shows or movies for the specified time period ",
	Long:  `returns information about all shows or movies for the specified time period `,
	Args:  cobra.MinimumNArgs(1),
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
			""+clientID+"",
			trakt.TokenAuth{AccessToken: "" + tokenDat.AccessToken + ""},
		)


		switch com := args[0]; com {

			case "myshows":
				if len(args) > 2 {
					
					showResults, err := client.Calendars().MyShows(args[1],args[2])
					if err != nil {
						fmt.Println(err)
					}
									

					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: myshows [start date] [days]")
				}
			case "mynewshows":
				if len(args) > 2 {
					
					showResults, err := client.Calendars().MyNewShows(args[1],args[2])
					if err != nil {
						fmt.Println(err)
					}
									

					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: mynewshows [start date] [days]")
				}
			case "myseasonpremiere":
				if len(args) > 2 {
					
					showResults, err := client.Calendars().MySeasonPremiere(args[1],args[2])
					if err != nil {
						fmt.Println(err)
					}
									

					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: myseasonpremiere [start date] [days]")
				}
			case "mymovies":
				if len(args) > 2 {
					
					showResults, err := client.Calendars().MyMovies(args[1],args[2])
					if err != nil {
						fmt.Println(err)
					}
									

					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: mymovies [start date] [days]")
				}
			case "mydvd":
				if len(args) > 2 {
					
					showResults, err := client.Calendars().MyDVD(args[1],args[2])
					if err != nil {
						fmt.Println(err)
					}
									

					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: mydvd [start date] [days]")
				}
			
			case "allshows":
				if len(args) > 2 {
					
					showResults, err := client.Calendars().AllShows(args[1],args[2])
					if err != nil {
						fmt.Println(err)
					}
									

					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: allshows [start date] [days]")
				}
			case "allnewshows":
				if len(args) > 2 {
					
					showResults, err := client.Calendars().AllNewShows(args[1],args[2])
					if err != nil {
						fmt.Println(err)
					}
									

					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: allnewshows [start date] [days]")
				}
			case "allpremieres":
				if len(args) > 2 {
					
					showResults, err := client.Calendars().AllPremieres(args[1],args[2])
					if err != nil {
						fmt.Println(err)
					}
									

					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: allpremieres [start date] [days]")
				}
			case "allmovies":
				if len(args) > 2 {
					
					showResults, err := client.Calendars().AllMovies(args[1],args[2])
					if err != nil {
						fmt.Println(err)
					}
									

					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: allmovies [start date] [days]")
				}
			case "alldvd":
				if len(args) > 2 {
					
					showResults, err := client.Calendars().AllDVD(args[1],args[2])
					if err != nil {
						fmt.Println(err)
					}
									

					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: alldvd [start date] [days]")
				}
			
			
			

			default:
				fmt.Println("available commands:")
				fmt.Println("  myshows [start date] [days]")
				fmt.Println("  mynewshows [start date] [days]")
				fmt.Println("  myseasonpremiere [start date] [days]")
				fmt.Println("  mymovies [start date] [days]")
				fmt.Println("  mydvd [start date] [days]")
				fmt.Println("  allshows [start date] [days]")
				fmt.Println("  allnewshows [start date] [days]")
				fmt.Println("  allpremieres [start date] [days]")
				fmt.Println("  allmovies [start date] [days]")
				fmt.Println("  alldvd [start date] [days]")

			}

	},
}
