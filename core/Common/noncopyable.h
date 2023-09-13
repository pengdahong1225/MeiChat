//
// Created by Messi on 2023/6/5.
//

#ifndef CREACTORSERVER_NONCOPYABLE_H
#define CREACTORSERVER_NONCOPYABLE_H

class noncopyable
{
public:
    noncopyable(const noncopyable &) = delete;
    void operator=(const noncopyable &) = delete;

protected:
    noncopyable() = default;
    ~noncopyable() = default;
};

#endif //CREACTORSERVER_NONCOPYABLE_H
