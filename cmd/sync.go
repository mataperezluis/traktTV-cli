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
	RootCmd.AddCommand(syncCmd)
	syncCmd.Flags().BoolP("extended", "e", false, "show extended info")
	syncCmd.SetHelpTemplate("use: sync [OPTIONS]\n"+
	"\navailable options:\n"+

	"  \nlastactivities"+
	"  \nplayback [type] [startDate] [endDate]"+
	"  \nplaybackremove [id]"+
	"  \ngetcollectionmovies [flag: extended]"+
	"  \ngetcollectionshows [flag: extended]"+
	"  \ngetwatchedshows [flag: extended]"+
	"  \n\n")

}

var syncCmd = &cobra.Command{
	Use:   "sync [OPTIONS]",
	Short: "Syncing with trakt ",
	Long:  `Syncing with trakt `,
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

		fstatus, _ := cmd.Flags().GetBool("extended")

		switch com := args[0]; com {

		
			case "lastactivities":
			
					showResults, err := client.Sync().LastActivities()
					if err != nil {
						fmt.Println(err)
					}
									

					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))

			case "playback":
				if len(args) > 3 {
						
					showResults, err := client.Sync().PlayBack(args[1],args[2],args[3])
					if err != nil {
						fmt.Println(err)
					}
					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: playback [type] [startDate] [endDate]")
				}

			case "playbackremove":
				if len(args) > 1 {
						
					_, err := client.Sync().PlayBackRemove(args[1])
					if err != nil {
						fmt.Println(err)
					}
					
				} else {
					fmt.Println("correct use: playbackremove [id]")
				}
			
			case "getcollectionmovies":
				
					extended:="false"
						
							if fstatus{
								extended="true"
							}
						
					showResults, err := client.Sync().GetCollectionMovies("movies",extended)
					if err != nil {
						fmt.Println(err)
					}
					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))

			case "getcollectionshows":
				
					extended:="false"
						
							if fstatus{
								extended="true"
							}
						
					showResults, err := client.Sync().GetCollectionShows("shows",extended)
					if err != nil {
						fmt.Println(err)
					}
					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))       
			
			case "getwatchedshows":
			
				extended:="false"
					
						if fstatus{
							extended="true"
						}
					
				showResults, err := client.Sync().GetCollectionShows("shows",extended)
				if err != nil {
					fmt.Println(err)
				}
				b, err2 := json.MarshalIndent(showResults, "", " ")
				if err2 != nil {
					fmt.Println(err2)
				}
				
				fmt.Println(string(b)) 
								

			default:
				fmt.Println("available commands:")
				fmt.Println("  lastactivities")
				fmt.Println("  playback [type] [startDate] [endDate]")
				fmt.Println("  playbackremove [id]")
				fmt.Println("  getcollectionmovies [flag: extended]")
				fmt.Println("  getcollectionshows [flag: extended]")
				fmt.Println("  getwatchedshows [flag: extended]")
				

			}

	},
}
