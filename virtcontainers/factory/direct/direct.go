// Copyright (c) 2018 HyperHQ Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// direct implements base vm factory without vm templating.

package direct

import (
	vc "github.com/kata-containers/runtime/virtcontainers"
	"github.com/kata-containers/runtime/virtcontainers/factory/base"
)

type direct struct {
	config vc.VMConfig
}

// New returns a new direct vm factory.
func New(config vc.VMConfig) base.FactoryBase {
	return &direct{config}
}

// Config returns the direct factory's configuration.
func (d *direct) Config() vc.VMConfig {
	return d.config
}

// GetBaseVM create a new VM directly.
func (d *direct) GetBaseVM() (*vc.VM, error) {
	vm, err := vc.NewVM(d.config)
	if err != nil {
		return nil, err
	}

	err = vm.Pause()
	if err != nil {
		vm.Stop()
		return nil, err
	}

	return vm, nil
}

// CloseFactory closes the direct vm factory.
func (d *direct) CloseFactory() {
}
