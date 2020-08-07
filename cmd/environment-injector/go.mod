module github.com/cloudfoundry-incubator/eirini-env-injector

go 1.14

require (
	github.com/SUSE/eirinix v0.2.1-0.20200420122346-85a6c535b0ad
	github.com/spf13/cobra v0.0.7
	go.uber.org/zap v1.14.1
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/sys v0.0.0-20200501145240-bc7a7d42d5c3 // indirect
	golang.org/x/tools v0.0.0-20200504193531-9bfbc385433f // indirect
	gopkg.in/yaml.v2 v2.2.8
	k8s.io/api v0.0.0-20200404061942-2a93acf49b83
	k8s.io/client-go v0.0.0-20200330143601-07e69aceacd6
	sigs.k8s.io/controller-runtime v0.4.0
)

replace code.cloudfoundry.org/cf-operator => code.cloudfoundry.org/quarks-operator v1.0.1-0.20200413083459-fb39a29ad746
