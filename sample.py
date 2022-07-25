from ctypes import POINTER, c_char_p, c_int, cdll

minica = cdll.LoadLibrary("./minica.dll")

# Simple single-domain certificate in python

generate_certificate = minica.generateCertificate
generate_certificate.argtypes = [
    c_char_p
]

print(generate_certificate(c_char_p("test.com".encode())))

# More complex certificate with list of domains and IPs
def to_c_str_array(strs: list[str]):
    ptr = (c_char_p * (len(strs)))()
    ptr[:] = [s.encode() for s in strs]
    return ptr

_generate_complex_certificate = minica.generateComplexCertificate
_generate_complex_certificate.argtypes = [
    POINTER(c_char_p),
    c_int,
    POINTER(c_char_p),
    c_int
]

def generate_complex_certificate(domains: list[str], ip_addresses: list[str]) -> int:
    return _generate_complex_certificate(to_c_str_array(domains), len(domains), to_c_str_array(ip_addresses), len(ip_addresses))

domains = ["*.one.com", "one.com"]
ip_addresses = ["10.0.0.1", "10.0.0.2", "10.0.0.3"]

print(generate_complex_certificate(domains, ip_addresses))
