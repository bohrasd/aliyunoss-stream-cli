package main

import (
	"flag"
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewDelCmd() *DelCmd {
	bucket, err := GetBucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	gc := &DelCmd{
		fs:     flag.NewFlagSet("del", flag.ContinueOnError),
		bucket: bucket,
	}
	gc.fs.StringVar(&gc.name, "channel", "", "name of the channel to be deleted")

	return gc
}

type DelCmd struct {
	fs     *flag.FlagSet
	bucket *oss.Bucket

	name string
}

func (g *DelCmd) Name() string {
	return g.fs.Name()
}

func (g *DelCmd) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *DelCmd) Run() error {
	if g.name == "" {
		fmt.Println("No channel name provided!")
		return nil
	}
	g.delLiveChannel()
	return nil
}

func (gc *DelCmd) delLiveChannel() {
	err := gc.bucket.DeleteLiveChannel(gc.name)
	if err != nil {
		HandleError(err)
	}

	fmt.Println("channel Deleted")

}
