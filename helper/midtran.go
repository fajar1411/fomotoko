package helper

type MidtransClient struct {
	ClientKey string
	ServerKey string
	BaseURL   string
}

func NewMidtransClient(clientKey, serverKey string) *MidtransClient {
	return &MidtransClient{
		ClientKey: clientKey,
		ServerKey: serverKey,
		BaseURL:   "https://api.midtrans.com/v1/",
	}
}

// func (mc *MidtransClient) CheckBankAccount(bank, noRekening string) (string, error) {
// 	url := fmt.Sprintf("%saccount_validation?bank=%s&account=%s", mc.BaseURL, bank, noRekening)

// 	// Create an HTTP client
// 	client := &http.Client{}

// 	// Create a request
// 	req, err := http.NewRequest("GET", url, nil)
// 	fmt.Printf("ini req", req)
// 	if err != nil {
// 		return "", err
// 	}
// 	req.Header.Add("Accept", "application/json")
// 	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(mc.ClientKey+":"+mc.ServerKey)))
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	// Check the response status
// 	if resp.StatusCode != http.StatusOK {
// 		// // Print the full response body for debugging
// 		// responseBody, _ := ioutil.ReadAll(resp.Body)
// 		// fmt.Printf("Raw response body: %s\n", responseBody)
// 		return "", fmt.Errorf("Failed to check bank account. Status Code: %d", resp.StatusCode)
// 	}

// 	var response map[string]interface{}
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		return "", err
// 	}

// 	// Extract the bank name from the response
// 	bankName, ok := response["bank_name"].(string)
// 	if !ok {
// 		return "", errors.New("Bank name not found in the response")
// 	}

// 	return bankName, nil
// }
