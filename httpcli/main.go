package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	headersFlag []string
	queriesFlag []string
	jsonFlag    string
)

// see more at http.MethodGet
var methods = []string{
	"get",
	"post",
	"put",
	"patch",
	"delete",
	// "head",
	// "connect",
	// "options",
	// "trace",
}

const rootMethod = "get"

var rootCmd = &cobra.Command{
	Use:   "httpcli <url>",
	Short: fmt.Sprintf("call http request with %s method", rootMethod),
	Args:  cobra.ExactArgs(1),
	RunE:  runCommandWithMethod(strings.ToUpper(rootMethod)),
}

var subCmds = []*cobra.Command{}

func init() {
	rootCmd.PersistentFlags().
		StringArrayVar(&headersFlag, "header", nil, "header of the request")
	rootCmd.PersistentFlags().
		StringArrayVar(&queriesFlag, "query", nil, "query parameters of the request")
	rootCmd.PersistentFlags().
		StringVar(&jsonFlag, "json", "", "json body of the request")

	for _, method := range methods {
		subCmds = append(subCmds, &cobra.Command{
			Use:   fmt.Sprintf("%s <url>", method),
			Short: fmt.Sprintf("call http request with %s method", method),
			Args:  cobra.ExactArgs(1),
			RunE:  runCommandWithMethod(strings.ToUpper(method)),
		})
	}
	rootCmd.AddCommand(subCmds...)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// -------------------------------------------------------------------------- //

func runCommandWithMethod(method string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		// there should be no error. if it happen, just panic.
		url := args[0]

		req, err := constructHttpRequest(method, url, jsonFlag)
		if err != nil {
			return err
		}

		if err = addHeaders(req, headersFlag); err != nil {
			return err
		}

		if err = addQueries(req, queriesFlag); err != nil {
			return err
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cmd.Println(string(body))
		return nil
	}
}

func constructHttpRequest(method, url, jsonBody string) (*http.Request, error) {
	if method == http.MethodPost {

		if !json.Valid([]byte(jsonBody)) {
			return nil, errors.New("invalid json body")
		}

		reqBody := bytes.NewBufferString(jsonBody)
		return http.NewRequest(method, url, reqBody)
	}
	return http.NewRequest(method, url, nil)
}

func addHeaders(req *http.Request, headers []string) error {
	for _, header := range headers {
		splited := strings.Split(header, "=")
		if len(splited) != 2 {
			return errors.New("invalid header flag value")
		}

		req.Header.Add(splited[0], splited[1])
	}

	return nil
}

func addQueries(req *http.Request, queries []string) error {
	myQuery := req.URL.Query()
	for _, query := range queries {
		splited := strings.Split(query, "=")
		if len(splited) != 2 {
			return errors.New("invalid query flag value")
		}

		myQuery.Add(splited[0], splited[1])
	}

	req.URL.RawQuery = myQuery.Encode()
	return nil
}
