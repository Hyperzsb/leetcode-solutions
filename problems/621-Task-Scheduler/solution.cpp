class Solution {
private:
    struct Task {
        char id;
        int count;
        int previous;
    };
    
    struct CompareTask {
        bool operator()(Task const &a, Task const &b) {
            return a.count < b.count;
        }
    };
    
public:
    int leastInterval(vector<char>& tasks, int n) {
        priority_queue<Task, vector<Task>, CompareTask> available;
        queue<Task> cooldown;
        
        int taskCount[26];
        memset(taskCount, 0, 26 * sizeof(int));
        
        int size = tasks.size();
        for(int i = 0; i < size; i++) {
            taskCount[tasks[i] - 'A']++;
        }
        
        for(int i = 0; i < 26; i++) {
            if(taskCount[i] > 0 ) {
                struct Task newTask = {(char)('A' + i), taskCount[i], 0 - n - 1};
                available.push(newTask);
            }
        }
        
        int clock = 0;
        struct Task current;
        while(!available.empty() || !cooldown.empty()) {
            while(!cooldown.empty()) {
                current = cooldown.front();
                if(current.previous + n < clock) {
                    cooldown.pop();
                    available.push(current);
                } else {
                    break;
                }
            }
            
            if(!available.empty()) {
                current = available.top();
                available.pop();
                current.count--;
                if(current.count > 0) {
                    current.previous = clock;
                    cooldown.push(current);
                }
            }
            
            clock++;
        }
        
        return clock;
    }
};

