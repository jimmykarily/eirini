module github.com/cloudfoundry-incubator/eirini-env-injector

go 1.14

require (
	github.com/SUSE/eirinix v0.2.1-0.20200420122346-85a6c535b0ad
	github.com/pelletier/go-toml v1.3.0 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.6.3 // indirect
	go.uber.org/zap v1.14.1
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.0.0-20200404061942-2a93acf49b83
	k8s.io/client-go v11.0.0+incompatible
	sigs.k8s.io/controller-runtime v0.4.0
)

replace code.cloudfoundry.org/cf-operator => code.cloudfoundry.org/quarks-operator v1.0.1-0.20200413083459-fb39a29ad746
