//
// Created by Messi on 2023/6/8.
//

#ifndef CREACTORSERVER_EPOLLPOLLER_H
#define CREACTORSERVER_EPOLLPOLLER_H

#include "Poller.h"

/*
 * epoll
 */
namespace
{
    const int kNew = -1;
    const int kAdded = 1;
    const int kDeleted = 2;
}
struct epoll_event;
namespace core::net
{
    class EpollPoller : public Poller
    {
    public:
        explicit EpollPoller(EventLoop *loop);
        ~EpollPoller() override;

        int poll(int timeout, ChannelList *activeChannels) override;
        void updateChannel(Channel *channel) override;
        void removeChannel(Channel *channel) override;

    private:
        static const int kInitEventListSize = 16;
        static const char *operationToString(int op);
        void fillActiveChannels(int activeNum, ChannelList *activeChannels);
        void update(int operation, Channel *channel);

    private:
        std::vector<struct epoll_event> eventList_;
        int epollfd_;
    };
}

#endif //CREACTORSERVER_EPOLLPOLLER_H
