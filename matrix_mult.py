import numpy as np
import time

a = np.array((
    (1,2,3),
    (4,5,6),
    (7,8,9)
    )
)

b = np.array((
    (1,2,3),
    (4,5,6),
    (7,8,9)
    )
)


# Time it from here
t = time.time()
c  = np.matmul(a,b)
d = time.time() - t

print("Time %s µs"%(d*10**6))
print(c)