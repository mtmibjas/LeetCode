class Solution {
public:
    int reverseDegree(string s) {
        int r = 0;
        for (size_t i = 0; i < s.size(); ++i) {
            char c = s[i];
            r += (i + 1) * (123 - static_cast<int>(c));
        }
        return r;
    }
};