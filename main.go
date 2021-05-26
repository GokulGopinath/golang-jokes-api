package main
import (
 "encoding/json"
 "fmt"
 "io/ioutil"
 "net/http"
)
type Response struct {   //made struct using https://mholt.github.io/json-to-go/
 ID     string `json:"id"`
 Joke   string `json:"joke"`
 Status int    `json:"status"`
}


func main() {
 fmt.Println("Calling API...")
client := &http.Client{}
 req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)  //creating http request 
 if err != nil {
  fmt.Print(err.Error())
 }
 req.Header.Add("Accept", "application/json")
 req.Header.Add("Content-Type", "application/json")
 resp, err := client.Do(req)  //sending the http request with resp, err := client.Do(req)
 if err != nil {
  fmt.Print(err.Error())
 }
defer resp.Body.Close()
 bodyBytes, err := ioutil.ReadAll(resp.Body)  //read the response
 if err != nil {
  fmt.Print(err.Error())
 }
var responseObject Response  //creating a variable of type response 
 json.Unmarshal(bodyBytes, &responseObject)   //convert json to data structures in GO  using unmarshalling
 fmt.Println("The API Response is \n", responseObject)
 fmt.Println("The joke is: ", responseObject.Joke)
}