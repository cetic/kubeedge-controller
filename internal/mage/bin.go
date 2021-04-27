package mage

import (
	"github.com/magefile/mage/mg"
)

type Bin mg.Namespace

// Compile all binaries
func (Bin) All() {
	mg.SerialDeps(
		Bin.MicroMetering,
		Bin.MicroConfiguration,
		Bin.MicroInvoicing,
		Bin.MicroIdentityaccess,
		Bin.FrontendApi,
		Bin.DsoApi,
		Bin.AdminApi,
		Bin.NetworkcostApi,
		Bin.Componentd,
		Bin.MicroSharedSettings,
	)
}

func (Bin) MicroMetering() error {
	return build("metering-srv", "Metering Service")
}

func (Bin) MicroConfiguration() error {
	return build("configuration-srv", "Configuration Service")
}

func (Bin) MicroInvoicing() error {
	return build("invoicing-srv", "Invoicing Service")
}

func (Bin) MicroSharedSettings() error {
	return build("shared_settings-srv", "Shared Settings Service")
}

func (Bin) MicroIdentityaccess() error {
	return build("identityaccess-srv", "IdentityAccess Service")
}

func (Bin) FrontendApi() error {
	return build("frontend-api", "Frontend API")
}
func (Bin) DsoApi() error {
	return build("dso-api", "DSO API")
}

func (Bin) AdminApi() error {
	return build("admin-api", "Admin API")
}

func (Bin) NetworkcostApi() error {
	return build("networkcost-api", "Networkcosts API")
}

func (Bin) Componentd() error {
	return build("componentd", "Componentd")
}
