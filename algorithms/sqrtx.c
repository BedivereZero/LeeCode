int mySqrt(int x) {
	int n = 0;
	while (n * n < x)
		n++;
	if (x < n * n)
		return n - 1;
	else
		return n;
}
