package main

import (
	"encoding/json"
	"log"

	"github.com/giantswarm/microerror"
	"github.com/skratchdot/open-golang/open"

	"github.com/giantswarm/default-apps-cmp/pkg/analysis"
	"github.com/giantswarm/default-apps-cmp/pkg/server"
)

var (
	// Define all default-apps apps to analyse schema for here.
	defaultApps = []analysis.DefaultApp{
		{
			ProviderName:  "AWS",
			RepositoryURL: "https://github.com/giantswarm/default-apps-aws",
			ValuesURL:     "https://raw.githubusercontent.com/giantswarm/default-apps-aws/master/helm/default-apps-aws/values.yaml",
		},
		{
			ProviderName:  "Azure",
			RepositoryURL: "https://github.com/giantswarm/default-apps-azure",
			ValuesURL:     "https://raw.githubusercontent.com/giantswarm/default-apps-azure/main/helm/default-apps-azure/values.yaml",
		},
		{
			ProviderName:  "Cloud Director",
			RepositoryURL: "https://github.com/giantswarm/default-apps-cloud-director",
			ValuesURL:     "https://raw.githubusercontent.com/giantswarm/default-apps-cloud-director/main/helm/default-apps-cloud-director/values.yaml",
		},
		{
			ProviderName:  "VSphere",
			RepositoryURL: "https://github.com/giantswarm/default-apps-vsphere",
			ValuesURL:     "https://raw.githubusercontent.com/giantswarm/default-apps-vsphere/main/helm/default-apps-vsphere/values.yaml",
		},
	}

	url = "http://localhost:8080/"
)

// Data is a big data structure we deliver to the web UI as JSON,
// containing all the information we want to display to users.
type Data struct {
	DefaultApps []analysis.DefaultApp
	Providers   []string

	// List of all properties with hierarchical name
	AppNames []string

	// Map of apps (key) and information per provider.
	AppsAndProviders map[string]map[string]analysis.ProviderAppsSummary
}

func main() {
	analyser, err := analysis.New(defaultApps)
	if err != nil {
		log.Fatal(microerror.Mask(err))
	}

	data := Data{
		DefaultApps:      defaultApps,
		Providers:        analyser.Providers(),
		AppsAndProviders: analyser.MergedApps(),
		AppNames:         analyser.AllApps(),
	}
	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Opening browser at %s", url)
	err = open.Start(url)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Serve(8080, dataJson)
	if err != nil {
		log.Fatal(err)
	}
}
