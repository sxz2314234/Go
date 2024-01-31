# Channel

##  Declare channel
The command method:
> var varname chan vartype

We also can initialize it, using `make`.For example:
> var eample := make(map[int]string)
>> if no initialize, the default is nil

## The channel with no buffer
The most important we should know is that only when there is a receiver, we can send a value to the channel.Otherwise, it would go to the deadlock.

## The channel with buffer
We should use `make` to allocate memory space for the channel.
> channel:=make(chan int, number of channels)

## Close the channel
If we want to close channel, we should use `close`.
> close(channel)

Then,how do we know whether the channel is closed: 
> value,ok:=<-chan

if closed, the ok is false;if not, the ok is true.

## Select multiplexing
In some case, we may need to receive data from multiple channels at the same time.When receiving data,if there is no data to be received,goroutine will be blocked.So we use the key word `select` to response at the same time.For example:
> select{
    case <-ch1:
    // ...
    case data:=<-ch2:
    // ...
    case ch3<-10:
    // ...
    default:
    // default action
}

## Mutex
### sync.Mutex
The `sync.Mutex` can promise us that there is only one goroutine access shared resources.

### sync.RWMutex
At such one case that we just need to access the resources instead of modifying it,we can use `RWMutex`.