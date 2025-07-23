.PHONY: generate_key

generate_key:
	openssl genpkey -algorithm RSA -out private-key.pem -pkeyopt rsa_keygen_bits:2048
	openssl rsa -pubout -in private-key.pem -out public-key.pem

