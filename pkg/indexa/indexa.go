package indexa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IndexaAPI interface {
	Me() (*User, error)
	Performance(account string) (*Performance, error)
}

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Document       string `json:"document"`
	DocumentType   string `json:"document_type"`
	IsActivated    bool   `json:"is_activated"`
	EmailActivated bool   `json:"email_activated"`
	AffiliateCode  string `json:"affiliate_code"`
	Preferences    struct {
		LastAccountVisited int    `json:"last_account_visited"`
		FinancialPlanning  string `json:"financial_planning"`
		Chart              struct {
			Type  string `json:"type"`
			Range struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"range"`
			Ignored string `json:"ignored"`
		} `json:"chart"`
	} `json:"preferences"`
	Profiles []interface{} `json:"profiles"`
	Accounts []struct {
		AccountNumber  string `json:"account_number"`
		CreatedAt      string `json:"created_at"`
		Status         string `json:"status"`
		Type           string `json:"type"`
		Path           string `json:"@path"`
		Funding        string `json:"funding"`
		StatusProvider string `json:"status_provider"`
		NumHolders     int    `json:"num_holders"`
		MainHolderName string `json:"main_holder_name"`
		Role           string `json:"role"`
		UserHolderType string `json:"user_holder_type"`
	} `json:"accounts"`
	AccountsRelations []struct {
		AccountNumber string `json:"account_number"`
		Relation      string `json:"relation"`
	} `json:"accounts_relations"`
	Person []interface{} `json:"person"`
}

type Performance struct {
	Return struct {
		TimeReturn          float64            `json:"time_return"`
		TimeReturnLastWeek  float64            `json:"time_return_last_week"`
		TimeReturnLastMonth float64            `json:"time_return_last_month"`
		TimeReturnLastYear  float64            `json:"time_return_last_year"`
		TimeReturnAnnual    float64            `json:"time_return_annual"`
		XIRR                float64            `json:"XIRR"`
		Investment          int                `json:"investment"`
		Pl                  float64            `json:"pl"`
		Average             float64            `json:"average"`
		MoneyReturn         float64            `json:"money_return"`
		MoneyReturnAnnual   float64            `json:"money_return_annual"`
		Inflows             int                `json:"inflows"`
		Outflows            int                `json:"outflows"`
		TaxOutflows         int                `json:"tax_outflows"`
		PlNetTax            float64            `json:"pl_net_tax"`
		TotalAmount         float64            `json:"total_amount"`
		Volatility          float64            `json:"volatility"`
		Index               map[string]float64 `json:"index"`
	} `json:"return"`
	Volatility float64 `json:"volatility"`
}

type IndexaClient struct {
	baseURL   string
	authToken string
	client    *http.Client
}

func NewIndexaClient(baseURL, authToken string) IndexaAPI {
	return &IndexaClient{
		baseURL:   baseURL,
		authToken: authToken,
		client:    &http.Client{},
	}
}

func (c *IndexaClient) Me() (*User, error) {
	url := fmt.Sprintf("%s/users/me", c.baseURL)

	body, err := c.makeRequest(http.MethodGet, url)

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, fmt.Errorf("couldn't unmarshall user: %v", err)
	}

	return &user, nil
}

func (c *IndexaClient) Performance(account string) (*Performance, error) {
	url := fmt.Sprintf("%s/accounts/%s/performance", c.baseURL, account)

	body, err := c.makeRequest(http.MethodGet, url)

	var p Performance
	err = json.Unmarshal(body, &p)
	if err != nil {
		return nil, fmt.Errorf("couldn't unmarshall performace: %v", err)
	}

	return &p, nil
}

func (c *IndexaClient) makeRequest(method, url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request to %s: %v", url, err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-AUTH-TOKEN", c.authToken)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error requesting to %s: %v", url, err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't read body: %v", err)
	}

	return body, nil
}
