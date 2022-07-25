# Avoid Using Value Receiver Type in Golang Structs Methods

## About this article
The main purpose of this article is as reminder for myself about the use of value/pointer receiver type in struct methods (Go programming language).
I decided to write about this, even though this is a very simple thing, after I spend more then 1 hour to debug an issue caused by the use of value receiver type. in struct methods.
In an example below I will show the issue caused by the use of valuke receiver but without showing any warnings.
In Golang we define methods for a struct similar to methods in object-oriented programming language.
The difference with OO language is in Golang there are 2 types of of receiver : `value` and `pointer` receiver types.

Pointer receiver type will copy the pointer of the 'object/struct' in memory. This is the same with methods in OO programming language where any changes made in the methods will change the actual 'object/struct' in memory. 
Example of pointer receiver type is as below :


Value receiver type will copy the 'object/struct' data. Any changes made to the object/struct' in methods will not change the actual 'object/struct' in memory.
Example of value receiver type is below : 
<code example>

In Golang we only use the value receiver if the method do not need to make any changes to the struct data.
However that is not true for all cases. In this article I will show an example where we need to use pointer receiver in a 'read-only'method.

## Value Pointer Type

### Example of `ineffective assignment` warning

### Example in `goroutine` implementation

## Pointer Receiver Type
The use of pointer receiver type in struct methods is mandatory for the sake of data integrity. Of course there are smaller number of cases where we need to use value receiver type.
Modification of earlier example, with pointer receiver type, can be seen below.

## Conclusion
