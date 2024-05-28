package cmd

import (
	"fmt"
	"os"

	"github.com/deduardolima/stress_tester/internal/usecase"
	"github.com/spf13/cobra"
)

var (
	url         string
	requests    int
	concurrency int
)

func newStressTesterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stress-tester",
		Short: "A CLI tool to perform stress testing on a web service",
		RunE:  runStressTest,
	}

	cmd.Flags().StringVar(&url, "url", "", "URL of the service to test")
	cmd.Flags().IntVar(&requests, "requests", 100, "Total number of requests")
	cmd.Flags().IntVar(&concurrency, "concurrency", 10, "Number of concurrent requests")
	cmd.MarkFlagRequired("url")

	return cmd
}

func runStressTest(cmd *cobra.Command, args []string) error {
	if url == "" {
		return fmt.Errorf("URL is required")
	}

	if !(len(url) >= 7 && (url[:7] == "http://" || url[:8] == "https://")) {
		url = "http://" + url
	}

	usecase.RunStressTest(url, requests, concurrency)
	return nil
}

func Execute() {
	rootCmd := newStressTesterCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
