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
	RootCmd.AddCommand(episodesCmd)
	episodesCmd.Flags().BoolP("extended", "e", false, "show extended info")
	episodesCmd.SetHelpTemplate("use: episodes [OPTIONS]\n"+
	"\navailable options:\n"+

				"\none [Trakt ID, Trakt slug, or IMDB ID] [number of the season] [number of the episode]"+
				"\ntranslations [Trakt ID, Trakt slug, or IMDB ID] [number of the season] [number of the episode] [language]"+
				"\ncomments [Trakt ID, Trakt slug, or IMDB ID] [season] [number of the episode] [sort]"+
				"\nlists [Trakt ID, Trakt slug, or IMDB ID] [season] [number of the episode] [type] [sort]"+
				"\npeople [Trakt ID, Trakt slug, or IMDB ID] [season] [number of the episode] [flag: extended]"+
				"\nratings [Trakt ID, Trakt slug, or IMDB ID] [season] [number of the episode]"+
				"\nstats [Trakt ID, Trakt slug, or IMDB ID] [season] [number of the episode]"+
				"\nwatching [Trakt ID, Trakt slug, or IMDB ID] [season] [number of the episode] \n")
}

var episodesCmd = &cobra.Command{
	Use:   "episodes [OPTIONS]",
	Short: "returns information about episodes",
	Long:  `returns information about episodes`,
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

			case "one":
				if len(args) > 3 {
					
					showResults, err := client.Episodes().ByNumber(args[1],args[2],args[3])
					if err != nil {
						fmt.Println(err)
					}
									

					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: one [Trakt ID, Trakt slug, or IMDB ID] [number of the season] [number of the episode]")
				}
			case "translations":
				if len(args) > 3 {
					languageData:=""
					if len(args) > 4{
						languageData=args[4]
					}
					showResults, err := client.Episodes().Translations(args[1],args[2],args[3],languageData)
					if err != nil {
						fmt.Println(err)
					}
					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: translations [Trakt ID, Trakt slug, or IMDB ID] [number of the season] [number of the episode] [language]")
				}
			case "comments":
				if len(args) > 3 {
					
					if !(len(args) > 4) {
						args=append(args, "newest")          
					}
					if args[4]!="newest" && args[4]!="oldest" && args[4]!="likes" &&  args[4]!="replies" &&  args[4]!="highest" && args[4]!="lowest" && args[4]!="plays" && args[4]!="watched"{
							args[4]="newest"    
					}
					showResults, err := client.Episodes().Comments(args[1],args[2],args[3],args[4])
					if err != nil {
						fmt.Println(err)
					}
					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: comments [Trakt ID, Trakt slug, or IMDB ID] [season] [number of the episode] [sort]")
				}
			case "lists":
				if len(args) > 3 {
					if !(len(args) > 4) {
						args=append(args, "personal")          
					}
					if !(len(args) > 5) {
						args=append(args, "popular")          
					}
					if args[4]!="personal" && args[4]!="official" && args[4]!="watchlists" &&  args[4]!="recommendations" &&  args[4]!="all"{
							args[4]="personal"    
					}
					if args[5]!="likes" && args[5]!="likes" && args[5]!="comments" &&  args[5]!="items" &&  args[5]!="added" &&  args[5]!="updated"{
						args[5]="popular"    
					}	
					showResults, err := client.Episodes().List(args[1],args[2],args[3],args[4],args[5])
					if err != nil {
						fmt.Println(err)
					}
					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: lists [Trakt ID, Trakt slug, or IMDB ID] [season] [number of the episode] [type] [sort]")
				}

			case "people":
				if len(args) > 3 {
					extended:="false"
					
					
						if fstatus{
							extended="true"
						}
					


					showResults, err := client.Episodes().People(args[1],args[2],args[3],extended)
					if err != nil {
						fmt.Println(err)
					}
					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: people [Trakt ID, Trakt slug, or IMDB ID] [season] [number of the episode] [optional: extended]")
				}
			
			case "ratings":
				if len(args) > 3 {
					showResults, err := client.Episodes().Ratings(args[1],args[2],args[3])
					if err != nil {
						fmt.Println(err)
					}
					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: ratings [Trakt ID, Trakt slug, or IMDB ID] [season] [number of the episode]")
				}
			
			case "stats":
				if len(args) > 3 {
					showResults, err := client.Episodes().Stats(args[1],args[2],args[3])
					if err != nil {
						fmt.Println(err)
					}
					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: stats [Trakt ID, Trakt slug, or IMDB ID] [season] [number of the episode]")
				}
			
			case "watching":
				if len(args) > 3{
					showResults, err := client.Episodes().Watching(args[1],args[2],args[3])
					if err != nil {
						fmt.Println(err)
					}
					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: watching [Trakt ID, Trakt slug, or IMDB ID] [season] [number of the episode]")
				}
			

			default:
				fmt.Println("available commands:")
				fmt.Println("  one [Trakt ID, Trakt slug, or IMDB ID] [number of the season] [number of the episode]")
				fmt.Println("  translations [Trakt ID, Trakt slug, or IMDB ID] [number of the season] [number of the episode] [language]")
				fmt.Println("  comments [Trakt ID, Trakt slug, or IMDB ID] [season] [number of the episode] [sort]")
				fmt.Println("  lists [Trakt ID, Trakt slug, or IMDB ID] [season] [number of the episode] [type] [sort]")
				fmt.Println("  people [Trakt ID, Trakt slug, or IMDB ID] [season] [number of the episode] [flag: extended]")
				fmt.Println("  ratings [Trakt ID, Trakt slug, or IMDB ID] [season] [number of the episode]")
				fmt.Println("  stats [Trakt ID, Trakt slug, or IMDB ID] [season] [number of the episode]")
				fmt.Println("  watching [Trakt ID, Trakt slug, or IMDB ID] [season] [number of the episode]")

			}

	},
}
