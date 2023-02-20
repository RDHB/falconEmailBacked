package search

import (
	zincsearchModels "falconEmailBackend/api/handler/zincsearch/models"
	"fmt"
	"reflect"
	"strings"
)

// GetHighlightedReponse replace response without highlight for highligtht response
func GetHighlightedReponse(zincSearchResponseJSON *zincsearchModels.ZincSearchResponseSearch, reqBody zincsearchModels.GetSearchEmails) {
	for i, data := range zincSearchResponseJSON.Hits.Hits {

		dSource := data.Source
		dHigh := data.Highlight
		valueHigh := reflect.ValueOf(dHigh)
		typeHigh := reflect.TypeOf(dHigh)

		for y := 0; y < valueHigh.NumField(); y++ {

			if valueHigh.Field(y).Len() != 0 {

				switch typeHigh.Field(y).Name {
				case "BCC":
					zincSearchResponseJSON.Hits.Hits[i].Source.BCC = getDataHighlighted(dHigh.BCC, dSource.BCC, reqBody.TagHighlightName, reqBody.ClassTagHighlight)
				case "CC":
					zincSearchResponseJSON.Hits.Hits[i].Source.CC = getDataHighlighted(dHigh.CC, dSource.CC, reqBody.TagHighlightName, reqBody.ClassTagHighlight)
				case "Date":
					zincSearchResponseJSON.Hits.Hits[i].Source.Date = getDataHighlighted(dHigh.Date, dSource.Date, reqBody.TagHighlightName, reqBody.ClassTagHighlight)
				case "From":
					zincSearchResponseJSON.Hits.Hits[i].Source.From = getDataHighlighted(dHigh.From, dSource.From, reqBody.TagHighlightName, reqBody.ClassTagHighlight)
				case "Message":
					zincSearchResponseJSON.Hits.Hits[i].Source.Message = getDataHighlighted(dHigh.Message, dSource.Message, reqBody.TagHighlightName, reqBody.ClassTagHighlight)
				case "Subject":
					zincSearchResponseJSON.Hits.Hits[i].Source.Subject = getDataHighlighted(dHigh.Subject, dSource.Subject, reqBody.TagHighlightName, reqBody.ClassTagHighlight)
				case "To":
					zincSearchResponseJSON.Hits.Hits[i].Source.To = getDataHighlighted(dHigh.To, dSource.To, reqBody.TagHighlightName, reqBody.ClassTagHighlight)
				}
			}
		}
	}
}

func getDataHighlighted(valueItemsHighlight []string, valueItemSource string, tagHighlightName string, classTagHighlight string) string {

	var (
		replacementText string
		replaceText     string
		itemHighlighted string
	)

	for _, valueItemHighlight := range valueItemsHighlight {

		// preparing replacementText
		replacementText = valueItemHighlight

		if strings.HasPrefix(replacementText, "…") {
			replacementText = strings.ReplaceAll(replacementText, "…", "")
		}

		if strings.HasSuffix(replacementText, "…") {
			replacementText = strings.ReplaceAll(replacementText, "…", "")
		}

		// preparing statement replaced
		replaceText = replacementText
		replaceText = strings.ReplaceAll(replaceText, fmt.Sprintf("<%s class='%s'>", tagHighlightName, classTagHighlight), "")
		replaceText = strings.ReplaceAll(replaceText, fmt.Sprintf("</%s>", tagHighlightName), "")

		itemHighlighted = strings.ReplaceAll(valueItemSource, replaceText, replacementText)
	}

	return itemHighlighted
}
