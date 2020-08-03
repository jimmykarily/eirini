package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"code.cloudfoundry.org/eirini"
	cmdcommons "code.cloudfoundry.org/eirini/cmd"
	"code.cloudfoundry.org/eirini/k8s/webhooks"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/tlsconfig"
	"github.com/jessevdk/go-flags"
	yaml "gopkg.in/yaml.v2"
)

type options struct {
	ConfigFile string `short:"c" long:"config" description:"Config for running event-reporter"`
}

func main() {
	var opts options
	_, err := flags.ParseArgs(&opts, os.Args)
	cmdcommons.ExitIfError(err)

	cfg, err := readConfig(opts.ConfigFile)
	cmdcommons.ExitIfError(err)

	logger := lager.NewLogger("env-injector")
	logger.RegisterSink(lager.NewPrettySink(os.Stdout, lager.DEBUG))

	webhook := webhooks.NewEnvInjectorWebhook(logger)

	serveTLS(logger, cfg, webhook.Handle)
}

func serveTLS(logger lager.Logger, cfg *eirini.EnvInjectorConfig, handler http.HandlerFunc) {
	tlsConfig, err := tlsconfig.Build(
		tlsconfig.WithInternalServiceDefaults(),
	).Server()

	server := &http.Server{
		Addr:      fmt.Sprintf("0.0.0.0:%d", cfg.TLSPort),
		TLSConfig: tlsConfig,
	}
	cmdcommons.ExitIfError(err)

	http.HandleFunc("/pods", handler)
	server.ListenAndServeTLS(cfg.ServerCertPath, cfg.ServerKeyPath)
}

func readConfig(path string) (*eirini.EnvInjectorConfig, error) {
	fileBytes, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}

	var conf eirini.EnvInjectorConfig
	err = yaml.Unmarshal(fileBytes, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
