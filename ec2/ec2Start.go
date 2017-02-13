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
// Date			:	Feb 11, 2017
//
// History	:
// 	Date:			Author:		Info:
//	Feb 11, 2017	LIS			First relase
//
// TODO:

package ec2

import (
	"fmt"
	"os/user"
	"time"

	myUtils "github.com/my10c/simple-aws-ec2/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Function to start an Instance by Name tag
func (ec2Ptr *Ec2) startInstanceByTag(instanceTag string) bool {
	if _, result := ec2Ptr.InstancesInfo[instanceTag]; result == false {
		fmt.Printf("Given instance Name (%s) not found.\n", instanceTag)
		return false
	}
	instanceStatus := ec2Ptr.GetStatusByTag(instanceTag)
	if instanceStatus == "running" {
		myUtils.StdOutAndLog(fmt.Sprintf("Instance %s is already running, state %s", instanceTag, instanceStatus))
		return false
	} else {
		userPtr, _ := user.Current()
		userName := userPtr.Username
		currTime := time.Now().Format(time.RFC850)
		instanceID := ec2Ptr.GetInstanceIDFromTag(instanceTag)
		params := &ec2.StartInstancesInput{
				InstanceIds: []*string{
					aws.String(instanceID),
				},
			AdditionalInfo: aws.String(fmt.Sprintf("Started by %s at %s", userName, currTime)),
			DryRun: aws.Bool(ec2Ptr.DryRun),
		}
		ec2Ptr.mu.Lock()
		defer ec2Ptr.mu.Unlock()
		_, err := ec2Ptr.session.StartInstances(params)
		if err != nil {
			myUtils.ExitIfError(err)
		}
		myUtils.StdOutAndLog(fmt.Sprintf("Instance %s is starting..", instanceTag))
		return true
	}
}

// Function to start an Instance by Name tag
func (ec2Ptr *Ec2) StartInstanceByTag(instanceTag string) bool {
	if ec2Ptr.WildCard == false {
		return ec2Ptr.startInstanceByTag(instanceTag)
	}
	wildCardMap :=  ec2Ptr.GetWildCardMap(instanceTag)
	if len(wildCardMap) == 0 {
		fmt.Printf("Given instance Name (%s) not found.\n", instanceTag)
		return false
	}
	for idx := range wildCardMap {
		ec2Ptr.startInstanceByTag(idx)
	}
	return true
}
