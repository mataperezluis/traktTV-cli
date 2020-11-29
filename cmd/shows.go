package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"traktTV-cli/trakt"
	"strconv"	

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(showsCmd)
}

var showsCmd = &cobra.Command{
	Use:   "shows [OPTIONS]",
	Short: "returns information about shows",
	Long:  `returns information about shows allpopular`,
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
				traktID,_:=strconv.Atoi(args[1])
				showResults, err := client.Shows().One(traktID)
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
				traktID,_:=strconv.Atoi(args[1])
				showResults, err := client.Shows().Alias(traktID)
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
				traktID,_:=strconv.Atoi(args[1])
				showResults, err := client.Shows().Certifications(traktID)
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
		}

	},
}
