/*
==================what are go routines ?
A goroutine is a lightweight thread managed by the Go runtime.
=====================================goroutines vs threads ?
GOROUTINE	                                                        THREAD
1.Goroutines are managed by the go runtime.	                   1.Operating system threads are managed by kernal.
2.Goroutine are not hardware dependent.	                       2.Threads are hardware dependent.
3.Goroutines have easy communication medium known as channel.  3.Thread doesnot have easy communication medium.
4.Due to the presence of channel one goroutine can             4.Due to lack of easy communication medium inter-threads communicate .
communicate with other goroutine with low latency.	             takes place with high latency
5.Goroutine doesnot have ID because go doesnot have             5.Threads have their own unique ID because they have Thread Local Storage.
Thread Local Storage.

6.Goroutines are cheaper than threads.	                          6.The cost of threads are higher than goroutine.
7.They are cooperatively scheduled.	                              7.They are preemptively scheduled.
8.They have fasted startup time than threads.                     8.They have slow startup time than goroutines.
9.Goroutine has growable segmented stacks.                    	9.Threads doesnot have growable segmented stacks.

===================why goroutines are lightweight process?
A goroutine is created with initial only 2KB of stack size.
Each function in go already has a check if more stack is needed or not and the stack can be
copied to another region in memory with twice the original size. This makes goroutine very
light on resources.

================why goroutines are hardware independent?
Goroutines exists only in the virtual space of go runtime and not in the OS

=======================Blocking in go routines?
a Go Runtime scheduler is needed which manages their lifecycle.
Go Runtime maintains three C structs for this purpose:
1.The G Struct : This represents a single go routine with it’s properties such as stack pointer, base of stack, it’s ID, it’s cache and it’s status
2.The M Struct : This represents an OS thread. It also contains a pointer to the global queue of runnable goroutines, the current running goroutine and the reference to the scheduler
3.The Sched Struct : It is a global struct and contains the queues free and waiting goroutines as well as threads.

=====================Goroutines c structs
If a goroutine blocks on system call, it blocks it’s running thread.
But another thread is taken from the waiting queue of Scheduler (the Sched struct) and
used for other runnable goroutines.
However, if you communicate using channels in go which exists only in virtual space,
 the OS doesn’t block the thread. Such goroutines simply go in the waiting state and other
  runnable goroutine (from the M struct) is scheduled in it’s place.

*/