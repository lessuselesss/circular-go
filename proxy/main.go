package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/circular-protocol/circular-go/circular_protocol_api"
)

var client *circular_protocol_api.Client

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Extract CEP action from query parameter
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "Missing 'cep' query parameter", http.StatusBadRequest)
		return
	}

	// Parse JSON body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var args map[string]string
	if len(body) > 0 {
		if err := json.Unmarshal(body, &args); err != nil {
			http.Error(w, "Error parsing JSON body", http.StatusBadRequest)
			return
		}
	}

	var result interface{}

	// Map CEP actions to SDK methods
	switch cep {
	case "Circular_CheckWallet_":
		result = client.CheckWallet(args["Blockchain"], args["Address"])
	case "Circular_GetWallet_":
		result = client.GetWallet(args["Blockchain"], args["Address"])
	case "Circular_GetWalletNonce_":
		result = client.GetWalletNonce(args["Blockchain"], args["Address"])
	case "Circular_GetLatestTransactions_":
		result = client.GetLatestTransaction(args["Blockchain"], args["Address"])
	case "Circular_GetWalletBalance_":
		result = client.GetWalletBalance(args["Blockchain"], args["Address"], args["Asset"])
	case "Circular_AddTransaction_":
		result = client.SendTransaction(
			args["ID"], args["From"], args["To"], args["Timestamp"],
			args["Type"], args["Payload"], args["Nonce"], args["Signature"],
			args["Blockchain"],
		)
	case "Circular_GetBlock_":
		num, _ := strconv.Atoi(args["BlockNumber"])
		result = client.GetBlock(args["Blockchain"], num)
	case "Circular_GetBlockHeight_":
		result = client.GetBlockCount(args["Blockchain"])
	case "Circular_GetBlockRange_":
		start, _ := strconv.Atoi(args["Start"])
		end, _ := strconv.Atoi(args["End"])
		result = client.GetBlockRange(args["Blockchain"], start, end)
	case "Circular_GetAssetList_":
		result = client.GetAssetList(args["Blockchain"])
	case "Circular_GetAsset_":
		result = client.GetAsset(args["Blockchain"], args["AssetName"])
	case "Circular_GetAssetSupply_":
		result = client.GetAssetSupply(args["Blockchain"], args["AssetName"])
	case "Circular_ResolveDomain_":
		result = client.GetDomain(args["Blockchain"], args["Domain"])
	case "Circular_GetVoucher_":
		result = client.GetVoucher(args["Blockchain"], args["Code"])
	case "Circular_GetAnalytics_":
		result = client.GetAnalytics(args["Blockchain"])
	case "Circular_GetPendingTransaction_":
		result = client.GetPendingTransaction(args["Blockchain"], args["ID"])
	case "Circular_GetTransactionbyID_":
		start, _ := strconv.Atoi(args["Start"])
		end, _ := strconv.Atoi(args["End"])
		result = client.GetTransactionByID(args["Blockchain"], args["ID"], start, end)
	case "Circular_GetTransactionbyNode_":
		start, _ := strconv.Atoi(args["Start"])
		end, _ := strconv.Atoi(args["End"])
		result = client.GetTransactionByNode(args["Blockchain"], args["NodeID"], start, end)
	case "Circular_GetTransactionbyAddress_":
		start, _ := strconv.Atoi(args["Start"])
		end, _ := strconv.Atoi(args["End"])
		result = client.GetTransactionByAddress(args["Blockchain"], args["Address"], start, end)
	case "Circular_GetTransactionbyDate_":
		result = client.GetTransactionByDate(args["Blockchain"], args["Address"], args["StartDate"], args["EndDate"])
	case "Circular_TestContract_":
		result = client.TestContract(args["Blockchain"], args["From"], args["Project"])
	case "Circular_CallContract_":
		result = client.CallContract(args["Blockchain"], args["From"], args["Address"], args["Request"])
	default:
		// Fallback for unknown actions or just proxying if needed, but for now error
		http.Error(w, "Unknown action: "+cep, http.StatusBadRequest)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	// Initialize SDK Client
	// Note: The URL here is for the SDK's internal use. 
	// Since we are proxying, we might want to use the production URL as default
	// or allow it to be configured. For now, we use the default.
	client = circular_protocol_api.NewClient("https://nag.circularlabs.io/NAG.php?cep=")

	http.HandleFunc("/NAG.php", handleRequest)

	fmt.Println("Go SDK Adapter listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
