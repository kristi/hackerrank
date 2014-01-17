import random

f = open('input', 'w')
n = 100000

print >>f, 3
print >>f, 5
print >>f, '2 1 3 1 2'
print >>f, n
print >>f, ' '.join(str(random.randint(1,1000000)) for _ in xrange(n))
n = 20000
print >>f, n
print >>f, ' '.join(str(x) for x in xrange(n))
