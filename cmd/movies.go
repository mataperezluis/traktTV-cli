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
	RootCmd.AddCommand(moviesCmd)
	moviesCmd.SetHelpTemplate("use: movies [OPTIONS]\n"+
	"\navailable options:\n"+
			"\n  allpopular"+
            "\n  trending"+
			"\n  search \"name of the show\""+
            "\n  recommended [period], periods:  daily , weekly , monthly , yearly , all"+
            "\n  played [period], periods:  daily , weekly , monthly , yearly , all"+
            "\n  watched [period], periods:  daily , weekly , monthly , yearly , all"+
			"\n  collected [period], periods:  daily , weekly , monthly , yearly , all"+
			"\n  updates [date_start], Example: 2020-11-27T00:00:00Z"+
			"\n  one [Trakt ID, Trakt slug, or IMDB ID]"+
			"\n  alias [Trakt ID, Trakt slug, or IMDB ID]"+
			"\n  boxoffice"+
			"\n  releases [Trakt ID, Trakt slug, or IMDB ID] [country]"+
			"\n  translations [Trakt ID, Trakt slug, or IMDB ID] [language]"+
			"\n  comments [Trakt ID, Trakt slug, or IMDB ID] [sort]"+
			"\n  lists [Trakt ID, Trakt slug, or IMDB ID] [type] [sort]"+
			"\n  people [Trakt ID, Trakt slug, or IMDB ID]"+
			"\n  ratings [Trakt ID, Trakt slug, or IMDB ID]"+
			"\n  stats [Trakt ID, Trakt slug, or IMDB ID]"+
			"\n  watching [Trakt ID, Trakt slug, or IMDB ID]"+
			"\n  related [Trakt ID, Trakt slug, or IMDB ID]\n")


}

