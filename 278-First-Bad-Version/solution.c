// The API isBadVersion is defined for you.
// bool isBadVersion(int version);

int firstBadVersion(int n) {
    long long start = 1, end = n, half = (start + end) / 2;
    while (start <= end) {
        if(isBadVersion(start)) {
            return start;
        } else if (!isBadVersion(end)) {
            return end + 1;
        } else if (isBadVersion(half)) {
            end = half - 1;
            half = (start + end) / 2;
        } else {
            start = half + 1;
        }
        half = (start + end) / 2;
    }
    return start;
}

