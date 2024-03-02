package render

import (
	b64 "encoding/base64"
	"github.com/edimarlnx/secure-templates/pkg/connectors"
	"github.com/edimarlnx/secure-templates/pkg/helpers"
	"log"
)

func RegisterSecret(connector connectors.Connector) func(secretName, keyName string) string {
	return func(secretName, keyName string) string {
		return connector.Secret(secretName, keyName)
	}
}

func Base64Encode(str string) string {
	return b64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(base64Str string) string {
	bytes, err := b64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		log.Fatalf("Error on decode base64 string %s. %v", base64Str, err)
	}
	return string(bytes)
}

func EnvVar(envName string) string {
	return helpers.GetEnv(envName, envName)
}
