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
// Date			:	Jan 27, 2017
//
// History	:
// 	Date:			Author:		Info:
//	Feb 10, 2017	LIS			Copied from other project and made it independent
//
// TODO:

package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"strconv"
)

// Function to exit if an error occured
func ExitIfError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: "+fmt.Sprint(err))
		log.Printf("-< %s >-\n", fmt.Sprint(err))
		os.Exit(1)
	}
}

// Function to exit if pointer is nill
func ExitIfNill(ptr interface{}) {
	if ptr == nil {
		fmt.Fprintln(os.Stderr, "Error: got a nil pointer.")
		log.Printf("-< Error: got a nil pointer. >-\n")
		os.Exit(1)
	}
}

// Function to print the given message to stdout and log file
func StdOutAndLog(message string) {
	fmt.Printf("-< %s >-\n", message)
	log.Printf("-< %s >-\n", message)
	return
}

// Function to check if given IP is a correct ip and whatever is in RFC1918 
func CheckRfc1918Ip(ip string) (bool, bool) {
	rfc1918ten := net.IPNet{IP: net.ParseIP("10.0.0.0"), Mask: net.CIDRMask(8, 32)}
	rfc1918oneninetwo := net.IPNet{IP: net.ParseIP("192.168.0.0"), Mask: net.CIDRMask(16, 32)}
	rfc1918oneseventwo := net.IPNet{IP: net.ParseIP("172.16.0.0"), Mask: net.CIDRMask(12, 32)}
	ip_type := net.ParseIP(ip)
	if ip_type == nil {
		return false, false
	}
	if rfc1918ten.Contains(ip_type) ||
		rfc1918oneninetwo.Contains(ip_type) ||
		rfc1918oneseventwo.Contains(ip_type) {
		return true, true
	}
	return true, false
}

// Function to create string array (slice) from the given strings
func MakeSliceFromStrings(args ...string) []string {
	return strings.Fields(strings.Join(args, " "))
}

// Function to create a string from a string array (slice)
func MakeStringFromSlice(array []*string) string {
	var buffer bytes.Buffer
	for idx := range array {
		buffer.WriteString(*array[idx])
	}
	return buffer.String()
}

// Function to check if the user that runs the app is root
func IsRoot(progName string) {
	if os.Geteuid() != 0 {
		StdOutAndLog(fmt.Sprintf("%s must be run as root.", progName))
		os.Exit(1)
	}
}

// Function to log any reveived signal
func SignalHandler() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt)
	go func() {
		sigId := <-interrupt
		StdOutAndLog(fmt.Sprintf("received %v",sigId))
		os.Exit(0)
	}()
}

// Function to get the MD5 value of the given file
func GetFileMD5(filePath string) (string, error) {
	var returnMD5String string
	// Open the passed argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	// Tell the program to call the following function when the current function returns
	defer file.Close()
	//Open a new hash interface to write to
	hash := md5.New()
	// Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	// Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]
	// Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

// Function to get the MD5 value of the given file pointer
func GetFilePtrMD5(filePtr *os.File) (string, error) {
	var returnMD5String string
	//Open a new hash interface to write to
	hash := md5.New()
	// Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, filePtr); err != nil {
		return returnMD5String, err
	}
	// Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]
	// Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

// Function to write a log if debug was enabled
func WriteDebug(debug string, messsage string) {
	debugMode, err := strconv.ParseBool(debug)
	if err != nil {
		return
	}
	if debugMode == true {
		log.Printf("Debug -< %s >-\n", messsage)
	}
	return
}

// Function to check if given file exist
func CheckFileExist(pathToFile string) bool {
	if _, err := os.Stat(pathToFile); os.IsNotExist(err) {
		return false
	}
	return true
}
