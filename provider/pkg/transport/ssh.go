package transport

import (
	"bytes"

	"github.com/tmc/scp"
	"golang.org/x/crypto/ssh"
)

type SSHTransport struct {
	Addr         string
	clientConfig *ssh.ClientConfig
}

func SSHInit(addr, user, key string) (*SSHTransport, error) {
	signer, err := ssh.ParsePrivateKey([]byte(key))
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	}
	config.HostKeyCallback = ssh.InsecureIgnoreHostKey() //nolint:gosec

	s := &SSHTransport{
		Addr:         addr,
		clientConfig: config,
	}

	return s, nil
}

func (s *SSHTransport) RunCmd(cmd string) ([]byte, error) {
	client, err := ssh.Dial("tcp", s.Addr, s.clientConfig)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var o bytes.Buffer
	session.Stdout = &o
	if err := session.Run(cmd); err != nil {
		return nil, err
	}

	return o.Bytes(), nil
}

func (s *SSHTransport) CopyFile(source, dest string) error {
	client, err := ssh.Dial("tcp", s.Addr, s.clientConfig)
	if err != nil {
		return err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	return scp.CopyPath(source, dest, session)
}
