class Solution {
    public void traverse(int idx, int[][] isConnected, boolean[] isVisited) {
        isVisited[idx] = true;

        for (int i = 0; i < isConnected[idx].length; i++) {
            if (isConnected[idx][i] == 1 && !isVisited[i]) {
                traverse(i, isConnected, isVisited);
            }
        }
    }

    public int findCircleNum(int[][] isConnected) {
        int n = isConnected.length;
        boolean[] isVisited = new boolean[n];

        int result = 0;
        for (int i = 0; i < n; i++) {
            if (!isVisited[i]) {
                traverse(i, isConnected, isVisited);
                result++;
            }
        }

        return result;
    }
}

