CC := gcc
SRC := ./src
INC := ./inc
OUT := ./out

Cflags := -I$(INC)
LIBS := -ltidy -lcurl
OBJflags := -I$(SRC) -c

PARSED_DIR := ./parsed

.PHONY : clean
clean:
	rm *.exe $(OUT)/* 2> /dev/null || true

$(OUT)/%.o : $(SRC)/%.c $(INC)/*.h
	mkdir $(OUT) 2> /dev/null || true
	$(CC) $(Cflags) $(OBJflags) $< -o $@ $(LIBS)

onethread.out : $(OUT)/onethreadmain.o $(OUT)/recipeparser.o $(OUT)/download.o $(OUT)/htmlnode.o $(OUT)/urlstorage.o $(OUT)/urlutils.o $(OUT)/gettime.o
	$(CC) $(Cflags) $^ -o $@ $(LIBS)

multithread1.out : $(SRC)/multithreadmain.c $(OUT)/recipeparser.o $(OUT)/download.o $(OUT)/htmlnode.o $(OUT)/urlstorage.o $(OUT)/urlutils.o $(OUT)/threadqueue.o $(OUT)/gettime.o
	mkdir $(PARSED_DIR) 2> /dev/null || true
	mkdir $(PARSED_DIR)/1 2> /dev/null || true
	$(CC) $(Cflags) $^ -o $@ $(LIBS) -DRESULT_DIR=\"$(PARSED_DIR)/1\" -DTHREADS=1

multithread2.out : $(SRC)/multithreadmain.c $(OUT)/recipeparser.o $(OUT)/download.o $(OUT)/htmlnode.o $(OUT)/urlstorage.o $(OUT)/urlutils.o $(OUT)/threadqueue.o $(OUT)/gettime.o
	mkdir $(PARSED_DIR) 2> /dev/null || true
	mkdir $(PARSED_DIR)/2 2> /dev/null || true
	$(CC) $(Cflags) $^ -o $@ $(LIBS) -DRESULT_DIR=\"$(PARSED_DIR)/2\" -DTHREADS=2

multithread4.out : $(SRC)/multithreadmain.c $(OUT)/recipeparser.o $(OUT)/download.o $(OUT)/htmlnode.o $(OUT)/urlstorage.o $(OUT)/urlutils.o $(OUT)/threadqueue.o $(OUT)/gettime.o
	mkdir $(PARSED_DIR) 2> /dev/null || true
	mkdir $(PARSED_DIR)/4 2> /dev/null || true
	$(CC) $(Cflags) $^ -o $@ $(LIBS) -DRESULT_DIR=\"$(PARSED_DIR)/4\" -DTHREADS=4

multithread8.out : $(SRC)/multithreadmain.c $(OUT)/recipeparser.o $(OUT)/download.o $(OUT)/htmlnode.o $(OUT)/urlstorage.o $(OUT)/urlutils.o $(OUT)/threadqueue.o $(OUT)/gettime.o
	mkdir $(PARSED_DIR) 2> /dev/null || true
	mkdir $(PARSED_DIR)/8 2> /dev/null || true
	$(CC) $(Cflags) $^ -o $@ $(LIBS) -DRESULT_DIR=\"$(PARSED_DIR)/8\" -DTHREADS=8

multithread16.out : $(SRC)/multithreadmain.c $(OUT)/recipeparser.o $(OUT)/download.o $(OUT)/htmlnode.o $(OUT)/urlstorage.o $(OUT)/urlutils.o $(OUT)/threadqueue.o $(OUT)/gettime.o
	mkdir $(PARSED_DIR) 2> /dev/null || true
	mkdir $(PARSED_DIR)/16 2> /dev/null || true
	$(CC) $(Cflags) $^ -o $@ $(LIBS) -DRESULT_DIR=\"$(PARSED_DIR)/16\" -DTHREADS=16

multithread32.out : $(SRC)/multithreadmain.c $(OUT)/recipeparser.o $(OUT)/download.o $(OUT)/htmlnode.o $(OUT)/urlstorage.o $(OUT)/urlutils.o $(OUT)/threadqueue.o $(OUT)/gettime.o
	mkdir $(PARSED_DIR) 2> /dev/null || true
	mkdir $(PARSED_DIR)/32 2> /dev/null || true
	$(CC) $(Cflags) $^ -o $@ $(LIBS) -DRESULT_DIR=\"$(PARSED_DIR)/32\" -DTHREADS=32