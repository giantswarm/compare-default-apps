# Default apps provider comparison 

Example screenshot

<img width="1437" alt="image" src="https://github.com/giantswarm/compare-default-apps/blob/main/image.png">

## Usage

Simply execute the program like this ...

```nohighlight
go run main.go
```

... and see the result in the opening browser.

## Background

default-apps are our means to provision workload clusters with basic capabilities.

| Provider | Cluster app repository | Schema |
|-|-|-|
| AWS | [default-apps-aws](https://github.com/giantswarm/default-apps-aws) | [Schema](https://raw.githubusercontent.com/giantswarm/default-apps-aws/master/helm/default-apps-aws/values.yaml) |
| Cloud Director | [default-apps-cloud-director](https://github.com/giantswarm/default-apps-cloud-director) | [Schema](https://raw.githubusercontent.com/giantswarm/default-apps-cloud-director/main/helm/default-apps-cloud-director/values.yaml) |
| GCP | [default-apps-gcp](https://github.com/giantswarm/default-apps-gcp) | [Apps](https://raw.githubusercontent.com/giantswarm/default-apps-gcp/main/helm/default-apps-gcp/values.yaml) |
| OpenStack | [default-apps-openstack](https://github.com/giantswarm/default-apps-openstack) | [Schema](https://raw.githubusercontent.com/giantswarm/default-apps-openstack/main/helm/default-apps-openstack/values.yaml) |
| VSphere | [default-apps-vsphere](https://github.com/giantswarm/default-apps-vsphere) | [Schema](https://raw.githubusercontent.com/giantswarm/default-apps-vsphere/main/helm/default-apps-vsphere/values.yaml) |
| Azure | [default-apps-azure](https://github.com/giantswarm/default-apps-azure) | [Schema](https://raw.githubusercontent.com/giantswarm/default-apps-vsphere/main/helm/default-apps-vsphere/values.yaml) |

This repository provides tooling to visualize which apps are default on each provider.
