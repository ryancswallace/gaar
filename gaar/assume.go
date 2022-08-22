package gaar

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const ExitSuccess = 0
const ExitRead = 10
const ExitParse = 11
const ExitValidate = 12
const ExitPrint = 13
const ExitDisplay = 14

type AssumedRoleUser struct {
	AssumedRoleId string
	Arn           string
}

type Credentials struct {
	AccessKeyId     string
	SecretAccessKey string
	SessionToken    string
	Expiration      string
}

type AssumeResponse struct {
	Credentials      Credentials
	AssumedRoleUser  AssumedRoleUser
	PackedPolicySize string
	SourceIdentity   string
}

func readAssumeResponse() (string, error) {
	var rawAssumeResponse strings.Builder

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rawAssumeResponse.WriteString(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return rawAssumeResponse.String(), nil
}

func makeAssumeResponse(rawAssumeResponse string) (AssumeResponse, error) {
	var assumeResponse AssumeResponse
	err := json.Unmarshal([]byte(rawAssumeResponse), &assumeResponse)
	return assumeResponse, err
}

func (assumeResponse AssumeResponse) validateAssumeResponse() error {
	if assumeResponse.Credentials.AccessKeyId == "" {
		return errors.New("found null value for AccessKeyId")
	}
	if assumeResponse.Credentials.SecretAccessKey == "" {
		return errors.New("found null value for SecretAccessKey")
	}
	if assumeResponse.Credentials.SessionToken == "" {
		return errors.New("found null value for SessionToken")
	}

	return nil
}

func (assumeResponse AssumeResponse) printEnvVars() error {
	fmt.Println(EnvVarSetKeyword + " AWS_ACCESS_KEY_ID=" + assumeResponse.Credentials.AccessKeyId)
	fmt.Println(EnvVarSetKeyword + " AWS_SECRET_ACCESS_KEY=" + assumeResponse.Credentials.SecretAccessKey)
	fmt.Println(EnvVarSetKeyword + " AWS_SESSION_TOKEN=" + assumeResponse.Credentials.SessionToken)

	return nil
}

func (assumeResponse AssumeResponse) dispResponse(dispConf DispConf) error {
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

func Run(dispConf DispConf) (int, error) {
	rawAssumeResponse, err := readAssumeResponse()
	if err != nil {
		msg := sanitizeMessage("gaar: Error reading standard input: " + err.Error())
		fmt.Fprintln(os.Stderr, msg)
		return ExitRead, err
	}

	assumeResponse, err := makeAssumeResponse(rawAssumeResponse)
	if err != nil {
		msg := sanitizeMessage("gaar: Error parsing response: " + err.Error())
		fmt.Fprintln(os.Stderr, msg)
		return ExitParse, err
	}

	err = assumeResponse.validateAssumeResponse()
	if err != nil {
		msg := sanitizeMessage("gaar: Error validating response: " + err.Error())
		fmt.Fprintln(os.Stderr, msg)
		return ExitValidate, err
	}

	err = assumeResponse.printEnvVars()
	if err != nil {
		msg := sanitizeMessage("gaar: Error printing environment variables: " + err.Error())
		fmt.Fprintln(os.Stderr, msg)
		return ExitPrint, err
	}

	err = assumeResponse.dispResponse(dispConf)
	if err != nil {
		msg := sanitizeMessage("gaar: Error displaying response: " + err.Error())
		fmt.Fprintln(os.Stderr, msg)
		return ExitDisplay, err
	}

	return ExitSuccess, nil
}
