package main

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
	"github.com/julienschmidt/httprouter"
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

	webhook := webhooks.NewEnvironmentInjectorWebhook(logger)

	serveTLS(logger, cfg, webhook.Handle)
}

func serveTLS(logger lager.Logger, cfg *eirini.EnvironmentInjectorConfig, handler httprouter.Handle) {
	tlsConfig, err := tlsconfig.Build(
		tlsconfig.WithInternalServiceDefaults(),
	).Server()

	httpHandler := httprouter.New()
	httpHandler.POST("/inject-instance-index", handler)

	server := &http.Server{
		Addr:      fmt.Sprintf("0.0.0.0:%d", cfg.TLSPort),
		TLSConfig: tlsConfig,
		Handler:   httpHandler,
	}
	cmdcommons.ExitIfError(err)

	logger.Fatal("server-crashed", server.ListenAndServeTLS(cfg.ServerCertPath, cfg.ServerKeyPath))
}

func readConfig(path string) (*eirini.EnvironmentInjectorConfig, error) {
	fileBytes, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}

	var conf eirini.EnvironmentInjectorConfig
	err = yaml.Unmarshal(fileBytes, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
