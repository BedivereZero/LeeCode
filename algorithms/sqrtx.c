int mySqrt(int x) {
	int a = 0;
	int b = x;
	int c = 0;
	while (a < b - 1) {
		c = (a + b) / 2;
		if (c * c < x)
			a = c;
		else
			b = c;
	}
	if (x < b * b)
		return a;
	else
		return b;
}
