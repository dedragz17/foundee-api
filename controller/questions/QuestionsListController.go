package questions

import (
	"encoding/json"
	FoundeeAPI "github.com/foundee-org/foundee-api"
	"net/http"
)

func QuestionsListController(writer http.ResponseWriter, request *http.Request) {

	db := FoundeeAPI.GetDB()

	type QuestionStruct struct {
		Title  string
		Author string
	}

	var result []QuestionStruct
	db.Raw("SELECT title, author FROM question ORDER BY id DESC LIMIT 20").Scan(&result)

	response := make(map[interface{}]interface{})

	counter := 0
	for question := range result {
		counter++

		response[counter] = result[question]
	}

	response["ok"] = true
	response["questions_count"] = counter

	json.NewEncoder(writer).Encode(response)
	writer.WriteHeader(http.StatusFound)
}
