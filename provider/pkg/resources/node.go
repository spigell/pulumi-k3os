package resources

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	"github.com/mitchellh/mapstructure"
	"github.com/spigell/pulumi-k3os/provider/pkg/transport"
	"gopkg.in/yaml.v3"
)

const (
	k3osConfigPath = "/var/lib/rancher/k3os/config.yaml"
)

type NodeConfig struct {
	SSHAuthorizedKeys []string          `json:"ssh_authorized_keys" yaml:"ssh_authorized_keys,omitempty"`
	WriteFiles        []CloudInitFile   `json:"write_files,omitempty" yaml:"write_files,omitempty"`
	Hostname          string            `json:"hostname,omitempty" yaml:"hostname,omitempty"`
	Runcmd            []string          `json:"run_cmd" yaml:"run_cmd,omitempty"`
	Bootcmd           []string          `json:"boot_cmd" yaml:"boot_cmd,omitempty"`
	Initcmd           []string          `json:"init_cmd" yaml:"init_cmd,omitempty"`
	DataSources       []string          `json:"data_sources,omitempty" yaml:"data_sources,omitempty"`
	Modules           []string          `json:"modules,omitempty" yaml:"modules,omitempty"`
	Sysctls           map[string]string `json:"sysctls,omitempty" yaml:"sysctls,omitempty"`
	K3os              K3OSConfig        `json:"k3os" yaml:"k3os"`
}

type K3OSConfig struct {
	ServerURL   string            `json:"server_url" yaml:"server_url,omitempty"`
	Password    string            `json:"password,omitempty" yaml:"password,omitempty"`
	Token       string            `json:"token,omitempty" yaml:"token,omitempty"`
	Labels      map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	K3sArgs     []string          `json:"k3s_args" yaml:"k3s_args"`
	Environment map[string]string `json:"environment,omitempty" yaml:"environment,omitempty"`
	Taints      []string          `json:"taints,omitempty" yaml:"taints,omitempty"`
}

type CloudInitFile struct {
	Encoding           string `json:"encoding" yaml:"encoding,omitempty"`
	Content            string `json:"content"`
	Owner              string `json:"owner" yaml:"owner,omitempty"`
	Path               string `json:"path"`
	RawFilePermissions string `json:"permissions" yaml:"permission,omitempty"`
}

func SetNodeConfig(r resource.PropertyMap) (map[string]interface{}, error) {
	addr, user, key := parseConnArgs(r["connection"])
	s, err := transport.SSHInit(addr, user, key)
	if err != nil {
		return nil, err
	}

	path, err := PutNodeConfigToDir("/tmp", r["nodeConfiguration"])
	if err != nil {
		return nil, err
	}

	tempTarget := fmt.Sprintf("/home/%s/%s", user, filepath.Base(path))
	err = s.CopyFile(path, tempTarget)
	if err != nil {
		return nil, err
	}

	_, err = s.RunCmd(fmt.Sprintf("sudo mv -v %s %s", tempTarget, k3osConfigPath))
	if err != nil {
		return nil, err
	}

	_, err = s.RunCmd("sudo reboot -d 1")
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	return result, nil
}

func DeleteNodeConfig(c resource.PropertyValue) error {
	addr, user, key := parseConnArgs(c)
	s, err := transport.SSHInit(addr, user, key)
	if err != nil {
		return err
	}

	_, err = s.RunCmd(fmt.Sprintf("sudo rm -rfv %s", k3osConfigPath))
	if err != nil {
		return err
	}

	_, err = s.RunCmd("sudo reboot -d 1")
	if err != nil {
		return err
	}

	return nil
}

func CheckNodeConnection(c resource.PropertyValue) (map[string]interface{}, error) {
	addr, user, key := parseConnArgs(c)
	s, err := transport.SSHInit(addr, user, key)
	if err != nil {
		return nil, err
	}

	_, err = s.RunCmd("sudo k3os config --dump")
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	result["success"] = true

	return result, nil
}

func parseConnArgs(c resource.PropertyValue) (string, string, string) {
	addr := c.ObjectValue()["addr"].StringValue()
	user := c.ObjectValue()["user"].StringValue()
	key := c.ObjectValue()["key"].StringValue()

	return addr, user, key
}

func PutNodeConfigToDir(dir string, r resource.PropertyValue) (string, error) {
	cfg := &NodeConfig{}

	mapstructure.Decode(r.Mappable(), &cfg)

	encoded, err := yaml.Marshal(&cfg)
	if err != nil {
		return "", err
	}

	path := filepath.Join(dir, "k3os-config.yaml")

	err = ioutil.WriteFile(path, encoded, 0600)
	if err != nil {
		return "", err
	}

	return path, nil
}
