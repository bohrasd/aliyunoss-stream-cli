package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewInfoCmd() *InfoCmd {
	bucket, err := GetBucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	gc := &InfoCmd{
		fs:     flag.NewFlagSet("info", flag.ContinueOnError),
		bucket: bucket,
	}
	gc.fs.StringVar(&gc.name, "channel", "", "channel name")

	return gc
}

type InfoCmd struct {
	fs     *flag.FlagSet
	bucket *oss.Bucket

	name string
}

func (g *InfoCmd) Name() string {
	return g.fs.Name()
}

func (g *InfoCmd) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *InfoCmd) Run() error {
	if g.name == "" {
		fmt.Println("No channel name provided!")
		return nil
	}
	_, err := g.bucket.GetLiveChannelInfo(g.name)
	if err != nil {
		HandleError(err)
	}
	g.liveChannelInfo()
	return nil
}

func (gc *InfoCmd) liveChannelInfo() {
	chInfoList, err := gc.bucket.ListLiveChannel(oss.Prefix(gc.name))
	if err != nil {
		HandleError(err)
	}

	if len(chInfoList.LiveChannel) > 0 {
		chInfo := chInfoList.LiveChannel[0]

		fmt.Printf("Name: %s\n", gc.name)
		fmt.Printf("Status: %s\n", chInfo.Status)
		fmt.Printf("Description: %s\n", chInfo.Description)
		fmt.Printf("Play Url: %s\n", chInfo.PlayUrls[0])
		fmt.Printf("Publish Url: %s\n", chInfo.PublishUrls[0])
	} else {
		fmt.Println("No channel founded")

		return
	}

	stat, err := gc.bucket.GetLiveChannelStat(gc.name)

	fmt.Println("Stats: ")

	statJson, err := json.MarshalIndent(stat, "", "  ")
	fmt.Printf("%s\n", string(statJson))
}
