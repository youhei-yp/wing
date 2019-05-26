// Copyright (c) 2018-2019 Xunmo All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2018/12/01   youhei         New version
// -------------------------------------------------------------------

package utils

import (
	"github.com/BurntSushi/toml"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"wing/logger"
)

// Configs struct to load configs from toml file
type Configs struct {
	once    sync.Once   // start sign moniter only once
	lock    sync.Mutex  // read and write sync lock
	cfgPath string      // configs file path
	cfgData interface{} // configs content
}

// configsMap configs map, key is config file path
var configsMap map[string]*Configs

// ReloadCallback configs reloaded event
type ReloadCallback func(last interface{})

// NewConfigs create singleton configures object
func NewConfigs(path string, container interface{}) *Configs {
	if configsMap == nil {
		configsMap = make(map[string]*Configs)
	}

	if c := configsMap[path]; c != nil {
		logger.W("Exust configs:", c, "for", path, "override it!!")
		delete(configsMap, path)
	}

	logger.I("Created a configs object")
	c := &Configs{cfgPath: path, cfgData: container}
	configsMap[path] = c
	return c
}

// GetConfigs return Config object by file path
func GetConfigs(path string) *Configs {
	if configsMap == nil {
		configsMap = make(map[string]*Configs)
	}

	if c := configsMap[path]; c != nil {
		logger.W("Got configs:", c, "for", path)
		return c
	}

	logger.W("Unexust configs for", path)
	return nil
}

// Get return configs content
func (c *Configs) Get() interface{} {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c := configsMap[c.cfgPath]; c != nil {
		return c.cfgData
	}
	return nil
}

// Load load configs from toml file
func (c *Configs) Load() error {
	c.lock.Lock()
	defer c.lock.Unlock()

	f, err := filepath.Abs(c.cfgPath)
	if err != nil {
		logger.E("Config file:", c.cfgPath, "err:", err)
		return err
	}

	if _, err := toml.DecodeFile(f, c.cfgData); err != nil {
		logger.E("Parse config file:", f, "err:", err)
		return err
	}

	logger.D("Parsed config file:", f, "data:", c.cfgData)
	return nil
}

// Start start moniter listen sign to reload config
// @notic :
//     you can send sign to reload as follow steps.
//     (1). login server pc and open terminal
//     (2). execute command
//          ~$: kill -SIGUSR1 {running port of go server}
func (c *Configs) Start(cb ReloadCallback) error {
	logger.D("Load configs before start monitor")
	if err := c.Load(); err != nil {
		return err
	}

	// start sign moniter to reload config as hot configuration
	c.once.Do(func() {
		sign := make(chan os.Signal, 1)
		signal.Notify(sign, syscall.SIGUSR1)
		callback := cb

		go func(cfg *Configs) {
			for {
				<-sign
				if cfg == nil {
					logger.E("Failed reloaded configs")
					return
				}
				logger.D("Reloading configs...")
				err := cfg.Load()
				if err == nil && callback != nil {
					callback(cfg.cfgData)
				}
			}
		}(c)
	})
	return nil
}
