package main

import "github.com/kdihalas/divan/pkg/provider"

func GetProvider(conf map[string]interface{}) provider.Provider {
	if conf["provider"] == "couchdb" {
		pro := provider.NewCouchDbProvider(conf)
		return pro
	}

	return nil
}