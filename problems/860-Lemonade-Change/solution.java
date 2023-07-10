class Solution {
    public boolean lemonadeChange(int[] bills) {
        int fiveCnt = 0, tenCnt = 0, twentyCnt = 0;
        for (int bill : bills) {
            switch (bill) {
                case 5:
                    fiveCnt++;
                    break;
                case 10:
                    tenCnt++;

                    if (fiveCnt <= 0) {
                        return false;
                    } else {
                        fiveCnt--;
                    }

                    break;
                case 20:
                    twentyCnt++;

                    if (tenCnt > 0 && fiveCnt > 0) {
                        fiveCnt--;
                        tenCnt--;
                    } else {
                        if (fiveCnt < 3) {
                            return false;
                        } else {
                            fiveCnt -= 3;
                        }
                    }

                    break;
            }
        }

        return true;
    }
}

