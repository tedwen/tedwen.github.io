import os

def main(filename, encoding = 'utf-8'):
    if not os.path.exists(filename):
        print 'Cannot open file'
        return
        
    cf = {}
    
    with open(filename) as fp:
        for line in fp:
            for c in line.decode(encoding):
                cf[c] = cf.get(c, 0) + 1
    
    cfv = sorted([(k,v) for k,v in cf.items()], key=lambda x:x[1], reverse=True)
    
    for c, v in cfv:
        print '%s: %d' % (c, v)

if __name__ == '__main__':
    import sys
    filename = sys.argv[1] if len(sys.argv) > 1 else 'page1.txt'
    main(filename)
    