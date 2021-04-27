package mage

import "github.com/magefile/mage/mg"

type Adapters mg.Namespace

// Compile all adapters
func (Adapters) All() {
	mg.SerialDeps(
		Adapters.Gac,
		Adapters.Meteringd,
		Adapters.Networkcostd,
		Adapters.IoEnergyPublisher,
		Adapters.IoEnergySubscriber,
		Adapters.Sftpd,
		Adapters.Haugazel,
	)
}

func (Adapters) Gac() error {
	return build("gac", "GAC")
}

func (Adapters) Meteringd() error {
	return build("meteringd", "GAC - Metering")
}

func (Adapters) Networkcostd() error {
	return build("networkcostd", "Networkcost")
}

func (Adapters) IoEnergyPublisher() error {
	return build("ioenergy-publisher", "IO.Energy Publisher")
}

func (Adapters) IoEnergySubscriber() error {
	return build("ioenergy-subscriber", "IO.Energy Subscriber")
}

func (Adapters) Sftpd() error {
	return build("sftpd", "Sftpd")
}

func (Adapters) Haugazel() error {
	return build("haugazel", "Haugazel")
}
