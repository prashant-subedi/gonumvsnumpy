import numpy as np
import time
x,k,y = 500,500,500
a = np.random.rand(x,k)
b = np.random.rand(k,y)

# Time it from here
t = time.time()
c  = np.matmul(a,b)
d = time.time() - t

print("Time %s s"%(d))
