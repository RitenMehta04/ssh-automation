// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"strings"
// )

// // Struct to hold parameters
// type Parameters struct {
// 	ReleaseBuild              string
// 	ReleaseLetter             string
// 	BuildNumber               string
// 	PrevReleaseLetter         string
// 	TwoReleasesPriorLetter    string
// 	Host                      string
// 	Database                  string
// 	PrefixLetter              string
// 	NewDisk                   string
// 	OldDisk                   string
// 	Port                      string
// 	Ctcws                     string
// 	Tma                       string
// }

// var params = Parameters{
// 	ReleaseBuild:              "S20",
// 	ReleaseLetter:             "T",
// 	BuildNumber:               "17",
// 	PrevReleaseLetter:         "U",
// 	TwoReleasesPriorLetter:    "T",
// 	Host:                      "\\SED1",
// 	Database:                  "020.36",
// 	PrefixLetter:              "S",
// 	NewDisk:                   "$DATA1",
// 	OldDisk:                   "$DATA2",
// 	Port:                      "13000",
// 	Ctcws:                     "36.64",
// 	Tma:                       "21.2",
// }

// func displayParameters() {
// 	fmt.Println("Release Build:", params.ReleaseBuild)
// 	fmt.Println("Release Letter:", params.ReleaseLetter)
// 	fmt.Println("Build Number:", params.BuildNumber)
// 	fmt.Println("Previous Release Letter:", params.PrevReleaseLetter)
// 	fmt.Println("2 Releases Prior Letter:", params.TwoReleasesPriorLetter)
// 	fmt.Println("Host:", params.Host)
// 	fmt.Println("Database:", params.Database)
// 	fmt.Println("Prefix Letter:", params.PrefixLetter)
// 	fmt.Println("New Disk:", params.NewDisk)
// 	fmt.Println("Old Disk:", params.OldDisk)
// 	fmt.Println("Port:", params.Port)
// 	fmt.Println("CTCWS:", params.Ctcws)
// 	fmt.Println("TMA:", params.Tma)
// }

// func updateParameters() {
// 	scanner := bufio.NewScanner(os.Stdin)

// 	fmt.Print("Enter Release Build [", params.ReleaseBuild, "]: ")
// 	input := getInput(scanner)
// 	params.ReleaseBuild = input

// 	fmt.Print("Enter Release Letter [", params.ReleaseLetter, "]: ")
// 	input = getInput(scanner)
// 	params.ReleaseLetter = input

// 	fmt.Print("Enter Build Number [", params.BuildNumber, "]: ")
// 	input = getInput(scanner)
// 	params.BuildNumber = input

// 	fmt.Print("Enter Previous Release Letter [", params.PrevReleaseLetter, "]: ")
// 	input = getInput(scanner)
// 	params.PrevReleaseLetter = input

// 	fmt.Print("Enter 2 Releases Prior Letter [", params.TwoReleasesPriorLetter, "]: ")
// 	input = getInput(scanner)
// 	params.TwoReleasesPriorLetter = input

// 	fmt.Print("Enter Host [", params.Host, "]: ")
// 	input = getInput(scanner)
// 	params.Host = input

// 	fmt.Print("Enter Database [", params.Database, "]: ")
// 	input = getInput(scanner)
// 	params.Database = input

// 	fmt.Print("Enter Prefix Letter [", params.PrefixLetter, "]: ")
// 	input = getInput(scanner)
// 	params.PrefixLetter = input

// 	fmt.Print("Enter New Disk [", params.NewDisk, "]: ")
// 	input = getInput(scanner)
// 	params.NewDisk = input

// 	fmt.Print("Enter Old Disk [", params.OldDisk, "]: ")
// 	input = getInput(scanner)
// 	params.OldDisk = input

// 	fmt.Print("Enter Port [", params.Port, "]: ")
// 	input = getInput(scanner)
// 	params.Port = input

// 	fmt.Print("Enter CTCWS [", params.Ctcws, "]: ")
// 	input = getInput(scanner)
// 	params.Ctcws = input

// 	fmt.Print("Enter TMA [", params.Tma, "]: ")
// 	input = getInput(scanner)
// 	params.Tma = input
// }

// func prestagingActivities() {
// 	// Implement pre-staging activities using SSH commands
// 	// SSH commands using dynamically set parameters
//     sh "ssh $remoteUser@$remoteHost ls -ld \"/L/ctc_data/ctc_data.${releaseBuild}.${database}\""
//     sh "ssh $remoteUser@$remoteHost ls -ld \"/L/ctcws/ctcws.${ctcws}\""
//     sh "ssh $remoteUser@$remoteHost ls -ld \"/L/gds/tma.${tma}\""
// }

// func additionalSSHCommands() {
// 	// Implement additional SSH commands
// 	// SSH commands for PRODSE.USER1
//     sh "ssh PRODSE.USER1@SED1 \"EMANT TP1; DE\""

//     // SSH commands for TRAIN.TRNING
//     sh "ssh TRAIN.TRNING@SED1 \"EMANT TT2; DE\""

//     // SSH commands for RC.MGR
//     sh "ssh RC.MGR@SED1 \"EMANT TRA; DE\""

//     // SSH commands for RX.FER
//     sh "ssh RX.FER@SED1 \"VOLUME $SYSTEM.EMANT; PURGE *; FI\""
// }

