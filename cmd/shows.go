package cmd

import (
  "fmt"
  "os"
 // "io/ioutil"
//  "encoding/json" 
  "traktTV-cli/trakt"   

  "github.com/spf13/cobra"
  "github.com/hashicorp/vault/api"
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

    var token = os.Getenv("VAULT_DEV_ROOT_TOKEN_ID")
    var vault_addr = os.Getenv("VAULT_ADDR")
        
    config := &api.Config{
		Address: vault_addr,
	}
	clientVault, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}
    
	clientVault.SetToken(token)
    secret, err := clientVault.Logical().Read("secret/data/token")
	if err != nil {
		fmt.Println(err)
		return
	}
	m, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		fmt.Printf("%T %#v\n", secret.Data["data"], secret.Data["data"])
		return
	}  

    access_token := fmt.Sprintf("%v", m["access_token"])

    client := trakt.NewClient(
		""+client_id+"",
		trakt.TokenAuth{AccessToken: ""+ access_token +""},
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
