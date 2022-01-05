#include "tallocation_block.h"

TAllocationBlock::TAllocationBlock(const size_t &size, const size_t &count) : size(size), count(count)
{
    used_blocks = (char *)malloc(size * count);
    for (size_t i = 0; i < count; ++i)
    {
        q_free_blocks.Push(used_blocks + i * size);
    }
    free_count = count;
}

void* TAllocationBlock::Allocate(const size_t &size_of_block)
{
    if (size != size_of_block)
    {
        std::cout << "Error" << std::endl;
    }
    void *result = nullptr;
    if (free_count == 0)
    {
        size_t old_count = count;
        count += 10;
        free_count += 10;
        used_blocks = (char*) realloc(used_blocks, size * count);
        for (size_t i = old_count; i < count; ++i)
        {
            q_free_blocks.Push(used_blocks + i * size);
        }
    }
    result = q_free_blocks.Top();
    q_free_blocks.Pop();
    --free_count;
    return result;
}

void TAllocationBlock::Deallocate(void *pointer)
{
    q_free_blocks.Push(pointer);
    ++free_count;
}

bool TAllocationBlock::HasFreeBlocks()
{
    return free_count > 0;
}

TAllocationBlock::~TAllocationBlock()
{
    free(used_blocks);
}
