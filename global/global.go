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
// Date			:	Feb 1, 2017
//
// History	:
//	Date:			Author:		Info:
//	Feb 1, 2017		LIS			First release
//

package global

import (
	"os"
	"path"
	"strconv"
	"time"
)
const (
	HR = "__________________________________________________"
)

var (
	now			= time.Now()
	MyProgname	= path.Base(os.Args[0])
	myAuthor	= "Luc Suryo"
	myCopyright	= "Copyright 2016 - " + strconv.Itoa(now.Year()) + " ©Badassops"
	myLicense	= "BSD, http://www.freebsd.org/copyright/freebsd-license.html ♥"
	// Global variables
	MyVersion	= "0.1"
	myEmail		= "<luc@badassops.com>"
	MyInfo		= MyProgname + " " + MyVersion + "\n" + myCopyright + "\nLicense" + myLicense +
					"\nWritten by " + myAuthor + " " + myEmail + "\n"

	// require and optional configuration keys
	RequiredKeys	[]string
	AWSKeys			[]string
	LoggingKeys		[]string
	DefaultValues	map[string]string
	Actions			string

	// for AWS setting
	defaultConfigPath	= os.Getenv("HOME") + "/.aws"
	defaultProfile		= "simple-aws-ec2"
	defaultRegion		= "us-east-1"
	defaultDebug		= "false"
	defaultTerminate	= "false"

	// for logging
	defaultLogFile			= os.Getenv("HOME") + "/.aws/" + MyProgname + ".log"
	defaultLogMaxSize		= 128	// megabytes
	defaultLogMaxBackups	= 3		// 3 files
	defaultLogMaxAge		= 10	// days
)

func init() {
	// supported action
	Actions =	"start stop status terminate"

	// this is need for viper to find the correct value in the correct tree
	RequiredKeys	= []string{}
	AWSKeys			= []string{"AwsFile", "Profile", "Region", "Debug", "AllowTerminate"}
	LoggingKeys		= []string{"LogFile", "LogMaxSize", "LogMaxBackups", "LogMaxAge"}

	// Hardcode default configuration file
	awsFile		:= defaultConfigPath + "/simple-ec2.cred"
	confFile	:= defaultConfigPath + "/simple-ec2.conf"

	// setup the default value, these are hardcoded.
	DefaultValues = make(map[string]string)
	DefaultValues["ConfigPath"]			=	defaultConfigPath
	DefaultValues["ConfFile"]			=	confFile
	DefaultValues["AwsFile"]			=	awsFile
	DefaultValues["Profile"]			=	defaultProfile
	DefaultValues["Region"]				=	defaultRegion
	DefaultValues["Debug"]				=	defaultDebug
	DefaultValues["AllowTerminate"]		=	defaultTerminate
	DefaultValues["LogFile"]			=	defaultLogFile
	DefaultValues["LogMaxSize"]			=	strconv.Itoa(defaultLogMaxSize)
	DefaultValues["LogMaxBackups"]		=	strconv.Itoa(defaultLogMaxBackups)
	DefaultValues["LogMaxAge"]			=	strconv.Itoa(defaultLogMaxAge)
}
