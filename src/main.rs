use std::{env, path::Path, process::Command};

use notify::{event::DataChange, Result, Watcher};
use walkdir::{DirEntry, WalkDir};



fn is_hidden(entry: &DirEntry) -> bool {
    entry.file_name()
         .to_str()
         .map(|s| vec!["target", ".git"].contains(&s))
         .unwrap_or(false)
}


fn main() {

    let mut command = Command::new("ls");

    let mut child = command.spawn()
        .expect("Error in spawning");

    let event_fn = move | res: Result<notify::Event>| {
        match res {
            Ok(event) => {
                match event.kind {
                    notify::EventKind::Modify(notify::event::ModifyKind::Data(_)) => {

                            child.kill().expect("command couldn't be killed");
                            child = command.spawn()
                                .expect("Failed to restart");
                            
                        println!("event: {:?}", event)
                    },
                    _ => ()
                }

                
            },
            Err(e) => println!("watch error: {:?}", e),
        }
    };

    let mut watcher = notify::recommended_watcher(event_fn)
        .unwrap();


    for entry in WalkDir::new(env::current_dir().unwrap()).into_iter().filter_entry(|e| !is_hidden(e)) {
        let path = entry.unwrap();
        println!("Path: {}", path.path().display());
        let _ = watcher.watch(path.path(), notify::RecursiveMode::NonRecursive);
    }

    loop {}


}