var moviesCmd = &cobra.Command{
	Use:   "movies [OPTIONS]",
	Short: "returns information about movies",
	Long:  `returns information about movies`,
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

		case "allpopular":
			shows, err := client.Movies().AllPopular()
			if err != nil {
				fmt.Println(err)
			}

              
                b, err2 := json.MarshalIndent(shows, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
			    fmt.Println(string(b))

        case "trending":
			shows, err := client.Movies().Trending()
			if err != nil {
				fmt.Println(err)
			}

              
                b, err2 := json.MarshalIndent(shows, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
				fmt.Println(string(b))
		case "boxoffice":
			shows, err := client.Movies().BoxOffice()
			if err != nil {
				fmt.Println(err)
			}

				
				b, err2 := json.MarshalIndent(shows, "", " ")
				if err2 != nil {
					fmt.Println(err2)
				}
				
				fmt.Println(string(b))	

        case "recommended":
			if !(len(args) > 1) {
                args=append(args, "weekly")     
			}
            if args[1]!="daily" && args[1]!="weekly" && args[1]!="monthly" && args[1]!="yearly" && args[1]!="all"{
                    args[1]="weekly"    
            }
			    showResults, err := client.Movies().Recommended(args[1])
			    if err != nil {
				    fmt.Println(err)
			    }
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
			        fmt.Println(err2)
		        }
                
		        fmt.Println(string(b))                
                    

        case "played":
			
            if !(len(args) > 1) {
				 args=append(args, "weekly")          
			}
            if args[1]!="daily" && args[1]!="weekly" && args[1]!="monthly" && args[1]!="yearly" && args[1]!="all"{
                    args[1]="weekly"    
            }

				    showResults, err := client.Movies().Played(args[1])
				    if err != nil {
					    fmt.Println(err)
				    }
                    b, err2 := json.MarshalIndent(showResults, "", " ")
                    if err2 != nil {
				        fmt.Println(err2)
			        }
                    
			        fmt.Println(string(b))                
                        


        case "watched":
			if !(len(args) > 1) {
				 args=append(args, "weekly")         
			}
            if args[1]!="daily" && args[1]!="weekly" && args[1]!="monthly" && args[1]!="yearly" && args[1]!="all"{
                    args[1]="weekly"    
            }
				    showResults, err := client.Movies().Watched(args[1])
				    if err != nil {
					    fmt.Println(err)
				    }
                    b, err2 := json.MarshalIndent(showResults, "", " ")
                    if err2 != nil {
				        fmt.Println(err2)
			        }
                    
			        fmt.Println(string(b))                
                        

        case "collected":
			if !(len(args) > 1) {
				 args=append(args, "weekly")          
			}
            if args[1]!="daily" && args[1]!="weekly" && args[1]!="monthly" && args[1]!="yearly" && args[1]!="all"{
                    args[1]="weekly"    
            }
				    showResults, err := client.Movies().Collected(args[1])
				    if err != nil {
					    fmt.Println(err)
				    }
                    b, err2 := json.MarshalIndent(showResults, "", " ")
                    if err2 != nil {
				        fmt.Println(err2)
			        }
                    
			        fmt.Println(string(b))                
                        

        case "anticipated":
			shows, err := client.Movies().Anticipated()
			if err != nil {
				fmt.Println(err)
			}

              
                b, err2 := json.MarshalIndent(shows, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
				fmt.Println(string(b))
				
		case "updates":
			if len(args) > 1 {
				showResults, err := client.Movies().Updates(args[1])
				if err != nil {
					fmt.Println(err)
				}
				b, err2 := json.MarshalIndent(showResults, "", " ")
				if err2 != nil {
					fmt.Println(err2)
				}
				
				fmt.Println(string(b))                
					
			} else {
				fmt.Println("correct use: updates date time (example:2020-11-27T00:00:00Z)")
			}
		case "updates-id":
			if len(args) > 1 {
				showResults, err := client.Movies().UpdatesId(args[1])
				if err != nil {
					fmt.Println(err)
				}
				b, err2 := json.MarshalIndent(showResults, "", " ")
				if err2 != nil {
					fmt.Println(err2)
				}
				
				fmt.Println(string(b))                
					
			} else {
				fmt.Println("correct use: updates-id date/time (example:2020-11-27T00:00:00Z)")
			}

		case "one":
			if len(args) > 1 {
				showResults, err := client.Movies().One(args[1])
				if err != nil {
					fmt.Println(err)
				}
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
			    fmt.Println(string(b))                
                    
			} else {
				fmt.Println("correct use: one [Trakt ID, Trakt slug, or IMDB ID]")
			}
		
		case "alias":
			if len(args) > 1 {
				showResults, err := client.Movies().Alias(args[1])
				if err != nil {
					fmt.Println(err)
				}
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
			    fmt.Println(string(b))                
                    
			} else {
				fmt.Println("correct use: alias [Trakt ID, Trakt slug, or IMDB ID]")
			}

		case "translations":
			if len(args) > 1 {
				languageData:=""
				if len(args) > 2{
					languageData=args[2]
				}
				showResults, err := client.Movies().Translations(args[1],languageData)
				if err != nil {
					fmt.Println(err)
				}
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
			    fmt.Println(string(b))                
                    
			} else {
				fmt.Println("correct use: translations [Trakt ID, Trakt slug, or IMDB ID] [language]")
			}
		
		case "releases":
			if len(args) > 1 {
				countryData:=""
				if len(args) > 2{
					countryData=args[2]
				}
				showResults, err := client.Movies().Releases(args[1],countryData)
				if err != nil {
					fmt.Println(err)
				}
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
			    fmt.Println(string(b))                
                    
			} else {
				fmt.Println("correct use: releases [Trakt ID, Trakt slug, or IMDB ID] [country]")
			}

		case "comments":
			if len(args) > 1 {
				
				if !(len(args) > 2) {
					args=append(args, "newest")          
				}
				if args[2]!="newest" && args[2]!="oldest" && args[2]!="likes" &&  args[2]!="replies" &&  args[2]!="highest" && args[2]!="lowest" && args[2]!="plays" && args[2]!="watched"{
						args[2]="newest"    
				}
				showResults, err := client.Movies().Comments(args[1],args[2])
				if err != nil {
					fmt.Println(err)
				}
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
			    fmt.Println(string(b))                
                    
			} else {
				fmt.Println("correct use: comments [Trakt ID, Trakt slug, or IMDB ID] [sort]")
			}
		
		case "lists":
			if len(args) > 1 {
				if !(len(args) > 2) {
					args=append(args, "personal")          
				}
				if !(len(args) > 3) {
					args=append(args, "popular")          
				}
				if args[2]!="personal" && args[2]!="official" && args[2]!="watchlists" &&  args[2]!="recommendations" &&  args[2]!="all"{
						args[2]="personal"    
				}
				if args[3]!="likes" && args[3]!="likes" && args[3]!="comments" &&  args[3]!="items" &&  args[3]!="added" &&  args[3]!="updated"{
					args[3]="popular"    
				}	
				showResults, err := client.Movies().List(args[1],args[2],args[3])
				if err != nil {
					fmt.Println(err)
				}
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
			    fmt.Println(string(b))                
                    
			} else {
				fmt.Println("correct use: lists [Trakt ID, Trakt slug, or IMDB ID] [type] [sort]")
			}
		

		case "people":
			if len(args) > 1 {
				
				showResults, err := client.Movies().People(args[1])
				if err != nil {
					fmt.Println(err)
				}
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
			    fmt.Println(string(b))                
                    
			} else {
				fmt.Println("correct use: people [Trakt ID, Trakt slug, or IMDB ID]")
			}
		
		case "ratings":
			if len(args) > 1 {
				showResults, err := client.Movies().Ratings(args[1])
				if err != nil {
					fmt.Println(err)
				}
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
			    fmt.Println(string(b))                
                    
			} else {
				fmt.Println("correct use: ratings [Trakt ID, Trakt slug, or IMDB ID]")
			}
		case "related":
			if len(args) > 1 {
				showResults, err := client.Movies().Related(args[1])
				if err != nil {
					fmt.Println(err)
				}
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
			    fmt.Println(string(b))                
                    
			} else {
				fmt.Println("correct use: related [Trakt ID, Trakt slug, or IMDB ID]")
			}
		
		case "stats":
			if len(args) > 1 {
				showResults, err := client.Movies().Stats(args[1])
				if err != nil {
					fmt.Println(err)
				}
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
			    fmt.Println(string(b))                
                    
			} else {
				fmt.Println("correct use: stats [Trakt ID, Trakt slug, or IMDB ID]")
			}
		
		case "watching":
			if len(args) > 1 {
				showResults, err := client.Movies().Watching(args[1])
				if err != nil {
					fmt.Println(err)
				}
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
			    fmt.Println(string(b))                
                    
			} else {
				fmt.Println("correct use: watching [Trakt ID, Trakt slug, or IMDB ID]")
			}

		case "search":
			if len(args) > 1 {
				showResults, err := client.Movies().Search(args[1])
				if err != nil {
					fmt.Println(err)
				}
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
			    fmt.Println(string(b))                
                    
			} else {
				fmt.Println("correct use: search \"name of the show\"")
			}
		default:
			fmt.Println("available commands:")
			fmt.Println("  allpopular")
            fmt.Println("  trending")
			fmt.Println("  search \"name of the show\"")
            fmt.Println("  recommended [period], periods:  daily , weekly , monthly , yearly , all")
            fmt.Println("  played [period], periods:  daily , weekly , monthly , yearly , all")
            fmt.Println("  watched [period], periods:  daily , weekly , monthly , yearly , all")
			fmt.Println("  collected [period], periods:  daily , weekly , monthly , yearly , all")
			fmt.Println("  updates [date_start], Example: 2020-11-27T00:00:00Z")	
			fmt.Println("  one [Trakt ID, Trakt slug, or IMDB ID]")
			fmt.Println("  alias [Trakt ID, Trakt slug, or IMDB ID]")
			fmt.Println("  boxoffice")
			fmt.Println("  releases [Trakt ID, Trakt slug, or IMDB ID] [country]")	
			fmt.Println("  translations [Trakt ID, Trakt slug, or IMDB ID] [language]")	
			fmt.Println("  comments [Trakt ID, Trakt slug, or IMDB ID] [sort]")
			fmt.Println("  lists [Trakt ID, Trakt slug, or IMDB ID] [type] [sort]")
			fmt.Println("  people [Trakt ID, Trakt slug, or IMDB ID]")
			fmt.Println("  ratings [Trakt ID, Trakt slug, or IMDB ID]")
			fmt.Println("  stats [Trakt ID, Trakt slug, or IMDB ID]")
			fmt.Println("  watching [Trakt ID, Trakt slug, or IMDB ID]")
			fmt.Println("  related [Trakt ID, Trakt slug, or IMDB ID]")

		}

	},
}
