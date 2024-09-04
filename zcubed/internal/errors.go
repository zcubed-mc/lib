package internal

import "fmt"

func PrefixErr(prefix string, err error) error {
	return fmt.Errorf("%s: %w", prefix, err)
}
