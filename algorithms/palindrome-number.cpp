class Solution {
	public:
		bool isPalindrome(int x) {
			int a = x;
			int b = 0;
			while(a) {
				b = b * 10 + x % 10;
				a = a / 10;
			}
			return x == b;
		}
};
