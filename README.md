# Bcrypt_hashing_algorithm
A basic Go program to generate a bcrypt hash value from  random words in the cli



**Bcrypt Features** 

-Key Derivation Function (KDF): bcrypt is primarily designed as a password hashing function and is intended for securely hashing passwords. It incorporates a Key Derivation Function (KDF) that is specifically designed to be slow and resource-intensive, making it computationally expensive to brute-force attack passwords.

-Adaptive Cost Factor: bcrypt allows for an adjustable cost factor, which determines the computational complexity of the hashing process. This cost factor can be increased over time to keep up with advances in hardware and computing power, making it more resistant to brute-force attacks as hardware improves.

-Salting: bcrypt automatically generates and manages a random salt for each hashed password. Salting helps prevent attackers from using precomputed tables (rainbow tables) to reverse-engineer passwords.

-Built-in Salt Management: bcrypt stores the salt value alongside the resulting hash, so there is no need to manage or store the salt separately.






	The output of a bcrypt hash is a byte slice , which typically contains characters that are not human-readable.
	As a result, the bcrypt hash is usually encoded in base64 format to make it 
	a printable string.  The base64 encoding scheme makes it easier to store and compare bcrypt 
	hashes.
	The length of a bcrypt hash is usually 60 characters when encoded in base64.  

 
For example, a bcrypt hash might look something like this (note that the actual value will vary each time the code is run):


`$2a$10$eoa9l1DqC8Y3v9.h1OJRTeSJoZ4m8Cw2VhsMmC34oNlyin9G47DSa`

Breaking down the components:

$2a$: This prefix indicates the bcrypt algorithm version used (2a in this case, which is the current version).

10$: This indicates the cost factor used for key derivation.

eoa9l1DqC8Y3v9.h1OJRTe: This part represents the salt used for hashing.

SJoZ4m8Cw2VhsMmC34oNlyin9G47DSa: The actual hash value.


**In bcrypt, the generated hash for the same input (e.g., the same name) will not be the same each time due to the inclusion of a random salt during the hashing process. The use of a salt is a critical security feature in password hashing, as it helps prevent the use of precomputed tables (rainbow tables) for reversing the hash back to the original password.**


 
