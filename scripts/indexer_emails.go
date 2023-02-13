package main

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	EmailModel "falconEmailBackend/pkg/models/zincsearchemail"
)

func main() {
	indexer()
}

func indexer() {
	filepath := "/home/rdhb/go/src/falconEmailBackend/data/enron_mail_20110402.tgz"
	indexerEmails(filepath)
}

func indexerEmails(filePath string) {

	createIndexZincSearchEnronCorpEmails()

	if filePath == "" {
		log.Println("filePath parameter not have been declared")
	}

	data, error := getDataFromFileTGZ(filePath)
	if error != nil {
		log.Println(error)
	}
	log.Println("Data was obtained from tgz file")

	// sendDataToZincSearchIndexer(data)
	sendBulkData(data)
}

func getDataFromFileTGZ(filePath string) ([]EmailModel.Email, error) {
	// reading file tgz
	file, error := os.Open(filePath)
	if error != nil {
		log.Println(error)
		return nil, error
	}

	defer file.Close()

	// reading gzf file
	gzf, error := gzip.NewReader(file)
	if error != nil {
		log.Println(error)
		return nil, error
	}

	// reading tar file
	tarReader := tar.NewReader(gzf)

	// getting emails
	emails := make([]EmailModel.Email, 0)
	for {
		header, error := tarReader.Next()
		if error == io.EOF {
			log.Println("The file have been readed")
			break
		}

		if error != nil {
			log.Println(error)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			continue
		case tar.TypeReg:
			if strings.Contains(header.Name, "all_documents") {
				emails = append(emails, getEmail(tarReader, header))
			}

		default:
			log.Println("Not was possible read file from tar")
		}
	}

	return emails, nil
}

func getEmail(tarReader *tar.Reader, header *tar.Header) EmailModel.Email {

	scanEmail := bufio.NewScanner(tarReader)
	scanEmail.Split(bufio.ScanLines)

	var date, from, subject, toS, ccS, bccS string
	var to, cc, bcc, messages []string

	readMessage := false

	for scanEmail.Scan() {
		line := scanEmail.Text()
		switch {
		case strings.HasPrefix(line, "Date:"):
			date = strings.ReplaceAll(line, "Date:", "")
			date = strings.TrimSpace(date)

		case strings.HasPrefix(line, "From:"):
			from = strings.ReplaceAll(line, "From:", "")
			from = strings.TrimSpace(from)

		case strings.HasPrefix(line, "Subject:"):
			subject = strings.ReplaceAll(line, "Subject:", "")
			subject = strings.TrimSpace(subject)

		case strings.HasPrefix(line, "To:"):
			toS = strings.ReplaceAll(line, "To:", "")
			toS = strings.ReplaceAll(toS, ",", "")
			toS = strings.ReplaceAll(toS, ";", "")
			toS = strings.TrimSpace(toS)
			to = strings.Split(toS, " ")

		case strings.HasPrefix(line, "Cc:"):
			ccS = strings.ReplaceAll(line, "Cc:", "")
			ccS = strings.ReplaceAll(ccS, ",", "")
			ccS = strings.ReplaceAll(ccS, ";", "")
			ccS = strings.TrimSpace(ccS)
			cc = strings.Split(ccS, " ")

		case strings.HasPrefix(line, "Bcc:"):
			bccS = strings.ReplaceAll(line, "Bcc:", "")
			bccS = strings.ReplaceAll(bccS, ",", "")
			bccS = strings.ReplaceAll(bccS, ";", "")
			bccS = strings.TrimSpace(bccS)
			bcc = strings.Split(bccS, " ")

		case strings.HasPrefix(line, "X-FileName:"):
			readMessage = true

		default:
			if readMessage {
				messageLine := strings.TrimSpace(line)
				messages = append(messages, messageLine)
			}
		}

	}

	dateS, _ := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700 (PST)", date)

	return EmailModel.Email{
		BCC:     strings.Join(bcc, ";"),
		CC:      strings.Join(cc, ";"),
		Date:    dateS.Format(time.RFC3339),
		From:    from,
		Message: strings.Join(messages, " "),
		Subject: subject,
		To:      strings.Join(to, ";"),
	}
}

func createIndexZincSearchEnronCorpEmails() {
	method := http.MethodPost
	url := "http://localhost:4080/api/index"
	user := os.Getenv("ZINCSEARCH_USER_ID_ADMIN")      //"admin"
	password := os.Getenv("ZINCSEARCH_PASSWORD_ADMIN") //"Complexpass#123"
	body := `{
		"name": "enronCorpEmails",
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
	}`

	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		log.Println(err)
	}

	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	resp.Body.Close()
}

func zincsearchSendDataAPI(method string, url string, user string, password string, body *bytes.Buffer) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println(err)
	}

	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	resp.Body.Close()
}

func sendDataToZincSearchIndexer(data []EmailModel.Email) {
	method := http.MethodPost
	url := "http://localhost:4080/api/enronCorpEmails/_doc"
	user := os.Getenv("ZINCSEARCH_USER_ID_ADMIN")      //"admin"
	password := os.Getenv("ZINCSEARCH_PASSWORD_ADMIN") //"Complexpass#123"

	for i := 0; i < len(data); i++ {

		payloadBuf := new(bytes.Buffer)

		error := json.NewEncoder(payloadBuf).Encode(data[i])
		if error != nil {
			log.Println(error)
		}

		zincsearchSendDataAPI(method, url, user, password, payloadBuf)
	}
}

func sendBulkData(data []EmailModel.Email) {
	bulkData := EmailModel.ZinckSearchBulkData{
		Index:   "enronCorpEmails",
		Records: data,
	}

	method := http.MethodPost
	url := "http://localhost:4080/api/_bulkv2"
	user := "admin"
	password := "Complexpass#123"

	payloadBuf := new(bytes.Buffer)

	error := json.NewEncoder(payloadBuf).Encode(bulkData)
	if error != nil {
		log.Println(error)
	}

	zincsearchSendDataAPI(method, url, user, password, payloadBuf)
}
