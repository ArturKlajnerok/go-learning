# Concurrency

Making progress on more than one task simultaneously is known as concurrency.  
Concurrency is the notion of multiple things happening at the same time.

Do not communicate by sharing memory; instead, share memory by communicating.

https://github.com/golang/go/wiki/LearnConcurrency
https://golang.org/doc/effective_go.html#concurrency

### Articles :

* https://blog.codeship.com/an-intro-to-concurrency-patterns-in-go/
* https://blog.codeship.com/concurrency-in-go-part-ii/
* https://medium.com/golangid/go-102-golang-concurrency-by-go-routine-3e0eab1a8ecb
* https://divan.github.io/posts/go_concurrency_visualize/
* http://mindbowser.com/golang-concurrency/
  <br /><br />
* https://blog.golang.org/concurrency-is-not-parallelism
* https://blog.golang.org/advanced-go-concurrency-patterns
* https://blog.golang.org/pipelines
* https://blog.golang.org/context
  <br /><br />
* https://hackernoon.com/concurrency-how-it-can-help-you-and-how-you-can-use-golang-to-reach-it-easily-ae3e070b3d2c
* https://medium.com/@tilaklodha/concurrency-and-parallelism-in-golang-5333e9a4ba64
* https://www.ardanlabs.com/blog/2014/01/concurrency-goroutines-and-gomaxprocs.html
* https://www.thepolyglotdeveloper.com/2017/05/concurrent-golang-applications-goroutines-channels/
* http://www.soroushjp.com/2015/02/07/go-concurrency-is-not-parallelism-real-world-lessons-with-monte-carlo-simulations/

### Books & Tutorials :

* [x] https://tour.golang.org/concurrency/1
* https://gobyexample.com/goroutines
* https://gobyexample.com/mutexes
* [x] https://www.golang-book.com/books/intro/10
* https://astaxie.gitbooks.io/build-web-application-with-golang/en/02.7.html
* [x] http://www.golangbootcamp.com/book/concurrency
* https://golangbot.com/concurrency/
* https://golangbot.com/goroutines/
* http://www.golangpatterns.info/concurrency/coroutines

### Documentation :

* https://golang.org/pkg/sync/

### Presentations :

* https://talks.golang.org/2012/waza.slide#1
* https://talks.golang.org/2012/concurrency.slide#1
* https://talks.golang.org/2013/advconc.slide#1
* https://www.slideshare.net/PramestiHattaK/golang-101-concurrency-vs-parallelism

### Videos :

* https://vimeo.com/49718712
* https://www.youtube.com/watch?v=cN_DpYBzKso
* https://www.youtube.com/watch?v=f6kdp27TYZs
* https://www.youtube.com/watch?v=QDDwwePbDtw

### General (not golang) :

* <https://en.wikipedia.org/wiki/Concurrency_(computer_science)>
* https://docs.oracle.com/javase/tutorial/essential/concurrency/
* https://developer.apple.com/library/content/documentation/General/Conceptual/ConcurrencyProgrammingGuide/Introduction/Introduction.html
* https://12factor.net/concurrency
* https://www.quora.com/What-is-concurrency-in-programming
* https://en.wikipedia.org/wiki/Concurrent_computing

### Definitions :

* Concurrency
* Goroutines
  * A goroutine is a function that is capable of running concurrently with other functions.
  * To create a goroutine we use the keyword `go` followed by a function invocation
  * A goroutine is a lightweight thread managed by the Go runtime
* Channels
  * Channels provide a way for two goroutines to communicate with one another and synchronize their execution.
  * This allows goroutines to synchronize without explicit locks or condition variables.
  * A channel type is represented with the keyword `chan` followed by the type of the things that are passed on the channel.
  * The `<-` (left arrow) operator is used to send and receive messages on the channel.
  * Channel Direction: We can specify a direction on a channel type thus restricting it to either sending or receiving. A channel that doesn't have these restrictions is known as bi-directional. A bi-directional channel can be passed to a function that takes send-only or receive-only channels.
  * Go has a special statement called `select` which works like a switch but for channels. `select` picks the first channel that is ready and receives from it (or sends to it). If more than one of the channels are ready then it randomly picks which one to receive from. If none of the channels are ready, the statement blocks until one becomes available.
  * `time.After` creates a channel and after the given duration will send the current time on it.
* Buffered channels
  * It's also possible to pass a second parameter to the make function when creating a channel. This creates a buffered channel with a capacity of n.
  * Normally channels are synchronous; both sides of the channel will wait until the other side is ready. A buffered channel is asynchronous; sending or receiving a message will not wait unless the channel is already full.
  * Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
* Synchronization
* Mutexes
  * If we just want to make sure only one goroutine can access a variable at a time to avoid conflicts.
  * This concept is called mutual exclusion, and the conventional name for the data structure that provides it is mutex.
  *
* WaitGroup
* Pipelines
* Context

- Deadlocks in Go (example)
- Race Conditions in Go (example)
- Data races cannot occur, by design
