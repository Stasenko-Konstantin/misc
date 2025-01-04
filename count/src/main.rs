use std::collections::HashMap;
use std::fs::File;
use std::io::{Error, Read};
use std::path::Path;
use encoding_rs::UTF_8;

fn main() {
    println!("counting...\n"); // TODO make index for files (maybe vec?) and print his length
    let curr_dir = std::env::current_dir().unwrap();
    let mut hmap = HashMap::new();
    let res = count(curr_dir, &mut hmap);
    match res {
        Ok(res) => println!("{:#?}", res),
        Err(e) => eprintln!("{:#?}", e),
    }
    println!("\ndone!");
}

fn count(
    path: std::path::PathBuf,
    map: &mut HashMap<String, i32>,
) -> Result<&mut HashMap<String, i32>, Error> {
    if path.to_str().unwrap().starts_with('.') {
        return Ok(map);
    }
    let dir = std::fs::read_dir(path.clone())?;
    for entry in dir {
        let entry = entry?;
        if entry.path().file_name().unwrap().to_str().unwrap().starts_with('.') {
            continue;
        }
        if entry.path().is_dir() {
            count(entry.path(), map)?;
        }
        if !entry.path().file_name().unwrap().to_str().unwrap().contains('.') {
            continue;
        }
        if !is_text_file(entry.path()) {
            continue;
        }
        let file = std::fs::read_to_string(entry.path().clone())?;
        let mut ext = ".".to_string();
        if let Some(ext_os_str) = entry.path().extension() {
            if let Some(ext_str) = ext_os_str.to_str() {
                ext = ext_str.to_string();
            }
        }
        let res = if let Some(i) = map.get_mut(&ext) {
            *i
        } else {
            0
        } + count_file_lines(file);
        if res != 0 {
            map.insert(ext, res);
        }
    }
    Ok(map)
}

fn count_file_lines(file: String) -> i32 {
    file.lines().count() as i32 // TODO working very strange
}

fn is_text_file<P: AsRef<Path>>(path: P) -> bool {
    let f = || {
        let mut file = File::open(path)?;
        let mut buffer = Vec::new();
        file.read_to_end(&mut buffer)?;
        let text_threshold = 0.7; // TODO very strange stuff
        let (cow, _, had_errors) = UTF_8.decode(&buffer);
        if had_errors {
            return Ok(false);
        }
        let printable_chars = cow.chars().filter(|c| c.is_alphanumeric() || c.is_whitespace()).count();
        let total_chars = cow.chars().count();
        Ok::<bool, Error>(printable_chars as f64 / total_chars as f64 > text_threshold)
    };
    f().unwrap_or(false)
}