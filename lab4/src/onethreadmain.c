#include <stdio.h>
#include <tidy.h>
#include <tidybuffio.h>
#include <curl/curl.h>
#include <time.h>
#include "gettime.h"
#include "recipeparser.h"
#include "urlstorage.h"
#include "urlutils.h"
 
#define RESULT_DIR "./parsed/one"
#define EXPRESSION ".html"

int main(void) {
  curl_global_init(CURL_GLOBAL_DEFAULT);
  url_storage_t storage;
  init_url_storage(&storage);
  int err = read_url_storage(&storage, "links.txt");
  if (err) {
    // fprintf(stderr, "Error reading URLs from file: %d\n", err);
    return 1;
  }
  printf("Found %d URLs.\n", storage.len);
  int failed = 0;
  struct timeval start;
  gettimeofday(&start, NULL);
  for (int i = 0; i < storage.len; i++) {
      char *basename = get_url_basename(storage.urls[i]);
      char *result_path = malloc(strlen(RESULT_DIR) + strlen(basename) + strlen(EXPRESSION) + 2);
      if(result_path == NULL) {
        failed++;
        // printf("can't process: %s", storage.urls[i]);
        continue;
      }
      sprintf(result_path, "%s/%s%s", RESULT_DIR, basename, EXPRESSION);
      free(basename);
      err = download_article(storage.urls[i], result_path);
      free(result_path);
      if (err) {
        failed++;
        // fprintf(stderr, "Error downloading article from %s: %d\n", storage.urls[i], err);
        continue;
      }
      // printf("Downloaded: %s\n", storage.urls[i]);
  }
  struct timeval end;
  gettimeofday(&end, NULL);
  printf("Done in %lf seconds\n", get_time(&start, &end));
  printf("Parsed %d articles.\n", storage.len);
  printf("Failed to download %d articles\n", failed);
  release_url_storage(&storage);
  curl_global_cleanup();
  return 0;
}