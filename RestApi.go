package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "os"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func fetchAllCurrency(w http.ResponseWriter,r *http.Request){
    
    response, err := http.Get("https://api.hitbtc.com/api/3/public/currency")
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
    fmt.Fprintf(w, string(responseData))
}

func fetchSingleCurrency(w http.ResponseWriter,r *http.Request){
    fmt.Println(r.Body)
     keys, ok := r.URL.Query()["symbol"]
    
    if !ok || len(keys[0]) < 1 {
        log.Println("Url Param 'symbol' is missing")
        return
    }

    // Query()["key"] will return an array of items, 
    // we only want the single item.
    symbol := keys[0]

    //log.Println("Url Param 'key' is: " + string(symbol))
    url:=fmt.Sprintf("%s%s","https://api.hitbtc.com/api/3/public/currency/",string(symbol))
    response, err := http.Get(url)
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
    fmt.Fprintf(w, string(responseData))
}

func handleRequests() {
    //http.HandleFunc("/", homePage)
    http.HandleFunc("/currency/all", fetchAllCurrency)
    http.HandleFunc("/currency/", fetchSingleCurrency)
    log.Fatal(http.ListenAndServe(":5001", nil))
}

func main() {
    handleRequests()
}