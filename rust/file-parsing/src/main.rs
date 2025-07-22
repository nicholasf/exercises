use std::env;

use std::io;
use std::io::BufReader;

use std::fs::File;
use std::fs::OpenOptions;

use std::path::PathBuf;

use std::process;

fn main() {
    let args: Vec<String> = env::args().collect();

    println!("args len {}", args.len());

    if args.len() < 2 {
        println!("\n\n Please supply a file to open!");
        process::exit(0);
    }

    let filePath: &str = args[1].trim();
    let path: PathBuf = PathBuf::from(filePath.to_string());
    

    let file = match open_file(path) {
        Ok(file) => file,
        Err(e) => panic!("This file is problematic. {}", e),    
    };

}

fn open_file(path: PathBuf) -> io::Result<File> {
    return OpenOptions::new().read(true).open(path);
}

// fn parse(term: &str) -> io::Result<()> {

//     for line in reader.lines() {
//         let l = line?;

//         if l.contains(term) {
//             println!("{}", l);
//         }
//     }

//     Ok(())
// }

