// Package analysis analyses schemas and creates comparison.
package analysis

import (
	"io/ioutil"
	"net/http"
	"sort"

	"gopkg.in/yaml.v3"
)

const (
	recursionLimit = 20
)

// ClusterApp holds information on a Giant Swarm cluster app for a Cluster API provider.
type DefaultApp struct {
	// User-friendly name of the Cluster API infrastructure provider.
	ProviderName string

	// URL of the GitHub repo landing page.
	RepositoryURL string

	// URL of the schema file in JSON for download.
	ValuesURL string
}

// Analyser is the agent that performs comparison and analysis on the schemas.
type Analyser struct {
	// The cluster apps handed over to the agent.
	DefaultApps []DefaultApp

	// Schemas holds the full schema information in the original hierarchical
	// form. Map key is the provider name.
	Values map[string]Values
}

type Values struct {
	ClusterName       string                 `yaml:"clusterName"`
	Organization      string                 `yaml:"organization"`
	ManagementCluster string                 `yaml:"managementCluster"`
	UserConfig        map[string]interface{} `yaml:"userConfig"`
	Apps              map[string]App         `yaml:"apps"`
}

type App struct {
	AppName       string                 `yaml:"appName"`
	ChartName     string                 `yaml:"chartName"`
	Catalog       string                 `yaml:"catalog"`
	ClusterValues map[string]interface{} `yaml:"clusterValues"`
	ForceUpgrade  bool                   `yaml:"forceUpgrade"`
	Namespace     string                 `yaml:"namespace"`
	Version       string                 `yaml:"version"`
}

func New(defaultApps []DefaultApp) (*Analyser, error) {
	a := &Analyser{
		DefaultApps: defaultApps,
		Values:      make(map[string]Values),
	}

	for _, defaultApp := range defaultApps {
		values, err := getValues(defaultApp.ValuesURL)
		if err != nil {
			return nil, err
		}

		a.Values[defaultApp.ProviderName] = *values
	}

	return a, nil
}

type ProviderAppsSummary struct {
	Version string
	Catalog string
}

func (a *Analyser) AllApps() (keys []string) {
	for key := range a.MergedApps() {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func (a *Analyser) MergedApps() map[string]map[string]ProviderAppsSummary {
	// Create complete list of all apps with a list of the providers having them.
	fullApps := make(map[string]map[string]ProviderAppsSummary)
	for _, defaultApp := range a.DefaultApps {
		// Collect all apps
		for _, v := range a.Values[defaultApp.ProviderName].Apps {
			if _, exists := fullApps[v.AppName]; !exists {
				fullApps[v.AppName] = make(map[string]ProviderAppsSummary)
			}

			fullApps[v.AppName][defaultApp.ProviderName] = ProviderAppsSummary{
				Version: v.Version,
				Catalog: v.Catalog,
			}
		}
	}

	return fullApps
}

func getValues(url string) (values *Values, err error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Copy data from the response to standard output
	yaml.Unmarshal(bytes, &values) //use package "io" and "os"
	return values, nil
}

// Providers returns the list of provider names in the order of definition.
func (a *Analyser) Providers() (providers []string) {
	for _, defaultApp := range a.DefaultApps {
		providers = append(providers, defaultApp.ProviderName)
	}

	return providers
}
