class Solution {
public:
    int minStoneSum(vector<int>& piles, int k) {
        priority_queue<int> q;
        for(int i = 0; i < piles.size(); i++) {
            q.push(piles[i]);
        }

        for(int i = 0; i < k; i++) {
            q.push(q.top() - q.top() / 2);
            q.pop();
        }

        int result = 0;
        while(!q.empty()) {
            result += q.top();
            q.pop();
        }

        return result;
    }
};

