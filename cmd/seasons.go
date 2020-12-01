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
	RootCmd.AddCommand(seasonsCmd)
}

var seasonsCmd = &cobra.Command{
	Use:   "seasons [OPTIONS]",
	Short: "returns information about seasons",
	Long:  `returns information about seasons`,
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
			""+client_id+"",
			trakt.TokenAuth{AccessToken: "" + tokenDat.AccessToken + ""},
		)

		switch com := args[0]; com {

		
			case "all":
				if len(args) > 1 {
					extension:=""


					for _,x:= range args{
						if x == "episodes"{
							extension="episodes"
							break;
						}else if x == "full"{
							extension="full"
							break;
						}
					}

					showResults, err := client.Seasons().All(args[1],extension)
					if err != nil {
						fmt.Println(err)
					}
									

					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: all [Trakt ID, Trakt slug, or IMDB ID] [extension optional: episodes, full]")
				}
			case "one":
				if len(args) > 2 {
					
					showResults, err := client.Seasons().ByNumber(args[1],args[2])
					if err != nil {
						fmt.Println(err)
					}
									

					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: one [Trakt ID, Trakt slug, or IMDB ID] [number of the season]")
				}
			case "comments":
				if len(args) > 2 {
					
					if !(len(args) > 3) {
						args=append(args, "newest")          
					}
					if args[3]!="newest" && args[3]!="oldest" && args[3]!="likes" &&  args[3]!="replies" &&  args[3]!="highest" && args[3]!="lowest" && args[3]!="plays" && args[3]!="watched"{
							args[3]="newest"    
					}
					showResults, err := client.Seasons().SeasonComments(args[1],args[2],args[3])
					if err != nil {
						fmt.Println(err)
					}
					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: comments [Trakt ID, Trakt slug, or IMDB ID] [season] [sort]")
				}
			case "lists":
				if len(args) > 2 {
					if !(len(args) > 3) {
						args=append(args, "personal")          
					}
					if !(len(args) > 4) {
						args=append(args, "popular")          
					}
					if args[3]!="personal" && args[3]!="official" && args[3]!="watchlists" &&  args[3]!="recommendations" &&  args[3]!="all"{
							args[3]="personal"    
					}
					if args[4]!="likes" && args[4]!="likes" && args[4]!="comments" &&  args[4]!="items" &&  args[4]!="added" &&  args[4]!="updated"{
						args[4]="popular"    
					}	
					showResults, err := client.Seasons().SeasonList(args[1],args[2],args[3],args[4])
					if err != nil {
						fmt.Println(err)
					}
					b, err2 := json.MarshalIndent(showResults, "", " ")
					if err2 != nil {
						fmt.Println(err2)
					}
					
					fmt.Println(string(b))                
						
				} else {
					fmt.Println("correct use: lists [Trakt ID, Trakt slug, or IMDB ID] [season] [type] [sort]")
				}
			

			default:
				fmt.Println("available commands:")
				fmt.Println("  all [Trakt ID, Trakt slug, or IMDB ID] [extension optional: episodes, full]")
				fmt.Println("  one [Trakt ID, Trakt slug, or IMDB ID] [number of the season]")
				fmt.Println("  comments [Trakt ID, Trakt slug, or IMDB ID] [season] [sort]")
				fmt.Println("  lists [Trakt ID, Trakt slug, or IMDB ID] [season] [type] [sort]")
			}

	},
}