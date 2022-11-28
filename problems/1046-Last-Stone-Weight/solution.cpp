class Solution {
public:
    int lastStoneWeight(vector<int>& stones) {
        priority_queue<int> stones_queue;
        
        int len = stones.size();
        for(int i = 0; i < len; i++) {
            stones_queue.push(stones[i]);
        }
        
        int stoneA = 0, stoneB = 0;
        while(stones_queue.size() > 1) {
            stoneA = stones_queue.top();
            stones_queue.pop();
            stoneB = stones_queue.top();
            stones_queue.pop();
            
            if(stoneA > stoneB) {
                stones_queue.push(stoneA - stoneB);
            }
        }
        
        if(stones_queue.size() == 0) {
            return 0;
        } else {
            return stones_queue.top();
        }
    }
};

