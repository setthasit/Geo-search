package config_test

import (
	"geo-search/internal/config"
	"geo-search/internal/config/mock_config"
	"os"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("App config", func() {
	var mockCtrl *gomock.Controller
	var originalEnvironmentType string

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		originalEnvironmentType = os.Getenv("APP_ENVIRONMENT")
	})

	AfterEach(func() {
		err := os.Setenv("APP_ENVIRONMENT", originalEnvironmentType)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("given production as EnvironmentType", func() {
		BeforeEach(func() {
			os.Setenv("APP_ENVIRONMENT", config.Production.String())
		})

		It("should valid", func() {
			loader := mock_config.NewMockConfigLoader(mockCtrl)
			loader.EXPECT().Load()

			cfg := config.WireAppConfig(loader)

			Expect(cfg.Environment).Should(Equal(config.Production))
			Expect(cfg.Environment.Valid()).Should(BeTrue())
		})
	})

	Context("given development as EnvironmentType", func() {
		BeforeEach(func() {
			os.Setenv("APP_ENVIRONMENT", config.Development.String())
		})

		It("should valid", func() {
			loader := mock_config.NewMockConfigLoader(mockCtrl)
			loader.EXPECT().Load()

			cfg := config.WireAppConfig(loader)

			Expect(cfg.Environment).Should(Equal(config.Development))
			Expect(cfg.Environment.Valid()).Should(BeTrue())
		})
	})

	Context("given invalid EnvironmentType", func() {
		BeforeEach(func() {
			os.Setenv("APP_ENVIRONMENT", "invalid")
		})

		It("should invalid, Then return production", func() {
			loader := mock_config.NewMockConfigLoader(mockCtrl)
			loader.EXPECT().Load()

			cfg := config.WireAppConfig(loader)

			Expect(cfg.Environment).Should(Equal(config.Production))
			Expect(cfg.Environment.Valid()).Should(BeTrue())
		})
	})
})
