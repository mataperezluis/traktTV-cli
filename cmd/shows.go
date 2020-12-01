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
	RootCmd.AddCommand(showsCmd)
}

var showsCmd = &cobra.Command{
	Use:   "shows [OPTIONS]",
	Short: "returns information about shows",
	Long:  `returns information about shows`,
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
			shows, err := client.Shows().AllPopular()
			if err != nil {
				fmt.Println(err)
			}

              
                b, err2 := json.MarshalIndent(shows, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
			    fmt.Println(string(b))

        case "trending":
			shows, err := client.Shows().Trending()
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
			    showResults, err := client.Shows().Recommended(args[1])
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

				    showResults, err := client.Shows().Played(args[1])
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
				    showResults, err := client.Shows().Watched(args[1])
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
				    showResults, err := client.Shows().Collected(args[1])
				    if err != nil {
					    fmt.Println(err)
				    }
                    b, err2 := json.MarshalIndent(showResults, "", " ")
                    if err2 != nil {
				        fmt.Println(err2)
			        }
                    
			        fmt.Println(string(b))                
                        

        case "anticipated":
			shows, err := client.Shows().Anticipated()
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
				showResults, err := client.Shows().Updates(args[1])
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
				showResults, err := client.Shows().UpdatesId(args[1])
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
				showResults, err := client.Shows().One(args[1])
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
				showResults, err := client.Shows().Alias(args[1])
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

		case "certifications":
			if len(args) > 1 {
				showResults, err := client.Shows().Certifications(args[1])
				if err != nil {
					fmt.Println(err)
				}
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
			    fmt.Println(string(b))                
                    
			} else {
				fmt.Println("correct use: certifications [Trakt ID, Trakt slug, or IMDB ID]")
			}
		case "translations":
			if len(args) > 1 {
				languageData:=""
				if len(args) > 2{
					languageData=args[2]
				}
				showResults, err := client.Shows().Translations(args[1],languageData)
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

		case "comments":
			if len(args) > 1 {
				
				if !(len(args) > 2) {
					args=append(args, "newest")          
				}
				if args[2]!="newest" && args[2]!="oldest" && args[2]!="likes" &&  args[2]!="replies" &&  args[2]!="highest" && args[2]!="lowest" && args[2]!="plays" && args[2]!="watched"{
						args[2]="newest"    
				}
				showResults, err := client.Shows().Comments(args[1],args[2])
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
				showResults, err := client.Shows().List(args[1],args[2],args[3])
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
		case "collection-progress":
			if len(args) > 1 {
				hidden:="false"
				specials:="false"
				count_specials:="false"

				for _,x:= range args{
					if x == "hidden"{
						hidden="true"
					}else if x == "specials"{
						specials="true"
					}else if x == "count_specials"{
						count_specials="true"
					}

				}

				showResults, err := client.Shows().CollectionProgress(args[1],hidden,specials,count_specials)
				if err != nil {
					fmt.Println(err)
				}
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
				fmt.Println(string(b))                
                    
			} else {
				fmt.Println("correct use: collection-progress [Trakt ID, Trakt slug, or IMDB ID] optionals: ")
				fmt.Println("hidden specials count_specials")
			}
		
		case "watched-progress":
			if len(args) > 1 {
				hidden:="false"
				specials:="false"
				count_specials:="false"

				for _,x:= range args{
					if x == "hidden"{
						hidden="true"
					}else if x == "specials"{
						specials="true"
					}else if x == "count_specials"{
						count_specials="true"
					}

				}

				showResults, err := client.Shows().WatchedProgress(args[1],hidden,specials,count_specials)
				if err != nil {
					fmt.Println(err)
				}
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
				fmt.Println(string(b))                
                    
			} else {
				fmt.Println("correct use: watched-progress [Trakt ID, Trakt slug, or IMDB ID] optionals: ")
				fmt.Println("hidden specials count_specials")
			}
		

		case "people":
			if len(args) > 1 {
				extended:="false"
					
					if len(args) > 2{
						if args[2]=="extended"{
							extended="true"
						}
					}
				showResults, err := client.Shows().People(args[1],extended)
				if err != nil {
					fmt.Println(err)
				}
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
			    fmt.Println(string(b))                
                    
			} else {
				fmt.Println("correct use: people [Trakt ID, Trakt slug, or IMDB ID] [optional: extended]")
			}
		
		case "ratings":
			if len(args) > 1 {
				showResults, err := client.Shows().Ratings(args[1])
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
				showResults, err := client.Shows().Related(args[1])
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
				showResults, err := client.Shows().Stats(args[1])
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
				showResults, err := client.Shows().Watching(args[1])
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
		
		case "next-episode":
			if len(args) > 1 {
				showResults, err := client.Shows().NextEpisode(args[1])
				if err != nil {
					fmt.Println(err)
				}
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
			    fmt.Println(string(b))                
                    
			} else {
				fmt.Println("correct use: next-episode [Trakt ID, Trakt slug, or IMDB ID]")
			}
		case "last-episode":
			if len(args) > 1 {
				showResults, err := client.Shows().LastEpisode(args[1])
				if err != nil {
					fmt.Println(err)
				}
                b, err2 := json.MarshalIndent(showResults, "", " ")
                if err2 != nil {
				    fmt.Println(err2)
			    }
                
			    fmt.Println(string(b))                
                    
			} else {
				fmt.Println("correct use: last-episode [Trakt ID, Trakt slug, or IMDB ID]")
			}



		case "search":
			if len(args) > 1 {
				showResults, err := client.Shows().Search(args[1])
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
			fmt.Println("  certifications [Trakt ID, Trakt slug, or IMDB ID]")
			fmt.Println("  translations [Trakt ID, Trakt slug, or IMDB ID] [language]")	
			fmt.Println("  comments [Trakt ID, Trakt slug, or IMDB ID] [sort]")
			fmt.Println("  lists [Trakt ID, Trakt slug, or IMDB ID] [type] [sort]")
			fmt.Println("  people [Trakt ID, Trakt slug, or IMDB ID] [optional: extended]")
			fmt.Println("  ratings [Trakt ID, Trakt slug, or IMDB ID]")
			fmt.Println("  stats [Trakt ID, Trakt slug, or IMDB ID]")
			fmt.Println("  watching [Trakt ID, Trakt slug, or IMDB ID]")
			fmt.Println("  related [Trakt ID, Trakt slug, or IMDB ID]")
			fmt.Println("  collection-progress [Trakt ID, Trakt slug, or IMDB ID] optionals: hidden specials count_specials")
			fmt.Println("  watched-progress [Trakt ID, Trakt slug, or IMDB ID] optionals: hidden specials count_specials")
			fmt.Println("  next-episode [Trakt ID, Trakt slug, or IMDB ID]")
			fmt.Println("  last-episode [Trakt ID, Trakt slug, or IMDB ID]")
		}

	},
}
