package search

import (
	zincsearchModels "falconEmailBackend/api/handler/zincsearch/models"
	"fmt"
)

// GetQueryGetAll obtain a valid string in orden to consume api: search/{index}/get_all
func GetQueryGetAll(reqBody zincsearchModels.GetAllRequest) string {
	return fmt.Sprintf(`{"search_type": "alldocuments", "from": %d, "max_results": %d, "sort_fields": ["-date"]}`, (reqBody.Page-1)*reqBody.MaxDataPage, reqBody.MaxDataPage)
}

// GetQuerySearchEmails obtain a valid string in orden to consume api: search/{index}/search_emails
func GetQuerySearchEmails(reqBody zincsearchModels.GetSearchEmails) string {
	return fmt.Sprintf(`
	{
		"search_type": "%s",
		"query": {
			"term": "%s"
		},
		"from": %d,
		"max_results": %d, 
		"sort_fields": ["-date"],
		"highlight": {
			"pre_tags": ["<%s class='%s'>"],
			"post_tags": ["</%s>"],
			"sort_fields": ["-date"],    
			"fields": {
				"bcc": {
					"pre_tags": [],
					"post_tags": []
				},
				"cc": {
					"pre_tags": [],
					"post_tags": []
				},
				"date": {
					"pre_tags": [],
					"post_tags": []
				},
				"from": {
					"pre_tags": [],
					"post_tags": []
				},
				"message": {
					"pre_tags": [],
					"post_tags": []
				},
				"subject": {
					"pre_tags": [],
					"post_tags": []
				},
				"to": {
					"pre_tags": [],
					"post_tags": []
				}
			}
		} 
	}
	`, reqBody.SearchType, reqBody.Term, (reqBody.Page-1)*reqBody.MaxDataPage, reqBody.MaxDataPage, reqBody.TagHighlightName, reqBody.ClassTagHighlight, reqBody.TagHighlightName)
}
