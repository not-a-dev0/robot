// Copyright (C) 2019 Ali Shanaakh, github@shanaakh.pro
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package autoria_api

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

type CarSearchResultResponse struct {
	CarsIDs []string `json:"ids"`
	Count   int64    `json:"count"`
	LastID  int64    `json:"last_id"`
}

type CarResultResponse struct {
	SearchResult CarSearchResultResponse `json:"search_result"`
}

type CarSearchResponse struct {
	Result CarResultResponse `json:"result"`
}

type CarSearchParam struct {
	Key   string
	Value string
}

func (param CarSearchParam) String() string {
	return fmt.Sprintf("%s=%s", param.Key, param.Value)
}

func (api *API) GetSearchCars(params ...CarSearchParam) (search CarSearchResponse) {
	searchParams := make([]string, len(params))

	for i, v := range params {
		searchParams[i] = fmt.Sprint(v)
	}

	fmt.Print(api.BuildURL("search", searchParams...))

	resp, err := http.Get(api.BuildURL("search", searchParams...))

	if err != nil {
		panic(err.Error())
	}

	err = json.NewDecoder(resp.Body).Decode(&search)

	if err != nil {
		panic(err.Error())
	}

	return search
}

func ParseCarSearchParams(url string) ([]CarSearchParam, error) {
	query := strings.TrimPrefix(url, "https://auto.ria.com/search/?")
	params := strings.Split(query, "&")

	searchParams := make([]CarSearchParam, len(params))

	for i, v := range params {
		lexemes := strings.Split(v, "=")

		if len(lexemes) != 2 {
			return nil, errors.New("Invalid params")
		}

		searchParams[i].Value = lexemes[0]
		searchParams[i].Key = lexemes[1]
	}

	return searchParams, nil
}
