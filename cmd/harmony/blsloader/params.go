package blsloader

import "time"

const (
	passExt     = ".pass"
	basicKeyExt = ".key"
	kmsKeyExt   = ".bls"
)

const (
	defPromptTimeout = 1 * time.Second
)

const (
	defWritePassDirMode  = 0600
	defWritePassFileMode = 0600
)
