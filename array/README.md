# Tips
## Debug
At the beginning, I use 'git run ./main.go' but I encounter one problem.So I want to use gdb to debug it.Then, I enter the next command:
> go build -gcflag 'N L' -o gdb_sandbox ./main.go

As you can see, the gdb_sandbox file I generate is to go to debug.
> gdb gdb_sandbox

## Array and Slice
As a beginner, I am aware of the difference between array and silce just now. Slice is dynamic array.When we will use the array we should:
> var arr1 = [number]type  // declare it, use a certain number

or
> var arr2 = [3]int{1,2,3}  // declare and initialize

When slice in turn, we can not specific a certain number:
> var slice1 = []int{1,2,3}

or
> var slice2 = make([]type,len,cap)  // the cap can be default

What superise me most is that the name of slice is one pointer but the name of array just is one value. So when passing an array to function, we just get
a duplication of that array.
