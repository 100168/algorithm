package main

/*
给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。

vector<int> prime;
bool is_prime[N];

	void Eratosthenes(int n) {
	  is_prime[0] = is_prime[1] = false;
	  for (int i = 2; i <= n; ++i) is_prime[i] = true;
	  for (int i = 2; i <= n; ++i) {
	    if (is_prime[i]) {
	      prime.push_back(i);
	      if ((long long)i * i > n) continue;
	      for (int j = i * i; j <= n; j += i)
	        // 因为从 2 到 i - 1 的倍数我们之前筛过了，这里直接从 i
	        // 的倍数开始，提高了运行速度
	        is_prime[j] = false;  // 是 i 的倍数的均不是素数
	    }
	  }
	}
*/
func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}
	ans := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			ans++
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return ans
}
