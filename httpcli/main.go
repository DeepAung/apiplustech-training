package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var (
	headersFlag []string
	queriesFlag []string
	jsonFlag    string
)

var subCommands = []string{
	http.MethodGet,
	// http.MethodHead,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	// http.MethodConnect,
	// http.MethodOptions,
	// http.MethodTrace,
}

var rootCmd = &cobra.Command{
	Use:   "httpcli",
	Short: "httpcli is a httpcli!!!",
	Long:  `httpcli is a httpcli!!!. nothing more i guess?`,
	RunE:  runCommand,
}

func init() {
	rootCmd.PersistentFlags().
		StringArrayVar(&headersFlag, "header", nil, "header of the request")
	rootCmd.PersistentFlags().
		StringArrayVar(&queriesFlag, "query", nil, "query parameters of the request")
	rootCmd.PersistentFlags().
		StringVar(&jsonFlag, "json", "", "json body of the request")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// -------------------------------------------------------------------------- //

func runCommand(cmd *cobra.Command, args []string) error {
	method, url, err := parseArguments(args)
	if err != nil {
		return err
	}

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

func parseArguments(args []string) (method string, url string, err error) {
	switch len(args) {
	case 0:
		err = errors.New("requires at least one argument")
		return
	case 1:
		method = http.MethodGet
		url = args[0]
		return
	case 2:
		if !slices.Contains(subCommands, strings.ToUpper(args[0])) {
			err = errors.New("invalid sub command")
			return
		}
		method = strings.ToUpper(args[0])
		url = args[1]
		return
	default:
		err = errors.New("requires at most two arguments")
		return
	}
}

func constructHttpRequest(method, url, jsonBody string) (*http.Request, error) {
	if method == http.MethodPost {
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
