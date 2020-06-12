Counting Cycles in a Digraph
=======

### Task Given
Design code that solves the following problem.

Given input directed graph G, your code should do the following:
1. Output all the vertices that have the same in and out degrees.
2. Output the average in-degree and the average out-degree of vertices.
3. Test if G has a cycle.
4. If G has a cycle you code should output at least one cycle.
5. Test if G has at most 3 cycles.
6. If the G has no cycles, then your code should output topological ordering of G.

The expected output of your code must be saved in another .txt file. The contents in the file
are listed as below:
To the output:
1. The first row represents the vertices that have the same in-degree and out-degree.
2. The second row represents the average in-degree and out-degree of the vertices.
3. The third row is ”Order:” if the digraph can be topological ordered. If not, it should be
”Cycle(s):”.
4. The fourth row represents the sorted topological order or a cycle. 
5. The last row represents if the digraph has at most three cycles.


### How to build

Copy folder into go path
```
copy /hw4/* to /home/username/go/src/hw4
```

Compile to binary

```
go build
```

Run without creating binary

```
go run hw4 input/test1 output/test2
```

You can also use prebuild hw4
```
hw4-linux           //linux
hw4-macos           //macos
hw4-windows.exe     //windows
```
Note - all binary are built for amd64

### How to use

##### General - use arguments

```
hw4     arg0        arg1      
hw4     input       output    // input.txt and output.txt doesn't require .txt
hw4.exe input.txt   output.txt
```

arg0 - input file name
arg1 - output file name

if filename does not include .txt extension
txt would be added to file name automatically

The following would be printed if the operation was successful
```
Success: input arg0 has been output to arg1
```

##### Single File - hard coded

Change constant variables into input.txt and output.txt
```go
var (
    inputFile  = "input.txt"    //input.txt file name
    outputFile = "output.txt"   //output.txt file name
)
```

##### Multiple - change main.go
Use multiple input and output files by changing internal main.go file
```go
// Uncomment the following lines to run multiple files with input
createTxt("input/test1", "output/test1")
createTxt("input/test2", "output/test2")
os.Exit(1)
```
Run hw4 without arguments/parameters
```
go run hw4
```

This creates output for 
- input/test1.txt to output/test1.txt
- input/test1.txt to output/test1.txt

##### Error Messages
Missing arguments and parameters when using General
```
Missing parameters
 e.g. hw4.exe test1.txt output1.txt
      hw4.exe test1 output1
      hw4.exe input/test1 output/test1
exit status 1
```

Cannot read file from input path
```
Reading file name: my/file/path/input.txt error
```

Cannot write file in output path
```
Writing to file name: my/file/path/output.txt error
```
usually caused by non-existent path or folder