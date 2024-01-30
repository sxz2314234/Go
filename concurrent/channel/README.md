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