package k8s_test

import (
	"fmt"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	. "code.cloudfoundry.org/eirini/k8s"
	"code.cloudfoundry.org/eirini/metrics"
	"code.cloudfoundry.org/eirini/route/routefakes"
)

var _ = FDescribe("Metrics", func() {

	var (
		collector      *MetricsCollector
		work           chan []metrics.Message
		fakeServer     *ghttp.Server
		scheduler      *routefakes.FakeTaskScheduler
		respondHandler http.HandlerFunc
	)

	BeforeEach(func() {
		respondHandler = ghttp.RespondWith(http.StatusOK, "")
	})

	JustBeforeEach(func() {
		fakeServer = ghttp.NewServer()
		fakeServer.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/you/name/it"),
				respondHandler,
			),
		)

		scheduler = new(routefakes.FakeTaskScheduler)
		work = make(chan []metrics.Message, 1)
		collector = NewMetricsCollector(work, scheduler, fmt.Sprintf("%s%s", fakeServer.URL(), "/you/name/it"))
	})

	Context("When collecting metrics", func() {

		var err error

		BeforeEach(func() {
			respondHandler = ghttp.RespondWith(http.StatusOK, `
{
	"metadata": {"name": "thor-9000", "namespace": "asgard"},
	"items": [
		{
			"metadata": {"name": "thor-9000", "namespace": "asgard", "labels": {
	  "guid": "app-guid"
	}
},
			"containers": [{
				"name": "bran-the-builder-9000",
				"usage": {"cpu": "420000m", "memory": "420Ki"}
			}]
		}
	]
}`)
		})

		JustBeforeEach(func() {
			collector.Start()
			task := scheduler.ScheduleArgsForCall(0)
			err = task()
		})

		It("should not return an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("should use the right source url", func() {
			Expect(fakeServer.ReceivedRequests()).To(HaveLen(1))
		})

		It("should send the received metrics", func() {
			Eventually(work).Should(Receive(Equal([]metrics.Message{
				{
					AppID:       "app-guid",
					IndexID:     "9000",
					CPU:         420,
					Memory:      430080,
					MemoryQuota: 10,
					Disk:        42000000,
					DiskQuota:   10,
				},
			})))
		})

		Context("there are no items", func() {

			BeforeEach(func() {
				respondHandler = ghttp.RespondWith(http.StatusOK, `
{
	"metadata": {"name": "thor-1000", "namespace": "myspace"},
	"items": []
}`)
			})

			It("should not return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should not send anything", func() {
				Consistently(work).ShouldNot(Receive())
			})
		})

		Context("there are no containers", func() {

			BeforeEach(func() {
				respondHandler = ghttp.RespondWith(http.StatusOK, `
{
	"metadata": {"name": "thor-9000", "namespace": "asgard"},
	"items": [
		{
			"metadata": {"name": "thor-9000", "namespace": "asgard", "labels": {
	  "guid": "app-guid"
	}
},
			"containers": []
		}
	]
}`)
			})

			It("should not return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should not send anything", func() {
				Consistently(work).ShouldNot(Receive())
			})
		})

		Context("memory metric does not have a unit", func() {
			BeforeEach(func() {
				respondHandler = ghttp.RespondWith(http.StatusOK, `
{
	"metadata": {"name": "thor-9000", "namespace": "asgard"},
	"items": [
		{
			"metadata": {"name": "thor-9000", "namespace": "asgard", "labels": {
	  "guid": "app-guid"
	}
},
			"containers": [{
				"name": "bran-the-builder-9000",
				"usage": {"cpu": "420000m", "memory": "420"}
			}]
		}
	]
}`)
			})

			It("should return not an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should send the received metrics", func() {
				Eventually(work).Should(Receive(Equal([]metrics.Message{
					{
						AppID:       "app-guid",
						IndexID:     "9000",
						CPU:         420,
						Memory:      430080,
						MemoryQuota: 10,
						Disk:        42000000,
						DiskQuota:   10,
					},
				})))
			})
		})

		Context("pod name doesn't have an index (eg staging tasks)", func() {
			BeforeEach(func() {
				respondHandler = ghttp.RespondWith(http.StatusOK, `
{
	"metadata": {"name": "thor-thunder0", "namespace": "asgard"},
	"items": [
		{
			"metadata": {"name": "thor-thunder0", "namespace": "asgard", "labels": {
	  "guid": "app-guid"
	}
},
			"containers": [{
				"name": "bran-the-builder",
				"usage": {"cpu": "420000m", "memory": "420M"}
			}]
		}
	]
}`)
			})

			It("should not return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should send the received metrics", func() {
				Eventually(work).Should(Receive(Equal([]metrics.Message{
					{
						AppID:       "app-guid",
						IndexID:     "0",
						CPU:         420,
						Memory:      430080,
						MemoryQuota: 10,
						Disk:        42000000,
						DiskQuota:   10,
					},
				})))
			})
		})

		Context("metrics source responds with an error", func() {

			BeforeEach(func() {
				respondHandler = ghttp.RespondWith(http.StatusBadRequest, "")
			})

			It("should return an error", func() {
				Expect(err).To(MatchError(ContainSubstring("metrics source responded with code: 400")))
			})
		})
	})
})
