package util

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"os"
	"reflect"
	"syscall"
	"unsafe"
)

const SO_REUSEPORT = 15
const backlog = 512

// addr should be in format '0.0.0.0:1235'
func NewTCPListener(address string, reusePort bool) (net.Listener, error) {
	if !reusePort {
		return net.Listen("tcp", address)
	} else {
		addr, e := net.ResolveTCPAddr("tcp", address)
		if e != nil {
			return nil, e
		}

		var fd int

		fd, e = syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
		if e != nil {
			return nil, e
		}

		e = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, SO_REUSEPORT, 1)
		if e != nil {
			return nil, e
		}

		addr4 := syscall.SockaddrInet4{
			Port: addr.Port,
			Addr: [4]byte{addr.IP[0], addr.IP[1], addr.IP[2], addr.IP[3]},
		}

		e = syscall.Bind(fd, &addr4)
		if e != nil {
			return nil, e
		}

		e = syscall.Listen(fd, backlog)

		if e != nil {
			return nil, e
		}

		f := os.NewFile(uintptr(fd), "")
		if f == nil {
			return nil, errors.New("fd is invalid")
		}

		defer f.Close()

		return net.FileListener(f)
	}
}

const testAddr = "www.baidu.com:443"

func init() {
	dummyTLS, err := tls.Dial("tcp", "www.baidu.com:443", nil)
	defer dummyTLS.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Test for getFd fail, connect to testAddr error: %v\n", err)
		panic("Test for getFd fail")
	}

	dtlsv := reflect.ValueOf(dummyTLS)
	if !dtlsv.IsValid() {
		fmt.Fprintf(os.Stderr, "Test for getFd fail, reflect error\n")
		panic("Test for getFd fail")
	}

	dtcpv := reflect.Indirect(dtlsv).FieldByName("conn")
	if !dtcpv.IsValid() {
		fmt.Fprintf(os.Stderr, "Test for getFd fail, reflect error, no conn field in tls.Conn\n")
		panic("Test for getFd fail")
	}

	dtcpi := dtcpv.InterfaceData()
	dtcpv = reflect.ValueOf((*net.TCPConn)(unsafe.Pointer(dtcpi[1])))
	if !dtcpv.IsValid() {
		fmt.Fprintf(os.Stderr, "Test for getFd fail, get TCPConn error\n")
		panic("Test for getFd fail")
	}

	fdv := reflect.Indirect(dtcpv).FieldByName("fd")
	if !fdv.IsValid() {
		fmt.Fprintf(os.Stderr, "Test for getFd fail, no fd field in *net.TCPConn\n")
		panic("Test for getFd fail")
	}

	fdv = reflect.Indirect(fdv)
	sysfd := fdv.FieldByName("sysfd")

	if sysfd.Kind() == reflect.Int || sysfd.Kind() == reflect.Int64 {
		fdCheck = 1
	} else {
		pfd := fdv.FieldByName("pfd")
		if !pfd.IsValid() {
			fmt.Fprintf(os.Stderr, "Test for getFd fail, no pfd field in *net.TCPConn\n")
			panic("Test for getFd fail")
		}

		sysfd := pfd.FieldByName("Sysfd")
		if sysfd.Kind() == reflect.Int || sysfd.Kind() == reflect.Int64 {
			fdCheck = 2
		} else {
			fmt.Fprintf(os.Stderr, "Test for getFd fail, panic\n")
			panic("Test for getFd fail")
		}
	}
}

func GetConnFd(c net.Conn) int64 {
	if t, ok := c.(*net.TCPConn); ok {
		return getTCPConnFd(t)
	}

	tlsv, _ := c.(*tls.Conn)
	cv := reflect.Indirect(reflect.ValueOf(tlsv)).FieldByName("conn").InterfaceData()
	cc := (*net.TCPConn)(unsafe.Pointer(cv[1]))
	return getTCPConnFd(cc)
}

var fdCheck int32 = 1

func getTCPConnFd(c *net.TCPConn) int64 {
	if fdCheck == 2 {
		fd := reflect.Indirect(reflect.ValueOf(c)).FieldByName("fd")
		return reflect.Indirect(fd).FieldByName("pfd").FieldByName("Sysfd").Int()
	}
	fd := reflect.ValueOf(c).FieldByName("fd")
	return reflect.Indirect(fd).FieldByName("sysfd").Int()
}

func DetectHost() string {
	if h, e := os.Hostname(); e == nil {
		addrs, e := net.LookupHost(h)
		if e == nil && len(addrs) > 0 {
			return addrs[0]
		}
	}

	if eth0, e := net.InterfaceByName("eth0"); e == nil {
		if addrs, e := eth0.Addrs(); e == nil && len(addrs) > 0 {
			return addrs[0].String()
		}
	}

	if eth1, e := net.InterfaceByName("eth1"); e == nil {
		if addrs, e := eth1.Addrs(); e == nil && len(addrs) > 0 {
			return addrs[0].String()
		}
	}

	if eth2, e := net.InterfaceByName("eth2"); e == nil {
		if addrs, e := eth2.Addrs(); e == nil && len(addrs) > 0 {
			return addrs[0].String()
		}
	}
	return ""
}
