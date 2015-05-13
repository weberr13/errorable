# errorable
A Maybe Error structure for encapsulating errors on basic types

Borrowing from functional paradigms, encapsulate a maybe error idea into basic types

currently implements int and slice of int

For instance if you have a slice of strings that you want to convert to 
a slice of integers, but you want to account for possible conversion issues:

s := []string{"1","2","three"}

ei := NewIntErrs(strconv.Atoi,s)

will return ei where

ei.GetFirstErr() will return the error converting "three"
and 

ei.Get(0) == 1
ei.Get(1) == 2
ei.Get(2) == 0 (or whatever the default for Atoi is, if you care)
