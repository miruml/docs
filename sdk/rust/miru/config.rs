// ... existing code ...
use std::fs::File;
use std::io::{self, Write};

pub struct Config<T> {
    data: Option<T>,
}

impl<T> Config<T> {
    pub fn new() -> Self {
        Config { data: None }
    }

    pub fn get(&self, obj: Option<T>, default: Option<T>) -> Option<&T> {
        self.data.as_ref().or(obj.as_ref()).or(default.as_ref())
    }

    pub fn write_to_file(&self, dest_file: &str) -> io::Result<()> {
        let mut file = File::create(dest_file)?;
        if let Some(ref data) = self.data {
            // Assuming T implements Display
            write!(file, "{:?}", data)?;
        }
        Ok(())
    }

    pub fn set(&mut self, obj: T) -> &T {
        self.data = Some(obj);
        self.data.as_ref().unwrap()
    }

    pub fn data(&self) -> Option<&T> {
        self.data.as_ref()
    }
}