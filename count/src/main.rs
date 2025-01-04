use std::any::Any;
use std::collections::HashMap;
use std::io::ErrorKind;
use std::io::Error;

fn main() {
    println!("counting...\n");
    let curr_dir = std::env::current_dir().unwrap();
    let res = count(curr_dir, HashMap::new());
    match res {
        Ok(res) => println!("{:#?}", res),
        Err(e) => eprintln!("{:#?}", e),
    }
    println!("\ndone!");
}

fn count(
    path: std::path::PathBuf,
    mut map: HashMap<String, i32>,
) -> Result<HashMap<String, i32>, Error> {
    if let Ok(dir) = std::fs::read_dir(path.clone()) {
        for entry in dir.flatten() {
            if entry.path().is_dir() {
                return count(entry.path(), map);
            }
            if let Ok(file) = std::fs::read_to_string(entry.path().clone()) {
                let mut ext = ".".to_string();;
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
                continue;
            }
            return Err(Error::new(ErrorKind::NotFound, format!("cannot read file: {}", entry.path().to_str().unwrap())));
        }
        return Ok(map);
    }
    Err(Error::new(ErrorKind::NotFound, format!("current directory not found: {}", path.to_str().unwrap())))
}

fn count_file_lines(file: String) -> i32 {
    file.lines().count() as i32
}