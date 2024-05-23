# Sieve of Eratosthenes

*Sift the Two's and Sift the Three's:*
*The Sieve of Eratosthenes.*
*When the multiples sublime,*
*The numbers that remain are Prime*

In mathematics, the sieve of Eratosthenes is an ancient algorithm for finding prime numbers. Its invention is attributed to Eratosthenes of Cyrene, a 3rd century BCE Greek mathematician.

The algorithm works to calculate primes by iteratively marking as composite (i.e., not prime) the multiples of each prime, starting with the first prime number, When the next identified prime exceeds the square root of the upper limit, all the remaining numbers in the list are prime.

*Euler's sieve* is a version of the sieve of Eratosthenes in which each composite number is eliminated exactly once. It, too, starts with a list of numbers from 2 to n in order.

On each step the first element is identified as the next prime. This prime is multiplied with each element of the list, including itself, and the results are marked in the list for subsequent deletion. The initial element and the marked elements are then removed from the working sequence, and the process is repeated. When the next identified prime exceeds the square root of the upper limit, all the remaining numbers in the list are prime.
