package gaar

import (
	"flag"
	"fmt"
	"os"
)

type DispConf struct {
	DispUser             bool
	DispCredentials      bool
	DispPackedPolicySize bool
	DispSourceIdentity   bool
}

func SetUsage() {
	flag.Usage = func() {
		usage := `Format AWS IAM role credentials.

gaar reads a JSON response from "aws sts assume-role" from stdin and prints environment variable
settings for the credentials received to stdout, formatted for the current shell.

Usage:
    gaar [-disp] [-disp-user] [-disp-credentials] [-disp-response]

		`

		example := `
To retrieve and export environment variables in one line, use
$ $(aws sts assume-role --role-arn foo --role-session-name bar | gaar)
`
		fmt.Fprintln(os.Stderr, usage)
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, example)
	}
}

func ParseCmd() DispConf {
	dispAll := flag.Bool("disp", false, "Display assume-role response to stderr (default false)")
	dispUser := flag.Bool("disp-user", false, "Display contents of AssumedRoleUser variable from assume-role response to stderr (default false)")
	dispCredentials := flag.Bool("disp-credentials", false, "Display contents of Credentials variable from assume-role response to stderr (default false)")
	dispPackedPolicySize := flag.Bool("disp-size", false, "Display value of PackedPolicySize variable from assume-role response to stderr (default false)")
	dispSourceIdentity := flag.Bool("disp-source-id", false, "Display value of SourceIdentity variable from assume-role response to stderr (default false)")

	flag.Parse()

	dispConf := DispConf{
		*dispUser || *dispAll,
		*dispCredentials || *dispAll,
		*dispPackedPolicySize || *dispAll,
		*dispSourceIdentity || *dispAll,
	}

	return dispConf
}
