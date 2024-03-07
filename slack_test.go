package slack_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	log "github.com/sirupsen/logrus"

	"github.com/bayugyug/slack"
)

var _ = Describe("Slack", func() {

	var ts *httptest.Server
	var alert slack.NotificationCreator
	var servicePath = fmt.Sprintf(`/services/%s/%s`, uuid.New().String(), uuid.New().String())

	BeforeEach(func() {
		_ = ts
	})

	AfterEach(func() {
	})

	Context("Push alert", func() {
		It("should return ok", func() {

			//inject
			ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				render.Status(r, http.StatusOK)
				fmt.Fprint(w, `ok`)
			}))
			defer ts.Close()
			slack.ServiceEndPoint = ts.URL

			// init
			alert = slack.NewNotification(servicePath)

			Expect(alert).NotTo(BeNil())

			err := alert.Notify([]*slack.Payload{
				{
					Title:       "Event Push",
					Message:     fake.Sentences(),
					WithDivider: true,
					Icon:        slack.IconTypeSpeechBalloon,
				},
				{
					Title:       "Notify",
					Message:     fake.Sentences(),
					WithDivider: true,
					Icon:        slack.IconTypeHeart,
				},
				{
					Title:       "Star",
					Message:     fake.Sentences(),
					WithDivider: true,
					Icon:        slack.IconTypeStar,
				},
				{
					Title:       "Warning",
					Message:     fake.Sentences(),
					WithDivider: true,
					Icon:        slack.IconTypeWarning,
				},
				{
					Title:       "Critical",
					Message:     fake.Sentences(),
					WithDivider: true,
					Priority:    slack.PriorityCritical,
					Icon:        slack.IconTypeCritical,
					Here:        "Please check ...",
				},
				{
					Title:       "Success",
					Message:     fake.Sentences(),
					WithDivider: true,
					Icon:        slack.IconTypeThumbsUp,
				},
			})

			Expect(err).To(BeNil())

			By("Push alert ok")
		})
	})

	Context("Push alert no divider", func() {
		It("should return ok", func() {
			//inject
			ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				render.Status(r, http.StatusOK)
				fmt.Fprint(w, `ok`)
			}))
			defer ts.Close()
			slack.ServiceEndPoint = ts.URL

			// init
			alert = slack.NewNotification(servicePath)

			Expect(alert).NotTo(BeNil())

			err := alert.Notify([]*slack.Payload{
				{
					Title:   "Star",
					Message: fake.Sentences(),
					Icon:    slack.IconTypeStar,
				},
			}, slack.Block{
				BlockID: uuid.New().String(),
				Type:    slack.SectionBlockType,
				Text: &slack.Text{
					Type: slack.MarkDownBlockType,
					Text: "HAHAHAHA: " + fake.Sentences(),
				},
			})

			Expect(err).To(BeNil())

			By("Push alert no divider ok")
		})
	})

	Context("Push alert invalid response", func() {
		It("should return ok", func() {

			//inject
			ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				render.Status(r, http.StatusInternalServerError)
				render.JSON(w, r,
					`Oops failed here`,
				)
			}))
			defer ts.Close()
			slack.ServiceEndPoint = ts.URL

			// init
			alert = slack.NewNotification(servicePath)

			Expect(alert).NotTo(BeNil())

			err := alert.Notify([]*slack.Payload{
				{
					Title:       "Event Push",
					Message:     fake.Sentences(),
					WithDivider: true,
					Icon:        slack.IconTypeSpeechBalloon,
				},
			})

			Expect(err).NotTo(BeNil())

			By("Push alert invalid response ok")
		})
	})

	Context("Push alert invalid reply", func() {
		It("should return ok", func() {

			//inject
			ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				render.Status(r, http.StatusOK)
				fmt.Fprint(w, `not-ok`)
			}))
			defer ts.Close()
			slack.ServiceEndPoint = ts.URL

			// init
			alert = slack.NewNotification(servicePath)

			Expect(alert).NotTo(BeNil())

			err := alert.Notify([]*slack.Payload{
				{
					Title:       "Event Push",
					Message:     fake.Sentences(),
					WithDivider: true,
					Icon:        slack.IconTypeSpeechBalloon,
				},
			})

			Expect(err).NotTo(BeNil())

			By("Push alert invalid reply ok")
		})
	})

	Context("Push alert with empty payload", func() {
		It("should return ok", func() {
			// init
			alert = slack.NewNotification(servicePath)

			Expect(alert).NotTo(BeNil())

			err := alert.Notify(nil)

			Expect(err).NotTo(BeNil())

			By("Push alert with empty payload ok")
		})
	})

	Context("Generic", func() {
		It("should return ok", func() {
			// init
			alert = slack.NewNotification(servicePath)

			Expect(alert).NotTo(BeNil())

			log.Println("critical", slack.PriorityCritical.String())
			By("Generic ok")
		})
	})
})
