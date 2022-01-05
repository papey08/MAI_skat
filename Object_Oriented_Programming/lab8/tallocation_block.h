#ifndef TALLOCATION_BLOCK_H
#define TALLOCATION_BLOCK_H

#include "tqueue.hpp"

class TAllocationBlock 
{
public:
    TAllocationBlock(const size_t &size, const size_t &count);
    void* Allocate(const size_t &size_of_block);
    void Deallocate(void *pointer);
    bool HasFreeBlocks();
    virtual ~TAllocationBlock();

private:
    size_t size;
    size_t count;
    size_t free_count;
    char *used_blocks;
    TQueue<void*> q_free_blocks;
};

#endif
