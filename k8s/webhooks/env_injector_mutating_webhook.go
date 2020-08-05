package webhooks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"code.cloudfoundry.org/lager"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

const (
	// TODO find the inex of the opi container form the raw pod input
	podPatchTemplate string = `[
      {
		   "op":"add",
		   "path":"/spec/containers[0]/env",
		   "name":"CF_INSTANCE_INDEX",
		   "value": %d
	  }
	]`
)

type EnvironmentInjectorMutatingWebhook struct {
	logger lager.Logger
}

func NewEnvironmentInjectorWebhook(logger lager.Logger) *EnvironmentInjectorMutatingWebhook {
	return &EnvironmentInjectorMutatingWebhook{
		logger: logger,
	}
}

func (h *EnvironmentInjectorMutatingWebhook) Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var body []byte
	if r.Body != nil {
		if data, err := ioutil.ReadAll(r.Body); err == nil {
			body = data
		}
	}

	// verify the content type is accurate
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		h.logger.Error("unexpected-content-type", fmt.Errorf("expected application/json, but was %s", contentType))
		return
	}

	h.logger.Info(fmt.Sprintf("handling request: %s", body))

	// The AdmissionReview that was sent to the webhook
	requestedAdmissionReview := admissionv1.AdmissionReview{}

	// The AdmissionReview that will be returned
	responseAdmissionReview := admissionv1.AdmissionReview{}

	scheme := runtime.NewScheme()
	codecs := serializer.NewCodecFactory(scheme)
	deserializer := codecs.UniversalDeserializer()
	if _, _, err := deserializer.Decode(body, nil, &requestedAdmissionReview); err != nil {
		h.logger.Error("failed-to-decode-body", err, lager.Data{"body": body})
		responseAdmissionReview.Response = toAdmissionResponse(err)
	} else {
		// pass to mutate func
		responseAdmissionReview.Response = h.mutatePod(requestedAdmissionReview)
	}

	// Return the same UID
	responseAdmissionReview.Response.UID = requestedAdmissionReview.Request.UID

	h.logger.Info(fmt.Sprintf("sending response: %v", responseAdmissionReview.Response))

	respBytes, err := json.Marshal(responseAdmissionReview)
	if err != nil {
		h.logger.Error("failed-to-marshal-admission-response", err, lager.Data{"admissionResponse": responseAdmissionReview})
	}
	if _, err := w.Write(respBytes); err != nil {
		h.logger.Error("failed-to-send-admission-response", err, lager.Data{"admissionResponse": responseAdmissionReview})
	}
}

func (h *EnvironmentInjectorMutatingWebhook) mutatePod(ar admissionv1.AdmissionReview) *admissionv1.AdmissionResponse {
	raw := ar.Request.Object.Raw
	pod := corev1.Pod{}
	scheme := runtime.NewScheme()
	codecs := serializer.NewCodecFactory(scheme)
	deserializer := codecs.UniversalDeserializer()
	if _, _, err := deserializer.Decode(raw, nil, &pod); err != nil {
		h.logger.Error("failed-to-decode-pod", err, lager.Data{"rawPodObject": raw})
		return toAdmissionResponse(err)
	}

	response := &admissionv1.AdmissionResponse{
		Allowed: true,
	}
	patch, err := calculatePodContainerPatch(pod.Name)
	if err != nil {
		h.logger.Error("error-patching-pod", err, lager.Data{"podName": pod.Name})
		return response
	}

	typeJSONPatch := admissionv1.PatchTypeJSONPatch
	response.PatchType = &typeJSONPatch
	response.Patch = patch

	return response
}

func calculatePodContainerPatch(podName string) ([]byte, error) {
	instanceIndex, err := parseInstanceIndex(podName)
	if err != nil {
		return nil, err
	}

	return []byte(fmt.Sprintf(podPatchTemplate, instanceIndex)), nil
}

func parseInstanceIndex(podName string) (int, error) {
	nameSegments := strings.Split(podName, "-")
	numSegments := len(nameSegments)
	if numSegments == 0 {
		return -1, fmt.Errorf("invalid pod name: %q. Pod names should contain dashes", podName)
	}
	instanceIndex, err := strconv.Atoi(nameSegments[numSegments-1])
	if err != nil {
		return -1, errors.Wrapf(err, "error parsing instanceIndex for pod %q", podName)
	}
	return instanceIndex, nil
}

// toAdmissionResponse is a helper function to create an AdmissionResponse
// with an embedded error
func toAdmissionResponse(err error) *admissionv1.AdmissionResponse {
	return &admissionv1.AdmissionResponse{
		Result: &metav1.Status{
			Message: err.Error(),
		},
	}
}
