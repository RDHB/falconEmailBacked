package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strings"

	chi "github.com/go-chi/chi"

	config "falconEmailBackend/api/handler/zincsearch/config"
	zincsearchModels "falconEmailBackend/api/handler/zincsearch/models"
	toolsFunctions "falconEmailBackend/pkg/tools/functions"
)

var zincsearchConfig = config.GetConfig()

// GetAll function that response the request when all messages are required
func GetAll(w http.ResponseWriter, r *http.Request) {

	// reviewing if url request have index
	index := chi.URLParam(r, "index")
	if index == "" {
		toolsFunctions.WriteErrorOne(w, http.StatusBadRequest, "You need speficy a value for index params", "")
		return
	}

	// reviewing and getting body request
	defer r.Body.Close()
	reqBody := zincsearchModels.GetAllRequest{}
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		toolsFunctions.WriteErrorOne(w, http.StatusInternalServerError, "Your request body not was readed correctly", "Error when intent reading bytes")
		return
	}

	// decoding body request into struct created for this
	err = json.Unmarshal(rBody, &reqBody)
	if err != nil {
		log.Println(err)
		toolsFunctions.WriteErrorOne(w, http.StatusInternalServerError, "Your request body not was readed correctly", "Error whent intent decode request body")
		return
	}

	// doing the query object in order to consume zincsearch
	query := GetQueryGetAll(reqBody)

	// creating request for consume zincsearch api
	zincSearchReq, err := http.NewRequest(http.MethodPost, fmt.Sprintf(zincsearchConfig["zincURL"], os.Getenv("ZINCSEARCH_ZINCHOST"), index), strings.NewReader(query))
	if err != nil {
		log.Println(err)
		toolsFunctions.WriteErrorOne(w, http.StatusInternalServerError, "An error occurred while trying to create a connection to the service", "Check the endpoint you wish to consume")
		return
	}

	// consuming zincsearch api
	zincSearchReq.SetBasicAuth(os.Getenv("ZINCSEARCH_USER_ID"), os.Getenv("ZINCSEARCH_PASSWORD"))
	zincSearchReq.Header.Set("Content-Type", "application/json")
	zincSearchResp, err := http.DefaultClient.Do(zincSearchReq)
	if err != nil {
		log.Println(err)
		toolsFunctions.WriteErrorOne(w, http.StatusBadGateway, "Not is possible get data", "Error occur conecting with service")
		return
	}

	// getting data returned by zincsearch api
	defer zincSearchResp.Body.Close()

	responseBody, err := ioutil.ReadAll(zincSearchResp.Body)
	if err != nil {
		log.Println(err)
		toolsFunctions.WriteErrorOne(w, http.StatusBadGateway, "Data not was readed successfully", "Error occur reading indexer service")
		return
	}

	// decoding data returned by zincsearch api
	zincSearchResponseJSON := zincsearchModels.ZincSearchResponseGelAll{}

	err = json.Unmarshal([]byte(string(responseBody)), &zincSearchResponseJSON)
	if err != nil {
		log.Println(err)
		toolsFunctions.WriteErrorOne(w, http.StatusBadGateway, "Data not was readed successfully", "Error occur decoding indexer service results")
		return
	}

	// wrinting a response for request
	totalData := int64(zincSearchResponseJSON.Hits.Total.Value)
	data := zincSearchResponseJSON.Hits.Hits
	totalPages := int64(math.Ceil(float64(totalData) / float64(reqBody.MaxDataPage)))
	toolsFunctions.WriteResponseOne(w, http.StatusOK, "Data were successfully recovered", totalData, reqBody.MaxDataPage, totalPages, reqBody.Page, data)
}

// GetSearchEmails get all emails after a search
func GetSearchEmails(w http.ResponseWriter, r *http.Request) {

	// reviewing if url request have index
	index := chi.URLParam(r, "index")
	if index == "" {
		toolsFunctions.WriteErrorOne(w, http.StatusBadRequest, "You need speficy a value for index params", "")
		return
	}

	// reviewing and getting body request
	defer r.Body.Close()
	reqBody := zincsearchModels.GetSearchEmails{}
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		toolsFunctions.WriteErrorOne(w, http.StatusInternalServerError, "Your request body not was readed correctly", "Error when intent reading bytes")
		return
	}

	// decoding body request into struct created for this
	err = json.Unmarshal(rBody, &reqBody)
	if err != nil {
		log.Println(err)
		toolsFunctions.WriteErrorOne(w, http.StatusInternalServerError, "Your request body not was readed correctly", "Error whent intent decode request body")
		return
	}

	// doing the query object in order to consume zincsearch
	query := GetQuerySearchEmails(reqBody)

	// creating request for consume zincsearch api
	zincSearchReq, err := http.NewRequest(http.MethodPost, fmt.Sprintf(zincsearchConfig["zincURL"], os.Getenv("ZINCSEARCH_ZINCHOST"), index), strings.NewReader(query))
	if err != nil {
		log.Println(err)
		toolsFunctions.WriteErrorOne(w, http.StatusInternalServerError, "An error occurred while trying to create a connection to the service", "Check the endpoint you wish to consume")
		return
	}

	// consuming zincsearch api
	zincSearchReq.SetBasicAuth(os.Getenv("ZINCSEARCH_USER_ID"), os.Getenv("ZINCSEARCH_PASSWORD"))
	zincSearchReq.Header.Set("Content-Type", "application/json")
	zincSearchResp, err := http.DefaultClient.Do(zincSearchReq)
	if err != nil {
		log.Println(err)
		toolsFunctions.WriteErrorOne(w, http.StatusBadGateway, "Not is possible get data", "Error occur conecting with service")
		return
	}

	// getting data returned by zincsearch api
	defer zincSearchResp.Body.Close()

	responseBody, err := ioutil.ReadAll(zincSearchResp.Body)
	if err != nil {
		log.Println(err)
		toolsFunctions.WriteErrorOne(w, http.StatusBadGateway, "Data not was readed successfully", "Error occur reading indexer service")
		return
	}

	// decoding data returned by zincsearch api
	zincSearchResponseJSON := zincsearchModels.ZincSearchResponseSearch{}

	err = json.Unmarshal([]byte(string(responseBody)), &zincSearchResponseJSON)
	if err != nil {
		log.Println(err)
		toolsFunctions.WriteErrorOne(w, http.StatusBadGateway, "Data not was readed successfully", "Error occur decoding indexer service results")
		return
	}

	if reqBody.TagHighlightName != "" {
		GetHighlightedReponse(&zincSearchResponseJSON, reqBody)
	}

	// wrinting a response for request
	totalData := int64(zincSearchResponseJSON.Hits.Total.Value)
	data := zincSearchResponseJSON.Hits.Hits
	totalPages := int64(math.Ceil(float64(totalData) / float64(reqBody.MaxDataPage)))
	toolsFunctions.WriteResponseOne(w, http.StatusOK, "Data were successfully recovered", totalData, reqBody.MaxDataPage, totalPages, reqBody.Page, data)
}
