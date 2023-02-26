package config_test

import (
	"geo-search/internal/config"
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config Loader", func() {
	Context("correct file path", func() {
		It("should success with file found", func() {
			loader := config.WireConfigLoader(&config.CommandConfig{
				EnvFile: "../../.env.example",
			})

			err := loader.Load()

			log.SetFlags(log.Llongfile)

			Expect(err).To(BeNil())
		})

		It("should return nil when loaded second time", func() {
			loader := config.WireConfigLoader(&config.CommandConfig{
				EnvFile: "../../.env.example",
			})

			err := loader.Load()
			Expect(err).To(BeNil())

			err = loader.Load()
			Expect(err).To(BeNil())
		})
	})

	Context("Incorrect file path", func() {
		It("should return error with no file found", func() {
			loader := config.WireConfigLoader(&config.CommandConfig{
				EnvFile: "./not/exist/.env",
			})

			err := loader.Load()

			Expect(err).ToNot(BeNil())
		})

		It("should return nil when loaded second time", func() {
			loader := config.WireConfigLoader(&config.CommandConfig{
				EnvFile: "./not/exist/.env",
			})

			err := loader.Load()
			Expect(err).ToNot(BeNil())

			err = nil
			err = loader.Load()
			Expect(err).To(BeNil())
		})
	})
})
