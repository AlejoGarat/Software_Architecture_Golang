package requests

import (
	"bytes"
	"consultant-service/models/read"
	"encoding/json"
	"fmt"
	"net/http"
)

type FilterRequest struct {
}

func ElectionBeginningFiltersRequest(filters read.ElectionBeginningFilters) (*http.Response, error) {
	client := &http.Client{}

	postBody, _ := json.Marshal(map[string]read.ElectionBeginningFilters{
		"eb_filters": filters,
	})

	responseBody := bytes.NewBuffer(postBody)

	req, err := http.NewRequest(http.MethodPut, "http://127.0.0.1:8080/consultant-api/v1/filters", responseBody)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)

	if err != nil {
		return resp, err
	}

	fmt.Println(resp)
	return resp, nil
}

func ElectionEndFiltersRequest(filters read.ElectionEndFilters) (*http.Response, error) {
	client := &http.Client{}

	postBody, _ := json.Marshal(map[string]read.ElectionEndFilters{
		"eb_filters": filters,
	})

	responseBody := bytes.NewBuffer(postBody)

	req, err := http.NewRequest(http.MethodPut, "http://127.0.0.1:8080/consultant-api/v1/filters", responseBody)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)

	if err != nil {
		return resp, err
	}

	fmt.Println(resp)
	return resp, nil
}

func VoteIssuanceFiltersRequest(filters read.VoteIssuanceFilters) (*http.Response, error) {
	client := &http.Client{}

	postBody, _ := json.Marshal(map[string]read.VoteIssuanceFilters{
		"eb_filters": filters,
	})

	responseBody := bytes.NewBuffer(postBody)

	req, err := http.NewRequest(http.MethodPut, "http://127.0.0.1:8080/consultant-api/v1/filters", responseBody)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)

	if err != nil {
		return resp, err
	}

	fmt.Println(resp)
	return resp, nil
}
