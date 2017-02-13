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

package initialize

import (
	"fmt"
	"flag"
	"log"
	"os"
	"strconv"

	myGlobal	"github.com/my10c/simple-aws-ec2/global"
	myHelp		"github.com/my10c/simple-aws-ec2/help"
	myUtils		"github.com/my10c/simple-aws-ec2/utils"

	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

type stringFlag struct {
	value	string
	set		bool
}

var (
	myAction	stringFlag
	myInstance	stringFlag
	myAWSFile	string
	myProfile	string
	myConfFile	string
	myRegion	string
)

// Function for the stringFlag struct, set the values
func (sf *stringFlag) Set(x string) error {
	sf.value = x
	sf.set = true
	return nil
}

// Function for the stringFlag struct, get the values
func (sf *stringFlag) String() string {
	return sf.value
}

// Function to return read the configuration file
// Also make sure the required keys are set
func GetConfig(argv map[string]string) map[string]string {
	var missingKeys []string
	argsMap := make(map[string]string)
	viper.SetConfigFile(argv["ConfFile"])
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	myUtils.ExitIfError(err)
	// get the required keys : AWS
	for confKey := range myGlobal.RequiredKeys {
		if viper.IsSet("AWS." + myGlobal.RequiredKeys[confKey]) {
			argsMap[myGlobal.RequiredKeys[confKey]] = viper.GetString("AWS." + myGlobal.RequiredKeys[confKey])
		} else {
			missingKeys = append(missingKeys, myGlobal.RequiredKeys[confKey])
		}
	}
	// get the optional keys: AWS, do not override if it was given from the CLI
	for optKey := range myGlobal.AWSKeys {
		keyName := myGlobal.AWSKeys[optKey]
		if _, hasVal := argv[keyName]; hasVal == false {
			if viper.IsSet("AWS." + keyName) {
				argsMap[keyName] = viper.GetString("AWS." + keyName)
			} else {
				// was not set, lets use the default from global
				argsMap[keyName] = myGlobal.DefaultValues[keyName]
			}
		} else {
			argsMap[keyName] =  argv[keyName]
		}
	}
	// get the optional keys: LoggingKeys, do not override if it was given from the CLI
	for optKey := range myGlobal.LoggingKeys {
		keyName := myGlobal.LoggingKeys[optKey]
		if _, hasVal := argv[keyName]; hasVal == false {
			if viper.IsSet("Logging." + keyName) {
				argsMap[keyName] = viper.GetString("Logging." + keyName)
			} else {
				// was not set, lets use the default from global
				argsMap[keyName] = myGlobal.DefaultValues[keyName]
			}
		} else {
			argsMap[keyName] =  argv[keyName]
		}
	}
	if len(missingKeys) != 0 {
		fmt.Printf("Following keys are missing in the configration files: %s\n", missingKeys)
		os.Exit(2)
	}
	return argsMap
}

// Function to initialize logging
func InitLog(logSettings map[string]string) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	MaxSize, _		:= strconv.Atoi(logSettings["LogMaxSize"])
	MaxBackups, _	:= strconv.Atoi(logSettings["LogMaxBackups"])
	MaxAge, _		:= strconv.Atoi(logSettings["LogMaxAge"])
	log.SetOutput(&lumberjack.Logger{
		Filename:	logSettings["LogFile"],
		MaxSize:	MaxSize,
		MaxBackups:	MaxBackups,
		MaxAge:		MaxAge,
	})
}

// Function to process the given args
func InitArgs() map[string]string {
	argsMap := make(map[string]string)
	actionList	:= myGlobal.Actions
	defAWSFile	:= myGlobal.DefaultValues["AwsFile"]
	defProfile	:= myGlobal.DefaultValues["Profile"]
	defRegion	:= myGlobal.DefaultValues["Region"]
	defConfig	:= myGlobal.DefaultValues["ConfFile"]
	// set the default force and dryrun, override it at the end if needed
	argsMap["Force"]	= "false"
	argsMap["DryRun"]	= "false"
	argsMap["WildCard"]	= "false"
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", myGlobal.MyProgname)
		flag.PrintDefaults()
	}
	version := flag.Bool("version", false, "Prints current version and exit.")
	setup := flag.Bool("setup", false, "Show the setup information and exit.")
	list := flag.Bool("list", false, "Show all instances in the set region.")
	wildCard := flag.Bool("wildcard", false, "Use name wildcard, this can be dangerous!")
	force := flag.Bool("force", false, "Force the stop action, works only with stop.")
	dryRun := flag.Bool("dryrun", false, "Dry run the given action.")
	flag.StringVar(&myAWSFile, "aws_config", defAWSFile, "AWS configuration file to use.")
	flag.StringVar(&myProfile, "profile", defProfile, "AWS profile to use.")
	flag.StringVar(&myConfFile, "config", defConfig, "Configuration file to use.")
	flag.Var(&myAction, "action", "Valid instance actions: "+actionList+".")
	flag.Var(&myInstance, "instance", "Instance tag name.")
	flag.StringVar(&myRegion, "region", defRegion, "EC2 region to use.")
	flag.Parse()
	if *version {
		fmt.Printf("%s\n", myGlobal.MyVersion)
		os.Exit(0)
	}
	if *setup {
		myHelp.SetupHelp()
	}
	// make sure config file exist
	if myUtils.CheckFileExist(myAWSFile) == false {
		myUtils.StdOutAndLog(fmt.Sprintf("AWS file %s does not exist, aborting.", myAWSFile))
		os.Exit(2)
	}
	if myUtils.CheckFileExist(myConfFile) == false {
		myUtils.StdOutAndLog(fmt.Sprintf("Configuration file %s does not exist, aborting.", myConfFile))
		os.Exit(2)
	}
	// make sure the correct action was given
	if !*list {
		actionArray := myUtils.MakeSliceFromStrings(actionList)
		actionOk := false
		for _, actionName := range actionArray {
				if myAction.value == actionName  {
					argsMap["action"] = myAction.value
					actionOk = true
					break
				}
		}
		if actionOk == false {
			myUtils.StdOutAndLog(fmt.Sprintf("Unsupported action %s, aborting.", myAction.value))
			os.Exit(2)
		}
	}
	argsMap["AwsFile"]		= myAWSFile
	argsMap["Profile"]		= myProfile
	argsMap["ConfFile"]		= myConfFile
	argsMap["Region"]		= myRegion
	if *wildCard {
		argsMap["WildCard"] = "true"
	}
	if *force {
		argsMap["Force"] = "true"
	}
	if *dryRun {
		argsMap["DryRun"] = "true"
	}
	if *list {
		argsMap["action"]	= "list"
		return argsMap
	}
	if !myInstance.set || !myAction.set {
		myHelp.Help()
	}
	argsMap["instance"]		= myInstance.value
	argsMap["action"]		= myAction.value
	return argsMap
}
