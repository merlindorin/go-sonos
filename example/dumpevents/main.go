package main

import (
	"context"
	"crypto/rand"
	"encoding/xml"
	"flag"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/hoomy-official/go-shared/pkg/net/do"
	"github.com/hoomy-official/go-sonos/api/v2/S19"
	"golang.org/x/sync/errgroup"
)

const (
	defaultTimeout         = 3 * time.Second
	defaultPortRangeMin    = 50000
	defaultPortRangeMax    = 55000
	defaultSonosDevicePort = "1400"
)

func main() {
	var publicAddress string
	var sonosAddress string
	var sonosPort string
	var host string
	var port string
	var timeout time.Duration

	nBig, err := rand.Int(rand.Reader, big.NewInt(defaultPortRangeMax-defaultPortRangeMin))
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&port, "http-port", strconv.Itoa(int(nBig.Int64())+defaultPortRangeMin), "http port")
	flag.StringVar(&host, "http-host", "", "http host, empty means all")
	flag.DurationVar(&timeout, "http-timeout", defaultTimeout, "http timeout")
	flag.StringVar(&publicAddress, "public-address", "", "public address, can be a domain name or an IP")
	flag.StringVar(&sonosAddress, "sonos-address", "", "sonos device address")
	flag.StringVar(&sonosPort, "sonos-port", defaultSonosDevicePort, "sonos device port")

	flag.Parse()

	if publicAddress == "" {
		log.Fatal("a public IP accessible by the sonos is mandatory")
	}

	ctx := context.Background()
	u := MustParse(net.JoinHostPort(sonosAddress, sonosPort))
	c := MustParse(net.JoinHostPort(publicAddress, port))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		all, er := io.ReadAll(request.Body)
		if er != nil {
			log.Println(er)
			return
		}

		var p Propertyset

		er = xml.Unmarshal(all, &p)
		if er != nil {
			log.Println(er)
			return
		}

		var e Event

		er = xml.Unmarshal([]byte(p.Property.LastChange), &e)
		if er != nil {
			log.Println(er)
			return
		}

		writer.WriteHeader(http.StatusOK)
	})

	group, ctx := errgroup.WithContext(ctx)

	server := &http.Server{
		Addr:              net.JoinHostPort(host, port),
		ReadHeaderTimeout: timeout,
	}

	group.Go(func() error {
		return server.ListenAndServe()
	})

	group.Go(func() error {
		return S19.NewService(do.NewDoer(u)).AVTransport.SubscribeEvent(ctx, c)
	})

	if err = group.Wait(); err != nil {
		log.Fatal(err)
	}
}

func MustParse(u string) *url.URL {
	parse, err := url.Parse(u)
	if err != nil {
		panic(err)
	}

	return parse
}

type Propertyset struct {
	XMLName  xml.Name `xml:"propertyset"`
	Text     string   `xml:",chardata"`
	E        string   `xml:"e,attr"`
	Property struct {
		Text       string `xml:",chardata"`
		LastChange string `xml:"LastChange"`
	} `xml:"property"`
}

type Event struct {
	XMLName    xml.Name `xml:"Event"`
	Text       string   `xml:",chardata"`
	Xmlns      string   `xml:"xmlns,attr"`
	R          string   `xml:"r,attr"`
	InstanceID struct {
		Text           string `xml:",chardata"`
		Val            string `xml:"val,attr"`
		TransportState struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"TransportState"`
		CurrentPlayMode struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"CurrentPlayMode"`
		CurrentCrossfadeMode struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"CurrentCrossfadeMode"`
		NumberOfTracks struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"NumberOfTracks"`
		CurrentTrack struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"CurrentTrack"`
		CurrentSection struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"CurrentSection"`
		CurrentTrackURI struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"CurrentTrackURI"`
		CurrentTrackDuration struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"CurrentTrackDuration"`
		CurrentTrackMetaData struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"CurrentTrackMetaData"`
		NextTrackURI struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"NextTrackURI"`
		NextTrackMetaData struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"NextTrackMetaData"`
		EnqueuedTransportURI struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"EnqueuedTransportURI"`
		EnqueuedTransportURIMetaData struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"EnqueuedTransportURIMetaData"`
		PlaybackStorageMedium struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"PlaybackStorageMedium"`
		AVTransportURI struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"AVTransportURI"`
		AVTransportURIMetaData struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"AVTransportURIMetaData"`
		NextAVTransportURI struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"NextAVTransportURI"`
		NextAVTransportURIMetaData struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"NextAVTransportURIMetaData"`
		CurrentTransportActions struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"CurrentTransportActions"`
		CurrentValidPlayModes struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"CurrentValidPlayModes"`
		DirectControlClientID struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"DirectControlClientID"`
		DirectControlIsSuspended struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"DirectControlIsSuspended"`
		DirectControlAccountID struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"DirectControlAccountID"`
		TransportStatus struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"TransportStatus"`
		SleepTimerGeneration struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"SleepTimerGeneration"`
		AlarmRunning struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"AlarmRunning"`
		SnoozeRunning struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"SnoozeRunning"`
		RestartPending struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"RestartPending"`
		TransportPlaySpeed struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"TransportPlaySpeed"`
		CurrentMediaDuration struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"CurrentMediaDuration"`
		RecordStorageMedium struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"RecordStorageMedium"`
		PossiblePlaybackStorageMedia struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"PossiblePlaybackStorageMedia"`
		PossibleRecordStorageMedia struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"PossibleRecordStorageMedia"`
		RecordMediumWriteStatus struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"RecordMediumWriteStatus"`
		CurrentRecordQualityMode struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"CurrentRecordQualityMode"`
		PossibleRecordQualityModes struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"PossibleRecordQualityModes"`
	} `xml:"InstanceID"`
}
