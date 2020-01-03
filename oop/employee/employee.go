package employee

import (
    "fmt"
)
type employee struct {
    firstName,lastName string
    totalLeaves,leavesTaken int
}

func New(firstName string,lastName string,totalLeaves int,leavesTaken int)employee{
    e := employee{firstName,lastName,totalLeaves,leavesTaken}
    return e
}

func (e employee) LeavesRemaining(){
    fmt.Printf("%s %s has %d leaves remaining",e.firstName,e.lastName,e.totalLeaves-e.leavesTaken)
}