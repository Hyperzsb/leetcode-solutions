class Solution {
private:
    struct Project {
        int capital;
        int profit;
    };

    struct CompareCapital {
        bool operator()(Project const &a, Project const &b) {
            if(a.capital > b.capital) {
                return true;
            } else { 
                return false;
            }
        }
    };

    struct CompareProfit {
        bool operator()(Project const &a, Project const &b) {
            if(a.profit < b.profit) {
                return true;
            } else {
                return false;
            }
        }
    };

public:
    int findMaximizedCapital(int k, int w, vector<int>& profits, vector<int>& capital) {
        priority_queue<Project, vector<Project>, CompareCapital> remaining;
        priority_queue<Project, vector<Project>, CompareProfit> valid;

        for(int i = 0; i < profits.size(); i++) {
            remaining.push(Project{capital[i], profits[i]});
        }

        for(int i = 0; i < k; i++) {
            while(!remaining.empty() && remaining.top().capital <= w) {
                valid.push(remaining.top());
                remaining.pop();
            }

            if(!valid.empty()) {
                w += valid.top().profit;
                valid.pop();
            } else {
                break;
            }
        }

        return w;
    }
};

