int climbStairs(int n) {
	int i = 0;
	int s = 0;
	int *p = (int *) malloc(n * sizeof(int));
	p[0] = 1;
	p[1] = 2;
	for (i = 2; i < n; i++)
		p[i] = p[i - 1] + p[i - 2];
	s = p[n - 1];
	free(p);
	return s;
}
