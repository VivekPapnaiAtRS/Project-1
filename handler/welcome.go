package handler

import (
	"encoding/json"
	"github.com/Project/models"
	"github.com/Project/utils"
	"net/http"
	"strings"
)

func CountFrequency(resp http.ResponseWriter, req *http.Request) {

	var textInputRequest models.TextInput

	err := json.NewDecoder(req.Body).Decode(&textInputRequest)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
	}

	textArray := strings.Split(textInputRequest.Text, " ")

	wordOccurrenceMap := make(map[string]int)

	for _, text := range textArray {
		wordOccurrenceMap[text] += 1
	}

	mostOccurrenceWordsMap := utils.SortMapByValue(wordOccurrenceMap)

	utils.EncodeJSONBody(resp, http.StatusOK, mostOccurrenceWordsMap)
}
