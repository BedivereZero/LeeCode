bool isHappy(int n) {
	int s = 0;
	if (n < 10)
		return n == 1;
	while(n) {
		s += (n % 10) * (n % 10);
		n = n / 10;
	}
	return isHappy(s);
}
