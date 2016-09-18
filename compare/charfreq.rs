use std::env;
use std::io::prelude::*;
use std::fs::File;
use std::io::BufReader;
use std::collections::HashMap;

fn main() {
    let args: Vec<_> = env::args().collect();
    let filename = if args.len() > 1 { args[1].as_str() } else { "page1.txt" };
    
    let mut cf = HashMap::new();
    
    if let Ok(f) = File::open(filename) {
        for line in BufReader::new(f).lines() {
            for c in line.unwrap().chars() {
                *cf.entry(c).or_insert(0) += 1;
            }
        }
        
        let mut cfv: Vec<(&char, &u32)> = cf.iter().collect();
        cfv.sort_by(|a,b| b.1.cmp(a.1));

        for v in cfv.iter() {
            println!("{}: {}", v.0, v.1);
        }
    } else {
        println!("Cannot open file");
    }
}
