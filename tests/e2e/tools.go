/*
This file is part of Cloud Native PostgreSQL.

Copyright (C) 2019-2020 2ndQuadrant Italia SRL. Exclusively licensed to 2ndQuadrant Limited.
*/

package e2e

import (
	"bytes"
	"log"
	"os/exec"

	"github.com/google/shlex"
)

//nolint:unparam,gosec
func run(command string) (stdout string, stderr string, err error) {
	tokens, err := shlex.Split(command)
	if err != nil {
		log.Printf("Error parsing command `%v`: %v\n", command, err)
		return "", "", err
	}

	var outBuffer, errBuffer bytes.Buffer
	cmd := exec.Command(tokens[0], tokens[1:]...)
	cmd.Stdout, cmd.Stderr = &outBuffer, &errBuffer
	if err = cmd.Run(); err != nil {
		log.Printf("Error executing command `%v`\nerr: %v\nstdout: %v\nstderr: %v\n",
			command, err, outBuffer.String(), errBuffer.String())
		return "", "", err
	}

	return outBuffer.String(), errBuffer.String(), nil
}
