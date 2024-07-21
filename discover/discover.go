package discover

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"slices"
	"time"

	s "github.com/koron/go-ssdp"
	"github.com/vanyda-official/go-shared/pkg/discover"
	"github.com/vanyda-official/go-shared/pkg/resolvers/ssdp"
	"golang.org/x/sync/errgroup"
)

const (
	httpRequestTimeout = 2 * time.Second
	discoverTimeout    = 15 * time.Second
)

type SonosDiscover struct {
	timeout     time.Duration
	retryCount  int
	stopOnCount int
}

type Option func(sonosDiscover *SonosDiscover)

func (o Option) apply(s *SonosDiscover) {
	o(s)
}

func WithDiscoveryPeriod(t time.Duration) Option {
	return func(sonosDiscover *SonosDiscover) {
		sonosDiscover.timeout = t
	}
}

// WithStopOnCount will end once the expected number of device is found. -1 means no limit.
func WithStopOnCount(count int) Option {
	return func(sonosDiscover *SonosDiscover) {
		sonosDiscover.stopOnCount = count
	}
}

func WithRetryCount(count int) Option {
	return func(sonosDiscover *SonosDiscover) {
		sonosDiscover.retryCount = count
	}
}

func NewSonosDiscover(opts ...Option) *SonosDiscover {
	defaultOptions := []Option{WithDiscoveryPeriod(discoverTimeout), WithStopOnCount(-1)}
	opts = append(defaultOptions, opts...)

	d := &SonosDiscover{}

	for _, opt := range opts {
		opt.apply(d)
	}

	return d
}

func (k SonosDiscover) Discover(ctx context.Context, discovered chan<- *SonosDevice) error {
	ch := make(chan *SonosService)
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return k.DiscoverService(ctx, ch)
	})

	g.Go(func() error {
		defer close(discovered)

		for service := range ch {
			if err := hydrate(ctx, service, discovered); err != nil {
				return err
			}
		}

		return nil
	})

	return g.Wait()
}

func hydrate(ctx context.Context, service *SonosService, discovered chan<- *SonosDevice) error {
	cl := http.Client{Timeout: httpRequestTimeout}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, service.Location, nil)
	if err != nil {
		return err
	}

	res, err := cl.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	all, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var device SonosDevice
	err = xml.Unmarshal(all, &device)
	if err != nil {
		return err
	}

	discovered <- &device
	return nil
}

func (k SonosDiscover) DiscoverService(ctx context.Context, discovered chan<- *SonosService) error {
	ctx, cancel := context.WithCancel(ctx)
	g, ctx := errgroup.WithContext(ctx)
	ch := make(chan interface{})

	g.Go(func() error {
		resolver := ssdp.New(
			ssdp.WithTransform(transformer),
			ssdp.WithWaitSecond(int(k.timeout.Seconds())/max(1, k.retryCount)),
			ssdp.WithRetry(max(1, k.retryCount)),
		)

		return discover.NewDiscover(resolver, discover.WithTimeout(k.timeout)).Discover(ctx, ch)
	})

	g.Go(func() error {
		defer close(discovered)

		var discoveredIDs []string

		for i := range ch {
			if gateway, ok := i.(*SonosService); ok {
				// deduplicate
				if !slices.Contains(discoveredIDs, gateway.USN) {
					discovered <- gateway
					discoveredIDs = append(discoveredIDs, gateway.USN)

					if k.stopOnCount != -1 && len(discoveredIDs) >= k.stopOnCount {
						cancel()
					}
				}
			}
		}

		return nil
	})

	return g.Wait()
}

func transformer(entry *s.Service) (interface{}, error) {
	header := entry.Header()
	return &SonosService{
		Type:                  entry.Type,
		USN:                   entry.USN,
		Location:              entry.Location,
		Server:                entry.Server,
		Bootseq:               header.Get("X-Rincon-Bootseq"),
		Household:             header.Get("X-Rincon-Household"),
		HouseholdSmartSpeaker: header.Get("Household.smartspeaker.audio"),
		WifiMode:              header.Get("X-Rincon-Wifimode"),
		Variant:               header.Get("X-Rincon-Variant"),
		SmartSpeaker:          header.Get("Location.smartspeaker.audio"),
		BootID:                header.Get("Bootid.upnp.org"),
	}, nil
}

type SonosService struct {
	USN                   string
	Location              string
	Server                string
	Bootseq               string
	Household             string
	HouseholdSmartSpeaker string
	WifiMode              string
	Variant               string
	SmartSpeaker          string
	Type                  string
	BootID                string
}
