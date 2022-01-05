#ifndef TQueue_HPP
#define TQueue_HPP

#include "tqueue_item.hpp"

template <typename T>
class TQueue
{
public:
    TQueue()
    {
        heap = new TQueueItem<T>[max_length];
    }

    TQueue(const TQueue &o)
        : heap(o.heap), element_size(o.element_size), max_length(o.max_length), length(o.max_length) {}

    void Push(const T &item)
    {
        if (length >= max_length - 1)
        {
            max_length += 100;
            TQueueItem<T> *heap2 = new TQueueItem<T>[max_length];
            for (size_t i = 0; i < length; ++i)
            {
                heap2[i] = heap[i];
            }
            free(heap);
            heap = heap2;
        }
        TQueueItem<T> n(item);
        int input_pos, parent_pos;
        input_pos = length;
        heap[input_pos] = n;
        parent_pos = (input_pos - 1) / 2;
        while (parent_pos >= 0 && input_pos > 0)
        {
            TQueueItem<T> temp = heap[input_pos];
            heap[input_pos] = heap[parent_pos];
            heap[parent_pos] = temp;
            input_pos = parent_pos;
            parent_pos = (input_pos - 1) / 2;
        }
        ++length;
    }

    void Pop()
    {
        if (length == 0)
        {
            return;
        }
        heap[0] = heap[length - 1];
        --length;
        Heapify(0);
    }

    T Top() const
    {
        if (length == 0)
        {
            std::cout << "\nError: Queue is empty" << std::endl;
            exit(EXIT_FAILURE);
        }
        return heap[0].GetObject();
    }

    bool Empty() const
    {
        return length == 0;
    }

    size_t Length() const
    {
        return length;
    }

    template <typename A>
    friend std::ostream& operator<<(std::ostream &os, const TQueue<A> &_queue)
    {
        size_t i = 0, k = 1;
        while (i < _queue.length)
        {
            while ((i < k) && (i < _queue.length))
            {
                os << _queue.heap[i] << "\t";
                ++i;
            }
            if (i != _queue.length)
            {
                os << std::endl;
            }
            k = k * 2 + 1;
        }
        return os;
    }

    void Clear()
    {
        while (length > 0)
        {
            Pop();
        }
    }

    ~TQueue() {}

private:
    void Heapify(const int &position)
    {
        size_t left = 2 * position + 1, right = 2 * position + 2;
        if (left < length)
        {
            TQueueItem<T> tmp = heap[position];
            heap[position] = heap[left];
            heap[left] = tmp;
            Heapify(left);
        }
        if (right < length)
        {
            TQueueItem<T> tmp = heap[position];
            heap[position] = heap[right];
            heap[right] = tmp;
            Heapify(right);
        }
    }
    TQueueItem<T> *heap;
    const int element_size = sizeof(TQueueItem<T>);
    size_t max_length = 100;
    size_t length = 0;
};

#endif
