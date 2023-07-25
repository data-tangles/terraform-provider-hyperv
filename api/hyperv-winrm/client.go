package hyperv_winrm

import (
	"github.com/qman-being/terraform-provider-hyperv/api"
	winrm_helper "github.com/qman-being/terraform-provider-hyperv/api/winrm-helper"
)

func New(clientConfig *ClientConfig) (*api.Provider, error) {
	return &api.Provider{
		Client: clientConfig,
	}, nil
}

type ClientConfig struct {
	WinRmClient winrm_helper.Client
}
