int cmp(const void *a, const void *b) {
	return *(char *)a - *(char *)b;
}

bool isAnagram(char* s, char* t) {
	qsort(s, strlen(s), sizeof(char), cmp);
	qsort(t, strlen(t), sizeof(char), cmp);
	return strcmp(s, t) == 0;
}
