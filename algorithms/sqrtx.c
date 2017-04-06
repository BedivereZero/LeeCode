int mySqrt(int x) {
	int a = 1;
	int b = 1;
	if (x == 1)
		return 1;
	while (a != b) {
		a = b;
		b = a / 2 + x / a / 2;
	}
	return a;
}
