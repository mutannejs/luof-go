package link

import (
    "time"
)

type Reader interface (
    // Get func (string) (Link, error)
    GetAll func () ([]Link, error)
)

type Writer interface (
    // Create func (Link) error
    // Update func (string, Link) error
    // Delete func (string) error
)

type Repository interface (
    Reader
    Writer
)
