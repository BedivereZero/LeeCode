int climbStairs(int n) {
	int a = 1;
	int b = 1;
	if (n < 2)
		return 1;
	else
		n = n - 1;
	while (n--) {
		a = a + b;
		b = a - b;
	}
	return a;
}
