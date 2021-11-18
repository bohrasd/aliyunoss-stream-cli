package main

import (
	"flag"
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewCreateCmd() *CreateCmd {
	bucket, err := GetBucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	gc := &CreateCmd{
		fs:     flag.NewFlagSet("create", flag.ContinueOnError),
		bucket: bucket,
	}
	gc.fs.StringVar(&gc.name, "channel", "", "channel name")
	gc.fs.StringVar(&gc.playlist, "playlist", "", "playlist name")

	return gc
}

type CreateCmd struct {
	fs     *flag.FlagSet
	bucket *oss.Bucket

	name, playlist string
}

func (g *CreateCmd) Name() string {
	return g.fs.Name()
}

func (g *CreateCmd) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *CreateCmd) Run() error {
	if g.name == "" {
		fmt.Println("No channel name provided!")
		return nil
	}
	_, err := g.bucket.GetLiveChannelInfo(g.name)
	if err == nil {
		fmt.Println("Channel Existed!")
		return nil
	}
	g.createLiveChannelSample()
	return nil
}

// createLiveChannelSample Samples for create a live-channel
func (gc *CreateCmd) createLiveChannelSample() {
	// Case 1 - Create live-channel with Completely configure
	config := oss.LiveChannelConfiguration{
		Description: gc.name,   //description information, up to 128 bytes
		Status:      "enabled", //enabled or disabled
		Target: oss.LiveChannelTarget{
			Type:         "HLS", //the type of object, only supports HLS, required
			FragDuration: 10,    //the length of each ts object (in seconds), in the range [1,100], default: 5
			FragCount:    4,     //the number of ts objects in the m3u8 object, in the range of [1,100], default: 3
			PlaylistName: gc.playlist,
		},
	}

	result, err := gc.bucket.CreateLiveChannel(gc.name, config)
	if err != nil {
		HandleError(err)
	}

	playURL := result.PlayUrls[0]
	publishURL := result.PublishUrls[0]
	fmt.Printf("create livechannel: %s\nplayURL: %s\npublishURL: %s\n", gc.name, playURL, publishURL)

	signedURL, err := gc.bucket.SignRtmpURL(gc.name, gc.playlist, 10*365*86400)

	fmt.Printf("Signed Url: %s\n", signedURL)
}
