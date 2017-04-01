class Solution {
	public:
		bool isValid(string s) {
			std::stack<char> st;
			for(int i = 0; i < s.size(); i++) {
				switch(s[i]) {
					case '(':
					case '[':
					case '{':
						st.push(s[i]);
						break;
					case ')':
						char top = '(';
					case ']':
						char top = '[';
					case '}':
						char top = ']';
						if(st.top != top)
							return false;
				}
			return st.empty();
		}
};
