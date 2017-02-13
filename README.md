# UNFINISHED

## simple-aws-ec2 : A Simple Go AWS EC2/ALB admin tool

### Background
The script is a simple script to start, stop and terminate a AWS2 EC2 instance
as well getting the instance status given `Name` tag. The script also support
wildcard `Name` tag, be carefull becuase it can be dangerous!


### Tag... again
The script is based by the fact every singe instance has the AWS 'Name' tag filled!
without this the script will break...


### Setup

To se how to setup the configuration, run the script with the --setup flag

```
simple-aws-ec2 0.1
Copyright 2016 - 2017 ©Badassops
LicenseBSD, http://www.freebsd.org/copyright/freebsd-license.html ♥
Written by Luc Suryo <luc@badassops.com>

	1. Get an AWS API key pair, and the region.
	2. Create the directory /home/momo/.aws, set permissions to 0700.
	3. Make sure to set both configuration files below, permissions to 0600.

Setup the AWS credential file: /home/momo/.aws/simple-ec2.cred
	Add the followong lines, all except regions are required
	[simple-aws-ec2]
	aws_access_key_id = {your-aws_access_key_id from 1 above}
	aws_secret_access_key = {your-aws_secret_access_key from 1 above}
	region = {your-aws-region from 1 above}

Setup the configuration file: /home/momo/.aws/simple-ec2.conf.
	# Optional For AWS, shown default values: **
	AWS:
	  AwsFile: /home/momo/.aws/simple-ec2.cred
	  Profile: simple-aws-ec2
	  Region: us-east-1
	  Debug: false
	  AllowTerminate: false
	# Optional For logging, shown default values: **
	Logging:
	  LogFile: /home/momo/.aws/admin.log
	  LogMaxSize: 128 [megabytes]
	  LogMaxBackups: 3 [files count]
	  LogMaxAge: 10 [n days]
```

### Options flags

with any argument the script will show some help information

```
```

Or with the --help a shorter version of the help page

```
Usage of simple-aws-ec2:
  -action value
    	Valid instance actions: start stop status terminate.
  -aws_config string
    	AWS configuration file to use. (default "/home/momo/.aws/simple-ec2.cred")
  -config string
    	Configuration file to use. (default "/home/momo/.aws/simple-ec2.conf")
  -dryrun
    	Dry run the given action.
  -force
    	Force the stop action, works only with stop.
  -instance value
    	Instance tag name.
  -list
    	Show all instances in the set region.
  -profile string
    	AWS profile to use. (default "simple-aws-ec2")
  -region string
    	EC2 region to use. (default "us-east-1")
  -setup
    	Show the setup information and exit.
  -version
    	Prints current version and exit.
  -wildcard
    	Use name wildcard, this can be dangerous!
exit status 2
```

### Examples

Getting status of instances with `Name` tag momo
```
simple-aws-ec2 --instance momo --wildcard --action status
```

Start an instance with `Name` tag momo1
```
simple-aws-ec2 --instance momo1 --action start
```

Get list of all instances
```
simple-aws-ec2 --list 
```

Enjoy, Momo
