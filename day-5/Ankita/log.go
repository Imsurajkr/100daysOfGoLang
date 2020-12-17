/*
=============Writing to Log Files============
	   what are log files?
	   they are the computer-generated files on your system which contains information
	   about usage patterns,activities operations within an OS,application,server or other device
	   why do we need to write output to log files?
	   because they are permanant ehile output on screen is temporary.so you can search them with some simple commands
	   like grep to view that output.
	   in UNIX OS ,you can find your log files in /var/log
	   in WINDOWS,you can find log files in control panel---->system and security----->Administrative tools--->event viewer

	   ##go offers log package to write into log files
	   Logging level --------------------------------
		it specifies the severity of the log entry
		debug,info,notice,warn,error are some of the logging levels
	  Logging Facilities --------------------------------
	  it is like a category  used for logging information
	  there value can be auth,local1,local2,local3,local4,mark,mail,kern,user,etc.
	  ###if there is no logging facility available then your logs will be ignored.
*/
