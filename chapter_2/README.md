# Communicating Sequential Processes

## Difference between Concurrency and Parallelism
### Concurrency
Concurrency is a programming paradigm that allows us to write pieces of code that could be executed at the same time.
Rob Pike mentions that concurrency is about the structure, how we write the code so that in case that multiples processes or functions are executed at the same time, everything works as expected.
Concurrency doesn't provide any warranty that the functions will be executed in a parallel way.

### Parallelism
Parallelism is when two processes or functions are actually executed at the same time. This has to be managed directly by the operating system and it has a strong relationship with the system when the application runs.
Rob Pike mentions that Parallelism is about the execution, how the code is executed in a specific environment.

## What Golang provides to simplify writing concurrency code
Golang provides a mechanism to abstract (separate) the parallelism and the concurrency, so we can write concurrent code without thinking too much about the implications in the system, that means, without thinking about managing the needed resources to execute parallel code, such as threads.

### GO Routines
Different to other programming languages, in Go we don't interact directly with the operating system to execute concurrent code but we rely on some abstraction provides by the go runtime. The principal abstraction are the go-routines. 
Go routines, in opposite to Threads, are managed directly by the go-runtime, who is in charge or multiplexing all the go-routines into the multiple OS threads created.

In the process of multiplexing the go-routines, the go-runtime will manage create some threads. Each thread will have a queue where the go-routines will be scheduled. 
Every time that a go-routine is created, a small stack will be assigned to that routine (1kb), and the stack size will dynamically change depending on the needed memory, allocating more memory or shrinking the stack.
In case that a thread is block for any reason, other threads can steal work from its queue, balancing the amount of work between all the threads.

All of this provides a more cheaper way to handle concurrent code in comparison to the more tradition way that involves directly creating and managing threads.
Each time a thread is created, it will have a fixed stack size (1mb). This: 
- limits the max number of threads that can be created
- limits the max memory that each thread can handle (more than 1 mb will be a memory overflow)
- If the thread needs less memory, the not needed memory will be wasted.

### Channels
Channels are very similar Pipes in Linux. They allow us to send information (communicate) from one point to another one.
Although Channels does not restrict who is the consumer and who is the producer, it's recommended to keep always one role and don't change it, that means, a consumer shouldn't change to be a producer.

Channels are concurrent-safe, so they can be used freely inside the go-routines.

Channels follow the ideas proposed in the Communicating Sequential Processes paper, where recommended creating a structure that allows passing information from one point to another one.


### When should we use channels and when should we use locks.
Some guidance to choose between using channels or more traditional mechanism to handle shared memory:
- If the performance of a function is critical, it's recommended to use traditional locks. Internally channels use locks, so always be an overhead. This only applies for real critical sections.
- If we want to change the data ownership from one routine to another one, channels should be use in those cases. It's recommended that only one routine owns a piece of data in a specific moment. If we want to share this piece of data, it is recommended, instead of sharing memory, send the data (communicate) to the next routine, so this one can now own the data. (don't communicate sharing memory, but share memory through communicating.)
- If we want keep enclosed the state of a struct (or other piece of data), traditional locks should be used, specially if multiple go routines will be handling this struct in a concurrent way.
- If we want to coordinate multiple go-routines, channels must be used for that.
