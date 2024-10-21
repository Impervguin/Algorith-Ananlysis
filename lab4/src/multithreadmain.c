#include <pthread.h>
#include <time.h>
#include <stdio.h>
#include <curl/curl.h>
#include <unistd.h>
#include "recipeparser.h"
#include "urlstorage.h"
#include "urlutils.h"
#include "gettime.h"
#include "threadqueue.h"

// #define THREADS 8
#define URLS_PER_THREAD 5

#ifndef THREADS 
#error No threads specified
#endif

// #define RESULT_DIR "./parsed"

#ifndef RESULT_DIR 
#error No RESULT_DIR specified
#endif


#define EXPRESSION ".html"

struct thread_arg
{
    char **urls;
    int start;
    int end;
    int thread_id;
    thqueue_t *queue;
    int failed;
};


void *thread_func(void *arg) {
    // printf(RESULT_DIR);
    int err = 0;
    struct thread_arg *ta = (struct thread_arg *)arg;
    for (int i = ta->start; i < ta->end; i++) {
        char *basename = get_url_basename(ta->urls[i]);
        char *result_path = malloc(strlen(RESULT_DIR) + strlen(basename) + strlen(EXPRESSION) + 2);
        if (result_path == NULL) {
            ta->failed++;
            free(basename);
            free(result_path);
            continue;
        }
        sprintf(result_path, "%s/%s%s", RESULT_DIR, basename, EXPRESSION);
    
        err = download_article(ta->urls[i], result_path);
        free(basename);
        free(result_path);
        if (err) {
            ta->failed++;
            continue;
        }
    }
    int res = thread_queue_enqueue(ta->thread_id, ta->queue);
    return NULL;
}



int main(void) {
    // printf(RESULT_DIR);
    // printf("%d", THREADS);
    thqueue_t queue;
    int err;
    err = thread_queue_init(&queue, THREADS);
    if (err!=OK) {
        fprintf(stderr, "Error initializing thread queue: %d\n", err);
        return 1;
    }
    struct thread_arg args[THREADS];
    
    pthread_t threads[THREADS];

    curl_global_init(CURL_GLOBAL_DEFAULT);
    url_storage_t storage;
    init_url_storage(&storage);
    err = read_url_storage(&storage, "links.txt");
    if (err) {
        return 1;
    }
    printf("Found %d URLs.\n", storage.len);
    int failed = 0;
    int done = 0;
    struct timeval start;
    gettimeofday(&start, NULL);
    for (int i = 0; i < THREADS; i++) {
        args[i].urls = storage.urls;
        args[i].start = i * URLS_PER_THREAD;
        args[i].end = (((i + 1) * URLS_PER_THREAD) > storage.len) ? storage.len : ((i + 1) * URLS_PER_THREAD);
        args[i].thread_id = i;
        args[i].queue = &queue;
        args[i].failed = 0;
        pthread_create(threads + i, NULL, &thread_func, args + i);
    }
    int current_url = (((THREADS) * URLS_PER_THREAD) > storage.len) ? storage.len : ((THREADS) * URLS_PER_THREAD);
    while (true)
    {
        int resp;
        int done_thread;
        for (;(resp = thread_queue_dequeue(&done_thread, &queue)) == EMPTY_ERROR;);
        if (resp == OK) {
            done += args[done_thread].end - args[done_thread].start;
            pthread_join(threads[done_thread], NULL);

            if (args[done_thread].failed != 0) {
                failed += args[done_thread].failed;
            }
            if (done == storage.len) {
                break;
            }
            args[done_thread].urls = storage.urls;
            args[done_thread].start = current_url;
            args[done_thread].end = (current_url + URLS_PER_THREAD > storage.len)? storage.len : (current_url + URLS_PER_THREAD);
            current_url = args[done_thread].end;
            args[done_thread].thread_id = done_thread;
            args[done_thread].queue = &queue;
            args[done_thread].failed = 0;
            pthread_create(threads + done_thread, NULL, &thread_func, args + done_thread);
        }
    }
    struct timeval end;
    gettimeofday(&end, NULL);
    printf("Done in %lf seconds\n", get_time(&start, &end));
    printf("Parsed %d articles.\n", done);
    printf("Failed to download %d articles\n", failed);
    release_url_storage(&storage);
    curl_global_cleanup();
    return 0;
}