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

	myGlobal "github.com/my10c/simple-aws-ec2/global"
)

// Function to get the instance status by Name tag
func (ec2Ptr *Ec2) GetStatusByTag(instanceTag string) string {
	return ec2Ptr.InstancesInfo[instanceTag]["state"]
}

// function internal to print status of given instance
func (ec2Ptr *Ec2) statusInstanceByTag(instanceTag string) bool {
	if _, result := ec2Ptr.InstancesInfo[instanceTag]; result == false {
		fmt.Printf("Given instance Name (%s) not found.\n", instanceTag)
		return false
	}
	fmt.Printf("%s%s\n", myGlobal.HR,myGlobal.HR)
	fmt.Printf("State %s |\tInstanceID %s |\tName %s\n",
		ec2Ptr.InstancesInfo[instanceTag]["state"],
		ec2Ptr.InstancesInfo[instanceTag]["instanceID"], instanceTag)
	fmt.Printf("%s%s\n", myGlobal.HR,myGlobal.HR)
	return true
}

// function internal to print status of given instance wildcard
func (ec2Ptr *Ec2) statusInstanceByTagWildcard(instanceTag string) bool {
	wildCardMap :=  ec2Ptr.GetWildCardMap(instanceTag)
	if len(wildCardMap) == 0 {
		fmt.Printf("Given instance Name (%s) not found.\n", instanceTag)
		return false
	}
	fmt.Printf("%s%s\n", myGlobal.HR,myGlobal.HR)
	for idx := range wildCardMap {
		fmt.Printf("State %s |\tInstanceID %s |\tName %s\n",
			wildCardMap[idx]["state"],
			wildCardMap[idx]["instanceID"], idx)
	}
	fmt.Printf("%s%s\n", myGlobal.HR,myGlobal.HR)
	return true
}

// Function to get the status of the given instance Name tag
func (ec2Ptr *Ec2) StatusInstanceByTag(instanceTag string) bool {
	if ec2Ptr.WildCard == false {
		return ec2Ptr.statusInstanceByTag(instanceTag)
	}
	return ec2Ptr.statusInstanceByTagWildcard(instanceTag)
}
