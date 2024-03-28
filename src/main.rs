use std::{env, path::Path, process::{Child, Command}, thread, time::Duration};

use config::Config;
use notify::{event::DataChange, Result, Watcher};
use walkdir::{DirEntry, WalkDir};

mod config;

fn is_hidden(entry: &DirEntry) -> bool {
    entry.file_name()
         .to_str()
         .map(|s| vec!["target", ".git", "bin"].contains(&s))
         .unwrap_or(false)
}


fn run_command(cmd: &str) -> Child {
    let argument:Vec<&str> = cmd.split(" ").collect::<Vec<&str>>();
    Command::new(&argument[0])
        .args(&argument[1..])
        .spawn()
        .unwrap()
}

fn main() {


    let cfg = config::new();

    run_command(&cfg.final_build_command);

    let mut command = run_command(&cfg.start_process);


    let event_fn = move | res: Result<notify::Event>| {
        match res {
            Ok(event) => {
                match event.kind {
                    notify::EventKind::Modify(notify::event::ModifyKind::Data(_)) => {

                            
                            command.kill().expect("command couldn't be killed");
                            command.wait().unwrap();


                            let mut build = run_command(&cfg.final_build_command);
                            let _ = build.wait();

                            command = run_command(&cfg.start_process); 
                            
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
