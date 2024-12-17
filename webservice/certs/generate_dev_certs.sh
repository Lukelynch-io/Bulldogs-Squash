#!/bin/bash

# Output file names
CERT_FILE="dev-cert.pem"
KEY_FILE="dev-key.pem"

# Certificate details
SUBJECT="/C=US/ST=California/L=San Francisco/O=DevCompany/OU=DevDepartment/CN=localhost"

# Generate the private key
openssl genpkey -algorithm RSA -out "$KEY_FILE" -pkeyopt rsa_keygen_bits:2048

TWENTY_TWO_YEARS=8250

# Generate the self-signed certificate
openssl req -new -x509 -key "$KEY_FILE" -out "$CERT_FILE" -days "$TWENTY_TWO_YEARS" -subj "$SUBJECT"

# Output the result
echo "Generated SSL certificate ($CERT_FILE) and key ($KEY_FILE) valid indefinitely."
