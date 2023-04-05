package node

import "io"

type DataModel interface {
	io.ReadWriter
}
