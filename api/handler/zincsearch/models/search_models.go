package models

// =====================================================================================================
// models for manage client request

// GetAllRequest model that define parameter that client should pass in orden to consume api: search/{index}/get_all
type GetAllRequest struct {
	Page        int64 `json:"page"`
	MaxDataPage int64 `json:"max_data_page"`
}

// GetSearchEmails model that define parameter that client should pass in orden to consume api: search/{index}/search_emails
type GetSearchEmails struct {
	Page              int64  `json:"page"`
	MaxDataPage       int64  `json:"max_data_page"`
	SearchType        string `json:"search_type"`
	Term              string `json:"term"`
	TagHighlightName  string `json:"tag_highlight_name,omitempty"`
	ClassTagHighlight string `json:"class_tag_highlight,omitempty"`
}

// ======================================================================================================
// models for manage zincsearch response api and response

// ZincSearchResponseGelAll model tata define the structure object in order to receive zinsearch response search api
type ZincSearchResponseGelAll struct {
	Took     int64  `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   shards `json:"_shards"`
	Hits     hits   `json:"hits"`
}

// ZincSearchResponseSearch model tata define the structure object in order to receive zinsearch response search api
type ZincSearchResponseSearch struct {
	Took     int64      `json:"took"`
	TimedOut bool       `json:"timed_out"`
	Shards   shards     `json:"_shards"`
	Hits     hitsSearch `json:"hits"`
}

type shards struct {
	Total      int64 `json:"total"`
	Successful int64 `json:"successful"`
	Skipped    int64 `json:"skipped"`
	Failed     int64 `json:"failed"`
}

type hits struct {
	Total    total   `json:"total"`
	MaxScore int64   `json:"max_score"`
	Hits     []hitss `json:"hits"`
}

type hitsSearch struct {
	Total    total             `json:"total"`
	MaxScore float64           `json:"max_score"`
	Hits     []hitshitsSearchs `json:"hits"`
}

type total struct {
	Value int `json:"value"`
}

type hitss struct {
	Index     string `json:"_index"`
	Type      string `json:"_type"`
	ID        string `json:"_id"`
	Score     int64  `json:"_score"`
	Timestamp string `json:"@timestamp"`
	Source    email  `json:"_source"`
}

type hitshitsSearchs struct {
	Index     string    `json:"_index"`
	Type      string    `json:"_type"`
	ID        string    `json:"_id"`
	Score     float64   `json:"_score"`
	Timestamp string    `json:"@timestamp"`
	Source    email     `json:"_source"`
	Highlight highlight `json:"highlight,omitempty"`
}

type email struct {
	BCC     string `json:"bcc,omitempty"`
	CC      string `json:"cc,omitempty"`
	Date    string `json:"date"`
	From    string `json:"from"`
	Message string `json:"message"`
	Subject string `json:"subject,omitempty"`
	To      string `json:"to"`
}

type highlight struct {
	BCC     []string `json:"bcc,omitempty"`
	CC      []string `json:"cc,omitempty"`
	Date    []string `json:"date,omitempty"`
	From    []string `json:"from,omitempty"`
	Message []string `json:"message,omitempty"`
	Subject []string `json:"subject,omitempty"`
	To      []string `json:"to,omitempty"`
}
