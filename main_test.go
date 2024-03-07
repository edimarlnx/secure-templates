package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func Test_initApp(t *testing.T) {
	workdir, err := os.Getwd()
	if err != nil {
		workdir = os.TempDir()
	}
	configFile := "test/local-file-cfg-test.json"
	tests := []struct {
		name            string
		args            []string
		requiredStrings []string
		envs            map[string]string
	}{
		{
			name: "init-config",
			args: []string{
				"secure-templates",
				"init-config",
				"-o",
				configFile,
				"-secret-file",
				fmt.Sprintf("%s/test/local-file-secret-test.json", workdir),
				"-private-key-passphrase",
				"test-pwd",
			},
			requiredStrings: []string{},
			envs: map[string]string{
				"SEC_TPL_CONFIG": configFile,
			},
		},
		{
			name: "put app user",
			args: []string{
				"secure-templates",
				"local-secret",
				"put",
				"core",
				"app_user",
				"dev_user",
			},
			requiredStrings: []string{
				"saved on secret",
			},
			envs: map[string]string{
				"SEC_TPL_CONFIG": configFile,
			},
		},
		{
			name: "put app passwd",
			args: []string{
				"secure-templates",
				"local-secret",
				"put",
				"core",
				"app_passwd",
				"2dabe3d7c66fb75f751202fdab19266b",
			},
			requiredStrings: []string{
				"saved on secret",
			},
			envs: map[string]string{
				"SEC_TPL_CONFIG": configFile,
			},
		},
		{
			name: "put client app user",
			args: []string{
				"secure-templates",
				"local-secret",
				"put",
				"client",
				"app_user",
				"dev_user",
			},
			requiredStrings: []string{
				"saved on secret",
			},
			envs: map[string]string{
				"SEC_TPL_CONFIG": configFile,
			},
		},
		{
			name: "put client app passwd",
			args: []string{
				"secure-templates",
				"local-secret",
				"put",
				"client",
				"app_passwd",
				"2dabe3d7c66fb75f751202fdab19266b",
			},
			requiredStrings: []string{
				"saved on secret",
			},
			envs: map[string]string{
				"SEC_TPL_CONFIG": configFile,
			},
		},
		{
			name: "Env file",
			args: []string{
				"secure-templates",
				"test/samples/.env",
			},
			requiredStrings: []string{
				"APP_USER=dev_user",
				"APP_PASSWORD=2dabe3d7c66fb75f751202fdab19266b",
				"CLIENT_APP_USER=dev_user",
				"CLIENT_APP_PASSWORD=2dabe3d7c66fb75f751202fdab19266b",
			},
			envs: map[string]string{
				"SEC_TPL_CONFIG": configFile,
			},
		},
		{
			name: "k8s secret yaml",
			args: []string{
				"secure-templates",
				"test/samples/k8s-secret.yaml",
			},
			requiredStrings: []string{
				"name: st-secret",
				"namespace: dev-ns",
				"APP_USER: ZGV2X3VzZXI=",
				"APP_PASSWORD: MmRhYmUzZDdjNjZmYjc1Zjc1MTIwMmZkYWIxOTI2NmI=",
				"CLIENT_APP_USER: \"dev_user\"",
				"CLIENT_APP_PASSWORD: \"2dabe3d7c66fb75f751202fdab19266b\"",
			},
			envs: map[string]string{
				"SEC_TPL_CONFIG":   configFile,
				"SECRET_NAME":      "st-secret",
				"SECRET_NAMESPACE": "dev-ns",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key, value := range tt.envs {
				t.Setenv(key, value)
			}
			buf := bytes.Buffer{}
			initApp(tt.args, &buf)
			str := buf.String()
			for _, requiredString := range tt.requiredStrings {
				if !strings.Contains(str, requiredString) {
					t.Fatalf("Required '%s' string not found.", requiredString)
				}
			}
		})
	}
}
