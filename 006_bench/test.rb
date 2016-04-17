
class Contract
end

Contract.new()

t0 = Time.now
n=0;
for i in 1..10000000
	n+=1
end
puts "loop time :" + (Time.now- t0).to_s

t0 = Time.now
n=0;
for i in 1..10000000
	n=Math.exp(i)
end
puts "exp time :"+  (Time.now- t0).to_s

t0 = Time.now
for i in  1..10000
	a=Array.new
	for j in  1..10000
		a.push(j)
	end
end
puts "append time :"+  (Time.now- t0).to_s

t0 = Time.now
for i in  1..10000000
	p=Contract.new;
end
puts "class time :"+  (Time.now- t0).to_s
