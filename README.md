# golang-blockchain
Building a blockchain in Go from scratch walkthrough part 1

1. go mod init github.com/monkrus/golang-blockchain.git 

2. Create a Block struct (Hash, Data, PrevHash)

3. Write a method to create a hash based on the previous hash.

3. We join the 2-dimensional array into one

4. Create an actual hash on info

5. Take created hash and push it into the hash field of the block
(where [:] is a slice referencing to the whole array)

6. NOTE! SHA256 is really simple comparing to the real way to calculate a hash in the blockchain 

7. Write a function to create an actual block

8. Create a blockchain type

9. Write a method to add block to the chain
- create a previous block
- create a new block
- append new block

10. Write a function to create a genesis block

11. Write a function to initialize blockchain 
- create an array of blocks using genesis block

12. Put functions into the main function
- intialization, adding blocks etc.

13. Run a for loop to see the blockchain

13. RUN go run main.go

