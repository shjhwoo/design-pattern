package protoType

type Inode interface {
	print(string)
	clone() Inode
}
