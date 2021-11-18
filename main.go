package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func main() {

	if err := root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

var (
	endpoint   = os.Getenv("OSS_ENDPOINT")
	accessID   = os.Getenv("OSS_AK")
	accessKey  = os.Getenv("OSS_SK")
	bucketName = os.Getenv("OSS_BUCKET")
)

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

func root(args []string) error {
	if len(args) < 1 {
		return errors.New("You must pass a sub-command")
	}

	cmds := []Runner{
		NewCreateCmd(),
		NewListCmd(),
		NewDelCmd(),
		NewInfoCmd(),
		NewSignCmd(),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}

	return fmt.Errorf("Unknown subcommand: %s", subcommand)
}

// GetBucket creates the test bucket
func GetBucket(bucketName string) (*oss.Bucket, error) {
	// New client
	client, err := oss.New(endpoint, accessID, accessKey)
	if err != nil {
		return nil, err
	}

	// Get bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

// HandleError is the error handling method in the sample code
func HandleError(err error) {
	fmt.Println("occurred error:", err)
	os.Exit(-1)
}
