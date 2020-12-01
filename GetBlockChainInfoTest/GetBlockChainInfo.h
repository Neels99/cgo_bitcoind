#include <stdio.h>

typedef struct RawGetBlockChainInfoResult{
    char* chain;
    int blocks;
    int headers;
    char* bestblockhash;
    double difficulty;
    long long mediantime;
    double verificationprogress;
    int pruned;
    int pruneheight;
    char* chainwork;
    
} RawGetBlockChainInfoResult;