package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/KarimovKamil/otus-go-final-project/internal/entity"
)

var ErrNetworkAlreadyExists = fmt.Errorf("network already exists in the list")

// GetNetworkPrefixBinary generates the binary network prefix based on the given IP address and subnet mask.
func GetNetworkPrefixBinary(ip, mask string) string {
	binaryIP := IPAddressToBinary(ip)
	maskInt, _ := strconv.Atoi(mask)
	return binaryIP[:maskInt]
}

// IPAddressToBinary converts an IP address from decimal format to binary format.
func IPAddressToBinary(ip string) string {
	ipParts := strings.Split(ip, ".")
	stringBuilder := strings.Builder{}
	for _, part := range ipParts {
		val, _ := strconv.Atoi(part)
		stringBuilder.WriteString(fmt.Sprintf("%08b", val))
	}
	return stringBuilder.String()
}

// GetNetwork retrieves a network entity based on the input string.
func GetNetwork(network string) entity.Network {
	parts := strings.Split(network, "/")
	return entity.Network{IP: parts[0], Mask: parts[1], BinaryPrefix: GetNetworkPrefixBinary(parts[0], parts[1])}
}
