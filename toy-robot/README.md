# Introduction

The below is an adaptation of the [original](./ToyRobotExam.md) that I decided to do to brush up on some Go skills. 

No external dependencies allowed, just the Go Standard Library.

It is the same except that 
* the IO layer should be via HTTP and TCP, and instead send JSON payloads to represent commands
* multiple tables should be supported, a HTTP GET or a query should return a list of them, including the position of the robot on each one

The JSON command structure will resemble:

```
{
    command: "PLACE X,Y,F | MOVE | LEFT | RIGHT | REPORT"

}
```

A response would resemble:

```
HTTP Status 200

Board Id.
```

For HTTP any move that would cause the robot to fall must return an appropriate 400 response code.



## Notes

The Clean Architecture.

The board lets you do things. The board holds the robot. The robot knows how to move and change where it faces. The board can call the robot to tell it to move but won't let it fall off the edge. 

The board also has functionality to support commands calling it to effect behaviours for PLACE, MOVE, LEFT, RIGHT and REPORT.

