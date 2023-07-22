# Prefixed ID Generator
Basing this distributed ID generator off of [MongoDB's ObjectId](https://www.mongodb.com/docs/manual/reference/method/ObjectId/) 

The ObjectId consists of the following:

- A timestamp measured in seconds - 4 bytes

- A random value generated from a cryptographically random number generator. The value is unique to the machine - 5 bytes

- An incrementing counter. Initialized to a random value - 3 bytes

Those elements will combine to generate a 12-byte unique ID from which a prefix can be added e.g.
