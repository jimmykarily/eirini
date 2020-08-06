package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	eirinix "github.com/SUSE/eirinix"
	injector "github.com/cloudfoundry-incubator/eirini-env-injector/cmd/envinjector"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc" // from https://github.com/kubernetes/client-go/issues/345
)

type EnvironmentInjectorConfig struct {
	ServiceName      string `yaml:"service_name"`
	ServiceNamespace string `yaml:"service_namespace"`
	ServicePort      int    `yaml:"service_port"`
	ConfigPath       string `yaml:"kube_config_path"`
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the eirini extension",
	Run: func(cmd *cobra.Command, args []string) {
		defer log.Sync()
		path, err := cmd.Flags().GetString("config")
		ExitIfError(err)

		if path == "" {
			Exitf("--config is missing")
		}

		cfg := setConfigFromFile(path)

		// webhookHost := viper.GetString("operator-webhook-host")
		// webhookPort := viper.GetInt32("operator-webhook-port")
		// serviceName := viper.GetString("operator-service-name")
		// webhookNamespace := viper.GetString("operator-webhook-namespace")
		// register := viper.GetBool("register")

		// if webhookHost == "" {
		// 	log.Fatal("required flag 'operator-webhook-host' not set (env variable: OPERATOR_WEBHOOK_HOST)")
		// }

		// RegisterWebhooks := true
		// if !register {
		// 	log.Info("The extension will start without registering")
		// 	RegisterWebhooks = false
		// }

		register := true
		x := eirinix.NewManager(
			eirinix.ManagerOptions{
				// Namespace:        namespace,
				// Host:             webhookHost,
				// Port:             cfg.ServicePort,
				Host:             "0.0.0.0",
				KubeConfig:       cfg.ConfigPath,
				ServiceName:      cfg.ServiceName,
				WebhookNamespace: cfg.ServiceNamespace,
				RegisterWebHook:  &register,
			})

		x.AddExtension(injector.New())

		log.Fatal(x.Start())
	},
}

func init() {
	startCmd.Flags().BoolP("register", "r", true, "Register the extension")

	rootCmd.AddCommand(startCmd)
}

func setConfigFromFile(path string) *EnvironmentInjectorConfig {
	fileBytes, err := ioutil.ReadFile(filepath.Clean(path))
	ExitIfError(err)

	var conf EnvironmentInjectorConfig
	err = yaml.Unmarshal(fileBytes, &conf)
	ExitIfError(err)

	return &conf
}

func ExitIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func Exitf(messageFormat string, args ...interface{}) {
	panic(fmt.Sprintf(messageFormat, args...))
}
