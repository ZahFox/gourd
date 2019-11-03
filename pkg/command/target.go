package command

// Target is what should be affected the action of a command
type Target int

const (
	// NOTARGET means no target
	NOTARGET Target = iota
	// HOST is a target that represents the local gourdd
	HOST
)
