// Node v5.0+
const fs = require('fs'),
      readline = require('readline');

const main = (filename, done) => {
    var cf = {};
    
    readline.createInterface({
        input: fs.createReadStream(filename),
        terminal: false
    }).on('line', (line)=>{
        for (var i in line) {
            var c = line[i];
            cf[c] = (cf[c] || 0) + 1;
        }
    }).on('close', ()=>{
        var cfv = [];
        for (var k in cf) {
            cfv.push([k, cf[k]]);
        }
        cfv.sort((a,b)=>{
            return b[1] - a[1];
        });
        done(cfv);
    });
}

main(process.argv[2] || 'page1.txt', (sorted)=>{
    for(var i in sorted) {
        console.log(sorted[i][0]+': '+sorted[i][1]);
    }
    process.exit(0);
});
