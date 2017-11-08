package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Read() *Settings {
	env, err := readEnv()
	fmt.Printf("Environment: %s\n", env)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	if env == EnvUnitTests {
		return nil
	}
	var result Settings
	filePath := "./.config/" + string(env) + ".json"
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err := json.Unmarshal(fileBytes, &result); err != nil {
		log.Fatalf("Unmarshalling error: %s", err)
		os.Exit(1)
	}

	return &result
}

func readEnv() (Env, error) {
	serviceEnv := os.Getenv(ServiceEnvVarName)
	if serviceEnv == "" {
		return EnvLocal, nil
	}
	switch Env(serviceEnv) {
	case Env(""):
		return EnvLocal, nil
	case EnvLocal:
		return EnvLocal, nil
	case EnvProd:
		return EnvProd, nil
	case EnvRC:
		return EnvRC, nil
	case EnvFeature1:
		return EnvFeature1, nil
	case EnvFeature2:
		return EnvFeature2, nil
	case EnvFeature3:
		return EnvFeature3, nil
	case EnvE2ETests:
		return EnvE2ETests, nil
	case EnvUnitTests:
		return EnvUnitTests, nil
	case EnvIntegrationTests:
		return EnvIntegrationTests, nil
	default:
		return "", fmt.Errorf("Unknown environment '%s'", os.Args[1])
	}
}
