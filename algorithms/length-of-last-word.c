int lengthOfLastWord(char* s) {
	char* p = s;
	int counter = 0;
	while (*p != '\0') {
		if (*p == ' ')
			counter = 0;
		else
			counter++;
		p++;
	}
	return counter;
}
