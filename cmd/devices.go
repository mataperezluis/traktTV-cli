package cmd

import (
    "fmt"
    "bytes"
    "time"

    "net/http"
    "encoding/json"
    "io/ioutil"
    "github.com/spf13/cobra"
    "traktTV-cli/trakt"
    "github.com/tj/go-spin"
)

const client_id = "88f5df64ae395414edfa783e5a62eaf8718e79d42eee8fe12306db3dd343240e"
const client_secret = "3b2fdf87b8805e38ce74fe31819d46859b6bfc35aec8262cc7dccbf69ac7debb"


type authData struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationURL string `json:"verification_url"`
	ExpiresIn       int    `json:"expires_in"`
	Interval        int    `json:"interval"`
}

type tokenData struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	CreatedAt    int    `json:"created_at"`
}

func init() {
  RootCmd.AddCommand(devicesCmd)
}

func show(s *spin.Spinner) {
	for {
		fmt.Printf("\r  \033[36mwaiting for the token\033[m %s ", s.Next())
		time.Sleep(100 * time.Millisecond)
	}
}


var devicesCmd = &cobra.Command{
  Use:   "devices",
  Short: "initiates connection to traktTV API",
  Long:  `Device authentication is for apps and services with limited input or display capabilities`,
  Run: func(cmd *cobra.Command, args []string) {
    

    var jsonStr = []byte(`{"Accept": "application/json",
               "client_id":"`+ client_id+ `"}`)

	client := http.Client{}
	request, err := http.NewRequest("POST", trakt.TraktAPIURL + "/oauth/device/code",  bytes.NewBuffer(jsonStr))
    request.Header.Set("Content-Type", "application/json")	
    if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

    body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()


    var dat authData

    err2 := json.Unmarshal(body, &dat)
    if err2 != nil {
        fmt.Println(err2)
    }
    fmt.Println("user_code: " + dat.UserCode)
    fmt.Println("visit this verification url: " + dat.VerificationURL)
    fmt.Println("and insert the user_code to generate an access token")

    
//----------------get the token ------------------------------------------------------------------

    var jsonStrToken = []byte(`{
        "code": "` + dat.DeviceCode + `",
        "client_id": "`+ client_id+ `",
        "client_secret": "`+ client_secret+ `"
    }`)
    
    s := spin.New()
    ticker := time.NewTicker(time.Duration(dat.Interval) * time.Second)
    done := make(chan bool)

    
    go show(s)

go func() {
        for {  
            select {
            case <-done:
                return
            case <-ticker.C:
                clientToken := http.Client{}
	    requestToken, err := http.NewRequest("POST", trakt.TraktAPIURL + "/oauth/device/token",  bytes.NewBuffer(jsonStrToken))
        requestToken.Header.Set("Content-Type", "application/json")	
        if err != nil {
		    fmt.Println(err)
	    }

	    respToken, err := clientToken.Do(requestToken)
	    if err != nil {
		    fmt.Println(err)
	    }
  
        if respToken.StatusCode == 200 {

            bodyToken, _ := ioutil.ReadAll(respToken.Body)
	        respToken.Body.Close()

            var tokenDat tokenData

            err3 := json.Unmarshal(bodyToken, &tokenDat)
            if err3 != nil {
                fmt.Println(err3)
            }
            var jsonData []byte
            
            jsonData, err := json.Marshal(tokenDat)
            if err != nil {
                fmt.Println(err)
            }

            errj := ioutil.WriteFile("apidata.txt", jsonData, 0644)
	        if errj != nil {
		        fmt.Println(err)
	        }

            fmt.Println("access_token succesfully generated ")

            ticker.Stop()
            done <- true
        }
            }
        }
    }()


    timerExpire := time.NewTimer(time.Duration(dat.ExpiresIn) * time.Second)
    go func() {
        <-timerExpire.C
        done <- true
        fmt.Println("The time has expired repeat the operation")
    } ()
	
	<- done
    ticker.Stop()
    
    


  },
}
