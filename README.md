# jsoncompare

jsoncompare allows comparing two json files based on its tree structure.
Two files are considers same if they contain equal and same leaf nodes in same 
tree structure. This comparison ignores the order of the nodes as they appear in
the tree, therefore it is less strict than comparison of json tree via other 
means as reflect.DeepEquals(..)

To build the command line tool, go inside diretory `jsoncompare` and run
`go build`.

```usage: ./jsoncompare  <first file path>  <second file path>```