// func additionalProfileSetup() {
// 	// Implement additional profile setup using SSH commands
// 	// SSH commands for RC.MGR
//     ssh RC.MGR@SED1 "VOLUME $AUDIT.EMANTS; BINSTALL; INSTALL EMAN{$release_letter}S EMAN{$release_letter}"

//     // SSH commands for RC.MGR
//     ssh RC.MGR@SED1 "VOLUME $AUDIT.EMANT; PURGE PROFILES; FUP DUP EMANT.PROFILES,*; EDIT PROFILES"

//     // SSH commands for RX.FER@SED1
//     ssh RX.FER@DEV2 "VOLUME $EMAN.EMAN{$release_letter}; EDIT PROFILES; LA; ADD <line number>"

//     //Paste the text copied previously. Review and save changes.

//     // SSH commands for RX.FER@DEV2
//     ssh RX.FER@DEV2 "EMAN{$release_letter} {$release_letter}RH; RE {$release_letter}RA ALL"
// }

// func getInput(scanner *bufio.Scanner) string {
// 	scanner.Scan()
// 	input := scanner.Text()
// 	if input == "" {
// 		return params.ReleaseBuild
// 	}
// 	return input
// }

// func main() {
// 	fmt.Println("Default Parameters:")
// 	displayParameters()

// 	fmt.Print("Do you want to update parameters? (y/n): ")
// 	var choice string
// 	fmt.Scanln(&choice)

// 	if strings.ToLower(choice) == "y" {
// 		updateParameters()
// 	}

// 	fmt.Println("Updated Parameters:")
// 	displayParameters()

// 	fmt.Println("Performing Pre-staging Activities:")
// 	prestagingActivities()

// 	fmt.Println("Performing Additional Profile Setup:")
// 	additionalProfileSetup()

// 	fmt.Println("Performing Additional SSH Commands:")
// 	additionalSSHCommands()
// }


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
	PrivateKeyPath         string
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
	PrivateKeyPath:         "/path/to/your/private/key",
}

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

func runSSHCommand(command string, userHost ...string) {
	cmdArgs := []string{"-i", params.PrivateKeyPath, "-p", params.RemotePort}

	if len(userHost) > 0 {
		remoteUserHost := strings.Split(userHost[0], "@")
		params.RemoteUser = remoteUserHost[0]
		params.RemoteHost = remoteUserHost[1]
	}

	cmdArgs = append(cmdArgs, fmt.Sprintf("%s@%s", params.RemoteUser, params.RemoteHost), command)

	cmd := exec.Command("ssh", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func prestagingActivities() {
	runSSHCommand(fmt.Sprintf("ls -ld \"/L/ctc_data/ctc_data.%s.%s\"", params.ReleaseBuild, params.Database))
	runSSHCommand(fmt.Sprintf("ls -ld \"/L/ctcws/ctcws.%s\"", params.Ctcws))
	runSSHCommand(fmt.Sprintf("ls -ld \"/L/gds/tma.%s\"", params.Tma))
}

func additionalSSHCommands() {
	runSSHCommand("EMANT TP1; DE", "PRODSE.USER1@SED1")
	runSSHCommand("EMANT TT2; DE", "TRAIN.TRNING@SED1")
	runSSHCommand("EMANT TRA; DE", "RC.MGR@SED1")
	runSSHCommand("VOLUME $SYSTEM.EMANT; PURGE *; FI", "RX.FER@SED1")
}

func additionalProfileSetup() {
	runSSHCommand("VOLUME $AUDIT.EMANTS; BINSTALL; INSTALL EMAN{$release_letter}S EMAN{$release_letter}", "RC.MGR@SED1")
	runSSHCommand("VOLUME $AUDIT.EMANT; PURGE PROFILES; FUP DUP EMANT.PROFILES,*; EDIT PROFILES", "RC.MGR@SED1")
	runSSHCommand("VOLUME $EMAN.EMAN{$release_letter}; EDIT PROFILES; LA; ADD <line number>", "RX.FER@DEV2")
	runSSHCommand("EMAN{$release_letter} {$release_letter}RH; RE {$release_letter}RA ALL", "RX.FER@DEV2")
}

func displayParameters() {
	fmt.Println("Release Build:", params.ReleaseBuild)
	fmt.Println("Release Letter:", params.ReleaseLetter)
	fmt.Println("Build Number:", params.BuildNumber)
	fmt.Println("Previous Release Letter:", params.PrevReleaseLetter)
	fmt.Println("2 Releases Prior Letter:", params.TwoReleasesPriorLetter)
	fmt.Println("Host:", params.Host)
	fmt.Println("Database:", params.Database)
	fmt.Println("Prefix Letter:", params.PrefixLetter)
	fmt.Println("New Disk:", params.NewDisk)
	fmt.Println("Old Disk:", params.OldDisk)
	fmt.Println("Port:", params.Port)
	fmt.Println("CTCWS:", params.Ctcws)
	fmt.Println("TMA:", params.Tma)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Default Parameters:")
	displayParameters()

	fmt.Print("Do you want to update parameters? (y/n): ")
	var choice string
	fmt.Scanln(&choice)

	if strings.ToLower(choice) == "y" {
		updateParameters(scanner)
	}

	fmt.Println("Updated Parameters:")
	displayParameters()

	fmt.Println("Performing Pre-staging Activities:")
	prestagingActivities()

	fmt.Println("Performing Additional Profile Setup:")
	additionalProfileSetup()

	fmt.Println("Performing Additional SSH Commands:")
	additionalSSHCommands()
}
