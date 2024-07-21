package discover

import "encoding/xml"

type SonosDevice struct {
	XMLName     xml.Name   `xml:"root"`
	Xmlns       string     `xml:"xmlns,attr"`
	SpecVersion Version    `xml:"specVersion"`
	Device      RootDevice `xml:"device"`
}

type Version struct {
	Major string `xml:"major"`
	Minor string `xml:"minor"`
}

type RootDevice struct {
	CommonDevice

	DeviceList struct {
		Device []SubDevice `xml:"device"`
	} `xml:"deviceList"`

	SoftwareVersion         string `xml:"softwareVersion"`
	SwGen                   string `xml:"swGen"`
	HardwareVersion         string `xml:"hardwareVersion"`
	SerialNum               string `xml:"serialNum"`
	MACAddress              string `xml:"MACAddress"`
	MinCompatibleVersion    string `xml:"minCompatibleVersion"`
	LegacyCompatibleVersion string `xml:"legacyCompatibleVersion"`
	APIVersion              string `xml:"apiVersion"`
	MinAPIVersion           string `xml:"minApiVersion"`
	DisplayVersion          string `xml:"displayVersion"`
	ExtraVersion            string `xml:"extraVersion"`
	NsVersion               string `xml:"nsVersion"`
	RoomName                string `xml:"roomName"`
	DisplayName             string `xml:"displayName"`
	ZoneType                string `xml:"zoneType"`
	Feature1                string `xml:"feature1"`
	Feature2                string `xml:"feature2"`
	Feature3                string `xml:"feature3"`
	SeriesID                string `xml:"seriesid"`
	Variant                 string `xml:"variant"`
	InternalSpeakerSize     string `xml:"internalSpeakerSize"`
	Memory                  string `xml:"memory"`
	Flash                   string `xml:"flash"`
	FlashRepartitioned      string `xml:"flashRepartitioned"`
	AmpOnTime               string `xml:"ampOnTime"`
	RetailMode              string `xml:"retailMode"`
	SSLPort                 string `xml:"SSLPort"`
	SecurehhSSLPort         string `xml:"securehhSSLPort"`
}

type CommonDevice struct {
	DeviceType       string `xml:"deviceType"`
	FriendlyName     string `xml:"friendlyName"`
	Manufacturer     string `xml:"manufacturer"`
	ManufacturerURL  string `xml:"manufacturerURL"`
	ModelNumber      string `xml:"modelNumber"`
	ModelDescription string `xml:"modelDescription"`
	ModelName        string `xml:"modelName"`
	ModelURL         string `xml:"modelURL"`
	UDN              string `xml:"UDN"`
	ServiceList      struct {
		Service []Service `xml:"service"`
	} `xml:"serviceList"`
	IconList struct {
		Icon []Icon `xml:"icon"`
	} `xml:"iconList"`
}

type SubDevice struct {
	CommonDevice

	XRhapsodyExtension struct {
		Xmlns              string `xml:"xmlns,attr"`
		DeviceID           string `xml:"deviceID"`
		DeviceCapabilities struct {
			Text               string `xml:",chardata"`
			InteractionPattern struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"interactionPattern"`
		} `xml:"deviceCapabilities"`
	} `xml:"X_Rhapsody-Extension"`
	XQPlaySoftwareCapability struct {
		Qq string `xml:"qq,attr"`
	} `xml:"X_QPlay_SoftwareCapability"`
}

type Icon struct {
	ID       string `xml:"id"`
	Mimetype string `xml:"mimetype"`
	Width    string `xml:"width"`
	Height   string `xml:"height"`
	Depth    string `xml:"depth"`
	URL      string `xml:"url"`
}

type Service struct {
	ServiceType string `xml:"serviceType"`
	ServiceID   string `xml:"serviceId"`
	ControlURL  string `xml:"controlURL"`
	EventSubURL string `xml:"eventSubURL"`
	SCPDURL     string `xml:"SCPDURL"`
}
