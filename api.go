package main

import (
	"C"
	"unsafe"

	"github.com/bjornsnoen/minica/certutils"
)

//export generateCertificate
func generateCertificate(domain *C.char) C.int {
	iss, err := certutils.GetIssuer("minica-key.pem", "minica.pem")
	if err != nil {
		return 1
	}
	_, err = certutils.Sign(iss, []string{C.GoString(domain)}, []string{})
	if err != nil {
		return 2
	}
	return 0
}

//export generateIPCertificate
func generateIPCertificate(ipAddress *C.char) C.int {
	iss, err := certutils.GetIssuer("minica-key.pem", "minica.pem")
	if err != nil {
		return 1
	}
	_, err = certutils.Sign(iss, []string{}, []string{C.GoString(ipAddress)})
	if err != nil {
		return 2
	}
	return 0
}

//export generateComplexCertificate
func generateComplexCertificate(domains **C.char, domainCount C.int, ipAddresses **C.char, ipAddressCount C.int) C.int {
	iss, err := certutils.GetIssuer("minica-key.pem", "minica.pem")
	if err != nil {
		return 1
	}

	domainSlice := unsafe.Slice(domains, domainCount)
	ipSlice := unsafe.Slice(ipAddresses, ipAddressCount)

	domainGoStrings := make([]string, domainCount)
	for i, cDomain := range domainSlice {
		domainGoStrings[i] = C.GoString(cDomain)
	}
	ipGoStrings := make([]string, ipAddressCount)
	for i, cIp := range ipSlice {
		ipGoStrings[i] = C.GoString(cIp)
	}

	_, err = certutils.Sign(iss, domainGoStrings, ipGoStrings)
	if err != nil {
		return 2
	}
	return 0
}
