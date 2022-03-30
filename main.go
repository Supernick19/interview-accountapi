package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "github.com/satori/go.uuid"
)

type Account struct {
    Data    *AccountData `json:"data,omitempty"`
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
  fetchAccounts()
  createAccount()
}

func fetchAccounts(){
    url := "http://localhost:8080/v1/organisation/accounts"
  method := "GET"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

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
}

func createAccount() {

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
        return
    }
    fmt.Println("b: " + string(b))

   
  payload := strings.NewReader(string(b))

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
  fmt.Println("143: " + string(body))
}