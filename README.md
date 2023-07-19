# bcrypt_hashing_algorithm
A basic Go program to generate a bcrypt hash value from  random words in the cli






















	The output of a bcrypt hash is a byte slice , which typically contains characters that are not human-readable.
	As a result, the bcrypt hash is usually encoded in base64 format to make it 
	a printable string.  The base64 encoding scheme makes it easier to store and compare bcrypt 
	hashes.
	The length of a bcrypt hash is usually 60 characters when encoded in base64.  
	
