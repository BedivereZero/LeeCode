/**
 * Return an array of arrays.
 * The sizes of the arrays are returned as *columnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
int** generate(int numRows, int** columnSizes) {
	int** p = columnSizes;
	int i = 0;
	int j = 0;
	p = (int *) malloc(numRows * sizeof(int));
	for (i=0; i < numRows; i++) {
		p[i] = (int) malloc((i + 1) * sizeof(int));
		p[i][0] = 1;
		p[i][i] = 1;
		for (j=1; j < i; j++)
			p[i][j] = p[i - 1][j] + p[i - 1][j - 1];
	}
	return p;
}
