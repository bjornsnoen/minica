from ctypes import cdll, c_char_p

print(
    cdll.LoadLibrary("./minica").generateCertificate(
        c_char_p("test.com".encode("utf-8"))
    )
)
