#ifndef TQueue_ITEM_HPP
#define TQueue_ITEM_HPP

#include <iostream>

template <typename T>
class TQueueItem 
{
public:
    TQueueItem() = default;

    TQueueItem(const T &item) : item(item) {}
    
    TQueueItem(const TQueueItem<T> &other) : item(other.item) {}

    T GetObject() const
    {
        return item;
    }

    TQueueItem<T> &operator=(const TQueueItem<T> &other) 
    {
        this->item = other.item;
        return *this;
    }

    bool operator==(const TQueueItem<T> &other) const
    {
        return (item == other.item);
    }

    bool operator!=(const TQueueItem<T> &other) const 
    {
        return (item != other.item);
    }

    ~TQueueItem() {}

private:
    T item;
};

#endif
