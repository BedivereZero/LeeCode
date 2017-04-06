int lengthOfLastWord(char* s) {
	char* p = s;
	int counter = 0;
	bool new_word = false;
	while (*p != '\0') {
		if (*p == ' ')
			new_world = false;
		else
			if (new_world) {
				new_world = false;
				counter = 1;
			}
			else
				counter++;
		p++;
	}
	return counter;
}
