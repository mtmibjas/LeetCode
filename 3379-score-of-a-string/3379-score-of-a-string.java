class Solution {
    

    public static int scoreOfString(String s) {
        int score = 0;

        for (int i = 0; i < s.length() - 1; i++) {
            char c1 = s.charAt(i);
            char c2 = s.charAt(i+1);

            int difference = Math.abs(c1 - c2);
            score = score + difference;
        }

        return score;
    }
}