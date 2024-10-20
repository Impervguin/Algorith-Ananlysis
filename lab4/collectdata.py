import os
import subprocess
import re

def ParseOutput(out: str):
    matches = re.findall(r"[-+]?(?:\d*\.*\d+)", out)
    res = []
    for m in matches:
        try:
            res.append(int(m))
        except ValueError:
            res.append(float(m))
    return res
        
multithreads = [1, 2, 4, 8, 16, 32, 64]
# multithreads = [1, 2]
multithreadFormat = "multithread{}.out"
onethread = "onethread.out"
linksCount = 1000
linksFile = "links.txt"
TestCount = 10


print("Parsing site Receipts...")
rc = subprocess.call(["python3", "parse.py", "-c", str(linksCount), "-s", linksFile])
if (rc != 0):
    print("Error parsing site Receipts.")
    exit(1)
print(f"Got {linksCount} site Receipts. Saved in {linksFile}")

rc = subprocess.call(["make", "clean"])

RunsInfo = dict()
RunsInfo['oneThread'] = []

print("Running onethreaded Tests...")
print("Building onethread program...")
rc = subprocess.call(["make", onethread])
print("Done building onethread program.")

for test in range(1, TestCount + 1):
    print(f"Running onethreaded test {test}...")
    process = subprocess.run(["./" + onethread], capture_output=True)
    if (process.returncode != 0):
        print(f"Error running onethreaded test {test}.")
        continue

    parsedValues = ParseOutput(str(process.stdout))
    runInfo = {"time": parsedValues[1], "done": parsedValues[2], "failed": parsedValues[3]}
    RunsInfo['oneThread'].append(runInfo)
    print(f"Onethreaded test {test} passed. Time: {runInfo['time']}s, Done: {runInfo['done']}, Failed: {runInfo['failed']}")

print("Done onethreaded tests.")

print("Running multithreaded Tests...")
for threads in multithreads:
    print(f"Running multithreaded test with {threads} threads...")
    print("Building multithreaded program...")
    rc = subprocess.call(["make", multithreadFormat.format(threads)])
    print("Done building multithreaded program.")

    for test in range(1, TestCount + 1):
        print(f"Running multithreaded test {test} with {threads} threads...")
        process = subprocess.run(["./" + multithreadFormat.format(threads)], capture_output=True)
        if (process.returncode != 0):
            print(f"Error running multithreaded test {test} with {threads} threads.")
            continue

        parsedValues = ParseOutput(str(process.stdout))
        runInfo = {"time": parsedValues[1], "done": parsedValues[2], "failed": parsedValues[3]}
        if threads not in RunsInfo:
            RunsInfo[threads] = []
        RunsInfo[threads].append(runInfo)
        print(f"Multithreaded test {test} with {threads} threads passed. Time: {runInfo['time']}s, Done: {runInfo['done']}, Failed: {runInfo['failed']}")

print("Done multithreaded tests.")

print("Writing results to results.txt...")
with open("results.txt", "w") as f:
    for threads, runs in RunsInfo.items():
        f.write(f"{threads} threads:\n")
        for run in runs:
            f.write(f"Time: {run['time']}s, Done: {run['done']}, Failed: {run['failed']}\n")

print("Results saved in results.txt.")

print("Calculating average results...")
averages = dict()
for name, runs in RunsInfo.items():
    averages[name] = sum([r['time'] for r in runs])
    averages[name] /= len(runs)

print("Average results:")
for name, average in averages.items():
    print(f"{name}: {average}s")

