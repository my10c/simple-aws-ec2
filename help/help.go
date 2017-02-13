// Copyright (c) 2017 Badassops
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//	* Redistributions of source code must retain the above copyright
//	notice, this list of conditions and the following disclaimer.
//	* Redistributions in binary form must reproduce the above copyright
//	notice, this list of conditions and the following disclaimer in the
//	documentation and/or other materials provided with the distribution.
//	* Neither the name of the <organization> nor the
//	names of its contributors may be used to endorse or promote products
//	derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSEcw
// ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Author		:	Luc Suryo <luc@badassops.com>
//
// Version		:	0.1
//
// Date			:	Feb 10, 2017
//
// History	:
// 	Date:			Author:		Info:
//	Feb 10, 2017	LIS			First release
//
// TODO:

package help

import (
	"fmt"
	"os"

	myGlobal "github.com/my10c/simple-aws-ec2/global"
)

// Function to show how to setup the configuration file
func SetupHelp() {
	fmt.Printf("%s", myGlobal.MyInfo)
	fmt.Printf("\n\t1. Get an AWS API key pair, and the region.\n")
	fmt.Printf("\t2. Create the directory %s, set permissions to 0700.\n", myGlobal.DefaultValues["ConfigPath"])
	fmt.Printf("\t3. Make sure to set both configuration files below, permissions to 0600.\n\n")
	fmt.Printf("Setup the AWS credential file: %s\n", myGlobal.DefaultValues["AwsFile"])
	fmt.Printf("\tAdd the followong lines, all except regions are required\n")
	fmt.Printf("\t[%s]\n", myGlobal.DefaultValues["Profile"])
	fmt.Printf("\taws_access_key_id = {your-aws_access_key_id from 1 above}\n")
	fmt.Printf("\taws_secret_access_key = {your-aws_secret_access_key from 1 above}\n")
	fmt.Printf("\tregion = {your-aws-region from 1 above}\n")
	fmt.Printf("\nSetup the configuration file: %s.\n", myGlobal.DefaultValues["ConfFile"])
	fmt.Printf("\t# Optional For AWS, shown default values: **\n")
	fmt.Printf("\tAWS:\n")
	fmt.Printf("\t  AwsFile: %s\n", myGlobal.DefaultValues["AwsFile"])
	fmt.Printf("\t  Profile: %s\n", myGlobal.DefaultValues["Profile"])
	fmt.Printf("\t  Region: %s\n", myGlobal.DefaultValues["Region"])
	fmt.Printf("\t  Debug: %s\n", myGlobal.DefaultValues["Debug"])
	fmt.Printf("\t  AllowTerminate: %s\n", myGlobal.DefaultValues["AllowTerminate"])
	fmt.Printf("\t# Optional For logging, shown default values: **\n")
	fmt.Printf("\tLogging:\n")
	fmt.Printf("\t  LogFile: %s\n", myGlobal.DefaultValues["LogFile"])
	fmt.Printf("\t  LogMaxSize: %s [megabytes]\n", myGlobal.DefaultValues["LogMaxSize"])
	fmt.Printf("\t  LogMaxBackups: %s [files count]\n", myGlobal.DefaultValues["LogMaxBackups"])
	fmt.Printf("\t  LogMaxAge: %s [n days]\n\n", myGlobal.DefaultValues["LogMaxAge"])
	os.Exit(0)
}

// Function to show the help information
func Help() {
	defAWSFile	:= myGlobal.DefaultValues["AwsFile"]
	defProfile	:= myGlobal.DefaultValues["Profile"]
	defRegion	:= myGlobal.DefaultValues["Region"]
	defConfig	:= myGlobal.DefaultValues["ConfFile"]
	fmt.Printf("%s", myGlobal.MyInfo)
	requiredList := "[--action=action] [--instance=instance 'Name' tag]"
	optionList := "optional: <--aws_config=aws credential file> <--config=configuration file>\n" +
			"\t<--profile=profile name to use> <--region=region-name>\n" +
			"\t<--wildcard> <--force> <--dryrun> <--setup> <--list> <--version> <--help>\n"
	fmt.Printf("Usage : %s %s\n\t%s\n", myGlobal.MyProgname, requiredList, optionList)
	fmt.Printf("\t\taction: action to be perform on the given instance to the given instance.\n")
	fmt.Printf("\t\t\tvalid actions: start, stop and terminate [1]\n")
	fmt.Printf("\t\tinstance: the instance 'Name' tag name. [2]\n")
	fmt.Printf("\t\taws_config: the AWS credential file to use, default: %s.\n", defAWSFile)
	fmt.Printf("\t\tprofile: the profile to use in the AWS credential file, default to %s.\n", defProfile)
	fmt.Printf("\t\tregion: the AWS region to use, default to %s.\n", defRegion)
	fmt.Printf("\t\tconfig: the configuration file to use, default: %s.\n", defConfig)
	fmt.Printf("\t\twildcard: wildcard the given instance 'Name' tag, this can be dangerous!\n")
	fmt.Printf("\t\tforce: force the stop. Works only for the stop action.\n")
	fmt.Printf("\t\tdryrun: dry run the given action.\n")
	fmt.Printf("\t\tsetup: show the setup guide.\n")
	fmt.Printf("\t\tlist: show all instances in the set region.\n")
	fmt.Printf("\t\tversion: show the program version.\n")
	fmt.Printf("\t[1] terminate requires the setting AllowTerminate to be set to true\n")
	fmt.Printf("\t[2] the program relies on the fact that all instances has their EC2 'Name' tag filled.\n")
	os.Exit(2)
}
