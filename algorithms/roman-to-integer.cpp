class Solution {
	public:
		int romanToInt(string s) {
			int result = 0;
			int pre = 0;
			for(int i = 0; i < s.length(); i++) {
				switch s[i] {
					case 'I':
						result += 1;
						pre = 1;
						break;
					case 'V':
						result += 5;
						if(pre) {
							result -= 2 * pre;
							pre = 0;
						}
						break;
					case 'X':
						result += 10;
						if(pre) {
							result -= 2 * pre;
							pre = 0;
						} else {
							pre = 10;
						}
						break;
					case 'L':
						result += 50;
						if(pre) {
							result -= 2 * pre;
							pre = 0;
						}
						break;
					case 'C':
						result += 100;
						if(pre) {
							result -= 2 * pre;
							pre = 0;
						} else {
							pre = 100;
						}
						break;
					case 'D':
						result += 500;
						if(pre) {
							result -= 2 * pre;
							pre = 0;
						}
						break;
					case 'M':
						result += 1000;
						if(pre) {
							result -= 2 * pre;
							pre = 0;
						} else {
							pre = 1000;
						}
						break;
				}
			}
			return result;
		}
};
