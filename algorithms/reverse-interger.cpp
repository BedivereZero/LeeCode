class Solution {
	public:
		int reverse(int x) {
			int y = 0;
			while(x) {
				if(y > 214748364) {
					return 0;
				}
				if(y == 214748364 && x % 10 > 7) {
					return 0;
				}
				if(y < -214748364) {
					return 0;
				}
				if(y == -214748364 && x % 10 < -8) {
					return 0;
				}
				y = y * 10 + x % 10;
				x = x / 10;
			}
			return y;
		}
};
