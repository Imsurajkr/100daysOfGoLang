/*
The Go garbage collector is responsible for collecting the memory that is not in use anymore.
 The implemented algorithm is a concurrent tri-color mark and sweep collector
 Marking phase
This phase performs a scan of the memory to know which blocks are still in use by our code
 and which ones should be collected.
However, since the garbage collector can run concurrently with our Go program, it needs a
way to detect potential changes in the memory while scanning. To tackle that potential issue,
an algorithm of write barrier is implemented and will allow Go to track any pointer changes.
The only condition to enable write barriers is to stop the program for a short time,
also called “Stop the World”:
Go also starts, at the beginning of the process, a marking worker per processor that help
with marking the memory.
Then, once the roots have been enqueued for processing, the marking phase can start
traversing and coloring the memory.
The garbage collector starts from the stack and follows pointers recursively to go through
 the memory. Spans that are marked as no scan stop the scanning. However, this process
  is not done by the same goroutine; each pointer is enqueued in a work pool. Then,
   the background mark workers seen previously
dequeue works from this pool, scan the objects and then enqueue the pointers found in it:
The garbage collector starts from the stack and follows pointers recursively to go
 through the memory. Spans that are marked as no scan stop the scanning. However,
 this process is not done by the same goroutine; each pointer is enqueued in a work pool.
  Then, the background mark workers seen previously dequeue works from this pool, scan the
   objects and then enqueue the pointers found in it:
   As a first state, all objects are considered white. Then, the objects are traversed and the ones reachable will turn grey. If an object is in a span marked as no scan, it can be painted in black since it does not need to be scanned:
   Grey objects are now enqueued to be scanned and turn black:
   The same thing happens for the objects enqueue until there are no more objects to process:

   t the end of the process, black objects are the ones in-use in memory when white objects are the ones to be collected. As we can see, since the instance of struct2 has been created in an anonymous function and is not reachable from the stack, it stays white and can be cleaned.
The colors are internally implemented thanks to a bitmap attribute in each span called gcmarkBits that traces the scan with setting to 1 the corresponding bit:
Image for post
color implementation with a bitmap
As we can see, the black and the grey color works the same way. The difference in the process is that the grey color enqueues an object to be scanned when black objects end the chain.
The garbage finally stops the world, flushes the changes made on each write barrier to the work pool and performs the remaining marking.
*/