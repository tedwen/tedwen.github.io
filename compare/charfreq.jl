filename = length(ARGS) > 0 ? ARGS[1] : "page1.txt"

cf = Dict()

f = open(filename)
for line in eachline(f)
    for c in line
        cf[c] = get(cf, c, 0) + 1
    end
end
close(f)

cf = sort(collect(cf), by=x->x[2], rev=true)

for (k, v) in cf
    println("$k: $v")
end
