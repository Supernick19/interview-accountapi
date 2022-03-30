package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "github.com/satori/go.uuid" // to generate a random uuid
)

type Account struct {
    Data            *AccountData    `json:"data,omitempty"`
    Links           *Link           `json:"links,omitempty"`
}

type Link struct {
    First           string      `json:"first,omitempty"`
    Last            string      `json:"last,omitempty"`
    Self            string      `json:"self,omitempty"`
}

type AccountData struct {
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
    Attributes     *AccountAttributes `json:"attributes,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

func main() {
    for i := 0; i < 3; i++ {
        createAccount() //create an account
    }

  acctId, err := createAccount() // create one more account
  fetchAccounts() // fetch the accounts

  // if account successfully created, fetch then delete it and show the results
  if err != nil {
      return
  }
  res :=  fetchAccount(acctId)
  fmt.Println(res)
  deleteAccount(acctId) 
  
  // fetch again to show the difference

}

// deletes the account with the specified ID
func deleteAccount(id string) {
    url := "http://localhost:8080/v1/organisation/accounts/" + id + "?version=0"
    method := "DELETE"

      payload := strings.NewReader(``)

      client := &http.Client {
      }
      req, err := http.NewRequest(method, url, payload)

      if err != nil {
        fmt.Println(err)
        return
      }

      res, err := client.Do(req)
      if err != nil {
        fmt.Println(err)
        return
      }
      defer res.Body.Close()

      body, err := ioutil.ReadAll(res.Body)
      if err != nil {
        fmt.Println(err)
        return
      }
      fmt.Println(string(body))
      fmt.Println("Deletion of account with ID of " + id + " successful")
}

//fetches the list of accounts and prints them to the console
func fetchAccount(id string) string{
      fmt.Println("Fetching Account with ID: " + id)
      url := "http://localhost:8080/v1/organisation/accounts/" + id + "?version=0"
      method := "GET"

      client := &http.Client {
      }
      req, err := http.NewRequest(method, url, nil)

      if err != nil {
        fmt.Println(err)
        return ""
      }
  
      res, err := client.Do(req)
      if err != nil {
        fmt.Println(err)
        return ""
      }
      defer res.Body.Close()

      body, err := ioutil.ReadAll(res.Body)
      if err != nil {
        fmt.Println(err)
        return ""
      }

      return string(body)
}

func fetchAccounts() string {
    fmt.Println("Fetching all accounts")
    return fetchAccount("")
}

// creates an account with a randomly-generated uuid, then returns the uuid
func createAccount() (response string, err error) {

  url := "http://localhost:8080/v1/organisation/accounts"
  method := "POST"
  var attributes AccountAttributes

  country := "GB"
  classification := "Personal"

  attributes.Country = &country
  attributes.BaseCurrency = "GBP"
  attributes.BankID = "400302"
  attributes.BankIDCode = "GBDSC"
  attributes.AccountNumber = "10000004"
  attributes.Iban = "GB28NWBK40030212764204"
  attributes.Name = []string{"Nick","Bury"}
  attributes.Bic = "NWBKGB42"
  attributes.AccountClassification = &classification

  // credit to Keithwachira
  myUuid := uuid.NewV4()
  acctId := myUuid.String()
 
 
  var accountData = &AccountData{
      ID: string(acctId), 
      OrganisationID: "3161a9d0-afc3-11ec-b909-0242ac120002", 
      Type: "accounts",
      Attributes: &attributes }
  // array of strings.

  var account = &Account{Data: accountData}

  b, err := json.Marshal(account)
    if err != nil {
        fmt.Println(err)
        return "", err
    }
   // fmt.Println("b: " + string(b))

   
  payload := strings.NewReader(string(b))

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return "", err
  }

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return "", err
  }
  defer res.Body.Close()

  /* for debugging only, not part of solution
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return "", err
  }
  fmt.Println("143: " + string(body)) */
  fmt.Println("Creation of account with ID of " + acctId + " successful")

  return acctId, nil
}