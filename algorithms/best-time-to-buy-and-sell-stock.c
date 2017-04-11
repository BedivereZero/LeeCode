int maxProfit(int* prices, int pricesSize) {
	int i = 0;
	int j = 0;
	int max = 0;
	for (i=0; i < pricesSize - 1; i++) {
		for (j=i + 1; j < pricesSize; j++) {
			if (max < prices[j] - prices[i])
				max = prices[j] - prices[i];
		}
	}
	return max;
}
