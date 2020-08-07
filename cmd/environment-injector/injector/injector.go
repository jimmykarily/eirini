package injector

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"

	eirinix "github.com/SUSE/eirinix"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// Extension changes pod definitions
type Extension struct{ Logger *zap.SugaredLogger }

// New returns the persi extension
func New() eirinix.Extension {
	return &Extension{}
}

func (ext *Extension) injectInstanceID(pod *corev1.Pod) error {
	index, err := parseInstanceIndex(pod.Name)
	if err != nil {
		return err
	}

	for i := range pod.Spec.Containers {
		c := &pod.Spec.Containers[i]
		if c.Name == "opi" {
			cfInstanceVar := v1.EnvVar{Name: "CF_INSTANCE_ID", Value: strconv.Itoa(index)}
			c.Env = append(c.Env, cfInstanceVar)

			return nil
		}
	}

	return nil
}

func parseInstanceIndex(podName string) (int, error) {
	nameSegments := strings.Split(podName, "-")
	numSegments := len(nameSegments)
	if numSegments == 0 {
		return -1, fmt.Errorf("invalid pod name: %q. Pod names should contain dashes", podName)
	}
	instanceIndex, err := strconv.Atoi(nameSegments[numSegments-1])
	if err != nil {
		return -1, fmt.Errorf("error parsing instanceIndex for pod %q: %w", podName, err)
	}
	return instanceIndex, nil
}

// Handle manages volume claims for ExtendedStatefulSet pods
func (ext *Extension) Handle(ctx context.Context, eiriniManager eirinix.Manager, pod *corev1.Pod, req admission.Request) admission.Response {

	if pod == nil {
		return admission.Errored(http.StatusBadRequest, errors.New("No pod could be decoded from the request"))
	}

	_, file, _, _ := runtime.Caller(0)
	log := eiriniManager.GetLogger().Named(file)

	ext.Logger = log
	podCopy := pod.DeepCopy()
	log.Debugf("Handling webhook request for POD: %s (%s)", podCopy.Name, podCopy.Namespace)

	err := ext.injectInstanceID(podCopy)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	return eiriniManager.PatchFromPod(req, podCopy)
}
