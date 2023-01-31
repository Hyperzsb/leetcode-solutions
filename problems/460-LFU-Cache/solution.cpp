class LFUCache {
    int capacity;
    int lfu;
    unordered_map<int,pair<int,int>> keyVal;
    unordered_map<int,list<int>> freqList;
    unordered_map<int,list<int>::iterator> pos;

public:
    LFUCache(int capacity) {
        this->capacity = capacity;
        lfu = 0;
    }
    
    int get(int key) {
        if(keyVal.find(key) == keyVal.end()) {
            return -1;
        }

        freqList[keyVal[key].second].erase(pos[key]);
        keyVal[key].second++;
        freqList[keyVal[key].second].push_back(key);
        pos[key] = --freqList[keyVal[key].second].end();

        if(freqList[lfu].empty()) {
            lfu++;
        }

        return keyVal[key].first;
    }
    
    void put(int key, int value) {
        if(capacity == 0) {
            return ;
        }

        if(keyVal.find(key) != keyVal.end()) {
            freqList[keyVal[key].second].erase(pos[key]);
            keyVal[key].second++;
            freqList[keyVal[key].second].push_back(key);
            pos[key] = --freqList[keyVal[key].second].end();

            if(freqList[lfu].empty()) {
                lfu++;
            }

            keyVal[key].first = value;

            return ;
        }

        if(keyVal.size() == capacity) {
            int delKey = freqList[lfu].front();
            keyVal.erase(delKey);
            pos.erase(delKey);
            freqList[lfu].pop_front();
        }

        keyVal[key] = {value, 1};
        freqList[1].push_back(key);
        pos[key] = --freqList[1].end();
        lfu = 1;
    }
};

/**
 * Your LFUCache object will be instantiated and called as such:
 * LFUCache* obj = new LFUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */

