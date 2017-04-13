char* normalize(char* s) {
	int i = 0;
	int l = 0;
	int m = 1;
	int ch[256];
	char* t = NULL;
	for (l=0; s[l]; l++);
	t = (char *) malloc(l * sizeof(char));
	for (i=0; i < 256; i++)
		ch[i] = 0;
	for (i=0; i < l; i++) {
		if (ch[s[i]] == 0) {
			ch[s[i]] = m;
			m++;
		}
		t[i] = ch[s[i]];
	}
	return t;
}

bool isIsomorphic(char* s, char* t) {
	return !strcmp(normalize(s), normalize(t));
}
