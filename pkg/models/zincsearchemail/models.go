package zincsearchemail

// Email model for structure emails in the indexer
type Email struct {
	BCC     string `json:"bcc"`
	CC      string `json:"cc"`
	Date    string `json:"date"`
	From    string `json:"from"`
	Message string `json:"message"`
	Subject string `json:"subject"`
	To      string `json:"to"`
}

type ZinckSearchBulkData struct {
	Index   string  `json:"index"`
	Records []Email `json:"records"`
}
