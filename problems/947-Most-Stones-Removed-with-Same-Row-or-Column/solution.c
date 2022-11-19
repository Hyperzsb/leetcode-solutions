int removeStones(int** stones, int stonesSize, int* stonesColSize){
    int *flags = (int *)malloc(1010 * sizeof(int));
    memset(flags, 0, 1010 * sizeof(int));
    
    int *trace = (int *)malloc(1010 * sizeof(int));
    memset(trace, 0, 1010 * sizeof(int));
    
    int index = -1;
    int height = 0;
    int islandsNum = 0;
    
    while(true) {
        index = -1;
        height = 0;
        for(int i = 0; i < stonesSize; i++) {
            if(flags[i] == 0) {
                index = i;
                break;
            }
        }
        
        if(index == -1) {
            break;
        }
        
        trace[height] = index;
        flags[trace[height]] = 1;
        
        int next = -1;
        while(height >= 0) {
            next = -1;
            for(int i = 0; i < stonesSize; i++) {
                if(flags[i] == 0 &&
                   (stones[trace[height]][0] == stones[i][0] ||
                    stones[trace[height]][1] == stones[i][1])) {
                    next = i;
                    break;
                }
            }
            
            if(next != -1) {
                height++;
                trace[height] = next;
                flags[trace[height]] = 1;
            } else {
                height--;
            }
        }
        
        islandsNum++;
    }
    
    return stonesSize - islandsNum;
}

