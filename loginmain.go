package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Struct to hold parameters
type Parameters struct {
	ReleaseBuild           string
	ReleaseLetter          string
	BuildNumber            string
	PrevReleaseLetter      string
	TwoReleasesPriorLetter string
	Host                   string
	Database               string
	PrefixLetter           string
	NewDisk                string
	OldDisk                string
	Port                   string
	Ctcws                  string
	Tma                    string
	RemoteHost             string
	RemotePort             string
	RemoteUser             string
}

var params = Parameters{
	ReleaseBuild:           "S20",
	ReleaseLetter:          "T",
	BuildNumber:            "17",
	PrevReleaseLetter:      "U",
	TwoReleasesPriorLetter: "T",
	Host:                   "\\SED1",
	Database:               "020.36",
	PrefixLetter:           "S",
	NewDisk:                "$DATA1",
	OldDisk:                "$DATA2",
	Port:                   "13000",
	Ctcws:                  "36.64",
	Tma:                    "21.2",
	RemoteHost:             "10.202.5.114",
	RemotePort:             "22",
	RemoteUser:             "psccqa",
}

const additionalShellScript = `
ssh -T rc.mgr@${tandem} << com
       gtacl
        eman${eman_env} c $build_env ebldrel
        eman${eman_env} c $build_env
        ccs
        scs
        exit
        exit
com
`

func getInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	input := scanner.Text()
	return input
}
func updateParameters(scanner *bufio.Scanner) {
	fmt.Print("Enter Release Build [", params.ReleaseBuild, "]: ")
	params.ReleaseBuild = getInput(scanner)

	fmt.Print("Enter Release Letter [", params.ReleaseLetter, "]: ")
	params.ReleaseLetter = getInput(scanner)

	fmt.Print("Enter Build Number [", params.BuildNumber, "]: ")
	params.BuildNumber = getInput(scanner)

	fmt.Print("Enter Previous Release Letter [", params.PrevReleaseLetter, "]: ")
	params.PrevReleaseLetter = getInput(scanner)

	fmt.Print("Enter 2 Releases Prior Letter [", params.TwoReleasesPriorLetter, "]: ")
	params.TwoReleasesPriorLetter = getInput(scanner)

	fmt.Print("Enter Host [", params.Host, "]: ")
	params.Host = getInput(scanner)

	fmt.Print("Enter Database [", params.Database, "]: ")
	params.Database = getInput(scanner)

	fmt.Print("Enter Prefix Letter [", params.PrefixLetter, "]: ")
	params.PrefixLetter = getInput(scanner)

	fmt.Print("Enter New Disk [", params.NewDisk, "]: ")
	params.NewDisk = getInput(scanner)

	fmt.Print("Enter Old Disk [", params.OldDisk, "]: ")
	params.OldDisk = getInput(scanner)

	fmt.Print("Enter Port [", params.Port, "]: ")
	params.Port = getInput(scanner)

	fmt.Print("Enter CTCWS [", params.Ctcws, "]: ")
	params.Ctcws = getInput(scanner)

	fmt.Print("Enter TMA [", params.Tma, "]: ")
	params.Tma = getInput(scanner)
}

func runSSHCommand(command string, password string, userHost ...string) {
	cmdArgs := []string{"-p", params.RemotePort}

	if len(userHost) > 0 {
		remoteUserHost := strings.Split(userHost[0], "@")
		params.RemoteUser = remoteUserHost[0]
		params.RemoteHost = remoteUserHost[1]
	}

	cmdArgs = append(cmdArgs, fmt.Sprintf("%s@%s", params.RemoteUser, params.RemoteHost))
	cmdArgs = append(cmdArgs, command) // Move command to the end

	cmd := exec.Command("sshpass", append([]string{"-p", password, "ssh"}, cmdArgs...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
