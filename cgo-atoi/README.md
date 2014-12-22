Quick test if it's worth to do anything about dealing with
strconv.Atoi(string(buffer)) or it's pretty much as fast as it gets


Apparetnly CGo is slower here. Even more, specially optimized function for
parsing 10-base integers of given type is probably most optimal (see an
example with "Direct" function)