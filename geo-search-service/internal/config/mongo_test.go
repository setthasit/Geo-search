package config_test

import (
	"os"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"geo-search/internal/config"
	"geo-search/internal/config/mock_config"
)

var _ = Describe("Mongo Config", func() {
	var mockCtrl *gomock.Controller

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
	})

	AfterEach(func() {
		err := os.Unsetenv("APP_MONGO_URI")
		Expect(err).NotTo(HaveOccurred())

		err = os.Unsetenv("MONGO_USERNAME")
		Expect(err).NotTo(HaveOccurred())

		err = os.Unsetenv("MONGO_PASSWORD")
		Expect(err).NotTo(HaveOccurred())

		err = os.Unsetenv("MONGO_DB")
		Expect(err).NotTo(HaveOccurred())
	})

	Context("Valid ENV", func() {
		const (
			uri      = "test:27017"
			username = "attest"
			password = "attest12345678"
			database = "test_geo_search"
		)

		BeforeEach(func() {
			err := os.Setenv("APP_MONGO_URI", uri)
			Expect(err).NotTo(HaveOccurred())
			err = os.Setenv("MONGO_USERNAME", username)
			Expect(err).NotTo(HaveOccurred())
			err = os.Setenv("MONGO_PASSWORD", password)
			Expect(err).NotTo(HaveOccurred())
			err = os.Setenv("MONGO_DB", database)
			Expect(err).NotTo(HaveOccurred())
		})

		It("Should have valid variable with valid ENV values", func() {
			loader := mock_config.NewMockConfigLoader(mockCtrl)
			loader.EXPECT().Load()

			cfg := config.WireMongoConfig(loader)

			Expect(cfg.URI).To(Equal(uri))
			Expect(cfg.Username).To(Equal(username))
			Expect(cfg.Password).To(Equal(password))
			Expect(cfg.Database).To(Equal(database))
		})
	})
})
