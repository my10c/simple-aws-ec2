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
// Version		:	0.2
//
// Date			:	Feb 11, 2017
//
// History	:
// 	Date:			Author:		Info:
//	Jan 30, 2017	LIS			First relase
//	Feb 11, 2017	LIS			Added support for Force, DryRun and WildCard
//
// TODO:

package ec2

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	myUtils "github.com/my10c/simple-aws-ec2/utils"
	myGlobal "github.com/my10c/simple-aws-ec2/global"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Ec2 struct {
	session			*ec2.EC2
	InstancesInfo	map[string]map[string]string
	Force			bool
	DryRun			bool
	WildCard		bool
	mu				sync.Mutex
}

// Function to create a new Ec2 object with a AWS session set
func New(awsSess *session.Session, force string, dryrun string, wildcard string) *Ec2 {
	forceBool, err:= strconv.ParseBool(force)
	if err != nil {
		forceBool = false
	}
	dryrunBool, err := strconv.ParseBool(dryrun)
	if err != nil {
		forceBool = false
	}
	wildCardBool, err := strconv.ParseBool(wildcard)
	if err != nil {
		wildCardBool = false
	}
	ec2PTR := &Ec2 {
		session: ec2Session(awsSess),
		Force: forceBool,
		DryRun: dryrunBool,
		WildCard: wildCardBool,
	}
	ec2PTR.setInstancesInfo()
	return ec2PTR
}

// Function to get a ec2 session
func ec2Session(sess *session.Session) *ec2.EC2 {
	ec2Sess := ec2.New(sess)
	myUtils.ExitIfNill(ec2Sess)
	return ec2Sess
}

// Function set all instanace info to the Ec2 object
func (ec2Ptr *Ec2) setInstancesInfo() {
	ec2Ptr.mu.Lock()
	defer ec2Ptr.mu.Unlock()
	ec2PtrMap := make(map[string]map[string]string)
	// Only running instances!
	params := &ec2.DescribeInstancesInput{
		MaxResults: aws.Int64(1024),
	}
	resp, err := ec2Ptr.session.DescribeInstances(params)
	if err != nil {
		myUtils.ExitIfError(err)
	}
	for idx := range resp.Reservations {
		for instIdx := range resp.Reservations[idx].Instances {
			theInstanceState := *resp.Reservations[idx].Instances[instIdx].State.Name
			theInstanceID := *resp.Reservations[idx].Instances[instIdx].InstanceId
			theInstancePrvIP := *resp.Reservations[idx].Instances[instIdx].PrivateIpAddress
			var theInstanceTag string
			for tagIdx := range resp.Reservations[idx].Instances[instIdx].Tags {
				if *resp.Reservations[idx].Instances[instIdx].Tags[tagIdx].Key == "Name" {
					theInstanceTag = *resp.Reservations[idx].Instances[instIdx].Tags[tagIdx].Value
				}
			}
			// trim the tag, all lowercase and remove the indenty part
			theInstanceTag = strings.ToLower(strings.Split(theInstanceTag, " ")[0])
			ec2PtrMap[theInstanceTag] = make(map[string]string)
			ec2PtrMap[theInstanceTag]["instanceID"]	= theInstanceID
			ec2PtrMap[theInstanceTag]["privateIP"]	= theInstancePrvIP
			ec2PtrMap[theInstanceTag]["state"]		= theInstanceState
		}
	}
	ec2Ptr.InstancesInfo = ec2PtrMap
}

// Function to get Instance ID from the given instance Name tag
func (ec2Ptr *Ec2) GetInstanceIDFromTag(tagName string) string {
	if len(ec2Ptr.InstancesInfo) == 0 {
		ec2Ptr.setInstancesInfo()
	}
	return ec2Ptr.InstancesInfo[tagName]["instanceID"]
}

// Function to get the instance private IP from the given instance Name tag
func (ec2Ptr *Ec2) GetIPrivateIPFromTag(tagName string) string {
	if len(ec2Ptr.InstancesInfo) == 0 {
		ec2Ptr.setInstancesInfo()
	}
	return ec2Ptr.InstancesInfo[tagName]["privateIP"]
}

// Function to get Instance Name tag from the given instance ID
func (ec2Ptr *Ec2) GetInstanceTagFromInstanceID(instanceID string) string {
	var returnValues string
	if len(ec2Ptr.InstancesInfo) == 0 {
		ec2Ptr.setInstancesInfo()
	}
	for key := range ec2Ptr.InstancesInfo {
		if ec2Ptr.InstancesInfo[key]["instanceID"] == instanceID {
			return key
		}
	}
	return returnValues
}

// Function to get Instance ID from the given private IP
func (ec2Ptr *Ec2) GetInstanceIDFromPrivateIP(privateIP string) string {
	var returnValues string
	if len(ec2Ptr.InstancesInfo) == 0 {
		ec2Ptr.setInstancesInfo()
	}
	for key := range ec2Ptr.InstancesInfo {
		if ec2Ptr.InstancesInfo[key]["privateIP"] == privateIP {
			return ec2Ptr.InstancesInfo[key]["instanceID"]
		}
	}
	return returnValues
}

// Function to get Instance ID from Name tag with wildcard
func (ec2Ptr *Ec2) GetInstanceIDs(instanceTag string) []string {
	var instancesID []string
	for  idx := range ec2Ptr.InstancesInfo {
		if strings.Contains(idx, instanceTag) {
			instancesID = append(instancesID, ec2Ptr.GetInstanceIDFromTag(idx))
		}
	}
	return instancesID
}

// Function to get instances map base the given Name tag with wildcard
func (ec2Ptr *Ec2) GetWildCardMap(instanceTag string) map[string]map[string]string {
	wildCardMap := make(map[string]map[string]string)
	for  idx := range ec2Ptr.InstancesInfo {
		if strings.Contains(idx, instanceTag) {
			wildCardMap[idx] = make(map[string]string)
			wildCardMap[idx]["instanceID"]	= ec2Ptr.InstancesInfo[idx]["instanceID"]
			wildCardMap[idx]["privateIP"]	= ec2Ptr.InstancesInfo[idx]["privateIP"]
			wildCardMap[idx]["state"]		= ec2Ptr.InstancesInfo[idx]["state"]
		}
	}
	return wildCardMap
}

// Function to get Instance id from either tag or ip
func (ec2Ptr *Ec2) GetInstanceID(instanceTag string, isIP bool) (string, bool) {
	var instanceID string
	if isIP == true {
		instanceID = ec2Ptr.GetInstanceIDFromPrivateIP(instanceTag)
	} else {
		instanceID = ec2Ptr.GetInstanceIDFromTag(instanceTag)
	}
	if len(instanceID) == 0 {
		return instanceID, false
	}
	return instanceID, true
}

// Function to print all instances info
func (ec2Ptr *Ec2) PrintInstances() {
	fmt.Printf("%s%s\n", myGlobal.HR,myGlobal.HR)
	for idx := range ec2Ptr.InstancesInfo {
		fmt.Printf("State %s |\tInstanceID %s |\tName %s\n",
			ec2Ptr.InstancesInfo[idx]["state"],
			ec2Ptr.InstancesInfo[idx]["instanceID"], idx)
	}
	fmt.Printf("%s%s\n", myGlobal.HR,myGlobal.HR)
}

// Function to release memory
func (ec2Ptr *Ec2) Release() {
	ec2Ptr = nil
}
