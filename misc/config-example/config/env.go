package config

const ServiceEnvVarName = "FOO_ENV"

const (
	EnvProd             = Env("prod")
	EnvLocal            = Env("local")
	EnvRC               = Env("rc")
	EnvFeature1         = Env("feature1")
	EnvFeature2         = Env("feature2")
	EnvFeature3         = Env("feature3")
	EnvE2ETests         = Env("e2e")
	EnvUnitTests        = Env("utests")
	EnvIntegrationTests = Env("itests")
)

type Env string
