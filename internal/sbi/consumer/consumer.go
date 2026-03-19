package consumer

import (
	"github.com/acore2026/ausf/pkg/app"
	Nnrf_NFDiscovery "github.com/acore2026/openapi/nrf/NFDiscovery"
	Nnrf_NFManagement "github.com/acore2026/openapi/nrf/NFManagement"
	Nudm_UEAuthentication "github.com/acore2026/openapi/udm/UEAuthentication"
)

type ConsumerAusf interface {
	app.App
}

type Consumer struct {
	ConsumerAusf

	*nnrfService
	*nudmService
}

func NewConsumer(ausf ConsumerAusf) (*Consumer, error) {
	c := &Consumer{
		ConsumerAusf: ausf,
	}

	c.nnrfService = &nnrfService{
		consumer:        c,
		nfMngmntClients: make(map[string]*Nnrf_NFManagement.APIClient),
		nfDiscClients:   make(map[string]*Nnrf_NFDiscovery.APIClient),
	}

	c.nudmService = &nudmService{
		consumer:    c,
		ueauClients: make(map[string]*Nudm_UEAuthentication.APIClient),
	}

	return c, nil
}
