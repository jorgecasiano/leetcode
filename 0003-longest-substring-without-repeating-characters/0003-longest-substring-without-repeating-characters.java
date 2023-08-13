class Solution {
    public int lengthOfLongestSubstring(String s) {
        Map<Character, Integer> lastIndexChar = new HashMap<>();

        int longestSub = 0;
        int currentSub = 0;
        int start;

        final var arr = s.toCharArray();
        for (int i = 0; i < arr.length; i++) {
          var lastIndex = lastIndexChar.get(arr[i]);
          if (lastIndex == null || lastIndex < i - currentSub) {
            currentSub++;
          } else {
            if (currentSub > longestSub) {
              longestSub = currentSub;
            }

            start = lastIndex + 1;
            currentSub = i - start + 1;
          }
          lastIndexChar.put(arr[i], i);
        }

        if (currentSub > longestSub) {
          longestSub = currentSub;
        }

        return longestSub;
    }
}