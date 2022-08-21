package cmd

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
		fmt.Fprintln(os.Stderr, "Format AWS IAM role credentials.")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "gaar reads a JSON response from `aws sts assume-role` from stdin and prints environment variable")
		fmt.Fprintln(os.Stderr, "settings for the credentials received to stdout, formatted for the current shell.")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Usage:")
		fmt.Fprintln(os.Stderr, "\tgaar [-disp] [-disp-user] [-disp-credentials] [-disp-response]")
		fmt.Fprintln(os.Stderr)
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "To retrieve and export environment variables in one line, use")
		fmt.Fprintln(os.Stderr, "$ $(aws sts assume-role --role-arn foo --role-session-name bar | gaar)")
		fmt.Fprintln(os.Stderr)
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
