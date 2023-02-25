class Solution {
public:
    int minimumDeviation(vector<int>& nums) {
        priority_queue<int> maxQ;

        int minVal = numeric_limits<int>::max(), maxVal = numeric_limits<int>::min();
        for(int i = 0; i < nums.size(); i++) {
            if(nums[i] % 2 != 0) {
                nums[i] *= 2;
            }

            minVal = min(minVal, nums[i]);
            maxQ.push(nums[i]);
        }

        int result = numeric_limits<int>::max();
        while(maxQ.top() % 2 == 0) {
            maxVal = maxQ.top();
            maxQ.pop();

            result = min(result, maxVal - minVal);

            maxVal /= 2;
            minVal = min(minVal, maxVal);

            maxQ.push(maxVal);
        }

        maxVal = maxQ.top();
        result = min(result, maxVal - minVal);
        return result;
    }
};

