Simple playground for experimenting with Go as a base for scripting some file management tasks.

hash.go
===
Standalone process that accepts a path (currently a single file) and generates a SHA-1 hash as output.

Step #2: Call the hashing process iteratively on an entire directory, computing and storing the hashes we generate.

Notes:
* Can the hashing process scale so that multiple hashes can be computed concurrently? (and does the overhead of doing so outweigh any performance gains?)
* The hash output should be configurable such that the output can be written to a flat file, to a database, etc.

-Owen <https://github.com/owen-datura>