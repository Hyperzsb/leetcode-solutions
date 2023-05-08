class Solution {
    public int distributeCandies(int[] candyType) {
        HashSet<Integer> candies = new HashSet<>();
        for (int candy : candyType) {
            candies.add(candy);
        }

        if (candies.size() > candyType.length / 2) {
            return candyType.length / 2;
        } else {
            return candies.size();
        }
    }
}

