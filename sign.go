package main

import (
	"flag"
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewSignCmd() *SignCmd {
	bucket, err := GetBucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	gc := &SignCmd{
		fs:     flag.NewFlagSet("sign", flag.ContinueOnError),
		bucket: bucket,
	}
	gc.fs.StringVar(&gc.name, "channel", "", "channel name")
	gc.fs.Int64Var(&gc.expiry, "expiry", 10*365*86400, "seconds to expire the signature")

	return gc
}

type SignCmd struct {
	fs     *flag.FlagSet
	bucket *oss.Bucket

	name   string
	expiry int64
}

func (g *SignCmd) Name() string {
	return g.fs.Name()
}

func (g *SignCmd) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *SignCmd) Run() error {
	if g.name == "" {
		fmt.Println("No channel name provided!")
		return nil
	}
	g.signUrl()
	return nil
}

func (g *SignCmd) signUrl() {

	chInfo, err := g.bucket.GetLiveChannelInfo(g.name)
	if err != nil {
		HandleError(err)
	}
	signedURL, err := g.bucket.SignRtmpURL(g.name, chInfo.Target.PlaylistName, g.expiry)

	fmt.Printf("%s\n", signedURL)
}
