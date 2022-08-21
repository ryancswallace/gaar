package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

func dispResponse(assumeResponse AssumeResponse, dispConf DispConf) error {
	if dispConf.DispUser {
		repr, err := json.MarshalIndent(assumeResponse.AssumedRoleUser, "", "\t")
		if err != nil {
			return err
		}
		msg := sanitizeMessage(string(repr))
		fmt.Fprintln(os.Stderr, msg)

	}
	if dispConf.DispCredentials {
		repr, err := json.MarshalIndent(assumeResponse.Credentials, "", "\t")
		if err != nil {
			return err
		}
		msg := sanitizeMessage(string(repr))
		fmt.Fprintln(os.Stderr, msg)
	}
	if dispConf.DispPackedPolicySize {
		msg := sanitizeMessage("PackedPolicySize: " + assumeResponse.PackedPolicySize)
		fmt.Fprintln(os.Stderr, msg)
	}
	if dispConf.DispSourceIdentity {
		msg := sanitizeMessage("SourceIdentity: " + assumeResponse.SourceIdentity)
		fmt.Fprintln(os.Stderr, msg)
	}

	return nil
}
