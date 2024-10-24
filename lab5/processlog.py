import numpy

unitData = dict()
with open("pipeline.log", "r") as logFile:
    for line in logFile:
        id, actType, start, end = line.split('\t')
        id = int(id)
        if actType not in ['loading', 'parsing', 'storage']:
            continue
        if id not in unitData.keys():
            unitData[id] = dict()
        unitData[id][actType] = (float(start), float(end))

loading = [d['loading'][1] - d['loading'][0] for _, d in unitData.items()]
parsing = [d['parsing'][1] - d['parsing'][0] for _, d in unitData.items()]
storage = [d['storage'][1] - d['storage'][0] for _, d in unitData.items()]

waitParsing = [d['parsing'][0] - d['loading'][1] for _, d in unitData.items()]
waitStorage = [d['storage'][0] - d['parsing'][1] for _, d in unitData.items()]

notWorkParsing = [unitData[id+1]['parsing'][0] - unitData[id]['parsing'][1] for id in unitData.keys() if id + 1 in unitData.keys()] 
notWorkStorage = [unitData[id+1]['storage'][0] - unitData[id]['storage'][1] for id in unitData.keys() if id + 1 in unitData.keys()]


processing = [d['storage'][1] - d['loading'][0] for _, d in unitData.items()]

print(f"Loading: min: {min(loading)*1000:.4g} msec, max: {max(loading)*1000:.4g} msec, avg: {numpy.average(loading)*1000:.4g} msec, med: {numpy.median(loading)*1000:.4g} msec")
print(f"Parsing: min: {min(parsing)*1000:.4g} msec, max: {max(parsing)*1000:.4g} msec, avg: {numpy.average(parsing)*1000:.4g} msec, med: {numpy.median(parsing)*1000:.4g} msec")
print(f"Storage: min: {min(storage)*1000:.4g} msec, max: {max(storage)*1000:.4g} msec, avg: {numpy.average(storage)*1000:.4g} msec, med: {numpy.median(storage)*1000:.4g} msec")
print(f"Wait parsing: min: {min(waitParsing)*1000:.4g} msec, max: {max(waitParsing)*1000:.4g} msec, avg: {numpy.average(waitParsing)*1000:.4g} msec, med: {numpy.median(waitParsing)*1000:.4g} msec")
print(f"Wait storage: min: {min(waitStorage)*1000:.4g} msec, max: {max(waitStorage)*1000:.4g} msec, avg: {numpy.average(waitStorage)*1000:.4g} msec, med: {numpy.median(waitStorage)*1000:.4g} msec")
print(f"Not work parsing: min: {min(notWorkParsing)*1000:.4g} msec, max: {max(notWorkParsing)*1000:.4g} msec, avg: {numpy.average(notWorkParsing)*1000:.4g} msec, med: {numpy.median(notWorkParsing)*1000:.4g} msec")
print(f"Not work storage: min: {min(notWorkStorage)*1000:.4g} msec, max: {max(notWorkStorage)*1000:.4g} msec, avg: {numpy.average(notWorkStorage)*1000:.4g} msec, med: {numpy.median(notWorkStorage)*1000:.4g} msec")
print(f"Processing: min: {min(processing)*1000:.4g} msec, max: {max(processing)*1000:.4g} msec, avg: {numpy.average(processing)*1000:.4g} msec, med: {numpy.median(processing)*1000:.4g} msec")





