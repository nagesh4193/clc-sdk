package dc

import (
	"fmt"

	"github.com/CenturyLinkCloud/clc-sdk/api"
)

func New(client api.HTTP) *Service {
	return &Service{
		client: client,
		config: client.Config(),
	}
}

type Service struct {
	client api.HTTP
	config *api.Config
}

func (s *Service) Get(id string) (*Response, error) {
	url := fmt.Sprintf("%s/datacenters/%s/%s?groupLinks=true", s.config.BaseURL, s.config.Alias, id)
	dc := &Response{}
	err := s.client.Get(url, dc)
	return dc, err
}

func (s *Service) GetAll() ([]*Response, error) {
	url := fmt.Sprintf("%s/datacenters/%s", s.config.BaseURL, s.config.Alias)
	dcs := make([]*Response, 0)
	err := s.client.Get(url, &dcs)
	return dcs, err
}

func (s *Service) GetCapabilities(id string) (*CapabilitiesResponse, error) {
	url := fmt.Sprintf("%s/datacenters/%s/%s/deploymentCapabilities", s.config.BaseURL, s.config.Alias, id)
	c := &CapabilitiesResponse{}
	err := s.client.Get(url, c)
	return c, err
}

func (s *Service) GetBareMetalCapabilities(dataCenterId string) (*BareMetalCapabilitiesResponse, error) {
	url := fmt.Sprintf("%s/datacenters/%s/%s/bareMetalCapabilities", s.config.BaseURL, s.config.Alias, dataCenterId)
	bm := &BareMetalCapabilitiesResponse{}
	err := s.client.Get(url, bm)
	return bm, err
}

type Response struct {
	ID    string    `json:"id"`
	Name  string    `json:"name"`
	Links api.Links `json:"links"`
}

type CapabilitiesResponse struct {
	SupportsPremiumStorage     bool `json:"supportsPremiumStorage"`
	SupportsBareMetalServers   bool `json:"supportsBareMetalServers"`
	SupportsSharedLoadBalancer bool `json:"supportsSharedLoadBalancer"`
	Templates                  []struct {
		Name               string   `json:"name"`
		Description        string   `json:"description"`
		StorageSizeGB      int      `json:"storageSizeGB"`
		Capabilities       []string `json:"capabilities"`
		ReservedDrivePaths []string `json:"reservedDrivePaths"`
	} `json:"templates"`
	DeployableNetworks []struct {
		Name      string `json:"name"`
		NetworkId string `json:"networkId"`
		Type      string `json:"type"`
		AccountID string `json:"accountID"`
	} `json:"deployableNetworks"`
}

type BareMetalCapabilitiesResponse struct {
	SKUs []struct {
		ID           string  `json:"id"`
		HourlyRate   float32 `json:"hourlyRate"`
		Availability string  `json:"availability"`
		Memory       []struct {
			CapacityInGB int `json:"capacityGB"`
		} `json:"memory"`
		Processor struct {
			Sockets        int    `json:"sockets"`
			CoresPerSocket int    `json:"coresPerSocket"`
			Description    string `json:"description"`
		} `json:"processor"`
		Storage []struct {
			Type         string `json:"type"`
			CapacityInGB int    `json:"capacityGB"`
			SpeedInRPM   int    `json:"speedRpm"`
		} `json:"storage"`
	} `json:"skus"`
	OperatingSystems []struct {
		Type                string  `json:"type"`
		Description         string  `json:"description"`
		HourlyRatePerSocket float32 `json:"hourlyRatePerSocket"`
	} `json:"operatingSystems"`
}
