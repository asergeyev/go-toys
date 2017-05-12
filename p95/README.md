Attempt to create faster p95 calculation than just "sort array", get Nth element.
Does not work on larger lists since building top (1-q) elements list is actually hard on CPU.

Short version is faster up to around 1000 elements in the list (leading to top-N below 50).
