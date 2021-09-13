import os
import time

res=[]
def benchmark(command,repeat):
    total=0
    for x in range(repeat):
        start = time.time()
        os.system(command)
        time_taken=(time.time()-start)
        res.append(time_taken)

    return res,sum(res)/len(res)