class Solution {
	public:
		int strStr(string haystack, string needle) {
			if (needle.size()) {
				int p = haystack.find(needle);
				if (p == haystack.npos)
					return -1;
				else
					return p;
			}
			else
				return 0;
		}
};
