//go:build !linux
// +build !linux

package netns

// NsHandle is a handle to a network namespace. It can be cast directly
// to an int and used as a file descriptor.
type NsHandle int

// Equal determines if two network handles refer to the same network
// namespace. This is done by comparing the device and inode that the
// file descriptors point to.
func (ns NsHandle) Equal(other NsHandle) bool {
	return ns == other
}

// String shows the file descriptor number and its dev and inode.
func (ns NsHandle) String() string {
	return "NS(none)"
}

// UniqueId returns a string which uniquely identifies the namespace
// associated with the network handle.
func (ns NsHandle) UniqueId() string {
	return "NS(none)"
}

// IsOpen returns true if Close() has not been called.
func (ns NsHandle) IsOpen() bool {
	return false
}

// Close closes the NsHandle and resets its file descriptor to -1.
// It is not safe to use an NsHandle after Close() is called.
func (ns *NsHandle) Close() error {
	return nil
}

// None gets an empty (closed) NsHandle.
func None() NsHandle {
	return NsHandle(-1)
}
