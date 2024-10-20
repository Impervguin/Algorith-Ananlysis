#include "threadqueue.h"
#include <stdlib.h>

int thread_queue_init(thqueue_t *queue, int amount) {
    if (!queue) return NULL_POINTER_ERROR;
    if (amount <= 0) return NULL_POINTER_ERROR;
    queue->queue = malloc(sizeof(int) * amount);
    if (!queue->queue) return MEMORY_ERROR;
    queue->len = amount;
    queue->head = 0;
    queue->tail = 0;
    queue->full = false;
    if (pthread_mutex_init(&queue->mutex, NULL) != 0) {
        free(queue->queue);
        return MEMORY_ERROR;
    }
    return OK;
}

// Not thread-safe
void thread_queue_release(thqueue_t *queue) {
    if (!queue) return;
    if (queue->queue) {
        free(queue->queue);
        pthread_mutex_destroy(&queue->mutex);
    }
}


int thread_queue_dequeue(int *thread_id, thqueue_t *queue) {
    if (!queue) {
        pthread_mutex_unlock(&queue->mutex);
        return NULL_POINTER_ERROR;
    }
    pthread_mutex_lock(&queue->mutex);
    if (queue->len == 0) {
        pthread_mutex_unlock(&queue->mutex);
        return EMPTY_ERROR;
    }

    if (queue->tail == queue->head && !queue->full) {
        pthread_mutex_unlock(&queue->mutex);
        return EMPTY_ERROR;
    }

    *thread_id = queue->queue[queue->tail];
    queue->tail++;
    if (queue->tail == queue->len)
        queue->tail = 0;
    queue->full = false;

    pthread_mutex_unlock(&queue->mutex);
    return OK;
}


int thread_queue_enqueue(int thread_id, thqueue_t *queue) {
    

    if (!queue) {
        pthread_mutex_unlock(&queue->mutex);
        return NULL_POINTER_ERROR;
    }

    pthread_mutex_lock(&queue->mutex);

    if (queue->len == 0) {
        pthread_mutex_unlock(&queue->mutex);
        return EMPTY_ERROR;
    }
    if (queue->full) {
        pthread_mutex_unlock(&queue->mutex);
        return FILLED_ERROR;
    }
    
    int next_head = queue->head + 1;
    if (next_head == queue->len) 
        next_head = 0;
    if (next_head == queue->tail)
        queue->full = true;
    queue->queue[queue->head] = thread_id;
    queue->head = next_head;

    pthread_mutex_unlock(&queue->mutex);
    return OK;
}