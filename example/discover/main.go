package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/hoomy-official/go-sonos/discover"
	"golang.org/x/sync/errgroup"
)

const (
	defaultDiscoveryDuration = 15 * time.Second
	defaultRetryCount        = 5
	defaultStopOnCount       = -1
)

func main() {
	var stopOnCount int
	var retryCount int
	var discoveryDuration time.Duration

	flag.DurationVar(&discoveryDuration, "duration", defaultDiscoveryDuration, "Duration for discovering devices")
	flag.IntVar(&retryCount, "retry", defaultRetryCount, "Number of retry for discovering new devices")
	flag.IntVar(&stopOnCount, "count", defaultStopOnCount, "Stop when the device count is reached")

	flag.Parse()

	ctx := context.Background()
	ch := make(chan *discover.SonosDevice)

	sonos := discover.NewSonosDiscover(
		discover.WithStopOnCount(stopOnCount),
		discover.WithRetryCount(retryCount),
		discover.WithDiscoveryPeriod(discoveryDuration),
	)
	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		for i := range ch {
			//nolint:forbidigo // demo code
			fmt.Printf("\n\n%#v\n\n", i.Device.UDN)
		}

		return nil
	})

	group.Go(func() error {
		return sonos.Discover(ctx, ch)
	})

	err := group.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
