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
// 	Date:			Author:		Info:
//	Feb 1, 2017		LIS			First release
//
// TODO:

package main

import (
	"fmt"
	"os"

	myInit "github.com/my10c/simple-aws-ec2/initialize"
	myAWS "github.com/my10c/simple-aws-ec2/aws"
	myEC2 "github.com/my10c/simple-aws-ec2/ec2"
)

func main() {
	args := myInit.InitArgs()
	config := myInit.GetConfig(args)
	myInit.InitLog(config)
	sess := myAWS.InitSession(config["AwsFile"], config["Profile"], config["Region"])
	ec2Ptr := myEC2.New(sess, args["Force"], args["DryRun"], args["WildCard"])
	switch args["action"] {
		case "list":
			ec2Ptr.PrintInstances()
		case "status":
			ec2Ptr.StatusInstanceByTag(args["instance"])
		case "start":
			//ec2Ptr.StartInstanceByTag(args["instance"])
		case "stop":
			//ec2Ptr.StopInstanceByTag(args["instance"])
		case "terminate":
			if config["AllowTerminate"] == "false" {
				fmt.Printf("AllowTerminate is set to false, aborting..")
				ec2Ptr.Release()
				os.Exit(2)
			}
			//ec2Ptr.TerminateInstanceByTag(args["instance"])
	}
	ec2Ptr.Release()
	os.Exit(0)
}
