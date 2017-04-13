bool isHappy(int n) {
	int s = 0;
	int r[] = {
		false,
		true,
		false,
		false,
		false,
		false,
		false,
		true,
		false,
		false
	};
	if (n < 10)
		return r[n];
	while(n) {
		s += (n % 10) * (n % 10);
		n = n / 10;
	}
	return isHappy(s);
}
