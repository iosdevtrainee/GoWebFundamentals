package main 
import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
  max := new(big.Int)Lsh(big.Int(1), 128)
  // Unique random number usually set by the CA for the certificate
  serialNumber, _ := rand.Int(rand.Reader, max)
  // Set up the receiver of the certificate i.e. description of company / entity
  subject := pkix.Name{
     Organization:       []string{"Manning Publications Co."},
     OrganizationalUnit: []string{"Books"},
     CommonName:         "Go Web Programming",
  }
  
 
  template := x509.Certificate{
     SerialNumber: serialNumber, 
     Subject:      subject, 
     // Certificates also need a expiration date and creation date i.e. NotBefore
     NotBefore:    time.Now(),
     NotAfter:     time.Now().Add(365 * 25 * time.Hour),
     // Certificate key usage and extkeyusage describe the environment the key will live in
     KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
     ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, 
     // IP Address describe the IP Addresses the key is valid for
     IPAddress:    []net.IP{net.ParseIP("127.0.0.1")},
  }
  
  // Generate a random private key with 2048 units of randomness
  pk, _ := rsa.GenerateKey(rand.Reader, 2048)
  
  // use the public and private keys to sign the certificate and generate a byte string
  derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, 
  &pk.PublicKey, pk)

  certOut, _ := os.Create("cert.pem")
  // write the bytes string of the certificate out to a cert.pem file
  pem.Encode(certOut, &pem.Block{Type:"CERTIFICATE", Bytes: derBytes})
  certOut.close()

  keyOut, _ := os.Create("key.pem")
  // write the key of the certificate out to a key.pem file
  // generally if the certificate is sign by a CA it will be concatenation of the 
  // sever's cert and the CA's cert
  pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)}) 
  keyOut.close()
