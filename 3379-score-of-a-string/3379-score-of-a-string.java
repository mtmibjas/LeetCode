class Solution {
    public int scoreOfString(String s) {
        int sum = 0;
        
        for (int i = 1; i < s.length(); i++) {
            sum += Abs(s.charAt(i-1), s.charAt(i)); 
        }

        return sum;
    }

    int Abs(int a, int b){

        if (a > b) {
            return a-b;
        }

        return b-a;
    }
}