package main

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/hoomy-official/go-shared/pkg/net/do"
	"github.com/hoomy-official/go-sonos/api/v2/S19"
)

func main() {
	ctx := context.Background()
	u := MustParse("http://192.168.1.100:1400")

	var mediainfo S19.AVTransportGetMediaInfoRes

	err := S19.NewService(do.NewDoer(u)).AVTransport.GetMediaInfo(ctx, &mediainfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", mediainfo) //nolint:forbidigo // demo code
}

func MustParse(u string) *url.URL {
	parse, err := url.Parse(u)
	if err != nil {
		panic(err)
	}

	return parse
}
