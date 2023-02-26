package config_test

import (
	"geo-search/internal/config"
	"geo-search/internal/config/mock_config"
	"os"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Elasticsearch Config", func() {
	var mockCtrl *gomock.Controller

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
	})

	AfterEach(func() {
		err := os.Unsetenv("APP_ELASTICSEARCH_URI")
		Expect(err).NotTo(HaveOccurred())

		err = os.Unsetenv("APP_ELASTICSEARCH_USERNAME")
		Expect(err).NotTo(HaveOccurred())

		err = os.Unsetenv("APP_ELASTICSEARCH_PASSWORD")
		Expect(err).NotTo(HaveOccurred())

		err = os.Unsetenv("APP_ELASTICSEARCH_CERT_PATH")
		Expect(err).NotTo(HaveOccurred())
	})

	Context("Valid ENV", func() {
		const (
			uri      = "https://test:9200"
			username = "attest"
			password = "attest12345678"
			cert     = "./certs/elasticsearch.crt"
		)

		BeforeEach(func() {
			err := os.Setenv("APP_ELASTICSEARCH_URI", uri)
			Expect(err).NotTo(HaveOccurred())
			err = os.Setenv("APP_ELASTICSEARCH_USERNAME", username)
			Expect(err).NotTo(HaveOccurred())
			err = os.Setenv("APP_ELASTICSEARCH_PASSWORD", password)
			Expect(err).NotTo(HaveOccurred())
			err = os.Setenv("APP_ELASTICSEARCH_CERT_PATH", cert)
			Expect(err).NotTo(HaveOccurred())
		})

		It("Should have valid variable with valid ENV values", func() {
			loader := mock_config.NewMockConfigLoader(mockCtrl)
			loader.EXPECT().Load()

			cfg := config.WireElasticsearchConfig(loader)

			Expect(cfg.URI).To(Equal(uri))
			Expect(cfg.Username).To(Equal(username))
			Expect(cfg.Password).To(Equal(password))
			Expect(cfg.CertPath).To(Equal(cert))
		})
	})
})
