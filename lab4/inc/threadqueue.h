#ifndef THREAD_QUEUE_H__
#define THREAD_QUEUE_H__

#include <pthread.h>
#include <stdbool.h>

// queue for finished threads
struct thread_queue {
    int *queue;
    int len;
    int head;
    int tail;
    pthread_mutex_t mutex;
    bool full;
};

typedef struct thread_queue thqueue_t;

#define OK 0
#define FILLED_ERROR 1
#define EMPTY_ERROR 2
#define MEMORY_ERROR 3
#define NULL_POINTER_ERROR 4

int thread_queue_init(thqueue_t *queue, int amount);
void thread_queue_release(thqueue_t *queue);
int thread_queue_dequeue(int *thread_id, thqueue_t *queue);
int thread_queue_enqueue(int thread_id, thqueue_t *queue);

#endif