package scripts

import "fmt"

// GetCreateIndexQuery query in orden to create index into zincsearch with highlight enabled
func GetCreateIndexQueryHighlightEnabled(index string) string {
	return fmt.Sprintf(`
	{
		"name": "%s",
		"storage_type": "disk",
		"shard_num": 3,
		"mappings": {
			"properties": {
				"@timestamp": {
					"type": "date",
					"index": true,
					"store": true,
					"sortable": true,
					"aggregatable": true,
					"highlightable": true
				},
				"_id": {
					"type": "keyword",
					"index": true,
					"store": false,
					"sortable": true,
					"aggregatable": true,
					"highlightable": false
				},
				"bcc": {
					"type": "text",
					"index": true,
					"store": true,
					"sortable": false,
					"aggregatable": false,
					"highlightable": true
				},
				"cc": {
					"type": "text",
					"index": true,
					"store": true,
					"sortable": false,
					"aggregatable": false,
					"highlightable": true
				},
				"date": {
					"type": "date",
					"format": "2006-01-02T15:04:05Z07:00",
					"index": true,
					"store": true,
					"sortable": true,
					"aggregatable": true,
					"highlightable": true
				},
				"from": {
					"type": "text",
					"index": true,
					"store": true,
					"sortable": false,
					"aggregatable": false,
					"highlightable": true
				},
				"message": {
					"type": "text",
					"index": true,
					"store": true,
					"sortable": false,
					"aggregatable": false,
					"highlightable": true
				},
				"subject": {
					"type": "text",
					"index": true,
					"store": true,
					"sortable": false,
					"aggregatable": false,
					"highlightable": true
				},
				"to": {
					"type": "text",
					"index": true,
					"store": true,
					"sortable": false,
					"aggregatable": false,
					"highlightable": true
				}
			}
		}
	}
	`, index)
}
