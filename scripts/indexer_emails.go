package scripts

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	EmailModel "falconEmailBackend/pkg/models/zincsearchemail"
)

const (
	index = "enronCorpEmails"
)

// Indexer function that allow send indexed data to zincsearch
func Indexer() {

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

	sendDataToZincSearchIndexer(data)
	// sendBulkData(data)
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
			// emails = append(emails, getEmail(tarReader, header))
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
		BCC:     strings.Join(bcc, "; "),
		CC:      strings.Join(cc, "; "),
		Date:    dateS.Format(time.RFC3339),
		From:    from,
		Message: strings.Join(messages, " "),
		Subject: subject,
		To:      strings.Join(to, "; "),
	}
}

func createIndexZincSearchEnronCorpEmails() {
	method := http.MethodPost
	url := fmt.Sprintf("%s/api/index", os.Getenv("ZINCSEARCH_ZINCHOST"))
	body := GetCreateIndexQueryHighlightEnabled(index)

	zincsearchSendDataAPI(method, url, bytes.NewBuffer([]byte(body)))
}

func zincsearchSendDataAPI(method string, url string, body *bytes.Buffer) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println(err)
	}

	req.SetBasicAuth(os.Getenv("ZINCSEARCH_USER_ID_ADMIN"), os.Getenv("ZINCSEARCH_PASSWORD_ADMIN"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	error := resp.Body.Close()
	if error != nil {
		log.Println(error)
	}
}

func sendDataToZincSearchIndexer(data []EmailModel.Email) {
	method := http.MethodPost
	url := fmt.Sprintf("%s/api/%s/_doc", os.Getenv("ZINCSEARCH_ZINCHOST"), index)

	for i := 0; i < len(data); i++ {

		payloadBuf := new(bytes.Buffer)

		error := json.NewEncoder(payloadBuf).Encode(data[i])
		if error != nil {
			log.Println(error)
		}

		zincsearchSendDataAPI(method, url, payloadBuf)
	}
}

func sendBulkData(data []EmailModel.Email) {
	bulkData := EmailModel.ZinckSearchBulkData{
		Index:   index,
		Records: data,
	}

	method := http.MethodPost
	url := fmt.Sprintf("%s/api/_bulkv2", os.Getenv("ZINCSEARCH_ZINCHOST"))

	payloadBuf := new(bytes.Buffer)

	error := json.NewEncoder(payloadBuf).Encode(bulkData)
	if error != nil {
		log.Println(error)
	}

	zincsearchSendDataAPI(method, url, payloadBuf)
}
