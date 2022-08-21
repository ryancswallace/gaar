package cmd

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

func parseAssumeResponse(rawAssumeResponse string) (AssumeResponse, error) {
	var assumeResponse AssumeResponse
	err := json.Unmarshal([]byte(rawAssumeResponse), &assumeResponse)
	return assumeResponse, err
}

func validateAssumeResponse(assumeResponse AssumeResponse) error {
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

func printEnvVars(assumeResponse AssumeResponse) error {
	fmt.Println(EnvVarSetKeyword + " AWS_ACCESS_KEY_ID=" + assumeResponse.Credentials.AccessKeyId)
	fmt.Println(EnvVarSetKeyword + " AWS_SECRET_ACCESS_KEY=" + assumeResponse.Credentials.SecretAccessKey)
	fmt.Println(EnvVarSetKeyword + " AWS_SESSION_TOKEN=" + assumeResponse.Credentials.SessionToken)

	return nil
}

func Run(dispConf DispConf) (int, error) {
	rawAssumeResponse, err := readAssumeResponse()
	if err != nil {
		msg := sanitizeMessage("Error reading standard input: " + err.Error())
		fmt.Fprintln(os.Stderr, msg)
		return ExitRead, err
	}

	assumeResponse, err := parseAssumeResponse(rawAssumeResponse)
	if err != nil {
		msg := sanitizeMessage("Error parsing response: " + err.Error())
		fmt.Fprintln(os.Stderr, msg)
		return ExitParse, err
	}

	err = validateAssumeResponse(assumeResponse)
	if err != nil {
		msg := sanitizeMessage("Error validating response: " + err.Error())
		fmt.Fprintln(os.Stderr, msg)
		return ExitValidate, err
	}

	err = printEnvVars(assumeResponse)
	if err != nil {
		msg := sanitizeMessage("Error printing environment variables: " + err.Error())
		fmt.Fprintln(os.Stderr, msg)
		return ExitPrint, err
	}

	err = dispResponse(assumeResponse, dispConf)
	if err != nil {
		msg := sanitizeMessage("Error displaying response: " + err.Error())
		fmt.Fprintln(os.Stderr, msg)
		return ExitDisplay, err
	}

	return ExitSuccess, nil
}
