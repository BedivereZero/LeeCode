char* normalize(char* s) {
	int i = 0;
	int l = 0;
	int m = 0;
	int ch[26];
	char* t = NULL;
	for (l=0; s[l]; l++);
	t = (char *) malloc(l * sizeof(char));
	for (i=0; i < 26; i++)
		ch[i] = 0;
	for (i=0; i < l; i++) {
		if (ch[s[i] - 'a'] == 0) {
			ch[s[i] - 'a'] = m;
			m++;
		}
		t[i] = 'a' + ch[s[i] - 'a'];
	}
	return t;
}

bool isIsomorphic(char* s, char* t) {
	return !strcmp(normalize(s), normalize(t));
}
