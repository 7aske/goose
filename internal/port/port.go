package port

// thanks to https://gist.github.com/montanaflynn/b59c058ce2adc18f31d6
import (
	"net"
	"strconv"
)

func New() (port int, err error) {

	server, err := net.Listen("tcp", ":0")

	if err != nil {
		return 0, err
	}

	defer server.Close()

	hostString := server.Addr().String()

	_, portString, err := net.SplitHostPort(hostString)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(portString)
}

func Check(port int) (status bool, err error) {

	host := ":" + strconv.Itoa(port)

	server, err := net.Listen("tcp", host)

	if err != nil {
		return false, err
	}

	_ = server.Close()

	return true, nil

}
