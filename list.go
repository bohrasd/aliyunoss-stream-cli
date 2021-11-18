package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewListCmd() *ListCmd {
	bucket, err := GetBucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	gc := &ListCmd{
		fs:     flag.NewFlagSet("list", flag.ContinueOnError),
		bucket: bucket,
	}
	gc.fs.StringVar(&gc.prefix, "prefix", "", "channel prefix")

	return gc
}

type ListCmd struct {
	fs     *flag.FlagSet
	bucket *oss.Bucket

	prefix string
}

func (g *ListCmd) Name() string {
	return g.fs.Name()
}

func (g *ListCmd) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *ListCmd) Run() error {
	g.listLiveChannel()
	return nil
}

func (gc *ListCmd) listLiveChannel() {
	list, err := gc.bucket.ListLiveChannel(oss.Prefix(gc.prefix))
	if err != nil {
		HandleError(err)
	}

	//fmt.Printf("%#v\n", list)
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "name\tstatus\turl\tlast changed")
	for _, aCh := range list.LiveChannel {
		fmt.Fprintln(w, strings.Join([]string{aCh.Name, aCh.Status, aCh.PlayUrls[0], aCh.LastModified.String()}, "\t"))
	}
	w.Flush()
}
