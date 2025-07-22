.PHONY: generate_key

generate_key:
	openssl genpkey -algorithm RSA -out private-key.pem -pkeyopt rsa_keygen_bits:2048
	openssl rsa -pubout -in private-key.pem -out public-key.pem
#	ssh-keygen -t rsa-sha2-256 -b 4096 -C "your_email@example.com" -f local-key

