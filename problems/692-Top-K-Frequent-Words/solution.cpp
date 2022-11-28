class Solution {
private:
    struct Word {
        string val;
        int count;
    };
    static bool cmp(Word a, Word b) {
        if(a.count > b.count) {
            return true;
        } else if(a.count == b.count) {
            if (a.val < b.val) {
                return true;
            } else {
                return false;
            }
        } else {
            return false;
        }
    }
    
public:
    vector<string> topKFrequent(vector<string>& words, int k) {
        map<string, int> words_map;
        vector<Word> words_queue;
        
        int len = words.size(), count = 0;
        for(int i = 0; i < len; i++) {
            if(words_map.find(words[i]) != words_map.end()) {
                words_queue[words_map[words[i]]].count++;
            } else {
                words_map.insert({words[i], count++});
                
                Word new_word;
                new_word.val = words[i];
                new_word.count = 1;
                
                words_queue.push_back(new_word);
            }
        }
        
        sort(words_queue.begin(), words_queue.end(), cmp);
        
        vector<string> result;
        for(int i = 0; i < k; i++) {
            result.push_back(words_queue[i].val);
        }
        
        return result;
    }
};

