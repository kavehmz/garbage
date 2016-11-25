import happybase
import time

connection = happybase.Connection('localhost')
table = connection.table('test')

b = table.batch()
start = time.time()

for x in range(0, 1000000):
	b.put(b'row-key-'+`x`, {b'cf:col': b'value', b'cf:col2': b'value2'})
b.send()

end = time.time()
elapsed = end - start
print("elapsed time = {:.12f} seconds".format(elapsed))

