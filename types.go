package gostatsd

// Nanotime is the number of nanoseconds elapsed since January 1, 1970 UTC.
// Get the value with time.Now().UnixNano().
type Nanotime int64

// IP is a v4/v6 IP address.
// We do not use net.IP because it will involve conversion to string and back several times.
type IP string

// UnknownIP is an IP of an unknown source.
const UnknownIP IP = ""

type Done func()

type Wait func()
