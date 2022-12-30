class Solution {
private:
    struct Task {
        int id;
        int qTime;
        int pTime;
    };

    struct CompareTaskAvail {
        bool operator()(Task const &a, Task const &b) {
            if(a.pTime > b.pTime) {
                return true;
            } else if (a.pTime == b.pTime) {
                return a.id > b.id;
            } else {
                return false;
            }
        }
    };

    static bool CompareTaskPending(Task const &a, Task const &b) {
        return a.qTime < b.qTime;
    };

public:
    vector<int> getOrder(vector<vector<int>>& tasks) {
        vector<Task> tasksPending;
        for(int i = 0; i < tasks.size(); i++) {
            tasksPending.push_back(Task{i, tasks[i][0], tasks[i][1]});
        }
        sort(tasksPending.begin(), tasksPending.end(), CompareTaskPending);
        priority_queue<Task, vector<Task>, CompareTaskAvail> tasksAvail;

        long long timer = 1;
        int taskIdx = 0;
        vector<int> result;
        while(taskIdx < tasksPending.size() || !tasksAvail.empty()) {
            while(taskIdx < tasksPending.size() && 
            tasksPending[taskIdx].qTime <= timer) {
                tasksAvail.push(tasksPending[taskIdx++]);
            }
            if(!tasksAvail.empty()) {
                // Jump to next idle time if busy
                timer += tasksAvail.top().pTime;
                result.push_back(tasksAvail.top().id);
                tasksAvail.pop();
            } else {
                // Jump to next busy time if idle
                timer = tasksPending[taskIdx].qTime;
            }
        }

        return result;
    }
};

